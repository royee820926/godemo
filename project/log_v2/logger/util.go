package logger

import (
    "fmt"
    "path"
    "runtime"
    "time"
)

type LogData struct {
	Message  string
	TimeStr  string
	LevelStr string
	FileName string
	FuncName string
	LineNo   int
	WarnAndFatal bool
}

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	// skip: 调用层级
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}

func writeLog(level int, format string, args ...interface{}) *LogData {

	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(level)

	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)

	//fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
	logData := &LogData{
	    Message: msg,
	    TimeStr: nowStr,
	    LevelStr: levelStr,
	    FileName: fileName,
	    FuncName: funcName,
        LineNo: lineNo,
        WarnAndFatal: false,
    }
	if level == LogLevelError || level == LogLevelWarn || level == LogLevelFatal {
		logData.WarnAndFatal = true
	}
	return logData
	//fmt.Fprintf(file, "%s %s () %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
}
