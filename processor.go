package main

import (
    "log"
)

func ProcessData(data map[string]interface{}) {
    if name, ok := data["name"]; ok {
        log.Printf("Name: %v", name)
    } else {
        log.Println("No 'name' field found in the JSON data")
    }
}

