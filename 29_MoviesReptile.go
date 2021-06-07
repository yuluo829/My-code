package main
 
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
    "time"
    "os"
 
    "github.com/PuerkitoBio/goquery"
)
 
/*
通过多个goroutine并发执行爬取操作，channel存放要爬取url内容和爬取结果
这样只需要设计爬取函数和提取函数
*/
 
func get_web_content(url string, chan_web chan string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("http get error", err)
        return
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("read error", err)
        return
    }
    chan_web <- string(body)
}
 
func extract_valid_content(body string, chan_r chan []byte) {
    dom, err := goquery.NewDocumentFromReader(strings.NewReader(body))
    if err != nil {
        fmt.Println(err)
    }
    dom.Find("ol.grid_view div.item").Each(func(i int, selection *goquery.Selection) {
        // extract result
        result := make(map[string]string)
        name := selection.Find("div.info span.title").First().Text()
        doctor_str := selection.Find("div.info div.bd p").First().Text()
        r := regexp.MustCompile(`导演:(?s:(.*?))(主演|主|&|\.\.\.)`)
        doctor := r.FindAllStringSubmatch(doctor_str, -1)[0][1]
        rating_num := selection.Find("div.star span.rating_num").First().Text()
        evaluation_str := selection.Find("div.star span").Last().Text()
        r = regexp.MustCompile(`(?s:(.*?))人评价`)
        evaluation := r.FindAllStringSubmatch(evaluation_str, -1)[0][1]
        ranking := selection.Find("div.pic em").First().Text()
        result["name"] = name
        result["doctor"] = doctor
        result["rating_num"] = rating_num
        result["evaluation"] = evaluation
        result["ranking"] = ranking
        json_str, err := json.Marshal(result)
        if err != nil {
            fmt.Println(err)
            return
        }
        chan_r <- json_str
    })
}
 
func main() {
    var (
        OutputFile = "./film_crawl.txt"
    )
    base_url := "https://movie.douban.com/top250?start=%d&filter="
    chan_web_content := make(chan string)
    defer close(chan_web_content)
    chan_r := make(chan []byte)
    defer close(chan_r)
    for i := 0; i < 10; i++ {
        url := fmt.Sprintf(base_url, i*25)
        go get_web_content(url, chan_web_content)
    }
 
    go func() {
        for {
            web_content, ok := <- chan_web_content
            if !ok {
                break
            }
            go extract_valid_content(web_content, chan_r)
        }
    }()
 
    flag := false
    to := time.NewTimer(time.Second * 5)
    file, err := os.OpenFile(OutputFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println("Failed to open the file", err.Error())
        return
    }
    defer file.Close()
    for {
        if flag {
            break
        }
        to.Reset(time.Second * 5)
        select {
        case res := <- chan_r:
            fmt.Printf("%s\n", res)
            file.Write(res)
            file.WriteString("\n")
        case <- to.C:
            flag = true
            break
        }
    }
    fmt.Println("end")
