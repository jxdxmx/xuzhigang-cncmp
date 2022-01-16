package main

import (
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"os"
)

// 模块3的作业（徐志刚）

// 构建本地镜像
// 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
// 将镜像推送至 docker 官方镜像仓库
// 通过 docker 命令本地启动 httpserver
// 通过 nsenter 进入容器查看 IP 配置

// 编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：
// 1、接收客户端 request，并将 request 中带的 header 写入 response header
// 2、读取当前系统的环境变量中的 VERSION 配置，并写入 response header
// 3、Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
// 4、当访问 localhost/healthz 时，应返回 200

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/healthz", healthZHandler)
	mux.HandleFunc("/homework", handler)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		glog.Error(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 1、接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		for i, str := range v {
			if i == 0 {
				w.Header().Set(k, str)
			} else {
				w.Header().Add(k, str)
			}
		}
	}

	// 2、读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	ver := os.Getenv("VERSION")
	w.Header().Set("VERSION", ver)

	// 3、Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	ip := r.RemoteAddr
	glog.Info("client ip ", ip, ",status code ", http.StatusOK)

	res := fmt.Sprintf("hello world handler, VERSION %s , client ip %s \n", ver, ip)
	_, _ = w.Write([]byte(res))
	w.WriteHeader(http.StatusOK)
}

func healthZHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world module3"))
	w.WriteHeader(http.StatusOK) // 4、当访问 localhost/healthz 时，应返回 200
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world module3 !!! index ."))
	w.WriteHeader(http.StatusOK) // 4、当访问 localhost/healthz 时，应返回 200
}
