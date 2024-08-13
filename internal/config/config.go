package config

import (
	"os"
	"strconv"
)

const (
	DEVELOPER    = "developer"
	HOMOLOGATION = "homologarion"
	PRODUCTION   = "production"
)

// Se a letra for maiúsula, ela é pública, se for privada é o contrário
// Quando acessar o objeto config ele compõe também os do banco de dados
type Config struct {
	SRV_PORT    string
	WEB_UI      bool
	Mode        string
	OpenBrowser bool

	DBConfig
}

type DBConfig struct {
	DB_DRIVE string
	DB_HOST  string
	DB_PORT  string
	DB_USER  string
	DB_PASS  string
	DB_NAME  string
	DB_DSN   string
}

// Assim que é contruído um constructor
// Aqui nós estamos utilizando o ponteiro para ele apontar o espaço na mamória que ele vai substituir, e não vai ficar criando cópias na memória
func NewConfig(config *Config) *Config {
	var conf *Config

	if config == nil || config.SRV_PORT == "" {
		conf = defaultConfig()
	} else {
		conf = config
	}

	// É um pacote padrão do GO que pega as variáveis de ambiente do sistema
	// Variáveis de ambiente devem ter prioridade para configs

	SRV_PORT := os.Getenv("SRV_PORT")
	if SRV_PORT != "" {
		conf.SRV_PORT = SRV_PORT
	}

	SRV_WEB_UI := os.Getenv("SRV_WEB_UI")
	if SRV_WEB_UI != "" {
		// WEB_UI é esperado um booleano, mas o Getenv pega apenas resultados em string então será usado a função strconv para fazer a conversão dos dados
		// É colocado o _ para ignorar o erro, pois se der errado nós já sabemos o motivo
		// ParseBool returns the boolean value represented by the string. It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.
		conf.WEB_UI, _ = strconv.ParseBool(SRV_WEB_UI)
	}

	SRV_OPEN_BROWSER := os.Getenv("SRV_OPEN_BROWSER")
	if SRV_OPEN_BROWSER != "" {
		conf.OpenBrowser, _ = strconv.ParseBool(SRV_OPEN_BROWSER)
	}

	SRV_MODE := os.Getenv("SRV_MODE")
	if SRV_MODE != "" {
		conf.Mode = SRV_MODE
	}

	// Configs do DB
	SRV_DB_DRIVE := os.Getenv("SRV_DB_DRIVE")
	if SRV_DB_DRIVE != "" {
		conf.DB_DRIVE = SRV_DB_DRIVE
	}

	SRV_DB_HOST := os.Getenv("SRV_DB_HOST")
	if SRV_DB_HOST != "" {
		conf.DB_HOST = SRV_DB_HOST
	}

	SRV_DB_PORT := os.Getenv("SRV_DB_PORT")
	if SRV_DB_PORT != "" {
		conf.DB_PORT = SRV_DB_PORT
	}

	SRV_DB_USER := os.Getenv("SRV_DB_USER")
	if SRV_DB_USER != "" {
		conf.DB_USER = SRV_DB_USER
	}

	SRV_DB_PASS := os.Getenv("SRV_DB_PASS")
	if SRV_DB_PASS != "" {
		conf.DB_PASS = SRV_DB_PASS
	}

	SRV_DB_NAME := os.Getenv("SRV_DB_NAME")
	if SRV_DB_NAME != "" {
		conf.DB_NAME = SRV_DB_NAME
	}

	return conf
}

func defaultConfig() *Config {
	default_config := Config{
		SRV_PORT:    "8080",
		WEB_UI:      true,
		OpenBrowser: true,
		Mode:        PRODUCTION,
		DBConfig: DBConfig{
			DB_DRIVE: "sqlite3",
			// DB_HOST: "postgres",
			// DB_PORT: "192.168.18.2",
			// DB_USER: "posgres",
			// DB_PASS: "123456",
			// DB_NAME: "stoq",
			DB_NAME: "db.sqlite3",
		},
	}

	return &default_config

}