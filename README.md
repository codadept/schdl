# schdl

An Interactive Terminal Task Scheduler Made in Golang

![Build Status](https://img.shields.io/github/actions/workflow/status/codadept/schdl/build.yml)
![License](https://img.shields.io/github/license/codadept/schdl)
![Go Version](https://img.shields.io/github/go-mod/go-version/codadept/schdl)

## Table of Contents

1. [Overview](#overview)
2. [Pre-Requisite](#prerequisites)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Contributing](#contributing)
6. [License](#license)

## Overview

Schdl is a terminal-based task scheduler written in Golang. It allows users to manage their tasks efficiently by storing tasks along with their due date and time. Users can create, delete, update, and retrieve tasks using an interactive interface.

![Demo](https://github.com/codadept/schdl/raw/main/.github/demo.webm)

## Prerequisites

- Go 1.23 or higher
- SQLite3

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/codadept/schdl.git
   ```

2. Navigate to the project directory:

   ```sh
   cd schdl
   ```

3. Install dependencies:

   ```sh
   go mod tidy
   ```

## Usage

1. Build the project:

   ```sh
   go build -o schdl ./cmd/schdl
   ```

2. Run the application:

   ```sh
   ./schdl
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
