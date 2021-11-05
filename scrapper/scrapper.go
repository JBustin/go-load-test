package scrapper

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/go-load-test/network"
	"github.com/go-load-test/utils"
	"golang.org/x/net/html"
)

type scrapper struct {
	client  network.HTTPClient
	urls    []string
	headers map[string]string
}

func New(urls []string, headers map[string]string) scrapper {
	return scrapper{urls: urls, headers: headers, client: &http.Client{}}
}

func (s scrapper) GetLinks() ([]string, error) {
	var wg sync.WaitGroup

	var links []string
	for _, pageUrl := range s.urls {
		wg.Add(1)
		go func(pageUrl string) {
			defer wg.Done()
			pageLinks, err := s.getLinkByUrl(pageUrl)
			if err != nil {
				log.Fatal(err)
			} else {
				links = append(links, pageLinks...)
			}
		}(pageUrl)
	}

	wg.Wait()
	return utils.Dedupe(links), nil
}

func (s scrapper) getLinkByUrl(pageUrl string) ([]string, error) {
	myUrl, err := url.Parse(pageUrl)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", pageUrl, nil)
	if err != nil {
		return nil, err
	}
	defer func() { req.Close = true }()
	for key, value := range s.headers {
		req.Header.Add(key, value)
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return extractLinks(string(bodyBuf), fmt.Sprintf("%s://%s", myUrl.Scheme, myUrl.Hostname()))
}

func extractLinks(content string, hostname string) ([]string, error) {
	var links []string

	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return nil, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if strings.HasPrefix(a.Val, "/") {
						links = append(links, fmt.Sprintf("%s%s", hostname, a.Val))
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links, nil
}
