package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)
type HTTPServer struct{
	Addr string  `yaml:"address" env-required:"true"`
}

//env-default:"production"   
type Config struct{
	Env string     `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string   `yaml:"storage_path" env-required:"true"`
	HTTPServer HTTPServer  `yaml:"http_server"`
}


func MustLoad() *Config{
	var configPath string 
	configPath = os.Getenv("CONFIG_PATH")   // as env variable

	if configPath == ""{     // passing through terminal as args
		flags:= flag.String("config", "", "path to config file")
		flag.Parse()
		configPath = *flags

		if configPath == ""{
			log.Fatal("config path is not set.")
		}
	}
	if _, err:= os.Stat(configPath); os.IsNotExist((err)){   // check if file exists at the path
		log.Fatalf("config file does not exist at path: %s", configPath)
	}

	var cfg Config
	err:= cleanenv.ReadConfig(configPath, &cfg)  // parse the config into the struct
	if err!= nil{
		log.Fatalf("failed to read config file: %s", err.Error())
	}

	return &cfg
}