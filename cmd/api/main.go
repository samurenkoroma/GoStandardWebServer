package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/srm-developer/standardWebServer/internal/app/api"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	flag.Parse()
	config := api.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("can not find configs file. using default values: ", err)
	}
	server := api.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	//log.Fatal(server.Start()) аналогично предыдущему блоку
}
