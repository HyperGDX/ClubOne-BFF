package middleware

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

var Proxy *httputil.ReverseProxy //NewReverseProxy(targetURL)
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 创建自定义的响应写入器
		writer := NewCustomResponseWriter(c.Writer)

		// 调用反向代理器的 ServeHTTP 方法来处理请求
		Proxy.ServeHTTP(writer, c.Request)

		// 从自定义响应写入器中获取转发服务器的响应
		//response := writer.GetResponse()

		// // 设置 Gin 服务器的响应
		// c.Status(response.StatusCode)
		// c.Header("Content-Type", response.Header.Get("Content-Type"))
		// c.Writer.Write(response.Body)
		//请求处理
		c.Next()
	}
}

// 自定义反向代理器
func NewReverseProxy(targetURL string) *httputil.ReverseProxy {
	target, _ := url.Parse(targetURL)

	// 创建自定义的 ReverseProxy
	proxy := httputil.NewSingleHostReverseProxy(target)

	// 修改请求的头部
	proxy.ModifyResponse = func(resp *http.Response) error {
		// 在这里可以对响应进行处理
		fmt.Println("Received response:", resp.Status)

		// 可以在这里修改转发服务器的响应

		return nil
	}

	return proxy
}

// 自定义响应写入器
type CustomResponseWriter struct {
	http.ResponseWriter
	response *http.Response
}

func NewCustomResponseWriter(w http.ResponseWriter) *CustomResponseWriter {
	return &CustomResponseWriter{
		ResponseWriter: w,
	}
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	// 创建转发服务器的响应对象
	w.response = &http.Response{
		StatusCode: statusCode,
		Header:     w.Header(),
		Body:       nil, // 在实际应用中，你可能需要从原始的 ResponseWriter 中获取转发服务器的响应体
	}

	// 调用原始的 ResponseWriter 的 WriteHeader 方法
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *CustomResponseWriter) Write(data []byte) (int, error) {
	// 在实际应用中，你可能需要将数据写入转发服务器的响应体

	// 调用原始的 ResponseWriter 的 Write 方法
	return w.ResponseWriter.Write(data)
}

func (w *CustomResponseWriter) GetResponse() *http.Response {
	return w.response
}
