package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ceph/go-ceph/rgw/admin"
	"github.com/robfig/cron/v3"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	klog "k8s.io/klog/v2"
)

var (
	rgwC *rgwCollector
)

func main() {

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	endpointURL := os.Getenv("CEPH_ENDPOINT_URL")
	queryEntriesEnv := os.Getenv("QUERY_ENTRIES")

	usageScheduleEnv := os.Getenv("USAGE_SCHEDULE")
	statsScheduleEnv := os.Getenv("STATS_SCHEDULE")

	if len(accessKey) == 0 {
		panic("must provide AWS_ACCESS_KEY_ID")
	}
	if len(secretKey) == 0 {
		panic("must provide AWS_SECRET_ACCESS_KEY")
	}
	if len(endpointURL) == 0 {
		panic("must provide CEPH_ENDPOINT_URL")
	}

	// default false
	queryEntries := false
	if len(queryEntriesEnv) != 0 && queryEntriesEnv == strings.ToLower("true") {
		queryEntries = true
	}

	// default to 1m
	usageSchedule := "@every 1m"
	if len(usageScheduleEnv) != 0 {
		usageSchedule = usageScheduleEnv
	}

	// default to 15m
	statsSchedule := "@every 15m"
	if len(statsScheduleEnv) != 0 {
		statsSchedule = statsScheduleEnv
	}

	co, err := admin.New(endpointURL, accessKey, secretKey, &http.Client{Timeout: time.Second * 60})

	if err != nil {
		panic(err)
	}
	rgwRegistry := prometheus.NewRegistry()
	rgwC = newrgwCollector(co, queryEntries, rgwRegistry)
	rgwC.init()

	chUsage := make(chan struct{}, 1)
	go func() {
		cronUsage := cron.New()
		_, err := cronUsage.AddFunc(usageSchedule, func() {
			if len(chUsage) == 0 {
				chUsage <- struct{}{}
			}
		})
		if err != nil {
			klog.Fatalf("Invalid usage schedule: %s: %s\n", usageSchedule, err)
		}
		klog.Infof("Configured usage schedule: %s\n", usageSchedule)

		cronUsage.Start()

		if len(chUsage) == 0 {
			// send initial message
			chUsage <- struct{}{}
		}

		for {
			<-chUsage
			rgwC.collectUsage()
		}
	}()


	chStats := make(chan struct{}, 1)
	go func() {
		cronStats := cron.New()
		_, err := cronStats.AddFunc(statsSchedule, func() {
			if len(chStats) == 0 {
				chStats <- struct{}{}
			}
		})
		if err != nil {
			klog.Fatalf("Invalid stats schedule: %s: %s\n", statsSchedule, err)
		}
		klog.Infof("Configured stats schedule: %s\n", statsSchedule)

		time.Sleep(30 * time.Second)
		cronStats.Start()

		if len(chStats) == 0 {
			// send initial message
			chStats <- struct{}{}
		}

		for {
			<-chStats
			rgwC.collectStats()
		}
	}()

	// http.Handle("/metrics", promhttp.Handler())
	http.Handle("/metrics", promhttp.HandlerFor(
		rgwRegistry,
		promhttp.HandlerOpts{},
	))

	klog.Info("Beginning to serve on port :9080")
	server := &http.Server{
		Addr:              ":9080",
		ReadHeaderTimeout: 300 * time.Second,
	}
	klog.Fatal(server.ListenAndServe())
}
