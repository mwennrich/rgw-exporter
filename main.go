package main

import (
	"net/http"
	"os"

	"github.com/ceph/go-ceph/rgw/admin"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	klog "k8s.io/klog/v2"
)

func main() {

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	//region := os.Getenv("CEPH_REGIOM")
	endpointURL := os.Getenv("CEPH_ENDPOINT_URL")

	if len(accessKey) == 0 {
		panic("must provide accessKey")
	}

	co, err := admin.New(endpointURL, accessKey, secretKey, nil)

	if err != nil {
		panic(err)
	}
	rgwCollector := newrgwCollector(co)
	prometheus.MustRegister(rgwCollector)

	http.Handle("/metrics", promhttp.Handler())
	klog.Info("Beginning to serve on port :9080")
	klog.Fatal(http.ListenAndServe(":9080", nil))
}
