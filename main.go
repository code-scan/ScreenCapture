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

var TaskQueue = make(chan string, 100)

func InitWorker() {
	for i := 0; i < 100; i++ {
		module.GoPool.Run(func() {
			dp := module.NewChrome()
			for t := range TaskQueue {
				if strings.Index(t, "http") != 0 {
					t = fmt.Sprintf("http://%s", t)
				}
				u, err := url.Parse(t)
				if err != nil {
					continue
				}
				dp.Capture(t, path.Join("result", fmt.Sprintf("%s.png", u.Host)))
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
	for _, t := range tasks {
		log.Println("[*] Push Task : ", t)
		TaskQueue <- strings.TrimSpace(t)
	}

}
func main() {
	go InitWorker()
	InsertTask()
	select {}
}
