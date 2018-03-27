package main

import (
    "fmt"
    "net"
    "os"
    "time"
)
func main() {

    ServerAddr,err := net.ResolveUDPAddr("udp",":8000")
    CheckError(err)

    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)

    defer ServerConn.Close()
    buf := make([]byte, 1024)

    for {
        _,addr,err := ServerConn.ReadFromUDP(buf)
        fmt.Println("Received ",string(buf[0:5]), " from ",addr)
        if string(buf[0:5]) == "Hello" {
          message := time.Now().String()
          buffer := make([]byte, 1024)
          copy(buffer[:], message)
          time.Sleep(time.Microsecond * 333333)
          _, err = ServerConn.WriteToUDP(buffer, addr)
				      CheckError(err)
              time.Sleep(time.Microsecond * 33333)
        }
        if err != nil {
            fmt.Println("Error: ",err)
        }
    }
}

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}
