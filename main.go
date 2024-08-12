package main

import (
    "log"
    "sync"
    "os"
)

func main() {
    files := []string{"file1.json", "file2.json", "file3.json"}
    var wg sync.WaitGroup
    errorCount := 0

    for _, file := range files {
        wg.Add(1)
        go func(filename string) {
            defer wg.Done()
            data, err := ReadJSONFile(filename)
            if err != nil {
                log.Printf("Failed to process file %s: %v", filename, err)
                errorCount++
                return
            }
            ProcessData(data)
        }(file)
    }

    wg.Wait()

    if errorCount > 0 {
        log.Printf("Processing completed with %d errors.", errorCount)
        os.Exit(1)
    } else {
        log.Println("All files processed successfully.")
    }
}
