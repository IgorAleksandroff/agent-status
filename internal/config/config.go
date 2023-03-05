package config

import (
	"flag"
	"log"
	"os"
)

const (
	ServerAddressEnv     = "RUN_ADDRESS"
	ServerAddressDefault = "localhost:8080"

	ServerGRPSSocketEnv     = "GRPS_SOCKET"
	ServerDefaultGRPSSocket = ":3200"

	DataBaseAddressEnv     = "DATABASE_URI"
	DataBaseAddressDefault = ""
)

type (
	config struct {
		DataBaseURI string
		serverConfig
	}

	serverConfig struct {
		Host       string
		GRPSSocket string
	}
)

func GetConfig() config {
	hostFlag := flag.String("a", ServerAddressDefault, "адрес и порт сервера")
	DBFlag := flag.String("d", DataBaseAddressDefault, "адрес подключения к БД")
	socketFlag := flag.String("g", ServerDefaultGRPSSocket, "порт gRPC сервера")
	flag.Parse()

	cfg := config{
		DataBaseURI: getEnvString(DataBaseAddressEnv, *DBFlag),
		serverConfig: serverConfig{
			Host:       getEnvString(ServerAddressEnv, *hostFlag),
			GRPSSocket: getEnvString(ServerGRPSSocketEnv, *socketFlag),
		},
	}

	log.Printf("Parsed config: %+v", cfg)

	return cfg
}

func getEnvString(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		log.Printf("empty env: %s, default: %s", envName, defaultValue)
		return defaultValue
	}
	return value
}
