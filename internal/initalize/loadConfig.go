package initalize

import (
	"go-simple-api/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	vp := viper.New()
	vp.AddConfigPath("./configs")
	vp.SetConfigName("local")
	vp.SetConfigType("yml")
	if err := vp.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := vp.Unmarshal(&global.Config); err != nil {
		panic(err)
	}

}
