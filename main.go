package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	link := "https://www.togai.com/blog"
	response, err := getHtmlPage(link)
	if err != nil {
		log.Fatalf("error occurred while getting html page for %s: %v", link, err)
	}

	parseHtml(response)
}

func getHtmlPage(link string) (string, error) {
	res, err := sendRequest(link)
	if err != nil {
		return "", fmt.Errorf("error occurred while sending request to %s: %v", link, err)
	}

	contentType := res.Header.Get("Content-Type")
	if contentType != "text/html" {
		return "", fmt.Errorf("error occured: expected HTML response got non HTML response - %s", contentType)
	}

	htmlResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error occurred while reading HTML response: %v", err)
	}

	return string(htmlResponse), nil
}

func sendRequest(link string) (*http.Response, error) {
	ctx := context.TODO()

	url, err := url.Parse(link)
	if err != nil {
		return nil, fmt.Errorf("error occurred while parsing link: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error occurred while forming request for %s: %v", link, err)
	}

	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error occurred while sending request to %s: %v", link, err)
	}

	return res, nil
}

func parseHtml(htmlContent string) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
