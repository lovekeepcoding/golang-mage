package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func simpleGet() {
	if resp, err := http.Get("http://127.0.0.1:5656"); err != nil {
		panic(err)
	} else {
		defer resp.Body.Close() //注意：一定要调用resp.Body.Close()，否则会协程泄漏（同时引发内存泄漏）
		/**
		具体看一下http协议
		*/
		fmt.Printf("response proto: %s\n", resp.Proto)
		fmt.Printf("response status: %s\n", resp.Status)
		fmt.Println("response header")
		for key, values := range resp.Header {
			fmt.Printf("%s: %v\n", key, values)
		}
		fmt.Println("response body")
		io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
		os.Stdout.WriteString("\n")
	}
}

func simplePost() {
	reader := strings.NewReader("Hello Server")                                            //把string转成io.Reader
	if resp, err := http.Post("http://127.0.0.1:5656", "text/plain", reader); err != nil { //Content-Type为text/plain，表示一个朴素的字符串
		panic(err)
	} else {
		defer resp.Body.Close() //注意：一定要调用resp.Body.Close()，否则会协程泄漏（同时引发内存泄漏）
		/**
		具体看一下http协议
		*/
		fmt.Printf("response proto: %s\n", resp.Proto)
		fmt.Printf("response status: %s\n", resp.Status)
		fmt.Println("response header")
		for key, values := range resp.Header {
			fmt.Printf("%s: %v\n", key, values)
		}
		fmt.Println("response body")
		io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
		os.Stdout.WriteString("\n")
	}
}

func postForm() {
	//通过form表单提交一些参数键值对
	if resp, err := http.PostForm("http://127.0.0.1:5656", url.Values{"name": []string{"zcy"}, "age": []string{"18"}}); err != nil {
		panic(err)
	} else {
		defer resp.Body.Close() //注意：一定要调用resp.Body.Close()，否则会协程泄漏（同时引发内存泄漏）
		/**
		具体看一下http协议
		*/
		fmt.Printf("response proto: %s\n", resp.Proto)
		fmt.Printf("response status: %s\n", resp.Status)
		fmt.Println("response header")
		for key, values := range resp.Header {
			fmt.Printf("%s: %v\n", key, values)
		}
		fmt.Println("response body")
		io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
		os.Stdout.WriteString("\n")
	}
}

func head() {
	//HEAD类似于GET，但HEAD方法只能取到http response报文头部，取不到resp.Body
	if resp, err := http.Head("http://127.0.0.1:5656"); err != nil {
		panic(err)
	} else {
		defer resp.Body.Close() //注意：一定要调用resp.Body.Close()，否则会协程泄漏（同时引发内存泄漏）
		/**
		具体看一下http协议
		*/
		fmt.Printf("response proto: %s\n", resp.Proto)
		fmt.Printf("response status: %s\n", resp.Status)
		fmt.Println("response header")
		for key, values := range resp.Header {
			fmt.Printf("%s: %v\n", key, values)
		}
		fmt.Println("response body")
		io.Copy(os.Stdout, resp.Body) //resp.Body为空
		os.Stdout.WriteString("\n")
	}
}

//(*http.Client).Do允许我们构造复杂的http request
func complexRequest() {
	reader := strings.NewReader("Hello Server")
	//HEAD、GET、POST 默认都属于简单请求 Simple Request，通过http.NewRequest可以支持全部的request method
	if req, err := http.NewRequest("DELETE", "http://127.0.0.1:5656", reader); err != nil {
		panic(err)
	} else {
		//自定义请求头
		req.Header.Add("User-Agent", "花果山")
		req.Header.Add("King", "孙悟空")
		//自定义Cookie
		//HTTP请求中的Cookie头只会包含name和value信息（服务端只能取到name和value），domain、path、expires等cookie属性是由浏览器使用的，对服务器来说没有意义
		req.AddCookie(&http.Cookie{
			Name:   "auth",
			Value:  "pass",
			Path:   "/",
			Domain: "localhost",
		})
		//设置请求超时
		client := &http.Client{
			Timeout: 500 * time.Millisecond,
		}
		if resp, err := client.Do(req); err != nil {
			fmt.Println(err)
		} else {
			defer resp.Body.Close()
			/**
			具体看一下http协议
			*/
			fmt.Printf("response proto: %s\n", resp.Proto)
			fmt.Printf("response status: %s\n", resp.Status)
			fmt.Println("response header")
			for key, values := range resp.Header {
				fmt.Printf("%s: %v\n", key, values)
			}
			fmt.Println("response body")
			io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
			os.Stdout.WriteString("\n")
		}
	}
}

func restful() {
	if resp, err := http.Post("http://127.0.0.1:5656/user/xiaoming/vip/bj/haidian", "text/plain", nil); err != nil { //Content-Type为text/plain，表示一个朴素的字符串
		panic(err)
	} else {
		defer resp.Body.Close()
	}
}

func requestBook() {
	if resp, err := http.Post("http://book.dianshang:5656", "text/plain", nil); err != nil { //Content-Type为text/plain，表示一个朴素的字符串
		panic(err)
	} else {
		defer resp.Body.Close()
		fmt.Println("response body")
		io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
		os.Stdout.WriteString("\n")
	}
}

func requestFood() {
	if resp, err := http.Post("http://food.dianshang:5656", "text/plain", nil); err != nil { //Content-Type为text/plain，表示一个朴素的字符串
		panic(err)
	} else {
		defer resp.Body.Close()
		fmt.Println("response body")
		io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
		os.Stdout.WriteString("\n")
	}
}

func requestPanic() {
	if resp, err := http.Get("http://127.0.0.1:5656/panic"); err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("response status: %s\n", resp.Status)
		fmt.Println("response header")
		for key, values := range resp.Header {
			fmt.Printf("%s: %v\n", key, values)
		}
		fmt.Println("response body")
		io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
		os.Stdout.WriteString("\n")
	}
}

func authLogin() {
	if resp, err := http.Post("http://127.0.0.1:5656/login", "text/plain", nil); err != nil {
		panic(err)
	} else {
		fmt.Println("response body")
		io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
		os.Stdout.WriteString("\n")
		loginCookies := resp.Cookies() //读取服务端返回的Cookie
		resp.Body.Close()
		if req, err := http.NewRequest("POST", "http://127.0.0.1:5656/center", nil); err != nil {
			panic(err)
		} else {
			//下次请求再带上cookie
			for _, cookie := range loginCookies {
				fmt.Printf("receive cookie %s = %s\n", cookie.Name, cookie.Value)
				cookie.Value += "1" //修改cookie后认证不通过
				req.AddCookie(cookie)
			}
			client := &http.Client{}
			if resp, err := client.Do(req); err != nil {
				fmt.Println(err)
			} else {
				defer resp.Body.Close()
				fmt.Println("response body")
				io.Copy(os.Stdout, resp.Body) //两个io数据流的拷贝
				os.Stdout.WriteString("\n")
			}
		}
	}
}

func main() {
	// simpleGet()
	// simplePost()
	// postForm()
	// head()
	// complexRequest()

	// restful()
	// requestPanic()
	// requestBook()
	// requestFood()

	//测试限流中间件
	// const P = 130
	// wg := sync.WaitGroup{}
	// wg.Add(P)
	// for i := 0; i < P; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		if resp, err := http.Get("http://127.0.0.1:5656"); err == nil {
	// 			resp.Body.Close()
	// 		}
	// 	}()
	// }
	// wg.Wait()

	authLogin()
}

//go run http/client/main.go
