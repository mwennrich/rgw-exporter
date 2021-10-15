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

func main() {

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	endpointURL := os.Getenv("CEPH_ENDPOINT_URL")
	queryCategories := os.Getenv("QUERY_CATEGORIES")

	if len(accessKey) == 0 {
		panic("must provide AWS_ACCESS_KEY_ID")
	}
	if len(secretKey) == 0 {
		panic("must provide AWS_SECRET_ACCESS_KEY")
	}
	if len(endpointURL) == 0 {
		panic("must provide CEPH_ENDPOINT_URL")
	}
	queryCategoriesBool := false
	if len(queryCategories) != 0 && queryCategories == strings.ToLower("true") {
		queryCategoriesBool = true
	}

	co, err := admin.New(endpointURL, accessKey, secretKey, &http.Client{Timeout: time.Second * 60})

	if err != nil {
		panic(err)
	}
	rgwCollector := newrgwCollector(co, queryCategoriesBool)
	prometheus.MustRegister(rgwCollector)

	http.Handle("/metrics", promhttp.Handler())
	klog.Info("Beginning to serve on port :9080")
	klog.Fatal(http.ListenAndServe(":9080", nil))
}
