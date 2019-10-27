// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package dao

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"sewadong-service/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"github.com/sirupsen/logrus"
)

const (
	dbTimeOutConnection = 30 * time.Second
)

// New creates new instance of DAO for database operations
func New(dbDriver string, conf *config.Config) (*gorm.DB, error) {
	connConfig := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s connect_timeout=5 sslmode=disable",
		conf.DBHost, conf.DBPort, conf.DBName, conf.DBUsername, conf.DBPassword)

	timeout := time.Now().Add(dbTimeOutConnection)
	var postgresORM *gorm.DB
	var err error
	retryCounter := 0

	for time.Now().Before(timeout) {
		postgresORM, err = gorm.Open(dbDriver, connConfig)
		if err == nil {
			break
		}
		retryCounter++
	}

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: timeout: %v", err)
	}

	if postgresORM == nil {
		return nil, fmt.Errorf("unable to initiate DAO: %v", err)
	}

	return postgresORM, nil
}

// MigrateDB for migrating db using Active Record
func MigrateDB() error {
	logger := logrus.WithFields(logrus.Fields{"module": "postgres"})

	env := os.Getenv("RAILS_ENV")
	if len(env) == 0 {
		env = "sewadong-service"
	}

	err := os.Setenv("RAILS_ENV", env)
	if err != nil {
		logger.Fatal("unable to set environment variables: ", err)
	}

	logger.Infoln("Migrating...")

	cmd := exec.Command("bundle", "exec", "rake", "db:migrate") // #nosec
	cmdOutput, err := cmd.CombinedOutput()
	logger.Infoln(string(cmdOutput))
	if err != nil {
		logger.Errorf("Failure during migration: %v\n", err)
	}

	logger.Infoln("Migration complete")
	return err
}
