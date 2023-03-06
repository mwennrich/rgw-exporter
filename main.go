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

func main() {

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	endpointURL := os.Getenv("CEPH_ENDPOINT_URL")
	queryEntriesEnv := os.Getenv("QUERY_ENTRIES")
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
	queryEntries := false
	if len(queryEntriesEnv) != 0 && queryEntriesEnv == strings.ToLower("true") {
		queryEntries = true
	}

	co, err := admin.New(endpointURL, accessKey, secretKey, &http.Client{Timeout: time.Second * 20})
	if err != nil {
		panic(err)
	}

	rgwCollector := newrgwCollector(co, queryEntries)
	prometheus.MustRegister(rgwCollector)

	// default to 15m
	statsSchedule := "@every 15m"
	if len(statsScheduleEnv) != 0 {
		statsSchedule = statsScheduleEnv
	}

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
			rgwCollector.collectStats()
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	klog.Info("Beginning to serve on port :9080")
	server := &http.Server{
		Addr:              ":9080",
		ReadHeaderTimeout: 300 * time.Second,
	}
	klog.Fatal(server.ListenAndServe())
}
