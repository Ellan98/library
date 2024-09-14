/*
 * @Date: 2024-08-29 10:41:10
 * @LastEditTime: 2024-08-29 10:41:18
 * @FilePath: \library_room\test\sse.go
 * @description: 注释
 */
package main

import (
	"fmt"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// 设置适当的头信息，表明这是一个 SSE 流
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 确保响应体自动刷新
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// 模拟连续发送事件到客户端
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "data: Message %d\n\n", i+1)
		flusher.Flush() // 立即将数据推送给客户端
		time.Sleep(1 * time.Second)
	}

	// 最后关闭连接
	
	fmt.Fprintf(w, "event: close\ndata: Connection will close\n\n")
	flusher.Flush()
}

func main() {
	http.HandleFunc("/sse", sseHandler)
	fmt.Println("Listening on http://localhost:8080/sse")
	http.ListenAndServe(":8080", nil)
}
