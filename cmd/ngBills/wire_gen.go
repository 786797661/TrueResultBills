// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"ngBills/internal/biz"
	"ngBills/internal/conf"
	"ngBills/internal/data"
	"ngBills/internal/server"
	"ngBills/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	newBillsRepo := data.NewGreeterRepo(dataData, logger)
	newBillsResultBizSer := biz.NewBillsResultBiz(newBillsRepo, logger)
	newBillsResultBizServiceSer := service.NewBillsResultBizService(newBillsResultBizSer)
	httpServer := server.NewHTTPServer(confServer, newBillsResultBizServiceSer, logger)
	grpcServer := server.NewGRPCServer(confServer, newBillsResultBizServiceSer, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}