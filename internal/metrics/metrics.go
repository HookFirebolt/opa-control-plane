package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	BundleBuildFailed = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ocp_bundle_build_failed",
			Help: "Number of times a bundle has failed to build",
		},
		[]string{"bundle", "reason"},
	)

	BundleBuildCount = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "ocp_bundle_build_count",
			Help: "Total number of times a bundle has been built",
		},
	)

	BundleBuildDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ocp_bundle_build_duration_seconds",
			Help:    "Bundle build duration in seconds",
			Buckets: []float64{0.1, 0.2, 0.5, 1, 1.5, 2, 5, 10, 30, 60},
		},
		[]string{"bundle"},
	)

	GitSyncFailed = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ocp_git_sync_failed_total",
			Help: "Total number of failed Git sync operations",
		},
		[]string{"source", "reason"},
	)

	GitSyncCount = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "ocp_git_sync_count_total",
			Help: "Total number of Git sync operations",
		},
	)

	GitSyncDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ocp_git_sync_duration_seconds",
			Help:    "Git sync duration in seconds",
			Buckets: []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1, 1.5, 2, 5, 10, 30, 60},
		},
		[]string{"source", "repo"},
	)
)
