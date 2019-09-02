package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"time"
)

func fetch (url string) *goquery.Document {
	fmt.Println("Fetch Url",url)
	client := &http.Client{}
	req,_ :=  http.NewRequest("GET",url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Http get err:",err)

	}

	if resp.StatusCode !=200 {
		log.Fatal("Http status code:", resp.StatusCode)

	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func parseUrls(url string) {

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
