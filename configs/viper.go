package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	WisperServer string `mapstructure:"wisper_server"`
	Name         string `mapstructure:"name"`
	ConfigPath   string `mapstructure:"config_path"`
}

var Config *config

func Init() {
	Config = LoadConfig()
}

func LoadConfig() (c *config) {
	home, _ := os.UserHomeDir()
	wisperConfigPath := filepath.Join(home, ".config", "wisper")
	wisperConfigName := "config"

	if _, err := os.Stat(wisperConfigPath); os.IsNotExist(err) {
		os.Mkdir(wisperConfigPath, 0755)
	}

	if _, err := os.Stat(filepath.Join(wisperConfigPath, wisperConfigName)); os.IsNotExist(
		err,
	) {
		file, err := os.Create(filepath.Join(wisperConfigPath, wisperConfigName))
		if err != nil {
			log.Fatalf("Error creating config file, %s", err)
		}
		defer file.Close()
	}

	viper.AddConfigPath(wisperConfigPath)
	viper.SetConfigName(wisperConfigName)
	viper.SetConfigType("yaml")

	viper.SetDefault("wisper_server", "http://127.0.0.1:3000/")
	viper.SetDefault("name", "wisper")
	viper.SetDefault("config_path", filepath.Join(wisperConfigPath, wisperConfigName))

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.WriteConfigAs(filepath.Join(wisperConfigPath, wisperConfigName))

	_ = viper.Unmarshal(&c)
	return
}
