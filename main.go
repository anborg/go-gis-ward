package main

import (
	"flag"
	"os"
	"log"
	"github.com/anborg/go-gis-ward/api"
	"github.com/anborg/go-gis-ward/repo"
	"github.com/anborg/go-gis-ward/util"
)

const (
	GEO_FILE = "gis-wards.json"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	//read cmdline
	var configFile string
	flag.StringVar(&configFile, "configFile", "", "Provide config file path,  e.g c:/my/dir/eftconf.yml")
	flag.Parse()
	if configFile == "" {
		log.Println("configFile mandatory: Eg -configFile=/var/conf/config2.yml")
		os.Exit(1)
	}
	//Read config
	var config util.Config
	if err := config.New(configFile); err != nil {
		log.Fatalf("Error reading config file :", configFile, err)
	} else {
		log.Println("App port: ", config.App.Port)
		log.Println("Ward polygons file:", config.App.WardGeoJson)
	}

	var wardRepo repo.Ward
	if err1 := wardRepo.New(config.App.WardGeoJson); err1 != nil {
		log.Fatalf("Error reading ward file: ", config.App.WardGeoJson, err1)
	}
	server, err := api.NewServer(config, wardRepo)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	if err := server.Start(config.App.Port); err != nil {
		log.Fatal("cannot start server:", err)
	}

} //main
