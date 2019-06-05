package logger

const (
    LogLevelDebug = iota
    LogLevelTrace
    LogLevelInfo
    LogLevelWarn
    LogLevelError
    LogLevelFatal
)

const (
    LogSplitTypeHour = iota
    LogSplitTypeSize
)

func getLevelText(level int) string {
    switch level {
    case LogLevelDebug:
        return "DEBUG"
    case LogLevelTrace:
        return "Trace"
    case LogLevelInfo:
        return "INFO"
    case LogLevelWarn:
        return "Warn"
    case LogLevelError:
        return "Error"
    case LogLevelFatal:
        return "Fatal"
    default:
        return "UNKNOWN"
    }
}

func getLogLevel(level string) int {
    switch level {
    case "debug":
        return LogLevelDebug
    case "trace":
        return LogLevelTrace
    case "info":
        return LogLevelInfo
    case "warn":
        return LogLevelWarn
    case "error":
        return LogLevelError
    case "fatal":
        return LogLevelFatal
    default:
        return LogLevelDebug
    }
}