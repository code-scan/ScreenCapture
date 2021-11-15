package module

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)

type Chrome struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewChrome() *Chrome {
	c := Chrome{}
	c.ctx, c.cancel = chromedp.NewContext(context.Background())
	return &c
}

func (this *Chrome) Capture(url, save string) {
	var buf []byte
	if url == "" {
		return
	}
	if err := chromedp.Run(this.ctx, this.fullScreenshot(url, 90, &buf)); err != nil {
		log.Println(url, err)
		return
	}
	if err := ioutil.WriteFile(save, buf, 0o644); err != nil {
		log.Println(url, err)
		return
	}
	log.Println("[*] Capture Successful , ", url)
}
func (this *Chrome) fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
