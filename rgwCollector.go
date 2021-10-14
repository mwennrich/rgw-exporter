package main

import (
	"context"

	"github.com/ceph/go-ceph/rgw/admin"
	"github.com/prometheus/client_golang/prometheus"
	ptr "k8s.io/utils/pointer"
)

type rgwCollector struct {
	rgwTotalBytes *prometheus.Desc
	rgw           *admin.API
}

func newrgwCollector(rgw *admin.API) *rgwCollector {
	return &rgwCollector{
		rgwTotalBytes: prometheus.NewDesc("rgw_total_bytes",
			"Shows rgw total usage in Bytes",
			[]string{"rgw_total_bytes_received", "user", "endpoint"}, nil,
		),
		rgw: rgw,
	}
}

func (collector *rgwCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.rgwTotalBytes
}

func (collector *rgwCollector) Collect(ch chan<- prometheus.Metric) {
	usage, err := collector.rgw.GetUsage(context.Background(), admin.Usage{ShowSummary: ptr.BoolPtr(true), ShowEntries: ptr.BoolPtr(false)})
	if err != nil {
		panic(err)
	}
	for _, user := range usage.Summary {
		ch <- prometheus.MustNewConstMetric(collector.rgwTotalBytes, prometheus.GaugeValue, float64(user.Total.BytesReceived), user.User, collector.rgw.Endpoint)
	}
}
