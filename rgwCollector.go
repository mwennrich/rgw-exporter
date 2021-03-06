package main

import (
	"context"

	"github.com/ceph/go-ceph/rgw/admin"
	"github.com/jinzhu/now"
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
	rgwBucketBytesReceived   *prometheus.Desc
	rgwBucketBytesSent       *prometheus.Desc
	rgwBucketSuccessfulOps   *prometheus.Desc
	rgwBucketOps             *prometheus.Desc
	rgw                      *admin.API
	queryEntries             bool
}

func newrgwCollector(rgw *admin.API, queryEntries bool) *rgwCollector {
	return &rgwCollector{
		rgwTotalBytesReceived: prometheus.NewDesc("rgw_total_bytes_received",
			"Shows rgw total received traffic in Bytes",
			[]string{"user", "s3_endpoint"}, nil,
		),
		rgwTotalBytesSent: prometheus.NewDesc("rgw_total_bytes_sent",
			"Shows rgw total sent traffic in Bytes",
			[]string{"user", "s3_endpoint"}, nil,
		),
		rgwTotalOps: prometheus.NewDesc("rgw_total_ops",
			"Shows rgw total ops",
			[]string{"user", "s3_endpoint"}, nil,
		),
		rgwTotalSuccessfulOps: prometheus.NewDesc("rgw_total_ops_successful",
			"Shows rgw total sucessfull ops",
			[]string{"user", "s3_endpoint"}, nil,
		),
		rgwCategoryBytesReceived: prometheus.NewDesc("rgw_category_bytes_received",
			"Shows rgw category received traffic in Bytes",
			[]string{"user", "s3_endpoint", "category"}, nil,
		),
		rgwCategoryBytesSent: prometheus.NewDesc("rgw_category_bytes_sent",
			"Shows rgw category sent traffic in Bytes",
			[]string{"user", "s3_endpoint", "category"}, nil,
		),
		rgwCategoryOps: prometheus.NewDesc("rgw_category_ops",
			"Shows rgw category ops",
			[]string{"user", "s3_endpoint", "category"}, nil,
		),
		rgwCategorySuccessfulOps: prometheus.NewDesc("rgw_category_ops_successful",
			"Shows rgw category sucessfull ops",
			[]string{"user", "s3_endpoint", "category"}, nil,
		),
		rgwBucketBytesReceived: prometheus.NewDesc("rgw_bucket_bytes_received",
			"Shows rgw bucket received traffic in Bytes",
			[]string{"user", "bucket", "owner", "s3_endpoint", "category"}, nil,
		),
		rgwBucketBytesSent: prometheus.NewDesc("rgw_bucket_bytes_sent",
			"Shows rgw bucket sent traffic in Bytes",
			[]string{"user", "bucket", "owner", "s3_endpoint", "category"}, nil,
		),
		rgwBucketOps: prometheus.NewDesc("rgw_bucket_ops",
			"Shows rgw bucket ops",
			[]string{"user", "bucket", "owner", "s3_endpoint", "category"}, nil,
		),
		rgwBucketSuccessfulOps: prometheus.NewDesc("rgw_bucket_ops_successful",
			"Shows rgw bucket sucessfull ops",
			[]string{"user", "bucket", "owner", "s3_endpoint", "category"}, nil,
		),
		rgw:          rgw,
		queryEntries: queryEntries,
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
	ch <- collector.rgwBucketBytesReceived
	ch <- collector.rgwBucketBytesSent
	ch <- collector.rgwBucketOps
	ch <- collector.rgwBucketSuccessfulOps
}

func (collector *rgwCollector) Collect(ch chan<- prometheus.Metric) {
	today := now.BeginningOfDay()
	usage, err := collector.rgw.GetUsage(context.Background(), admin.Usage{ShowSummary: ptr.BoolPtr(true), ShowEntries: ptr.BoolPtr(collector.queryEntries), Start: today.String()})
	if err != nil {
		panic(err)
	}
	if collector.queryEntries {
		for _, entry := range usage.Entries {
			for _, bucket := range entry.Buckets {
				for _, category := range bucket.Categories {
					ch <- prometheus.MustNewConstMetric(collector.rgwBucketBytesReceived, prometheus.CounterValue, float64(category.BytesReceived), entry.User, bucket.Bucket, bucket.Owner, collector.rgw.Endpoint, category.Category)
					ch <- prometheus.MustNewConstMetric(collector.rgwBucketBytesSent, prometheus.CounterValue, float64(category.BytesSent), entry.User, bucket.Bucket, bucket.Owner, collector.rgw.Endpoint, category.Category)
					ch <- prometheus.MustNewConstMetric(collector.rgwBucketOps, prometheus.CounterValue, float64(category.Ops), entry.User, bucket.Bucket, bucket.Owner, collector.rgw.Endpoint, category.Category)
					ch <- prometheus.MustNewConstMetric(collector.rgwBucketSuccessfulOps, prometheus.CounterValue, float64(category.SuccessfulOps), entry.User, bucket.Bucket, bucket.Owner, collector.rgw.Endpoint, category.Category)
				}
			}
		}
	}
	for _, user := range usage.Summary {
		ch <- prometheus.MustNewConstMetric(collector.rgwTotalBytesReceived, prometheus.CounterValue, float64(user.Total.BytesReceived), user.User, collector.rgw.Endpoint)
		ch <- prometheus.MustNewConstMetric(collector.rgwTotalBytesSent, prometheus.CounterValue, float64(user.Total.BytesSent), user.User, collector.rgw.Endpoint)
		ch <- prometheus.MustNewConstMetric(collector.rgwTotalOps, prometheus.CounterValue, float64(user.Total.Ops), user.User, collector.rgw.Endpoint)
		ch <- prometheus.MustNewConstMetric(collector.rgwTotalSuccessfulOps, prometheus.CounterValue, float64(user.Total.SuccessfulOps), user.User, collector.rgw.Endpoint)

		for _, category := range user.Categories {
			ch <- prometheus.MustNewConstMetric(collector.rgwCategoryBytesReceived, prometheus.CounterValue, float64(category.BytesReceived), user.User, collector.rgw.Endpoint, category.Category)
			ch <- prometheus.MustNewConstMetric(collector.rgwCategoryBytesSent, prometheus.CounterValue, float64(category.BytesSent), user.User, collector.rgw.Endpoint, category.Category)
			ch <- prometheus.MustNewConstMetric(collector.rgwCategoryOps, prometheus.CounterValue, float64(category.Ops), user.User, collector.rgw.Endpoint, category.Category)
			ch <- prometheus.MustNewConstMetric(collector.rgwCategorySuccessfulOps, prometheus.CounterValue, float64(category.SuccessfulOps), user.User, collector.rgw.Endpoint, category.Category)
		}
	}
}
