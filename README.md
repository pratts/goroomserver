# Go Room Server
Go Room Server is a [Golang](http://golang.org/) based server with websocket functionality using [Gorilla-Websocket](https://github.com/gorilla/websocket).

### Description
The server provides the functionality to create multiple applications acting as namespaces. Each application consists of list of rooms and connected users.
Users connect using either custom websocket implementation or the goroomserver libraries [Go-Room-Client-JS](https://github.com/pratts/goroomclient-js) and [Go-Room-Client-Go](https://github.com/pratts/goroomclient-go) coming soon.

### Installation
    go get github.com/pratts/goroomserver

### Concept
The server consists of multiple components:

    1. A main server that will contain a map of applications associated with the application name.
    2. Each application will contain room and user services consisting of rooms created and 
       logged in users.
    3. Each room will contain a map of users against their names who have joined the room.
    4. Each user will have a connection reference to use for communication and the rooms user 
       has joined.
    5. Each application and room will have an extension with init method that'll be called after 
       successful initialization of application and room respectively.
    6. Server works with various events that are triggered either from server or client.
    7. Each application and room will have a map of event handlers. On each
       event getting triggered, the respective event handler's handleEvent method will be called 
       with appropriate parameters.
    8. Following is the list of event that are currently handled on server end:
        a.) Connection
        b.) Disconnection
        c.) Login
        d.) Logout
        e.) Join Room
        f.) Leave Room
        g.) Message

### More feature coming
    1. Trigerring events from server side
    2. Thread pool to handle the events triggered. Currently each websocket message is handled in a 
       separate goroutine
    3. Exception handling
    4. Client side libraries
    5. Server side application example

### Note:
    The logic to start websocket server is blocking. Either start the websocket at the end, after you've 
    initiated everything or start it in a separate goroutine.

### Credits
The server users the [Gorilla-Websocket](https://github.com/gorilla/websocket) library for websocket implementation.
