package config

const (
	configFilePath = "./credentials/config.json"
)

type Config struct {
	Environment string         `json:"environment"`
	Database    DatabaseConfig `json:"database"`
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
