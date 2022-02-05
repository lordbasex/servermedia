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

func Start() {

        serverMediaPort = goDotEnvVariable("SERVER_MEDIA_PORT")
        if serverMediaPort == "" {
                serverMediaPort = "8011"
        }

        fs := http.FileServer(http.Dir("./audios"))
        http.Handle("/", fs)

        log.Println("Server Media: http://0.0.0.0:" + serverMediaPort + " running...")

        err := http.ListenAndServe(":"+serverMediaPort, nil)
        if err != nil {
                log.Fatal(err)
        }
}
