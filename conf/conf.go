package conf 

import (
        "io/ioutil"
        "encoding/json"
        "log"
        "os"
)


type configuration struct {
    Server string `json:"server"`
    Channel string `json:"channel"`
    Nickname string `json:"nickname"`
}

const jsonFile = "conf/conf.json"


func load() *configuration {

    f, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err == nil {

        log.SetOutput(f)

        defer f.Close()

        file, err := ioutil.ReadFile(jsonFile)
        if err == nil {

            conf := new(configuration)
            if err := json.Unmarshal(file, conf); err == nil {
                return conf
            }
        }
    }

    return nil
}

var LogFile = "/tmp/toby.log"

var Config = load()
