package logger

import (
    "fmt"
    "os"
    "strconv"
)

type FileLogger struct {
    level       int
    logPath     string
    logName     string

    file        *os.File
    warnFile    *os.File

    // 定义LogData队列，存放日志
    LogDataChan chan *LogData
}

func NewFileLogger(config map[string]string) (log LogInterface, err error) {
    // 日志存储路径
    logPath, ok := config["log_path"]
    if !ok {
        err = fmt.Errorf("not found log_path")
        return
    }
    // 日志文件名
    logName, ok := config["log_name"]
    if !ok {
        err = fmt.Errorf("not found log_name")
        return
    }
    // 日志级别
    logLevel, ok := config["log_level"]
    if !ok {
        err = fmt.Errorf("not found log_level")
        return
    }
    level := getLogLevel(logLevel)
    // 日志队列
    logChanSize, ok := config["log_chan_size"]
    if !ok {
        logChanSize = "50000"
    }
    chanSize, err := strconv.Atoi(logChanSize)
    if err != nil {
        chanSize = 50000
    }

    log = &FileLogger{
        level: level,
        logPath: logPath,
        logName: logName,

        LogDataChan: make(chan *LogData, chanSize),
    }
    log.Init()
    return
}

func (f *FileLogger) Init() {
    // 普通日志
    filename := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
    file, err := os.OpenFile(filename, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0755)
    if err != nil {
        panic(fmt.Sprintf("open file %s failed, err:%v", filename, err))
    }
    f.file = file

    // 错误日志
    filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
    file, err = os.OpenFile(filename, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0755)
    if err != nil {
        panic(fmt.Sprintf("open file %s failed, err:%v", filename, err))
    }

    f.warnFile = file

    // 后台写入日志
    go f.writeLogBackground()
}

func (f *FileLogger) writeLogBackground() {
    // todo
}

func (f *FileLogger) SetLevel(level int)  {
    if level < LogLevelDebug || level > LogLevelFatal {
        level = LogLevelDebug
    }
    f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
    if f.level > LogLevelDebug {
        return
    }
    logData := writeLog(LogLevelDebug, format, args...)
    select {
    case f.LogDataChan <- logData:
    default:
    }
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
    if f.level > LogLevelTrace {
        return
    }
    logData := writeLog(LogLevelTrace, format, args...)
    select {
    case f.LogDataChan <- logData:
    default:
    }
}

func (f *FileLogger) Info(format string, args ...interface{}) {
    if f.level > LogLevelInfo {
        return
    }
    logData := writeLog(LogLevelInfo, format, args...)
    select {
    case f.LogDataChan <- logData:
    default:
    }
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
    if f.level > LogLevelWarn {
        return
    }
    logData := writeLog(LogLevelWarn, format, args...)
    select {
    case f.LogDataChan <- logData:
    default:
    }
}

func (f *FileLogger) Error(format string, args ...interface{}) {
    if f.level > LogLevelError {
        return
    }
    logData := writeLog(LogLevelError, format, args...)
    select {
    case f.LogDataChan <- logData:
    default:
    }
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
    if f.level > LogLevelFatal {
        return
    }
    logData := writeLog(LogLevelFatal, format, args...)
    select {
    case f.LogDataChan <- logData:
    default:
    }
}

func (f *FileLogger) Close()  {
    f.file.Close()
    f.warnFile.Close()
}
