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
	rgwTotalBytes            *prometheus.Desc
	rgwTotalObjects          *prometheus.Desc
	rgwCategoryBytesReceived *prometheus.Desc
	rgwCategoryBytesSent     *prometheus.Desc
	rgwCategorySuccessfulOps *prometheus.Desc
	rgwCategoryOps           *prometheus.Desc
	rgwBucketBytesReceived   *prometheus.Desc
	rgwBucketBytesSent       *prometheus.Desc
	rgwBucketSuccessfulOps   *prometheus.Desc
	rgwBucketOps             *prometheus.Desc
	rgwBucketBytes           *prometheus.Desc
	rgwBucketObjects         *prometheus.Desc
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
		rgwTotalBytes: prometheus.NewDesc("rgw_total_bytes",
			"Shows rgw total bytes",
			[]string{"user", "s3_endpoint"}, nil,
		),
		rgwTotalObjects: prometheus.NewDesc("rgw_total_objects",
			"Shows rgw total number of objects",
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
			"Shows rgw category successfull ops",
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
		rgwBucketBytes: prometheus.NewDesc("rgw_bucket_bytes",
			"Shows rgw bucket bytes",
			[]string{"user", "bucket", "s3_endpoint"}, nil,
		),
		rgwBucketObjects: prometheus.NewDesc("rgw_bucket_objects",
			"Shows rgw bucket number of objects",
			[]string{"user", "bucket", "s3_endpoint"}, nil,
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
	ch <- collector.rgwTotalBytes
	ch <- collector.rgwTotalObjects
	ch <- collector.rgwCategoryBytesReceived
	ch <- collector.rgwCategoryBytesSent
	ch <- collector.rgwCategoryOps
	ch <- collector.rgwCategorySuccessfulOps
	ch <- collector.rgwBucketBytesReceived
	ch <- collector.rgwBucketBytesSent
	ch <- collector.rgwBucketOps
	ch <- collector.rgwBucketSuccessfulOps
	ch <- collector.rgwBucketBytes
	ch <- collector.rgwBucketObjects
}

func (collector *rgwCollector) Collect(ch chan<- prometheus.Metric) {
	today := now.BeginningOfDay()
	usage, err := collector.rgw.GetUsage(context.Background(), admin.Usage{ShowSummary: ptr.Bool(true), ShowEntries: ptr.Bool(collector.queryEntries), Start: today.String()})
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
		stats, err := collector.rgw.ListUsersBucketsWithStat(context.Background(), user.User)
		if err != nil {
			panic(err)
		}
		var size, numObjects uint64
		for _, bucket := range stats {
			size += *bucket.Usage.RgwMain.SizeActual
			numObjects += *bucket.Usage.RgwMain.NumObjects
			if collector.queryEntries {
				ch <- prometheus.MustNewConstMetric(collector.rgwBucketBytes, prometheus.GaugeValue, float64(*bucket.Usage.RgwMain.SizeActual), user.User, bucket.Bucket, collector.rgw.Endpoint)
				ch <- prometheus.MustNewConstMetric(collector.rgwBucketObjects, prometheus.GaugeValue, float64(*bucket.Usage.RgwMain.NumObjects), user.User, bucket.Bucket, collector.rgw.Endpoint)
			}
			ch <- prometheus.MustNewConstMetric(collector.rgwTotalBytes, prometheus.GaugeValue, float64(size), user.User, collector.rgw.Endpoint)
			ch <- prometheus.MustNewConstMetric(collector.rgwTotalObjects, prometheus.GaugeValue, float64(numObjects), user.User, collector.rgw.Endpoint)
		}
	}
}
