//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeApp() *App {
	wire.Build(NewConfig, NewDB, NewService, NewApp)
	return new(App)
}
