# Log Generator and Memcache Loader

This project is a Go application that generates log data, writes it to a file, and then loads it into Memcache. The application is structured into two main components: a `log_generator` package and a `main` application that utilizes this package.

## Components

### log_generator Package

- **Log Struct**: Represents a log entry with fields for `Time`, `Level`, and `Message`.
- **Functions**:
  - `JsonLog(file string)`: Generates JSON log entries and appends them to the specified file.
  - `Is_file_empty(file string) bool`: Checks if the specified file is empty.
  - `Get_file_data(file string) []Log`: Reads and decodes log entries from the specified file into a slice of `Log` structs.

### Main Application

- Initializes a Memcache client with a specified timeout.
- Checks if the log file (`test.log`) is empty and generates log entries if necessary.
- Reads log entries from the file and stores them in Memcache using the log level as the key.
- Retrieves and prints the stored log entries from Memcache.

## Usage

1. Ensure Memcache is running locally on port 11211.
2. Run the Go application:
   ```bash
   go run main.go
   ```
3. The application will generate log entries if `test.log` is empty, store them in Memcache, and print the stored entries.

## Dependencies

- Go 1.16 or higher
- [gomemcache](https://github.com/bradfitz/gomemcache) - Memcache client for Go

## License

This project is licensed under the MIT License.
