package main

import (
	"ghuettoize"
	"github.com/sirupsen/logrus"
)

func main() {
	ghettoizeHook, err := ghuettoize.NewGhettoize()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.AddHook(ghettoizeHook)

	logrus.Info("Starting application")

	logrus.Info("Creating entry in the database")

	logrus.Info("Entry created")

	logrus.Error("Can't send a message to the user")

}
