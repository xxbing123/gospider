package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func fetch (url string) string {
	fmt.Println("Fetch Url",url)
	client := &http.Client{}
	req,_ :=  http.NewRequest("GET",url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Http get err:",err)
		return ""
	}

	if resp.StatusCode !=200 {
		fmt.Println("Http status code:", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Read error", err)
		return ""
	}

	return string(body)
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func parseUrls(url string) {
	body := fetch(url)
	body = strings.Replace(body, "\n", "", -1)
	rp := regexp.MustCompile(`<div class="hd">(.*?)</div>`)
	titleRe := regexp.MustCompile(`<span class="title">(.*?)</span>`)
	idRe := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(\d+)/"`)
	items := rp.FindAllStringSubmatch(body, -1)
	for _, item := range(items) {
		//fmt.Println(idRe.FindStringSubmatch(item[1])[1],
		//	titleRe.FindStringSubmatch(item[1])[1])
		//fmt.Println(item[1]) // item[1]是匹配到的字符串
		id := idRe.FindStringSubmatch(item[1])  // 获取ID
		title := titleRe.FindStringSubmatch(item[1])
		fmt.Println(id[1],title[1])
	}
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	for i :=0; i < 10; i++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()

			parseUrls("https://movie.douban.com/top250?start=" + strconv.Itoa(25 * page))
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Took %s\n",elapsed)

}
