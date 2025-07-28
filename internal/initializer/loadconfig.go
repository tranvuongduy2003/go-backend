package initializer

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tranvuongduy2003/go-backend/global"
)

func LoadConfig() {
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

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Unable to decode configuration %w", err)
	}
}