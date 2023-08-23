package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入登录地址（默认为http://192.168.31.1/cgi-bin/luci/web）: ")
	host, _ := reader.ReadString('\n')
	host = strings.TrimSpace(host)
	if host == "" {
		host = "http://192.168.31.1/cgi-bin/luci/web"
	}

	resp, err := http.Get(host)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "    key: ") {
			fmt.Println(strings.TrimPrefix(line, "key: "))
		} else if strings.HasPrefix(line, "    iv: ") {
			fmt.Println(strings.TrimPrefix(line, "iv: "))
		} else if strings.HasPrefix(line, "        hardwareVersion: ") {
			fmt.Println(strings.TrimPrefix(line, "hardwareVersion: "))
		}
	}
	fmt.Println("注意查看是否为你的设备")
	if err := scanner.Err(); err != nil {
		fmt.Println("读取失败:", err)
	}
}
