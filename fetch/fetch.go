package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Fetch1(url string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	//从命令行参数中获取url
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
}

func Fetch2(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	//从url获取resp，如果出现错误将错误发送出去
	//同时会阻塞该协程直到主协程从ch中把错误信息读出
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	fmt.Println(resp.Status)
	nbytes, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}

	ch <- fmt.Sprintf("done and nbytes: %7d", nbytes)

}
