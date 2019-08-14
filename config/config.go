package config

import(
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// init parse config file
// 2019/08/08 10:40:52
func init() {
	configFile := pflag.StringP("config", "c", "config.toml", "config file path")
	pflag.Parse()

	viper.SetConfigFile(*configFile)
	viper.ReadInConfig()
}
