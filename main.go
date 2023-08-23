package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
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
	re := regexp.MustCompile(`^\s*(key|iv|hardwareVersion):\s*'([^']*)'`)
	for scanner.Scan() {
		line := scanner.Text()
		match := re.FindStringSubmatch(line)
		if match != nil {
			fmt.Printf("%s: %s\n", match[1], match[2])
		}
	}
	fmt.Println("注意查看是否为你的设备")
	if err := scanner.Err(); err != nil {
		fmt.Println("读取失败:", err)
	}
}
