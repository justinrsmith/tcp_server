package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "net"
)

var (
    ip string
    port string
)

func init() {
    flag.StringVar(&ip, "ip", "127.0.0.1", "IP Address to assign the server")
    flag.StringVar(&port, "port", "3333", "Port to assign to the server")
    flag.Parse()
}

func main() {
    fmt.Println("Launching server...")
    // Listen on TCP port 2000 on all interfaces.
    l, err := net.Listen("tcp", ip+":"+port)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Listening on %s:%s...\n", ip, port)
    defer l.Close()
    for {
        // Wait for a connection
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }
        // Handle the connection in a new goroutine.
        // The loop then returns to accepting, so that
        // multiple connections may be served concurrently.
        go handleConnection(conn)
    }
}

func handleConnection(c net.Conn){
    // Read connection into default sized buffer
    r := bufio.NewReader(c)
    input, err := r.ReadString('\n')
    if err != nil {
        c.Close()
        return
    }
    message := "The received message was: "
    message += input
    c.Write([]byte(message))
    c.Close()
}

