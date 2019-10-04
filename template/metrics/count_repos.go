// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/gorp.v1"
)

func registerRepoCount(db *gorp.DbMap) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_repo_count",
			Help: "Total number of registered repositories.",
		}, func() float64 {
			i, _ := db.SelectFloat("SELECT count(*) FROM repos WHERE repo_active = true")
			return i
		}),
	)
}
