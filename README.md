## General
____________

### Author
* Josh McIntyre

### Website
* jmcintyre.net

### Overview
* GoGetData is a simple certificate generation and static data web server demo

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
* Serve static files requested via HTTP(S)
* Generates a self-signed certificate with private key
* Writes certificate and private key to disk, PEM encoded
* Uses existing certificate and key if they exist

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
* Default port is 443 - request files such as https://localhost/<filename>
* Generated certificates are self-signed, so you will need to configure your browswer/application to accept them
* Certificate and key files are placed up one directory from the server binary, and are PEM encoded