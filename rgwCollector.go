package main

import (
	"context"

	"github.com/ceph/go-ceph/rgw/admin"
	"github.com/jinzhu/now"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/klog/v2"
	ptr "k8s.io/utils/pointer"
)

type rgwCollector struct {
	rgwTotalBytesReceived    *prometheus.GaugeVec
	rgwTotalBytesSent        *prometheus.GaugeVec
	rgwTotalSuccessfulOps    *prometheus.GaugeVec
	rgwTotalOps              *prometheus.GaugeVec
	rgwTotalBytes            *prometheus.GaugeVec
	rgwTotalObjects          *prometheus.GaugeVec
	rgwCategoryBytesReceived *prometheus.GaugeVec
	rgwCategoryBytesSent     *prometheus.GaugeVec
	rgwCategorySuccessfulOps *prometheus.GaugeVec
	rgwCategoryOps           *prometheus.GaugeVec
	rgwBucketBytesReceived   *prometheus.GaugeVec
	rgwBucketBytesSent       *prometheus.GaugeVec
	rgwBucketSuccessfulOps   *prometheus.GaugeVec
	rgwBucketOps             *prometheus.GaugeVec
	rgwBucketBytes           *prometheus.GaugeVec
	rgwBucketObjects         *prometheus.GaugeVec
	rgw                      *admin.API
	queryEntries             bool
	registry                 *prometheus.Registry
}

func newrgwCollector(rgw *admin.API, queryEntries bool, registry *prometheus.Registry) *rgwCollector {
	return &rgwCollector{
		rgwTotalBytesReceived: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_total_bytes_received",
			Help: "Shows rgw total received traffic in Bytes",
		},
			[]string{"user", "s3_endpoint"},
		),
		rgwTotalBytesSent: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_total_bytes_sent",
			Help: "Shows rgw total sent traffic in Bytes",
		},
			[]string{"user", "s3_endpoint"},
		),
		rgwTotalOps: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_total_ops",
			Help: "Shows rgw total ops",
		},
			[]string{"user", "s3_endpoint"},
		),
		rgwTotalBytes: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_total_bytes",
			Help: "Shows rgw total bytes",
		},
			[]string{"user", "s3_endpoint"},
		),
		rgwTotalObjects: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_total_objects",
			Help: "Shows rgw total number of objects",
		},
			[]string{"user", "s3_endpoint"},
		),
		rgwTotalSuccessfulOps: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_total_ops_successful",
			Help: "Shows rgw total successful ops",
		},
			[]string{"user", "s3_endpoint"},
		),
		rgwCategoryBytesReceived: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_category_bytes_received",
			Help: "Shows rgw category received traffic in Bytes",
		},
			[]string{"user", "s3_endpoint", "category"},
		),
		rgwCategoryBytesSent: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_category_bytes_sent",
			Help: "Shows rgw category sent traffic in Bytes",
		},
			[]string{"user", "s3_endpoint", "category"},
		),
		rgwCategoryOps: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_category_ops",
			Help: "Shows rgw category ops",
		},
			[]string{"user", "s3_endpoint", "category"},
		),
		rgwCategorySuccessfulOps: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_category_ops_successful",
			Help: "Shows rgw category successful ops",
		},
			[]string{"user", "s3_endpoint", "category"},
		),
		rgwBucketBytesReceived: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_bucket_bytes_received",
			Help: "Shows rgw bucket received traffic in Bytes",
		},
			[]string{"user", "bucket", "owner", "s3_endpoint", "category"},
		),
		rgwBucketBytesSent: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_bucket_bytes_sent",
			Help: "Shows rgw bucket sent traffic in Bytes",
		},
			[]string{"user", "bucket", "owner", "s3_endpoint", "category"},
		),
		rgwBucketOps: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_bucket_ops",
			Help: "Shows rgw bucket ops",
		},
			[]string{"user", "bucket", "owner", "s3_endpoint", "category"},
		),
		rgwBucketSuccessfulOps: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_bucket_ops_successful",
			Help: "Shows rgw bucket successful ops",
		},
			[]string{"user", "bucket", "owner", "s3_endpoint", "category"},
		),
		rgwBucketBytes: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_bucket_bytes",
			Help: "Shows rgw bucket bytes",
		},
			[]string{"user", "bucket", "s3_endpoint"},
		),
		rgwBucketObjects: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "rgw_bucket_objects",
			Help: "Shows rgw bucket number of objects",
		},
			[]string{"user", "bucket", "s3_endpoint"},
		),
		rgw:          rgw,
		queryEntries: queryEntries,
		registry:     registry,
	}
}

func (collector *rgwCollector) init() {
	collector.registry.MustRegister(
		collector.rgwTotalBytesReceived,
		collector.rgwTotalBytesSent,
		collector.rgwTotalOps,
		collector.rgwTotalSuccessfulOps,
		collector.rgwTotalBytes,
		collector.rgwTotalObjects,
		collector.rgwCategoryBytesReceived,
		collector.rgwCategoryBytesSent,
		collector.rgwCategoryOps,
		collector.rgwCategorySuccessfulOps,
		collector.rgwBucketBytesReceived,
		collector.rgwBucketBytesSent,
		collector.rgwBucketOps,
		collector.rgwBucketSuccessfulOps,
		collector.rgwBucketBytes,
		collector.rgwBucketObjects,
	)
}

func (collector *rgwCollector) collectUsage() {
	today := now.BeginningOfDay()
	usage, err := collector.rgw.GetUsage(context.Background(), admin.Usage{ShowSummary: ptr.Bool(true), ShowEntries: ptr.Bool(collector.queryEntries), Start: today.String()})
	if err != nil {
		klog.Errorf("failed to fetch usage date: %w", err)
		return
	}
	if collector.queryEntries {
		for _, entry := range usage.Entries {
			for _, bucket := range entry.Buckets {
				for _, category := range bucket.Categories {
					collector.rgwBucketBytesReceived.WithLabelValues(entry.User, bucket.Bucket, bucket.Owner, collector.rgw.Endpoint, category.Category).Set(float64(category.BytesReceived))
					collector.rgwBucketBytesSent.WithLabelValues(entry.User, bucket.Bucket, bucket.Owner, collector.rgw.Endpoint, category.Category).Set(float64(category.BytesSent))
					collector.rgwBucketOps.WithLabelValues(entry.User, bucket.Bucket, bucket.Owner, collector.rgw.Endpoint, category.Category).Set(float64(category.Ops))
					collector.rgwBucketSuccessfulOps.WithLabelValues(entry.User, bucket.Bucket, bucket.Owner, collector.rgw.Endpoint, category.Category).Set(float64(category.SuccessfulOps))
				}
			}
		}
	}
	for _, user := range usage.Summary {
		collector.rgwTotalBytesReceived.WithLabelValues(user.User, collector.rgw.Endpoint).Set(float64(user.Total.BytesReceived))
		collector.rgwTotalBytesSent.WithLabelValues(user.User, collector.rgw.Endpoint).Set(float64(user.Total.BytesSent))
		collector.rgwTotalOps.WithLabelValues(user.User, collector.rgw.Endpoint).Set(float64(user.Total.Ops))
		collector.rgwTotalSuccessfulOps.WithLabelValues(user.User, collector.rgw.Endpoint).Set(float64(user.Total.SuccessfulOps))

		for _, category := range user.Categories {
			collector.rgwCategoryBytesReceived.WithLabelValues(user.User, collector.rgw.Endpoint, category.Category).Set(float64(category.BytesReceived))
			collector.rgwCategoryBytesSent.WithLabelValues(user.User, collector.rgw.Endpoint, category.Category).Set(float64(category.BytesSent))
			collector.rgwCategoryOps.WithLabelValues(user.User, collector.rgw.Endpoint, category.Category).Set(float64(category.Ops))
			collector.rgwCategorySuccessfulOps.WithLabelValues(user.User, collector.rgw.Endpoint, category.Category).Set(float64(category.SuccessfulOps))
		}
	}
}

func (collector *rgwCollector) collectStats() {
	users, err := collector.rgw.GetUsers(context.Background())
	if err != nil || users == nil {
		klog.Errorf("failed to fetch stats: %w", err)
		return
	}
	for _, user := range *users {
		stats, err := collector.rgw.ListUsersBucketsWithStat(context.Background(), user)
		if err != nil {
			klog.Errorf("failed to fetch stats: %w", err)
			continue
		}

		var size, numObjects uint64
		for _, bucket := range stats {

			if bucket.Usage.RgwMain.SizeActual == nil || bucket.Usage.RgwMain.NumObjects == nil {
				continue
			}

			size += *bucket.Usage.RgwMain.SizeActual
			numObjects += *bucket.Usage.RgwMain.NumObjects
			if collector.queryEntries {
				collector.rgwBucketBytes.WithLabelValues(user, bucket.Bucket, collector.rgw.Endpoint).Set(float64(*bucket.Usage.RgwMain.SizeActual))
				collector.rgwBucketObjects.WithLabelValues(user, bucket.Bucket, collector.rgw.Endpoint).Set(float64(*bucket.Usage.RgwMain.NumObjects))
			}
		}
		collector.rgwTotalBytes.WithLabelValues(user, collector.rgw.Endpoint).Set(float64(size))
		collector.rgwTotalObjects.WithLabelValues(user, collector.rgw.Endpoint).Set(float64(numObjects))
	}
}
