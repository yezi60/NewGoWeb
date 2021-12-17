package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)

// 重启服务
func reLanuch() {
	// 名称，参数
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func deployPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1> Hello, this is my depoly page! </h1>")
	reLanuch()
}

func main() {
	http.HandleFunc("/", deployPage)
	http.ListenAndServe(":5000", nil)
}
