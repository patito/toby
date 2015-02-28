package settings 

import (
        "io/ioutil"
        "encoding/json"
        "log"
        "os"
)


type Settings struct {
    Server string `json:"server"`
    Channel string `json:"channel"`
    Nickname string `json:"nickname"`
}

const jsonFile = "settings.json"

var settings *Settings = nil

var logFile = "/tmp/toby.log"

func get() *Settings {

    if settings == nil {

        f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
        if err != nil {
            log.Println("Settings::get() - Error creating logfile.")
            return nil
        }

        log.SetOutput(f)

        defer f.Close()

        file, err := ioutil.ReadFile(jsonFile)
        if err != nil {
            log.Println("Settings::get() - Error loading json file.")

            return nil
        }

        settings = new(Settings)
        if err := json.Unmarshal(file, settings); err != nil {
            log.Println("Settings::get() - Error unmarshal json.")

            return nil
        }
    }

    return settings
}
