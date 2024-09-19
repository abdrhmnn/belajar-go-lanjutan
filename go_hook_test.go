package belajargolanjutan

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

// hook merupakan callback yang akan di eksekusi setelah logging dari level tertentu selesai dilakukan
// hook itu interface jadi harus bikin struct dulu untuk implement interface Hook di logrus

// implement interface Hook
type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Berhasil berjalan!", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Warn("Testing")
}

// singleton
func TestSingleton(t *testing.T) {
	// ini akan mengakses global objek New() dari logrus
	logrus.Info("Testing")
	logrus.Warn("Testing")
}
