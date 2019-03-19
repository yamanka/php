package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")); err != nil {
		panic(err)
	}
	if err := gz.Flush(); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	fmt.Println(str)
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "text/html")
	// 往w里写入内容，就会在浏览器里输出
	//fmt.Fprintf(w, string(b.Bytes()[:]))
	//fmt.Println(string(b.Bytes()[:]))
	w.Write(b.Bytes())
}

func main() {
	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/", index)

	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
