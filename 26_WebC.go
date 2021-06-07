package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	//	"strings"
)

//爬取博主名author，博客名title，博客地址blogSite，发布日期publishTime，浏览量browingTimes

//封装一个函数，爬取一页内容
func SpiderOnePage(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		fmt.Println("http.Get err1 = ", err1)
		return
	}
	defer resp.Body.Close()

	//读取网页的内容
	buf := make([]byte, 4*1024)

	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			//如果读取接收，直接break
			if err2 == io.EOF {
				break
				//如果是其他错误，就打印出来
			} else {
				fmt.Println("resp.Body.Read err2 = ", err2)
				break
			}
		}
		result += string(buf[:n]) //读取多少，写多少
	}

	return
}

//获取博客作者Id
func getAuthor(srcResult string) (author string) {
	//原文格式 var username = "ijxr1a64ji53l"
	re := regexp.MustCompile(`var username = "(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
	}
	tmpt := re.FindAllStringSubmatch(srcResult, 1)

	for _, data := range tmpt {
		author = data[1]
		break
	}
	return
}

//获取博客title
func getBlogTitle(blogResult string) (title string) {
	// 原文格式 <title>【资讯】福布斯：旅行积分计划是区块链主要目标，对旅行者来说是好消息 - CSDN博客</title>
	re := regexp.MustCompile(`<title>(?s:(.*?)) -`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
	}
	tempTitle := re.FindAllStringSubmatch(blogResult, -1)

	//
	for _, data := range tempTitle {
		title = data[1]
		break
	}
	return

}

//获取博客发布时间
func getBlogPTime(blogResult string) (publishTime string) {
	//原文格式 <span class="link_postdate">2018年01月30日 16:14:25</span>
	re := regexp.MustCompile(`<span class="link_postdate">(?s:(.*?))</span`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
	}
	temp := re.FindAllStringSubmatch(blogResult, 1)

	for _, data := range temp {
		publishTime = data[1]
		//如果爬取到内容，则直接返回
		if publishTime != "" {
			return
		}
		break
	}

	//第二种原文格式 <span class="time">2018年01月12日 00:00:00</span>
	re1 := regexp.MustCompile(`<span class="time">(?s:(.*?))</span`)
	if re1 == nil {
		fmt.Println("regexp.MustCompile err")
	}
	temp1 := re1.FindAllStringSubmatch(blogResult, 1)
	var publishTime1 string

	for _, data1 := range temp1 {
		publishTime1 = data1[1]
		break
	}

	publishTime = publishTime + publishTime1

	return

}

//获取博客阅读次数
func getReadTimes(blogResult string) (readTimes string) {
	//原文格式 <span class="link_view" title="阅读次数">5485人阅读</span>
	re := regexp.MustCompile(`<span class="link_view" title="阅读次数">(?s:(.*?))人阅读</span`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
	}
	temp := re.FindAllStringSubmatch(blogResult, 1)

	for _, data := range temp {
		readTimes = data[1]
		if readTimes != "" {
			return
		}

		break
	}
	//原文原代码第二种格式 <i class="icon iconfont icon-read"></i><span class="txt">428</span></button></li>
	re1 := regexp.MustCompile(`<i class="icon iconfont icon-read"></i><span class="txt">(?s:(.*?))</span`)
	if re1 == nil {
		fmt.Println("regexp.MustCompile err")
	}
	temp1 := re1.FindAllStringSubmatch(blogResult, 1)
	var readTimes1 string
	for _, data1 := range temp1 {
		readTimes1 = data1[1]
		break
	}

	//如第一种匹配不到，则为空，两种格式的结果相加，得到最终结果
	readTimes = readTimes + readTimes1

	return
}

//获取具体博客页码结果
func getBlogResult(srcResult string) (result string) {
	//获取博客网址
	// <a href="https://blog.csdn.net/IJXR1A64JI53L/article/details/78047489"  target=
	re := regexp.MustCompile(`<a href="(?s:(.*?))"  target=`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
	}
	tmpt := re.FindAllStringSubmatch(srcResult, -1)

	for i, data := range tmpt {

		//按照获取的网址地址规律，取出i=1,4,7,...等的数据
		if i%3 == 1 {
			blogUrl := string(data[1])
			/			fmt.Printf("blogUrl #%d#= #%v#\n", i, blogUrl)
			//获取博客详情
			blogdetail, err := getBlogDetail(blogUrl)
			if err != nil {
				fmt.Println("getBlogDetail err=", err)
				return
			}
			if i > 28 {
				break
			}

			result += blogdetail + "\n"

			//分隔每个博客
			result += "----------------------------------------------------------\n\n"
		}
	}
	return
}

func getBlogDetail(blogUrl string) (blogDetail string, err error) {

	//获取具体博客页码结果
	blogPage, err2 := SpiderOnePage(blogUrl) //得到页面全部内容
	if err2 != nil {
		err = err2
		fmt.Println("blogResult SpiderOnePage(blogUrl) err2 = ", err2)
		return
	}

	//根据得到的页面全部内容，提取博客详情
	//获取博客title
	title := getBlogTitle(blogPage)
	//	fmt.Println("getBlogTitle title = ", title)
	blogDetail = "title=" + title + "\n"

	//获取博客作者Id
	author := getAuthor(blogPage)
	//	fmt.Printf("blogAuthor = #%v#\n", author)
	blogDetail += "author =" + author + "\n"

	blogDetail += "blogUrl =" + blogUrl + "\n"

	//获取博客发布时间
	publishTime := getBlogPTime(blogPage)
	//	fmt.Println("getBlog publishTime = ", publishTime)
	blogDetail += "publishTime =" + publishTime + "\n"

	//获取博客阅读量
	readTimes := getReadTimes(blogPage)
	//	fmt.Println("getBlog readTimes = ", readTimes)
	blogDetail += "readTimes =" + readTimes + "\n"

	return
}

//爬取博客内容，得到每一页的结果
func GetBlog(url string) (result string, err error) {
	//获取搜索的页面内容
	srcResult, err1 := SpiderOnePage(url)
	if err1 != nil {
		err = err1
		return
	}

	//爬取搜索的页面
	result = getBlogResult(srcResult)

	fmt.Println("getBlogResult result= \n", result)

	return
}

func SpiderPage(i int, page chan int) {
	//搜索返回得到的网址为三部分，中间部分为页码
	urlHead := "https://so.csdn.net/so/search/s.do?p="
	urlEnd := "&q=%E5%8C%BA%E5%9D%97%E9%93%BE&t=blog&domain=&o=&s=&u=&l=&f=&rbg=0"
	url := urlHead + strconv.Itoa(i) + urlEnd
	fmt.Println("url=", url)
	//爬取博客内容，得到每一页的结果
	result, err := GetBlog(url)
	if err != nil {
		fmt.Println("StartSpider err=", err)
		return
	}

	//把爬取的内容写入到文件

	fileName := "csdn区块链博客" + strconv.Itoa(i) + ".txt"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err ", err)
		return
	}

	f.WriteString(result)
	defer f.Close()
	page <- i
}

//在csdn博客主页输入区块链，进行搜索
func StartSpider(startPage, endPage int) {
	//for循环，遍历每一页的地址
	page := make(chan int)
	for i := startPage; i <= endPage; i++ {
		go SpiderPage(i, page) //并发的爬取每一页内容
	}
	for i := startPage; i <= endPage; i++ {
		fmt.Printf("第%d个网页爬取完成\n", <-page)
	}

}

func main() {
	var startPage, endPage int
	fmt.Println("请输入爬取的第一页")
	fmt.Scan(&startPage)
	fmt.Println("请输入爬取的最后一页")
	fmt.Scan(&endPage)
	StartSpider(startPage, endPage)
}

