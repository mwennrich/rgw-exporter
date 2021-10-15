package main

import (
	"context"

	"github.com/ceph/go-ceph/rgw/admin"
	"github.com/prometheus/client_golang/prometheus"
	ptr "k8s.io/utils/pointer"
)

type rgwCollector struct {
	rgwTotalBytesReceived    *prometheus.Desc
	rgwTotalBytesSent        *prometheus.Desc
	rgwTotalSuccessfulOps    *prometheus.Desc
	rgwTotalOps              *prometheus.Desc
	rgwCategoryBytesReceived *prometheus.Desc
	rgwCategoryBytesSent     *prometheus.Desc
	rgwCategorySuccessfulOps *prometheus.Desc
	rgwCategoryOps           *prometheus.Desc
	rgw                      *admin.API
	queryCategories          bool
}

func newrgwCollector(rgw *admin.API, queryCategories bool) *rgwCollector {
	return &rgwCollector{
		rgwTotalBytesReceived: prometheus.NewDesc("rgw_total_bytes_received",
			"Shows rgw total received traffic in Bytes",
			[]string{"user", "endpoint"}, nil,
		),
		rgwTotalBytesSent: prometheus.NewDesc("rgw_total_bytes_sent",
			"Shows rgw total sent traffic in Bytes",
			[]string{"user", "endpoint"}, nil,
		),
		rgwTotalOps: prometheus.NewDesc("rgw_total_ops",
			"Shows rgw total ops",
			[]string{"user", "endpoint"}, nil,
		),
		rgwTotalSuccessfulOps: prometheus.NewDesc("rgw_total_ops_successful",
			"Shows rgw total sucessfull ops",
			[]string{"user", "endpoint"}, nil,
		),
		rgwCategoryBytesReceived: prometheus.NewDesc("rgw_category_bytes_received",
			"Shows rgw category received traffic in Bytes",
			[]string{"user", "endpoint", "category"}, nil,
		),
		rgwCategoryBytesSent: prometheus.NewDesc("rgw_category_bytes_sent",
			"Shows rgw category sent traffic in Bytes",
			[]string{"user", "endpoint", "category"}, nil,
		),
		rgwCategoryOps: prometheus.NewDesc("rgw_category_ops",
			"Shows rgw category ops",
			[]string{"user", "endpoint", "category"}, nil,
		),
		rgwCategorySuccessfulOps: prometheus.NewDesc("rgw_category_ops_successful",
			"Shows rgw category sucessfull ops",
			[]string{"user", "endpoint", "category"}, nil,
		),
		rgw:             rgw,
		queryCategories: queryCategories,
	}
}

func (collector *rgwCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.rgwTotalBytesReceived
	ch <- collector.rgwTotalBytesSent
	ch <- collector.rgwTotalOps
	ch <- collector.rgwTotalSuccessfulOps
	ch <- collector.rgwCategoryBytesReceived
	ch <- collector.rgwCategoryBytesSent
	ch <- collector.rgwCategoryOps
	ch <- collector.rgwCategorySuccessfulOps
}

func (collector *rgwCollector) Collect(ch chan<- prometheus.Metric) {
	usage, err := collector.rgw.GetUsage(context.Background(), admin.Usage{ShowSummary: ptr.BoolPtr(true), ShowEntries: ptr.BoolPtr(false)})
	if err != nil {
		panic(err)
	}
	for _, user := range usage.Summary {
		ch <- prometheus.MustNewConstMetric(collector.rgwTotalBytesReceived, prometheus.GaugeValue, float64(user.Total.BytesReceived), user.User, collector.rgw.Endpoint)
		ch <- prometheus.MustNewConstMetric(collector.rgwTotalBytesSent, prometheus.GaugeValue, float64(user.Total.BytesSent), user.User, collector.rgw.Endpoint)
		ch <- prometheus.MustNewConstMetric(collector.rgwTotalOps, prometheus.GaugeValue, float64(user.Total.Ops), user.User, collector.rgw.Endpoint)
		ch <- prometheus.MustNewConstMetric(collector.rgwTotalSuccessfulOps, prometheus.GaugeValue, float64(user.Total.SuccessfulOps), user.User, collector.rgw.Endpoint)

		if collector.queryCategories {
			for _, category := range user.Categories {
				ch <- prometheus.MustNewConstMetric(collector.rgwCategoryBytesReceived, prometheus.GaugeValue, float64(category.BytesReceived), user.User, collector.rgw.Endpoint, category.Category)
				ch <- prometheus.MustNewConstMetric(collector.rgwCategoryBytesSent, prometheus.GaugeValue, float64(category.BytesSent), user.User, collector.rgw.Endpoint, category.Category)
				ch <- prometheus.MustNewConstMetric(collector.rgwCategoryOps, prometheus.GaugeValue, float64(category.Ops), user.User, collector.rgw.Endpoint, category.Category)
				ch <- prometheus.MustNewConstMetric(collector.rgwCategorySuccessfulOps, prometheus.GaugeValue, float64(category.SuccessfulOps), user.User, collector.rgw.Endpoint, category.Category)
			}
		}
	}
}
