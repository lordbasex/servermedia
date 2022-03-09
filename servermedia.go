package servermedia

import (
	"log"
	"net/http"
	"os"
)

/**
* Get env variables
**/
func goDotEnvVariable(key string) string {
	return os.Getenv(key)
}

var serverMediaPort string

func Start(serverMediaDir string) {

	serverMediaPort = goDotEnvVariable("SERVER_MEDIA_PORT")
	if serverMediaPort == "" {
		serverMediaPort = "8011"
	}

	if serverMediaDir == "" {
		serverMediaDir = "/audios"
	}

	fs := http.FileServer(http.Dir(serverMediaDir))
	http.Handle("/", fs)

	log.Println("Server Media: DIR: " + serverMediaDir + " | URL: http://0.0.0.0:" + serverMediaPort + " running...")

	err := http.ListenAndServe(":"+serverMediaPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
