# schdl - An Interactive Terminal Scheduler Made in Golang

## Table of Contents

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Project Structure](#project-structure)
4. [Contributing](#contributing)
5. [License](#license)

## Introduction

Schdl is a terminal-based task scheduler written in Golang. It allows users to manage their tasks efficiently by storing tasks along with their due date and time. Users can create, delete, update, and retrieve tasks using an interactive interface.

## Installation

To install schdl, make sure you have Go installed on your machine. Then, follow these steps:

```sh
go get github.com/codadept/schdl
cd $GOPATH/src/github.com/codadept/schdl
go build -o schdl
```

## Project Structure

The project is structured as follows:

```
schdl/
├── cmd/
│   └── schdl/
│       └── main.go
├── internal/
│   ├── storage/
│   │   └── storage.go
│   ├── task/
│   │   └── task.go
│   ├── ui/
│   │   └── ui.go
│   └── util/
├── LICENSE
├── README.md
├── go.mod
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

## License

This project is licensed under the GNU GENERAL PUBLIC LICENSE Version 3 - see the LICENSE file for details.
