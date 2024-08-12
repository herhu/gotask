package main

import (
    "log"
    "sync"
)

func main() {
    files := []string{"file1.json", "file2.json", "file3.json"} // Replace with actual file paths
    var wg sync.WaitGroup

    for _, file := range files {
        wg.Add(1)
        go func(filename string) {
            defer wg.Done()
            data, err := ReadJSONFile(filename)
            if err != nil {
                log.Printf("Error reading file %s: %v", filename, err)
                return
            }
            ProcessData(data)
        }(file)
    }

    wg.Wait()
    log.Println("All files processed.")
}
