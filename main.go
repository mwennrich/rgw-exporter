package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ceph/go-ceph/rgw/admin"

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
	queryEntries := os.Getenv("QUERY_ENTRIES")

	if len(accessKey) == 0 {
		panic("must provide AWS_ACCESS_KEY_ID")
	}
	if len(secretKey) == 0 {
		panic("must provide AWS_SECRET_ACCESS_KEY")
	}
	if len(endpointURL) == 0 {
		panic("must provide CEPH_ENDPOINT_URL")
	}
	queryEntriesBool := false
	if len(queryEntries) != 0 && queryEntries == strings.ToLower("true") {
		queryEntriesBool = true
	}

	co, err := admin.New(endpointURL, accessKey, secretKey, &http.Client{Timeout: time.Second * 60})

	if err != nil {
		panic(err)
	}
	rgwRegistry := prometheus.NewRegistry()
	rgwC = newrgwCollector(co, queryEntriesBool,rgwRegistry)
	rgwC.init()

	// prometheus.MustRegister(rgwRegistry)

	go func() {
		for {
			rgwC.collect()
			time.Sleep(time.Second * 60)
		}
	}()

	go func() {
		for {
			rgwC.collectStats()
			time.Sleep(time.Minute * 15)
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
