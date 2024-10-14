package config

const (
	developmentConfigFilePath = "./config/config.json"
)

type Config struct {
	Environment string         `json:"environment"`
	Server      ServerConfig   `json:"server"`
	Database    DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type DatabaseConfig struct {
	DatabaseType             string `json:"database_type"`
	DatabaseConnectionFormat string `json:"database_connection_format"`
	DatabaseName             string `json:"database_name"`
	Username                 string `json:"username"`
	Password                 string `json:"password"`
	Host                     string `json:"host"`
	Port                     string `json:"port"`
}
