package iniconfig

import (
    "fmt"
    "io/ioutil"
    "testing"
)

type Config struct {
    ServerConf ServerConfig `ini:"server"`
    MysqlConf  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
    Ip   string `ini:"ip"`
    Port int    `ini:"port"`
}

type MysqlConfig struct {
    Username string  `ini:"username"`
    Passwd   string  `ini:"passwd"`
    Database string  `ini:"database"`
    Host     string  `ini:"host"`
    Port     int     `ini:"port"`
    Timeout  float32 `ini:"timeout"`
}

func TestIniConfig(t *testing.T) {
    fmt.Println("hello")
    data, err := ioutil.ReadFile("./config.ini")
    if err != nil {
        t.Error("read file failed")
    }

    // 测试UnMarshal
    var conf Config
    err = UnMarshal(data, &conf)
    if err != nil {
        t.Errorf("unmarshal failed, err:%v", err)
    }

    t.Logf("unmarshal success, conf:%#v", conf)

    // 测试Marshal
    //fmt.Printf("%v\n", conf)
    confData, err := Marshal(conf)
    if err != nil {
       t.Errorf("marshal failed, err:%v", err)
    }

    t.Logf("marshal success, conf:%#v", string(confData))

    // 测试Marshal写入文件
    //MarshalFile(conf, "d:/logs/test.conf")
}

func TestIniConfigFile(t *testing.T) {

    filename := "d:/logs/test.conf"
    var conf Config
    conf.ServerConf.Ip = "localhost"
    conf.ServerConf.Port = 8888

    // 测试UnMarshal
    err := MarshalFile(filename, conf)
    if err != nil {
        t.Errorf("marshal failed, err:%v", err)
    }

    var conf2 Config

    // 测试Marshal
    //fmt.Printf("%v\n", conf)
    err = UnMarshalFile(filename, &conf2)
    if err != nil {
        t.Errorf("unmarshal failed, err:%v", err)
    }

    t.Logf("unmarshal success, conf:%#v", conf2)
}