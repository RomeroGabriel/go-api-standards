package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBDriver      string `mapstrucute:"DB_DRIVER"`
	DBHost        string `mapstrucute:"DB_HOST"`
	DBPort        string `mapstrucute:"DB_PORT"`
	DBUser        string `mapstrucute:"DB_USER"`
	DBPassword    string `mapstrucute:"DB_PASSWORD"`
	DBName        string `mapstrucute:"DB_NAME"`
	WebServerPort string `mapstrucute:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstrucute:"JWT_SECRET"`
	JWTExperesIn  int    `mapstrucute:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}
