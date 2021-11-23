package gologger

import (
	"log"
	"os"
)

type Logger struct {
	LogName string
	LogPath string
	LogChan chan Log
	quit    chan bool
}

func (l *Logger) Initialize() error {
	l.LogChan = make(chan Log)
	l.quit = make(chan bool)
	// Check to see if the log directory exists
	if _, err := os.Stat(l.LogPath); os.IsNotExist(err) {
		// Create the log directory
		err = os.MkdirAll(l.LogPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Check to see if the log file exists
	if _, err := os.Stat(l.LogPath + l.LogName); os.IsNotExist(err) {
		// Create the log file
		f, err := os.Create(l.LogPath + l.LogName)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	go l.LogWorker()

	return nil
}

func (l *Logger) Shutdown() {
	l.quit <- true
}

func (l *Logger) LogWorker() {
	// Open the log file
	f, err := os.OpenFile(l.LogPath+l.LogName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	// Check to see if there was an error opening the file
	if err != nil {
		log.Fatal(err)
	}

	// Close the file when the function exits
	defer f.Close()

	for {
		select {
		case logMsg := <-l.LogChan:
			// Display the log message to console
			log.Println(logMsg.ConsoleString())

			// Write the log message to the log file
			f.WriteString(logMsg.String() + "\n")
		case <-l.quit:
			return
		}
	}
}

func (l *Logger) Debug(msg string) {
	l.LogChan <- Log{DEBUG, msg}
}

func (l *Logger) Info(msg string) {
	l.LogChan <- Log{INFO, msg}
}

func (l *Logger) Warn(msg string) {
	l.LogChan <- Log{WARN, msg}
}

func (l *Logger) Success(msg string) {
	l.LogChan <- Log{SUCC, msg}
}

func (l *Logger) Error(msg string) {
	l.LogChan <- Log{ERROR, msg}
}
