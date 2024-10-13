package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// base64编码的favicon数据
const faviconBase64 = "AAABAAEAEBAAAAAAAABoAwAAFgAAACgAAAAQAAAAIAAAAAEAGAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAAD////////////////////////////////////////////////////////////////Z/Pzb/f3b/f0ategPpuEmwe2rzdLb/f3b/f2qzdImwewQp+Mateja/f3a/f3b/f3a/f3Z/f3Z/f0ivusVr+YPpeINoeAivushveoNoeAQpuMVruUivuvb/f3a/f3a/f3b/f3a/f3A3fItx+4pw+0kveoZr+QUpuETpeAZr+Qkvusow+wsxe0AeM3b/f3b/f3a/f3b/f3A3fI1y+80w+s1vOgxvOcwwOowv+ozveg1vegzxOs2zfAAeM3a/f3a/f3a/f3a/f0AeM01wuo1welY2/N77Ppd3/Zc3/Z77PpX2fMzvug2xesAeM3a/f3b/f3b/f3a/f0Ad8kgtuc/0fBo3u8lJSVv5PZv5PcmJyho3e9A0fEgs+YBd8ja/f3Z/Pza/f2p3OYasOUqxexL1/Ny5PQxMTF/6vh+6fgxMTFy4vNL1/MqxewZruSo2+ba/f3a/f0owuwpwutM2PJl4fZ86fk/Pz+N7fmM7Pg/Pz986fln4vZM2fMpwOsowuza/f00y+9C0/FV3PVu4PWT7fmK7ftQUFCb8vua8fpOTk6K7PmW7vpr4fVU2/VB0vE0zO/b/f0nwetT2vJt5PaJ6fik8Pqh8Pqq8/up8vqg7/mh7vmF6vht5PdV3fUnweva/f3a/Pzb/f3A3fL///////+B6fif8Pq39fu29fuf8ft/6ff///////8AeM3b/f3b/v7b/f3a/f3v/v72/v72/v72/v566PeS7vmS7/l55/f2/v72/v72/v7b/f3a/f3b/f3b/v7b/f3a/f3b/v7a/f3a/f1e3fSC6/iD7Pld3fTa/f3a/f3a/f3b/f3b/v7b/f3b/f3b/f3Z/f3a/f3a/f3b/f3b/f1q4/Zq4/ba/f3a/f3a/f3a/f3b/f3b/f3a/f38/v78/v78/v78/v78/v78/v78/v78/v78/v78/v78/v78/v78/v78/v78/v78/v7//wAA48cAAOAHAADgBwAA4AcAAOAHAADgBwAAwAMAAIABAAAAAAAAwAMAAPgfAAD8PwAA/D8AAP5/AAD//wAA"

// 解码base64数据为字节切片
var faviconData, _ = base64.StdEncoding.DecodeString(faviconBase64)

func main() {
	// 创建一个默认的 Gin 引擎
	r := gin.Default()

	// 设置一个 GET 路由处理函数
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "你好，世界！")
	})

	// 添加 /favicon.ico 路由
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Data(http.StatusOK, "image/x-icon", faviconData)
	})

	// 运行服务器在 8080 端口
	fmt.Println("服务器正在运行，地址为 http://0.0.0.0:8080")
	r.Run("0.0.0.0:8080")
}
