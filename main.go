package main

/*
#cgo CFLAGS: -Wno-nullability-completeness
#include <stdint.h>
*/
import (
	"C"
	"fmt"
	"net"
	"sync"
	"time"
)

// fastestConnection 尝试连接多个地址，返回最快连接的地址
func fastestConnection(addresses []string, port string, timeout time.Duration) (string, error) {
	var wg sync.WaitGroup
	result := make(chan string, 1) // 用于接收最快连接的地址
	var once sync.Once             // 确保只写入一次
	defer close(result)

	for _, address := range addresses {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			conn, err := net.DialTimeout("tcp", net.JoinHostPort(addr, port), timeout)
			if err == nil {
				once.Do(func() {
					result <- addr // 发送第一个成功连接的地址
				})
				conn.Close()
			}
		}(address)
	}

	// 等待所有连接完成
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done) // 所有goroutine完成后关闭done通道
	}()

	select {
	case fastest := <-result:
		return fastest, nil
	case <-done:
		return "", fmt.Errorf("no address could be connected")
	case <-time.After(timeout):
		return "", fmt.Errorf("operation timed out")
	}
}

//export DnsCheck
func DnsCheck() *C.char {
	addresses := []string{"223.5.5.5", "114.114.114.114", "1.1.1.1"}
	port := "53" // 通常DNS服务使用端口53
	timeout := 3 * time.Second

	fastest, err := fastestConnection(addresses, port, timeout)
	if err != nil {
		fmt.Println("Error:", err)
		return C.CString("")
	}
	return C.CString(fastest)
}

func main() {
	result := DnsCheck()
	println(result)
}
