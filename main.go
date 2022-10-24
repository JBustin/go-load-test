package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-load-test/config"
	"github.com/go-load-test/scrapper"
	"github.com/go-load-test/utils"
	"github.com/go-load-test/worker"
)

func main() {
	fmt.Println("- Go Load Test -")

	f := flag.String("f", "test.json", "JSON filepath.")
	flag.Parse()

	conf, err := config.New(*f)
	handlerErr(err)

	fmt.Println(conf)

	var urls []string

	if conf.Scrap {
		s := scrapper.New(conf.Urls, conf.Headers)
		urls, err = s.GetLinks()
		handlerErr(err)
	} else {
		urls = conf.Urls
	}

	if len(urls) == 0 {
		fmt.Println("No found url, exit.")
		os.Exit(0)
	}

	urls = utils.Fill(urls, conf.Hits)
	worker := worker.New(urls, conf)
	err = worker.Process()
	handlerErr(err)

	fmt.Println(worker.Report())
}

func handlerErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
