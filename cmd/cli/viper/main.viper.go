package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
	Security struct {
		JWT struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config file: %w \n", err))
	}

	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Port::", viper.GetInt("security.jwt.key"))

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Unable to decode configuration %w", err)
	}

	fmt.Println("Config Port::", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("Database User: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)
	}
}