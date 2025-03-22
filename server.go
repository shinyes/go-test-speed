package main

import (
    "fmt"
    "net"
    "os"
)

const defaultPort = "54312"

func main() {
    port := defaultPort
    if len(os.Args) > 1 {
        port = os.Args[1]
    }

    addr := ":" + port
    listener, err := net.Listen("tcp", addr)
    if err != nil {
        fmt.Println("服务端监听失败:", err)
        return
    }
    fmt.Println("服务端监听端口:", port)

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("接受连接失败:", err)
            continue
        }
        fmt.Println("客户端已连接:", conn.RemoteAddr())
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer func() {
        fmt.Println("客户端已断开:", conn.RemoteAddr())
        conn.Close()
    }()

    buffer := make([]byte, 32*1024)
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            return
        }
        // 回传数据给客户端，实现下载测速
        _, err = conn.Write(buffer[:n])
        if err != nil {
            return
        }
    }
}
