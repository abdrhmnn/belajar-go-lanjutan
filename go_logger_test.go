package belajargolanjutan

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestCreateLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello logger!")
	fmt.Println("Hello logger!")
}

// level logging
func TestLevelLogger(t *testing.T) {
	logger := logrus.New()

	// by default informasi yg keluar itu dari logger info ke atas
	// bisa set level mulainya
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("Test")
	logger.Debug("Test")
	logger.Info("Test")
	logger.Warn("Test")
	logger.Error("Test")
}

// logger output
func TestOutput(t *testing.T) {
	// output nya merupakan tipe io.Writer

	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	// by default format output dalam bentuk text
	// bisa diubah menjadi format json
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello!")
	logger.Warn("Hello!")
	logger.Error("Hello!")
}

// logger with field
func TestLoggerWithField(t *testing.T) {
	// field digunakan untuk menambahkan informasi logging
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("key1", "abdu").Info("DUAR!!!!")
	logger.WithField("key1", "eunha").
		WithField("key2", "sowon").Warn("DUARR WARN!!!")

	// atau bisa sekaligus input
	logger.WithFields(logrus.Fields{
		"key3": "yerin",
		"key4": "umji",
	}).Error("DUARRR ERROR!!!")
}
