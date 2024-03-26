package models

type DBConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSLMODE  string
	DB_TIMEZONE string
}

type ServerConfig struct {
	SERVER_PORT   string
	SERVER_PREFIX string
}
