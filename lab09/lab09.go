package main

import (
	"fmt"
	"flag"
	"github.com/gocolly/colly"
)

func main() {
	maximum := flag.Int("max", 10, "Max number of comments to show")
	flag.Parse()

	c := colly.NewCollector()

	var cnt = 1
	c.OnHTML(".push", func(e *colly.HTMLElement) {
		if cnt > *maximum {
			return 
		}
		fmt.Printf("%d. ", cnt)	
		e.ForEach(".f3.hl.push-userid", func(_ int, el *colly.HTMLElement) {
			fmt.Printf("名字：%s", el.Text)
		})
		e.ForEach(".f3.push-content", func(_ int, el *colly.HTMLElement) {
			fmt.Printf("，留言%s", el.Text)
		})
		e.ForEach(".push-ipdatetime", func(_ int, el *colly.HTMLElement) {
			fmt.Printf("，時間：%s", el.Text)
		})
		cnt++
	
	})

	c.Visit("https://www.ptt.cc/bbs/joke/M.1481217639.A.4DF.html")
}
