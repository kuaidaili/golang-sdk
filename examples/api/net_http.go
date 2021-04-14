package main

//GO版本 GO1
import (
    "compress/gzip"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {

    // api链接
    api_url := "http://svip.kdlapi.com/api/getproxy/?orderid=96518362xxxxxx&num=100&protocol=1&method=2&an_an=1&an_ha=1&sep=1"

    // 请求api链接
    req, _ := http.NewRequest("GET", api_url, nil)
    req.Header.Add("Accept-Encoding", "gzip") //使用gzip压缩传输数据让访问更快
    client := &http.Client{}
    res, err := client.Do(req)

    // 处理返回结果
    if err != nil {
        // 请求发生异常
        fmt.Println(err.Error())
    } else {
        defer res.Body.Close() //保证最后关闭Body

        fmt.Println("status code:", res.StatusCode) // 获取状态码

        // 有gzip压缩时,需要解压缩读取返回内容
        if res.Header.Get("Content-Encoding") == "gzip" {
            reader, _ := gzip.NewReader(res.Body) // gzip解压缩
            defer reader.Close()
            io.Copy(os.Stdout, reader)
            os.Exit(0) // 正常退出
        }

        // 无gzip压缩, 读取返回内容
        body, _ := ioutil.ReadAll(res.Body)
        fmt.Println(string(body))
    }
}