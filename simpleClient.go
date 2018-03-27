package main

import (
    "fmt"
    "net"
    "time"
  
)

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}

func main() {
    ServerAddr,err := net.ResolveUDPAddr("udp","***.***.***.***:8000")
    CheckError(err)

    LocalAddr, err := net.ResolveUDPAddr("udp", "***.***.***.***:2000")
    CheckError(err)

    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)

    defer Conn.Close()
    fmt.Println("Hello Will be Sent")
    msg := "Hello"
    buf := make([]byte, 1024)
    copy(buf[:], msg)
    fmt.Println("Actually sending it")
    _,err = Conn.Write(buf)
    fmt.Println("Sent")
    if err != nil {
      fmt.Println(msg, err)
    }
    time.Sleep(time.Microsecond * 33333)
    _,_,err = Conn.ReadFromUDP(buf)
    fmt.Println(string(buf[0:]))
}
