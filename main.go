package main

import (
	"Capture/module"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"path"
	"strings"
)

type request struct {
	ID  int
	URL string
}

var TaskQueue = make(chan request, 100)

func InitWorker() {
	for i := 0; i < 100; i++ {
		module.GoPool.Run(func() {
			dp := module.NewChrome()
			for req := range TaskQueue {
				t := req.URL
				if strings.Index(t, "http") != 0 {
					t = fmt.Sprintf("http://%s", t)
				}
				u, err := url.Parse(t)
				if err != nil {
					continue
				}
				dp.Capture(t, path.Join("result", fmt.Sprintf("%d_%s.png", req.ID, u.Host)))
			}
		})
	}

}
func InsertTask() {
	task, err := ioutil.ReadFile("task.txt")
	if err != nil {
		panic(err)
	}
	tasks := strings.Split(string(task), "\n")
	for i, t := range tasks {
		log.Println("[*] Push Task : ", t)

		TaskQueue <- request{
			ID:  i,
			URL: t,
		}
	}

}
func main() {
	go InitWorker()
	InsertTask()
	select {}
}
