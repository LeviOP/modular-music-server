package main

import (
	"fmt"
	"net"

	"example.com/modular-music-server/handlers"
	"example.com/modular-music-server/util"
    "example.com/modular-music-server/config"
)

func main() {
    config.Config()
    return
    // Listen for incoming connections
    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server is listening on port 8080")

    for {
        // Accept incoming connections
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        // Handle client connection in a goroutine
        go handleClient(conn)
    }
}

// type Client struct {
//     Connection net.Conn
// }
//
// var client Client;

func handleClient(conn net.Conn) {
    defer conn.Close()

    // client.Connection = conn;

    for {
        messageType, data, err := util.ReadMessage(conn)
        if err != nil {
            if err.Error() == "EOF" {
                fmt.Println("EOF from client. Closing connection.")
                break
            }
            fmt.Println("Error: ", err);
            continue
        }

        switch messageType {
        case util.MESSAGE_HANDSHAKE_REQUEST:
            handlers.HandshakeRequest(conn, data)
        case util.MESSAGE_REQUESTLIST:
            handlers.RequestList(conn, data)
        case util.MESSAGE_REQUESTPROVIDER:
            handlers.RequestProvider(conn, data)
        }
    }

}
