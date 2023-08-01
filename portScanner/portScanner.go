package main

import (
    "net"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) == 1 {
        fmt.Println("Please specify a host\nportScanner <host>")
        os.Exit(1)
    }
    host := os.Args[1]
    startPort := 1
    endPort := 65535
    for ;startPort <= endPort; startPort++ {
        connection, err := net.Dial("tcp", fmt.Sprintf("%v:%v", host,startPort))
        if err != nil {
            continue
        }
        fmt.Printf("%v is open\n", startPort)
        connection.Close()
    }

}
