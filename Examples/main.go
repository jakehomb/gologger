package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jakehomb/gologger"
)

var logger gologger.Logger

func main() {
	SetupCloseHandler()

	logger = gologger.Logger{LogPath: "C:\\Temp\\", LogName: "LogFile.txt"}

	logger.Initialize()
	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			logger.Debug("This is a debug message")

			time.Sleep(time.Second * time.Duration(rand.Intn(2)))
		}
	}()

	go func() {
		for {
			logger.Info("This is an information message")

			time.Sleep(time.Second * time.Duration(rand.Intn(3)))
		}
	}()

	go func() {
		for {
			logger.Warn("This is a warning message")

			time.Sleep(time.Second * time.Duration(rand.Intn(4)))
		}
	}()

	go func() {
		for {
			logger.Success("This is a success message")

			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		}
	}()

	go func() {
		for {
			logger.Error("This is an error message")

			time.Sleep(time.Second * time.Duration(rand.Intn(6)))
		}
	}()

	for {
		time.Sleep(time.Second * 1)
	}
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		logger.Info("Ctrl+C pressed in terminal")
		logger.Shutdown()
		os.Exit(0)
	}()
}
