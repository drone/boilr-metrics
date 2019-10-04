// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/gorp.v1"
)

func registerUserCount(db *gorp.DbMap) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "drone_user_count",
			Help: "Total number of active users.",
		}, func() float64 {
			i, _ := db.SelectFloat("SELECT count(*) FROM users")
			return i
		}),
	)
}
