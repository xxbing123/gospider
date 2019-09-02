package main

import (
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"time"
)

/*
css选择器11
*/

func fetch(url string) *goquery.Document {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transCfg,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Http get err:", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Http status code:", resp.StatusCode)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func main() {
	//fmt.Println("https://top.chinaz.com/all/")
	doc := fetch("https://top.chinaz.com/all/")
	doc.Find("#content > div.Wrapper > div.TopListCent > div li").Each(func(index int, ele *goquery.Selection) {
		e := ele.Find(".clearfix > div.CentTxt > h3.rightTxtHead")
		if title, ok := e.Find("a").Attr("title"); ok {
			fmt.Print(title + "   ")
		}

		fmt.Print(e.Find(".col-gray").Eq(0).Text())
		fmt.Println()

	})
	//fmt.Println(res)

}
