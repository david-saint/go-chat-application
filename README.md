# Go Chat Application DƧ

## Introduction

This is just a simple chat application that makes use of gorilla/websocket.
I shows the best way to handle websockets with goroutines and channels.

## Getting Started

### Install Packages

The only package used in this project, that needs installing is Gorilla/websocket.
You can install it like this

```bash
go get github.com/gorilla/websocket
```


### Demonstration

To demo this application, all you have to do is clone and build.
You can specify the port you want to run the application on with the -addr flag (it defaults to 8080).

* Clone this repo: `git clone github.com/david-saint/go-chat-application`
* Cd into the project: `cd go-chat-application`
* Build the application: `go build -o release/chat`
* Run the application: `./release/chat -addr=":8080"`
* Open Browser and view the application: [`Localhost Port 8080`](http://localhost:8080/)

|Endpoint|Description|
|:------------- | :----------: |
|`/`|displays the landing page and a link to the actual chat application|
|`/chat`|shows the chat application, a textarea, a button and list of messages|
|`/room`|This is the websocket endpoint, which can be accessed at `ws://localhost:8080/room`|
