package main

import "fmt"

type Config struct {
	dsn string
}

func NewConfig() *Config {
	return &Config{
		dsn: "数据库连接",
	}
}

type DataBase struct {
	config *Config
}

func NewDB(config *Config) *DataBase {
	return &DataBase{
		config: config,
	}
}

func (b DataBase) Find() {
	fmt.Println("hahahah")
}

type Service struct {
	dateBase *DataBase
}

func NewService(dateBase *DataBase) *Service {
	return &Service{
		dateBase: dateBase,
	}
}

type App struct {
	s *Service
}

func NewApp(service *Service) *App {
	return &App{
		s: service,
	}
}

func main() {
	app := InitializeApp()
	app.s.dateBase.Find()
}
