// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package metrics

import (
	"context"

	"gopkg.in/gorp.v1"
)

var noContext = context.Background()

// Register all prometheus metrics.
func Register(db *gorp.DbMap) {
	registerBuildCount(db)
	registerPendingBuildCount(db)
	registerRepoCount(db)
	registerRunningBuildCount(db)
	registerUserCount(db)
}
