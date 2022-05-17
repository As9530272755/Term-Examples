package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	sshHost := "10.0.0.200"
	sshPort := 22
	sshUser := "root"
	sshPassword := "666666"
	sshType := "password"

	//  创建ssh登录配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second,
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 登录类型 password
	if sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		return
	}

	// 定义链接地址
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)

	// 链接 ssh 客户端
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("创建连接失败：", err)
	}
	defer sshClient.Close()

	// 客户端 session
	session, _ := sshClient.NewSession()

	// 调用命令
	comdb, err := session.CombinedOutput("hostname")
	if err != nil {
		log.Fatal("远程执行cmd失败", err)
	}

	// 输出命令
	fmt.Print(string(comdb))
}
