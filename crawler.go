package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}

func getUrls(url string) ([]string, error) {
	return getListOfTag(url, "a")
}

func getListOfTag(url string, tag string) ([]string, error) {
	//fmt.Println("URL: ", url)

 	var urls []string

  	response, err := http.Get(url)

	if err == nil {
		tokens := html.NewTokenizer(response.Body)

		for {
			nextToken := tokens.Next()

			switch {

			case nextToken == html.ErrorToken://end of the function
				response.Body.Close()
				return urls, nil

			case nextToken == html.StartTagToken:
				token := tokens.Token()

			isAnchor := token.Data == tag
				if isAnchor {

					for _, a := range token.Attr {
						if a.Key == "href" {
							urls = append(urls, a.Val)
						}
					}
				}
			}
		}
	}else{
		return urls, &errorString{"couldn't open " + url}
	}
}


func getUrlsInListOfUrls(slice []string)([]string) {

	var urlsList []string

	for _, url := range slice {
		urls, err := getUrls(url)
		if err == nil {
			urlsList = append(urlsList, urls...)
		}

	}

	return urlsList
}

func printUrls(slice []string){
	for _, url := range slice {
		fmt.Println(url)
	}
}

func Crawler(firstUrl string){
	urls, _ := getUrls(firstUrl)

	fmt.Println(len(urls))

	for true{
		urls = getUrlsInListOfUrls(urls)
		fmt.Println(len(urls))
	}
}

func main() {

	Crawler("https://github.com/Lucasfrota")
	/*
  	firstUrl := "https://github.com/Lucasfrota"//"https://pt.wikipedia.org/wiki/Android"//"https://www.google.com.br/"

	urls, _ := getUrls(firstUrl)
	fmt.Println(len(urls))//printUrls(urls)

	fmt.Println("---------------------------\n\n")

	urls = getUrlsInListOfUrls(urls)
	fmt.Println(len(urls))//printUrls(urls)

	fmt.Println("---------------------------\n\n")

	urls = getUrlsInListOfUrls(urls)
	fmt.Println(len(urls))//printUrls(urls)
	*/
}
