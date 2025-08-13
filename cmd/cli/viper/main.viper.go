package viper

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
}

func Test() {
	vp := viper.New()
	vp.AddConfigPath("./configs")
	vp.SetConfigName("production")
	vp.SetConfigType("yml")
	if err := vp.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := vp.Unmarshal(&config); err != nil {
		panic(err)
	}

}
