package logger

import (
    "testing"
)

func TestFileLogger(t *testing.T) {
    logger := NewFileLogger(LogLevelDebug, "d:/logs/", "test")
    logger.Debug("user id[%d] is come from China", 123123)
    logger.Warn("test warn log")
    logger.Fatal("test fatal log")
    logger.Close()
}

func TestConsoleLogger(t *testing.T) {
    logger := NewConsoleLogger(LogLevelDebug)
    logger.Debug("user id[%d] is come from China", 123123)
    logger.Warn("test warn log")
    logger.Fatal("test fatal log")
    logger.Close()
}