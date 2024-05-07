# Go Mongo Excel Generator

This project is a Go application that demonstrates how to generate Excel files from MongoDB data using the Gin web framework and the Excelize library. It provides different approaches for generating Excel files and serves them to clients via HTTP endpoints.

## Installation

1. Install dependencies:

```bash
go mod tidy
```

2. Build and run the application:

```bash
go run main.go
```

## Usage

Once the application is running, you can access the following endpoints:

- `/ping/student-info`: Ping endpoint for student information.
- `/ping/teacher-info`: Ping endpoint for teacher information.
- `/get/student-info`: Retrieve all student information from MongoDB.
- `/generate/student-info/excel/save-new`: Generate a new Excel file with student information and save it locally.
- `/generate/student-info/excel/memory-new`: Generate a new Excel file with student information and serve it from memory.
- `/generate/student-info/excel/stream-new`: Generate a new Excel file with student information and stream it to the client.
- `/generate/student-info/excel/stream-random`: Generate a large Excel file with random data and stream it to the client.
- `/generate/teacher-info/excel`: Generate an Excel file with teacher information.

## Approach Comparison

This project demonstrates three different approaches for generating Excel files:

1. **Save New Excel File Approach**:

   - Generates Excel files and saves them to the local filesystem.
   - Suitable for storing generated files for later use or archival purposes.
   - Consumes disk space for storage.

2. **Memory New Excel File Approach**:

   - Generates Excel files and serves them directly from memory.
   - Ideal for generating files on-the-fly without saving them to disk.
   - Avoids disk storage but does not persist files.

3. **Stream New Excel File Approach**:
   - Generates Excel files and streams them directly to the client.
   - Efficient for handling large datasets without consuming excessive memory.
   - Suitable for real-time generation and streaming.
