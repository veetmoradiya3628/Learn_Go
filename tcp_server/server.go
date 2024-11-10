package main

import (
    "fmt"
    "net"
)

func main() {
    listener, err := net.Listen("tcp", "127.0.0.1:9999")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()
    fmt.Println("Server is listening on 127.0.0.1:9999")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        fmt.Println("Client connected")
        conn.Close()
    }
}
