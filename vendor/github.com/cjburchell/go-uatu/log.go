package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/cjburchell/tools-go/trace"
)

// Level of the log
type Level struct {
	// Text representation of the log
	Text string
	// Severity value of the log
	Severity int
}

var (
	// DEBUG log level
	DEBUG = Level{Text: "Debug", Severity: 0}
	// INFO log level
	INFO = Level{Text: "Info", Severity: 1}
	// WARNING log level
	WARNING = Level{Text: "Warning", Severity: 2}
	// ERROR log level
	ERROR = Level{Text: "Error", Severity: 3}
	// FATAL log level
	FATAL = Level{Text: "Fatal", Severity: 4}
)

var levels = []Level{DEBUG,
	INFO,
	WARNING,
	ERROR,
	FATAL,
}

// GetLogLevel gets the log level for input text
func GetLogLevel(levelText string) Level {
	for i := range levels {
		if levels[i].Text == levelText {
			return levels[i]
		}
	}

	return INFO
}

// Warnf Print a formatted warning level message
func Warnf(format string, v ...interface{}) {
	printLog(fmt.Sprintf(format, v...), WARNING)
}

// Warn Print a warning message
func Warn(v ...interface{}) {
	printLog(fmt.Sprint(v...), WARNING)
}

// Error Print a error level message
func Error(err error, v ...interface{}) {
	printErrorLog(err, fmt.Sprint(v...), ERROR)
}

// Errorf Print a formatted error level message
func Errorf(err error, format string, v ...interface{}) {
	printErrorLog(err, fmt.Sprintf(format, v...), ERROR)
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func printErrorLog(err error, msg string, level Level) {
	if err == nil {
		printLog(msg, level)
	}

	if msg == "" {
		msg = fmt.Sprintf("Error: %s\n", err.Error())
	} else {
		msg = fmt.Sprintf("%s\nError: %s\n", msg, err.Error())
	}

	if err, ok := err.(stackTracer); ok {
		msg += "Stack Trace -----------------------------------------------------------------------------------------\n"
		for _, f := range err.StackTrace() {
			msg += fmt.Sprintf("%+v\n", f)
		}
		msg += "-----------------------------------------------------------------------------------------------------"
	} else {
		msg += trace.GetStack(2)
	}

	printLog(msg, level)
}

// Fatal print fatal level message
func Fatal(err error, v ...interface{}) {
	printErrorLog(err, fmt.Sprint(v...), FATAL)
	log.Panic(v...)
}

// Fatalf print formatted fatal level message
func Fatalf(err error, format string, v ...interface{}) {
	printErrorLog(err, fmt.Sprintf(format, v...), FATAL)
	log.Panicf(format, v...)
}

// Debug print debug level message
func Debug(v ...interface{}) {
	printLog(fmt.Sprint(v...), DEBUG)
}

// Debugf print formatted debug level  message
func Debugf(format string, v ...interface{}) {
	printLog(fmt.Sprintf(format, v...), DEBUG)
}

// Print print info level message
func Print(v ...interface{}) {
	printLog(fmt.Sprint(v...), INFO)
}

// Printf print info level message
func Printf(format string, v ...interface{}) {
	printLog(fmt.Sprintf(format, v...), INFO)
}

var hostname, _ = os.Hostname()

type Publisher interface {
	Publish(messageBites []byte) error
}

//var natsConn *nats.Conn
//var restClient *http.Client

// Settings for sending logs
type Settings struct {
	ServiceName  string
	MinLogLevel  Level
	LogToConsole bool
}

var settings = Settings{
	MinLogLevel:  DEBUG,
	LogToConsole: true,
}
var publishers []Publisher

// Setup the logging system
func Setup(newSettings Settings, newPublishers []Publisher) (err error) {
	settings = newSettings
	publishers = newPublishers
	return err
}

// Message to be sent to centralized logger
type Message struct {
	Text        string `json:"text"`
	Level       Level  `json:"level"`
	ServiceName string `json:"serviceName"`
	Time        int64  `json:"time"`
	Hostname    string `json:"hostname"`
}

func (message Message) String() string {
	return fmt.Sprintf("[%s] %s %s - %s", message.Level.Text, time.Unix(message.Time/1000, 0).Format("2006-01-02 15:04:05 MST"), message.ServiceName, message.Text)
}

func printLog(text string, level Level) {
	message := Message{
		Text:        text,
		Level:       level,
		ServiceName: settings.ServiceName,
		Time:        time.Now().UnixNano() / 1000000,
		Hostname:    hostname,
	}

	if level.Severity >= settings.MinLogLevel.Severity && settings.LogToConsole {
		if strings.HasSuffix(message.String(), "\n") {
			fmt.Print(message.String())
		} else {
			fmt.Println(message.String())
		}
	}

	if publishers == nil {
		return
	}

	messageBites, err := json.Marshal(message)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, publisher := range publishers {
		err = publisher.Publish(messageBites)
		if err != nil {
			fmt.Printf("Unable to send log to publisher (%s): %s", err.Error(), message.String())
		}
	}

}

type Writer struct {
	Level Level
}

func (w Writer) Write(p []byte) (n int, err error) {
	printLog(string(p), w.Level)
	return len(p), nil
}
