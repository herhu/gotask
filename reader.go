package main

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "log"
    "os"
    "time"
)

const maxRetries = 3

func ReadJSONFile(filename string) (map[string]interface{}, error) {
    var data map[string]interface{}
    var err error

    for attempt := 1; attempt <= maxRetries; attempt++ {
        data, err = tryReadFile(filename)
        if err == nil {
            return data, nil
        }
        log.Printf("Attempt %d: Error reading file %s: %v", attempt, filename, err)
        time.Sleep(2 * time.Second) // Wait before retrying
    }

    return nil, errors.New("max retries reached: " + err.Error())
}

func tryReadFile(filename string) (map[string]interface{}, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    bytes, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    var data map[string]interface{}
    if err := json.Unmarshal(bytes, &data); err != nil {
        return nil, err
    }

    return data, nil
}
