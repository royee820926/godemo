package logger

const (
    LogLevelDebug = iota
    LogLevelTrace
    LogLevelInfo
    LogLevelWarn
    LogLevelError
    LogLevelFatal
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