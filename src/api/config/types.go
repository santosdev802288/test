package config

import (
	"time"

	kafka "dev.azure.com/SiigoDevOps/Siigo/_git/go-slim.git/kafka"
)

type Configuration struct {
	// configuration for kafka broker
	// https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
	AppName        string             `yaml:"appName"`
	Kafka          map[string]string  `yaml:"kafka"`
	Mongo          MongoConfiguration `yaml:"mongo"`
	Redis          RedisConfiguration `yaml:"redis"`
	EventStore     EventStore         `yaml:"eventStore"`
	BusinessLogger BusinessLogger     `yaml:"businessLogger"`
	KeyVault       *AzureVaultConfig  `yaml:"keyVault"`
	Grpc           Grpc               `yaml:"grpc"`
	HttpServer     HttpServer         `yaml:"httpServer"`
}

//
type MongoConfiguration struct {
	ConnectionString string `yaml:"connectionString" env:"MONGO_CONNECTION_STRING"`
	Database         string `yaml:"database" env:"MONGO_DATABASE"`
	Collection       string `yaml:"collection" env:"MONGO_COLLECTION"`
}

type RedisConfiguration struct {
	Addr     string        `yaml:"addr" env:"REDIS_ADDR"`
	Password string        `yaml:"password" env:"REDIS_PASSWORD"`
	Db       int           `yaml:"db" env:"REDIS_DB"`
	TimeOut  time.Duration `env:"REDIS_TIMEOUT" env-default:"0s" yaml:"timeOut"`
}

type EventStore struct {
	Url string `yaml:"url"`
}

type BusinessLogger struct {
	Enabled       bool              `yaml:"enabled"`
	LokiUrl       string            `yaml:"lokiUrl"`
	DefaultLabels map[string]string `yaml:"defaultLabels"`
}

type AzureVaultConfig struct {
	AzureClientId     *string `yaml:"azureClientId"`
	AzureClientSecret *string `yaml:"azureClientSecret"`
	AzureTenantId     *string `yaml:"azureTenantId"`
	VaultName         string  `yaml:"vaultName"`

	Kafka *kafka.KafkaCerts `yaml:"kafka"`
}

type Grpc struct {
	Host string `env:"HOST"      env-default:"[::]" yaml:"host"`
	Port int    `env:"GRPC_PORT" env-default:"10000" yaml:"port"`

	ServerMinTime time.Duration `env:"GRPC_SERVER_MIN_TIME" env-default:"5m" yaml:"serverMinTime"` // if a client pings more than once every 5 minutes (default), terminate the connection
	ServerTime    time.Duration `env:"GRPC_SERVER_TIME" env-default:"2h" yaml:"serverTime"`        // ping the client if it is idle for 2 hours (default) to ensure the connection is still active
	ServerTimeout time.Duration `env:"GRPC_SERVER_TIMEOUT" env-default:"20s" yaml:"serverTimeout"` // wait 20 second (default) for the ping ack before assuming the connection is dead
	ConnTime      time.Duration `env:"GRPC_CONN_TIME" env-default:"10s" yaml:"connTime"`           // send pings every 10 seconds if there is no activity
	ConnTimeout   time.Duration `env:"GRPC_CONN_TIMEOUT" env-default:"20s" yaml:"connTimeout"`     // wait 20 second for ping ack before considering the connection dead
}

type HttpServer struct {
	Port int `env:"HTTP_PORT" env-default:"11000" yaml:"port"`
}
