package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron/v3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var barkKey string

func main() {
	barkKey = "test"
	var wg sync.WaitGroup
	wg.Add(1)
	nyc, _ := time.LoadLocation("Asia/Shanghai")
	saveLog(barkPush("测试发送", barkKey, "测试"))
	c := cron.New(cron.WithLocation(nyc))
	_, err := c.AddFunc("0 7-23 * * MON-FRI", getlubu)
	if err != nil {
		saveLog(err.Error())
		return
	}
	getlubu()
	c.Start()
	wg.Wait()
	saveLog("已结束运行")
}

func getlubu() {
	// Request the HTML page.
	res, err := http.Get("https://www.boc.cn/sourcedb/whpj/")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	gong := "body > div.wrapper > div.BOC_main > div.publish > div:nth-child(3) > table > tbody > tr:nth-child("
	namegps := ") > td:nth-child(1)"
	salegps := ") > td:nth-child(2)"
	buygps := ") > td:nth-child(4)"
	timegps := ") > td:nth-child(8)"
	for i := 2; i < 30; i++ {
		name := gong + strconv.Itoa(i) + namegps
		nameText := doc.Find(name).Text()
		if nameText == "卢布" {
			saleName := gong + strconv.Itoa(1) + ") > th:nth-child(2)"
			buyName := gong + strconv.Itoa(1) + ") > th:nth-child(4)"
			timeName := gong + strconv.Itoa(1) + ") > th:nth-child(8)"
			saleDom := gong + strconv.Itoa(i) + salegps
			buyDom := gong + strconv.Itoa(i) + buygps
			timeDom := gong + strconv.Itoa(i) + timegps
			saleNameText := doc.Find(saleName).Text()
			buyNameText := doc.Find(buyName).Text()
			timeNameText := doc.Find(timeName).Text()
			saleText := doc.Find(saleDom).Text()
			buyText := doc.Find(buyDom).Text()
			timeText := doc.Find(timeDom).Text()
			body := saleNameText + ": " + saleText + "\n" + buyNameText + ": " + buyText + "\n" + timeNameText + ": " + timeText
			saveLog(barkPush(body, barkKey, nameText))
			break
		}
		if i == 29 {
			result, _ := ioutil.ReadAll(res.Body)
			saveLog(string(result))
			saveLog(barkPush("貌似被封了", barkKey, nameText))

		}
	}
}
