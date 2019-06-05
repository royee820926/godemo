package logger

import (
    "fmt"
    "os"
    "strconv"
    "time"
)

type FileLogger struct {
    level       int
    logPath     string
    logName     string

    file        *os.File
    warnFile    *os.File

    // 定义LogData队列，存放日志
    LogDataChan chan *LogData

    // 日志切分方式或大小
    logSplitType int
    logSplitSize int64
    lastSplitHour int
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
    // 日志拆分方式
    var logSplitType int = LogSplitTypeHour
    var logSplitSize int64
    logSplitStr, ok := config["log_split_type"]
    if !ok {
        logSplitStr = "home"
    }else {
        if logSplitStr == "size" {
            logSplitSizeStr, ok := config["log_split_size"]
            if !ok {
                logSplitSizeStr = "104857600"
            }
            logSplitSize, err := strconv.ParseInt(logSplitSizeStr, 10, 64)
            if err != nil {
                logSplitSize = 104857600
            }

            logSplitType = LogSplitTypeSize
        }else {
            logSplitType = LogSplitTypeHour
        }
    }

    log = &FileLogger{
        level: level,
        logPath: logPath,
        logName: logName,

        // 管道队列初始化
        LogDataChan: make(chan *LogData, chanSize),

        // 日志切分方式初始化
        logSplitSize: logSplitSize,
        logSplitType: logSplitType,
        lastSplitHour: time.Now().Hour(),
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

/**
 * 按小时拆分日志
 */
func (f *FileLogger) splitFileHour(warnFile bool) {
    now := time.Now()
    hour := now.Hour()
    if hour == f.lastSplitHour {
        return
    }

    f.lastSplitHour = hour
    var backupFileName string
    var fileName string

    if warnFile {
        backupFileName = fmt.Sprintf("%s/%s.log.wf_%04d%02d%02d%02d", f.logPath, f.logName,
            now.Year(), now.Month(), now.Day(), f.lastSplitHour)

        fileName = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
    }else {
        backupFileName = fmt.Sprintf("%s/%s.log_%04d%02d%02d%02d", f.logPath, f.logName,
            now.Year(), now.Month(), now.Day(), f.lastSplitHour)

        fileName = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
    }

    file := f.file
    if warnFile {
        file = f.warnFile
    }
    file.Close()
    // 备份当前日志文件
    os.Rename(fileName, backupFileName)

    file, err := os.OpenFile(fileName, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0755)
    if err != nil {
        return
    }

    // 更新当前日志
    if warnFile {
        f.warnFile = file
    }else {
        f.file = file
    }
}

/**
 * 按大小拆分日志
 */
func (f *FileLogger) splitFileSize(warnFile bool) {
    file := f.file
    if warnFile {
        file = f.warnFile
    }

    // 获取文件大小
    statInfo, err := file.Stat()
    if err != nil {
        return
    }
    fileSize := statInfo.Size()
    if fileSize <= f.logSplitSize {
        return
    }

    var backupFileName string
    var fileName string

    now := time.Now()
    if warnFile {
        backupFileName = fmt.Sprintf("%s/%s.log.wf_%04d%02d%02d%02d%02d%02d", f.logPath, f.logName,
            now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

        fileName = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
    }else {
        backupFileName = fmt.Sprintf("%s/%s.log_%04d%02d%02d%02d%02d%02d", f.logPath, f.logName,
            now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

        fileName = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
    }

    file.Close()
    // 备份当前日志文件
    os.Rename(fileName, backupFileName)

    file, err = os.OpenFile(fileName, os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0755)
    if err != nil {
        return
    }

    // 更新当前日志
    if warnFile {
        f.warnFile = file
    }else {
        f.file = file
    }
}

/**
 * 检查是否需要拆分日志
 */
func (f *FileLogger) checkSplitFile(warnFile bool) {
    if f.logSplitType == LogSplitTypeHour {
        f.splitFileHour(warnFile)
        return
    }

    f.splitFileSize(warnFile)
}

func (f *FileLogger) writeLogBackground() {
    // 从管道中把取出数据
    for logData := range f.LogDataChan {
        var file = f.file
        if logData.WarnAndFatal {
            file = f.warnFile
        }
        f.checkSplitFile(logData.WarnAndFatal)

        fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n",
            logData.TimeStr, logData.LevelStr, logData.FileName,
            logData.FuncName, logData.LineNo, logData.Message)
    }
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
