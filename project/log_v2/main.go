package main

import (
    "project/log_v2/logger"
    "time"
)

func main() {
    initLogger("console", "d:/logs/", "user_server", "debug")
    Run()
    return
}

func initLogger(name, logPath, logName string, level string) (err error) {
    config := make(map[string]string, 8)
    config["log_path"] = logPath
    config["log_name"] = logName
    config["log_level"] = level

    // 选择输出到文件或终端
    //err = logger.InitLogger("file", config)
    //err = logger.InitLogger("console", config)
    err = logger.InitLogger(name, config)

    if err != nil {
        return
    }
    //log = logger.NewFileLogger(level, logPath, logName)
    //log = logger.NewConsoleLogger(level)
    logger.Debug("init logger success")
    return
}

func Run() {
    for {
        logger.Debug("user server is running")
        time.Sleep(time.Second)
    }
}