package main

import (
    "project/log/lib"
)

func main() {
    //file := lib.NewFileLog(fmt.Sprintf("c:%sa.log", os.PathSeparator))
    //file.LogDebug("this is a debug log")
    //file.LogWarn("this is a warn log")

    //console := lib.NewConsoleLog("xxxx")
    //console.LogDebug("this is a debug log")
    //console.LogWarn("this is a warn log")

    //log := lib.NewFileLog("c:/a.log")
    log := lib.NewConsoleLog("xxxx")
    log.LogDebug("this is a debug log")
    log.LogWarn("this is a warn log")
}
