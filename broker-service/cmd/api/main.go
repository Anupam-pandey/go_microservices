package main

import (
	"log"
	"net/http"
	"fmt"
)

const webPort = "8080"

type Config struct {}

func main()  {

	app := Config{}
	log.Printf("Starting broker services on port %s",webPort)
	
	// define http server
	
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s",webPort),
		Handler: app.routes(),
	}

	//start server

	err := srv.ListenAndServe()
	if err!=nil{
		log.Panic(err)
	}
}