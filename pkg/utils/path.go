package utils

import (
	"log"
	"os"
	"os/user"
)

func BasePath() string {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return basePath
}

func ConfigPath() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return currentUser.HomeDir + "/.go-clean"
}
