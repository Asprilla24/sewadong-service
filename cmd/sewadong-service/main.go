// Copyright (c) 2019 Braggart Inc. All Rights Reserved.
// This is licensed software from Braggart Inc, for limitations
// and restrictions contact your company contract manager.

package main

import (
	"os"
	"os/signal"
	"syscall"

	"sewadong-service/pkg/config"
	"sewadong-service/pkg/dao"
	"sewadong-service/pkg/dao/postgres"
	"sewadong-service/pkg/server"
	"sewadong-service/pkg/service"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetOutput(os.Stdout)

	serverConfig := config.GetConfig()

	dbConn, err := dao.New("postgres", serverConfig)
	if err != nil {
		logrus.WithFields(logrus.Fields{"module": "main"}).
			Fatal("unable to connect database: ", err)
	}
	defer dbConn.Close() // nolint : errcheck, used in defer

	err = dao.MigrateDB()
	if err != nil {
		logrus.WithFields(logrus.Fields{"module": "main"}).
			Fatal("unable to migrate database: ", err)
	}

	dbClient := postgres.NewDB(dbConn)

	restService := service.New(dbClient)
	restServer := server.New(serverConfig, restService)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		logrus.Info("received os quit signal")
		restServer.Stop()
	}()

	restServer.Serve()
}
