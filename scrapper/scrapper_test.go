package scrapper

import (
	"testing"

	"github.com/go-load-test/mocks"
	"github.com/stretchr/testify/assert"
)

var htmlContent string = `<html>
	<head></head>
	<body>
		<div id="main-container">
			<a href="/homepage"><h1>Title</h1></a>
		</div>
		<div id="footer>
			<ul>
				<li><a href="https://external-link.com">entry 1</a></li>
				<li><a href="/page1">entry 2</a></li>
				<li><a href="/page2">entry 3</a></li>
				<li><span href="/page1">not a link</a></li>
			</ul>
		</div>
	</body>
</html>`

func Test_ExtractLinks(t *testing.T) {
	hostname := "https://mysite.com"
	links, err := extractLinks(htmlContent, hostname)

	assert.Equal(
		t,
		nil,
		err,
		"Extract links should not return an error",
	)

	assert.Equal(
		t,
		[]string{
			"https://mysite.com/homepage",
			"https://mysite.com/page1",
			"https://mysite.com/page2",
		},
		links,
		"Extract internal links from html",
	)
}

func Test_Scrapper(t *testing.T) {
	s := scrapper{
		urls: []string{"https://mysite.com/homepage"},
		headers: map[string]string{
			"x-custom-header-1": "abcd",
			"x-custom-header-2": "1234",
		},
		client: mocks.HTTPresponse(200, htmlContent, nil),
	}
	links, err := s.GetLinks()

	assert.Equal(
		t,
		nil,
		err,
		"Get links from urls should not return an error",
	)

	assert.Equal(
		t,
		[]string{
			"https://mysite.com/homepage",
			"https://mysite.com/page1",
			"https://mysite.com/page2",
		},
		links,
		"Get links from urls should work as expected",
	)
}
