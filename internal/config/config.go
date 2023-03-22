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

	MailAddressEnv  = "MAIL_HOST"
	MailUsernameEnv = "MAIL_USER"
	MailPasswordEnv = "MAIL_PASSWORD"
	MailFromEnv     = "MAIL_FROM"
	MailToEnv       = "MAIL_TO"

	AutoAssigmentGRPSSocketEnv     = "GRPS_SOCKET"
	DefaultAutoAssigmentGRPSSocket = ":3220"

	QueueSizeDefault = 100
)

type (
	config struct {
		DataBaseURI             string
		QueueSize               int64
		AutoAssigmentGRPSSocket string
		ServerConfig
		MessengerConfig
	}

	ServerConfig struct {
		Host       string
		GRPSSocket string
	}

	MessengerConfig struct {
		Host, Username, Password, From, To string
	}
)

func GetConfig() config {
	hostFlag := flag.String("a", ServerAddressDefault, "адрес и порт сервера")
	DBFlag := flag.String("d", DataBaseAddressDefault, "адрес подключения к БД")
	socketFlag := flag.String("g", ServerDefaultGRPSSocket, "порт gRPC сервера")

	mailHostFlag := flag.String("h", "", "адрес почтового сервера")
	mailUsernameFlag := flag.String("u", "", "пользователь от почтового ящика")
	mailPwdFlag := flag.String("p", "", "пароль от почтового ящика")
	mailFromFlag := flag.String("f", "", "почтовый ящик отправителя")
	mailToFlag := flag.String("t", "", "почтовый ящик получателя")
	autoAssigmentSocketFlag := flag.String("s", DefaultAutoAssigmentGRPSSocket, "gRPC сервер:порт")
	flag.Parse()

	cfg := config{
		DataBaseURI:             getEnvString(DataBaseAddressEnv, *DBFlag),
		QueueSize:               QueueSizeDefault,
		AutoAssigmentGRPSSocket: getEnvString(AutoAssigmentGRPSSocketEnv, *autoAssigmentSocketFlag),
		ServerConfig: ServerConfig{
			Host:       getEnvString(ServerAddressEnv, *hostFlag),
			GRPSSocket: getEnvString(ServerGRPSSocketEnv, *socketFlag),
		},
		MessengerConfig: MessengerConfig{
			Host:     getEnvString(MailAddressEnv, *mailHostFlag),
			Username: getEnvString(MailUsernameEnv, *mailUsernameFlag),
			Password: getEnvString(MailPasswordEnv, *mailPwdFlag),
			From:     getEnvString(MailFromEnv, *mailFromFlag),
			To:       getEnvString(MailToEnv, *mailToFlag),
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
