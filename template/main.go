// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"net/http"

	"{{GoModule}}/handler"
	"{{GoModule}}/metrics"
	
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gopkg.in/gorp.v1"

	_ "github.com/joho/godotenv/autoload"
)

type spec struct {
	Bind  string `envconfig:"DRONE_BIND" default:":3000"`
	Debug bool   `envconfig:"DRONE_DEBUG"`
	Token string `envconfig:"DRONE_TOKEN"`

	Database struct {
		Datasource string `envconfig:"DRONE_DATABASE_DATASOURCE"`
		Driver     string `envconfig:"DRONE_DATABASE_DRIVER"`
	}
}

func main() {
	spec := new(spec)
	err := envconfig.Process("", spec)
	if err != nil {
		logrus.WithError(err).
			Fatalln("cannot read configuration from environment")
	}

	if spec.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// opens a connection with the database.
	db, err := sql.Open(spec.Database.Driver, spec.Database.Datasource)
	if err != nil {
		logrus.WithError(err).
			WithField("database", spec.Database.Driver).
			WithField("datasource", spec.Database.Datasource).
			Fatalln("cannot open database connection")
	}

	// creates an orm instance and configures the dialect
	// based on the database driver.
	dbm := &gorp.DbMap{Db: db}
	switch spec.Database.Driver {
	case "sqlite3":
		dbm.Dialect = gorp.SqliteDialect{}
	case "mysql":
		dbm.Dialect = gorp.MySQLDialect{}
	default:
		dbm.Dialect = gorp.PostgresDialect{}
	}

	// registers all prometheus metrics.
	metrics.Register(dbm)

	// registers the promethes handler.
	http.Handle("/", handler.New(spec.Token))

	logrus.Infof("server listening on address %s", spec.Bind)
	logrus.Fatal(http.ListenAndServe(spec.Bind, nil))
}
