package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// nolint: gochecknoglobals
var (
	promNamespace = "apc_exporter"

	statusGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "status",
		Help:      "0=Normal / 1=Power Failure",
	}, []string{"model_name"})

	linevGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "linev",
		Help:      "Current input line voltage",
	}, []string{"model_name"})

	bchargeGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "bcharge",
		Help:      "Percentage of UPS load capacity used as estimated by UPS",
	}, []string{"model_name"})

	remainingRuntimeGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "remaining_runtime",
		Help:      "Remaining runtime left on battery as estimated by the UPS",
	}, []string{"model_name"})

	mbattchgGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "mbattchg",
		Help:      "Min battery charge % (BCHARGE) required for system shutdown",
	}, []string{"model_name"})

	mintimelGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "mintimel",
		Help:      "Min battery runtime (MINUTES) required for system shutdown",
	}, []string{"model_name"})

	maxtimelGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "maxtimel",
		Help:      "Max battery runtime (TIMEOUT) after which system is shutdown",
	}, []string{"model_name"})

	lotrans = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "lotrans",
		Help:      "Input line voltage below which UPS will switch to battery",
	}, []string{"model_name"})

	hitrans = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "hitrans",
		Help:      "Input line voltage above which UPS will switch to battery",
	}, []string{"model_name"})

	battv = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "battv",
		Help:      "Current battery voltage",
	}, []string{"model_name"})

	tonbatt = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "tonbatt",
		Help:      "Seconds currently on battery",
	}, []string{"model_name"})

	cumonbatt = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "cumonbatt",
		Help:      "Cumulative seconds on battery since apcupsd startup",
	}, []string{"model_name"})

	nominv = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "nominv",
		Help:      "Nominal input voltage",
	}, []string{"model_name"})

	nombattv = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "nombattv",
		Help:      "Nominal battery voltage",
	}, []string{"model_name"})

	nompower = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "nompower",
		Help:      "Nominal power output in watts",
	}, []string{"model_name"})

	loadPctGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: promNamespace,
		Name:      "load_pct",
		Help:      "current load as %",
	}, []string{"model_name"})
)
