package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
)

var dateFormat = "2006/01/02 15:04:05"

func main() {

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: dateFormat,
	})

	var sigChannel = make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	// get user opts
	var cmdPath, promAddr string
	var pollInterval time.Duration
	var v bool
	flag.StringVar(&cmdPath, "cmd-path", "/usr/sbin/apcaccess", "absolute path to pwstat command")
	flag.StringVar(&promAddr, "prom-addr", ":9400", "bind address of the prom http server")
	flag.DurationVar(&pollInterval, "poll-interval", time.Second*5, "time interval to gather power stats")
	flag.BoolVar(&v, "version", false, "print version")
	flag.BoolVar(&v, "v", false, "print version")

	flag.Parse()

	if v {
		var verPrinter = printer.New()
		var info = version.Get()
		if err := verPrinter.PrintInfo(os.Stdout, info); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())

		var server = &http.Server{
			Addr:         promAddr,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		if err := server.ListenAndServe(); err != nil {
			log.Fatal("http server error: ", err)
		}
	}()
	fmt.Println("started, go to grafana to monitor")

	gatherAndSaveStats(cmdPath)

	var ticker = time.NewTicker(pollInterval)
	for {
		select {
		case <-ticker.C:
			gatherAndSaveStats(cmdPath)

		case <-sigChannel:
			log.Info("shutting down")
			return
		}
	}
}

func gatherAndSaveStats(cmdPath string) {
	out, err := getPowerStats(cmdPath)
	if err != nil {
		log.Error(err)
	}

	batteryData, err := parse(out)
	if err != nil {
		log.Error(err)
		fmt.Printf("[%s]", out)
		fmt.Println()
	}

	if batteryData.Status == "ONLINE" {
		statusGauge.WithLabelValues(batteryData.Model).Set(0)
	} else if batteryData.Status == "ONBATT" {
		statusGauge.WithLabelValues(batteryData.Model).Set(1)
	}

	linevGauge.WithLabelValues(batteryData.Model).Set(float64(batteryData.Linev))
	loadPctGauge.WithLabelValues(batteryData.Model).Set(float64(batteryData.Loadpct))
	bchargeGauge.WithLabelValues(batteryData.Model).Set(float64(batteryData.Bcharge))
	remainingRuntimeGauge.WithLabelValues(batteryData.Model).Set(float64(batteryData.Timeleft.Seconds()))
	mbattchgGauge.WithLabelValues(batteryData.Model).Set(float64(batteryData.Mbattchg))
	mintimelGauge.WithLabelValues(batteryData.Model).Set(float64(batteryData.Mintimel.Seconds()))
	maxtimelGauge.WithLabelValues(batteryData.Model).Set(float64(batteryData.Maxtime.Seconds()))
	lotrans.WithLabelValues(batteryData.Model).Set(float64(batteryData.Lotrans))
	hitrans.WithLabelValues(batteryData.Model).Set(float64(batteryData.Hitrans))
	battv.WithLabelValues(batteryData.Model).Set(float64(batteryData.Battv))
	tonbatt.WithLabelValues(batteryData.Model).Set(float64(batteryData.Tonbatt.Seconds()))
	cumonbatt.WithLabelValues(batteryData.Model).Set(float64(batteryData.Cumonbatt.Seconds()))
	nominv.WithLabelValues(batteryData.Model).Set(float64(batteryData.Nominv))
	nombattv.WithLabelValues(batteryData.Model).Set(float64(batteryData.Nombattv))
	nompower.WithLabelValues(batteryData.Model).Set(float64(batteryData.Nompower))
}
