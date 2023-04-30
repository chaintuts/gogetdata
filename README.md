## General
____________

### Author
* Josh McIntyre

### Website
* jmcintyre.net

### Overview
* GoGetData is a simple static data web server demo

## Development
________________

### Git Workflow
* master for releases (merge development)
* development for bugfixes and new features

### Building
* make build
Build the application - wraps `go build`
* make clean
Clean the build directory

### Features
* Serve static files requested via HTTP

### Requirements
* Requires Go language build tools

### Platforms
* Windows
* MacOSX
* Linux

## Usage
____________

### Command Line Usage
* Run `./gogetdata`
* Use a web browser to request a file in the same directory as the server binary
* Default port is 8080 - request files such as localhost:8080/<filename>