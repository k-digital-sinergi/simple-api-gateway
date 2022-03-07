package config

import "github.com/caarlos0/env/v6"

type Environment struct {
	DbEnvironment
	RedisEnvironment
	GRPCEnvironment
	ServerAddress string `env:"ServerAddress"`
}

type DbEnvironment struct {
	CustomerDbEnvironment
	OrderDbEnvironment
}

type CustomerDbEnvironment struct {
	CustomerDbHost     string `env:"CustomerDbHost"`
	CustomerDbPort     string `env:"CustomerDbPort"`
	CustomerDbName     string `env:"CustomerDbName"`
	CustomerDbUser     string `env:"CustomerDbUser"`
	CustomerDbPassword string `env:"CustomerDbPassword"`
}

type OrderDbEnvironment struct {
	OrderDbHost     string `env:"OrderDbHost"`
	OrderDbPort     string `env:"OrderDbPort"`
	OrderDbName     string `env:"OrderDbName"`
	OrderDbUser     string `env:"OrderDbUser"`
	OrderDbPassword string `env:"OrderDbPassword"`
}

type RedisEnvironment struct {
	RedisAddress       string `env:"RedisAddress"`
	RedisPassword      string `env:"RedisPassword"`
	RedisSlave         string `env:"RedisSlave"`
	RedisSlavePassword string `env:"RedisSlavePassword"`
}

type GRPCEnvironment struct {
	CustomerGrpcAddr string `env:"CustomerGrpcAddr" envDefault:"127.0.0.1:50011"`
	OrderGrpcAddr    string `env:"OrderGrpcAddr" envDefault:"127.0.0.1:50012"`
}

var Env = Environment{}

func LoadEnv() error {
	return env.Parse(&Env)
}
