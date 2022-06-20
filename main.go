package main

import (
	"github.com/qwertyqq2/apiServer"
	"log"
	"flag"
	"github.com/BurntSushi/toml"
)

var(
	configPath string
)

func init(){
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "conf path")
}

func main(){

	flag.Parse()
	

	config := apiServer.NewConfig()
	_,err := toml.DecodeFile(configPath, config)
	if err!=nil{
		log.Fatal(err)
	}

	server:=apiServer.New(config)
	if err:=server.Start();err!=nil{
		log.Fatal(err)
	}
}
