package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "time"
)

func main() {

    message("Launching server...")
    ln, err := net.Listen("tcp", ":4956")
    exitOnError(err)
    for {
        message("Wait for connection...")
        conn, err := ln.Accept()
        exitOnError(err)
        go handleConnection(conn)
    }

    message("Done")

}

func handleConnection(conn net.Conn) {

    message("")
    message("Handle connection...")
    str, err := bufio.NewReader(conn).ReadString('\n')
    exitOnError(err)
    if len(str) > 0 {
        t := time.Now()

        conn.Write([]byte(t.Format(time.RFC3339Nano) + "\n"))

        received, err := time.Parse(time.RFC3339Nano, str[:len(str)-1])
        exitOnError(err)

        message("Received:", received)
        message("Sent:    ", t)
    }

    conn.Close()
}

func message(a ...interface{}) (n int, err error) {
    return fmt.Print("[S] ", fmt.Sprintln(a...))
}

func exitOnError(err error) {
    if err != nil {
        message("ERR:", err)
        os.Exit(1)
    }
}
