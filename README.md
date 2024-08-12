```markdown
# JSON Processor

This Go application reads data from multiple JSON files concurrently and processes them by printing a specific part of the JSON (the `name` field).

## Project Structure

- **main.go**: The entry point of the application. Handles the concurrent processing of JSON files.
- **reader.go**: Contains the logic for reading and parsing JSON files.
- **processor.go**: Contains the logic for processing the JSON data.
- **models.go**: (Optional) Can be used to define data structures depending on your JSON schema.

## JSON Files

This project processes three JSON files:
- **file1.json**: Contains information about a cat.
- **file2.json**: Contains information about a dog.
- **file3.json**: Contains information about a bird.

### Example JSON Content

Here are the contents of the JSON files:

- **file1.json**
  ```json
  {
      "name": "Whiskers",
      "species": "Cat",
      "raza": "Siamese"
  }
  ```

- **file2.json**
  ```json
  {
      "name": "Buddy",
      "species": "Dog",
      "raza": "Golden Retriever"
  }
  ```

- **file3.json**
  ```json
  {
      "name": "Tweety",
      "species": "Bird",
      "raza": "Canary"
  }
  ```

## How to Run the Application

### Running with `go run`

You can run the application directly using `go run`:

```sh
go run main.go reader.go processor.go models.go
```

### Compiling the Application

You can compile the application into a standalone binary using `go build`:

```sh
go build -o json_processor main.go reader.go processor.go models.go
```

After compiling, you can run the executable:

```sh
./json_processor
```

### Cross-Compiling

To compile the application for a different operating system, use Go’s cross-compiling feature. For example, to compile for Windows:

```sh
GOOS=windows GOARCH=amd64 go build -o json_processor.exe main.go reader.go processor.go models.go
```

## Output

When you run the application, it will process the JSON files and print the `name` field for each one:

```sh
2024/08/12 10:00:00 Name: Whiskers
2024/08/12 10:00:00 Name: Buddy
2024/08/12 10:00:00 Name: Tweety
2024/08/12 10:00:00 All files processed.
```
