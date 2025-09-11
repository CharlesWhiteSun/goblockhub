package handler_test

import (
	"os"
	"testing"

	"goblockhub/internal/logger"
)

func TestMain(m *testing.M) {
    logger.InitLogger(logger.DebugLevel)
    code := m.Run()
    os.Exit(code)
}