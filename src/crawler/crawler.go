package crawler

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

//PUBLIC FUNCTIONS

func Crawler(firstUrl string){

	urls, _ := getUrls(firstUrl)

	fmt.Println(len(urls))

	for true{
		urls = getUrlsInListOfUrls(urls)
		fmt.Println(len(urls))
	}
}

func GetListOfTag(url string, tag string) ([]html.Token, error) {

 	var content []html.Token

  	response, err := http.Get(url)

	if err == nil {
		tokens := html.NewTokenizer(response.Body)

		for {
			nextToken := tokens.Next()

			switch {
				case nextToken == html.ErrorToken://end of the function
					response.Body.Close()
					return content, nil

				case nextToken == html.StartTagToken:
					token := tokens.Token()

					isTheTag := token.Data == tag

					if isTheTag {
						content = append(content, token)
					}
			}
		}
	}else{
		return content, &errorString{"something went wrong with " + url}
	}
}

func GetAttr(token html.Token, attr string) (string){
	for _, token := range token.Attr {
		if token.Key == attr {
			return token.Val
		}
	}
	return ""
}

//PRIVATE FUNCTIONS

func getUrls(url string) ([]string, error) {
	var urls []string

	tokens, err := GetListOfTag(url, "a")
	for _, token := range tokens {
		urls = append(urls, GetAttr(token, "href"))
	}
	return urls, err
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

type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
