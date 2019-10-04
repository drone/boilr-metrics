// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/gorp.v1"
)

func registerBuildCount(db *gorp.DbMap) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_build_count",
			Help: "Total number of builds.",
		}, func() float64 {
			i, _ := db.SelectFloat("SELECT count(*) FROM builds")
			return i
		}),
	)
}

func registerPendingBuildCount(db *gorp.DbMap) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_pending_builds",
			Help: "Total number of pending builds.",
		}, func() float64 {
			i, _ := db.SelectFloat("SELECT count(*) FROM builds WHERE build_status = 'pending'")
			return i
		}),
	)
}

func registerRunningBuildCount(db *gorp.DbMap) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_running_builds",
			Help: "Total number of running builds.",
		}, func() float64 {
			i, _ := db.SelectFloat("SELECT count(*) FROM builds WHERE build_status = 'running'")
			return i
		}),
	)
}
