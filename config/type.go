package config

import "strings"

type EnvMode string

func (e EnvMode) IsEqual(comparator string) bool {
	return strings.EqualFold(string(e), comparator)
}

func GetEnvMode(mode string) EnvMode {
	switch mode {
	case "staging":
		return Staging
	case "prod":
		return Production
	default:
		return Development
	}
}

const (
	Development EnvMode = "development"
	Production  EnvMode = "production"
	Staging     EnvMode = "staging"
)

type Config struct {
	Server struct {
		Host    string `json:"host" yaml:"host"`
		Version string `json:"version" yaml:"version"`
	} `json:"server" yaml:"server"`

	DBManagementSystem DBConnection `json:"db_management_system" yaml:"db_management_system"`
}

type DBConnection struct {
	Host           string `json:"host" yaml:"host"`
	Port           string `json:"port" yaml:"port"`
	Username       string `json:"username" yaml:"username"`
	Password       string `json:"password" yaml:"password"`
	DBName         string `json:"db_name" yaml:"db_name"`
	MigrationsFile string `json:"migrations_file" yaml:"migrations_file"`
}
