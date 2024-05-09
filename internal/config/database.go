package config

type DatabaseConfig struct {
	DBName            string
	Host              string
	MaxIdleConnection int
	MaxOpenConnection int
	Password          string
	Port              int
	Username          string
}

func getDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		DBName:            fatalGetString("SQL_DB_NAME"),
		Host:              fatalGetString("SQL_HOST"),
		MaxIdleConnection: fatalGetInt("SQL_MAX_IDLE_CONNECTION"),
		MaxOpenConnection: fatalGetInt("SQL_MAX_OPEN_CONNECTION"),
		Password:          fatalGetString("SQL_PASSWORD"),
		Port:              fatalGetInt("SQL_PORT"),
		Username:          fatalGetString("SQL_USERNAME"),
	}
}
