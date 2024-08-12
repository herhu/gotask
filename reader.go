package main

import (
    "encoding/json"
    "io/ioutil"
    "os"
)

func ReadJSONFile(filename string) (map[string]interface{}, error) {
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

