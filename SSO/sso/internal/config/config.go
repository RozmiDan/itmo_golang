package config

import (
	"io"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	env         string 				`yaml:"env" env-default:"local"`
	storagePath string				`yaml:"storage_path" env-requried:"true"`
	tokenTTL    time.Duration `yaml:"token_ttl" env-default:"10s"`
	grpcConf 		GRPCConfig		`yaml:"grpc"`
}

type GRPCConfig struct {
	port        int						`yaml:"port"`
	timeout     time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	readFile := ""
	
	io.Rea
	cleanenv.ParseYAML()
	return 
}