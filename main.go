package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod"
)

var url string

func main() {
	fmt.Println("Enter url")
	fmt.Scanf("%s", &url)
	startTime := time.Now()
	page := rod.New().MustConnect().MustPage(url).MustWaitLoad()
	fmt.Println("Time taken to load", time.Now().Sub(startTime))

	os.Mkdir(".screenshots", 0755)
	fileName := fmt.Sprintf(".screenshots/%s.webp", page.MustInfo().Title)
	page.MustScreenshot(fileName)
}
