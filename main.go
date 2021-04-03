package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//how to handle concurrency
//we want small, synchronous, iterator based, fault tolerant with zero external dependencies based html parser
//Amazon can temporarily block the IP from which automated requests go. Different means can be used for it.
//For example, Amazon may show a captcha or a page with an error. Therefore, for the scraper to work successfully,
//we need to think about how it will catch and bypass these cases.
//can we use go channels
//it is not guaranteed that the io.Writer is safe to use concurrent
//try accessing data using proxies
//what is user agent in headers
//running benchmarks
//can we use regex for this?
//css selector vs XPath
//remove all /n, /t
//make a complete path of the document
//what other libraries can be used?


func main3() {
	dataBytes, err := ioutil.ReadFile("B00KY1U7GM.txt")
	if err != nil{
		fmt.Print(err)
	}
	fmt.Print(dataBytes)
	node, err := html.Parse(strings.NewReader(string(dataBytes)))
	fmt.Print(node)
}

func main() {

	//'Cookie: i18n-prefs=EUR; session-id=257-6417912-1788422; session-id-time=2082787201l; session-token=aKgjiosGVB1i4pr/2qc6Ova6WcrTbxX0F1ZGHi8BZYez0Teb5/Srlv7/oKQlY1QzOHCxvtpyfTJ7lnQonf5CQYt2QoRzHt4SmBkksyMzXNElVf1h8vXrj7Wk6rCXX7XozsINre4d3dYPYWF2lv81Zmi0L++99dHbYLAMS9TVs1DQgOgo9ChDc5sjRCQvY8Jg; ubid-acbde=258-3360747-3990251'
	headers := map[string]string{
		//"content-type": "text/html",
		//"accept-encoding":"gzip, deflate, br",
		//"accept": "*/*"
		"Cookie": "i18n-prefs=EUR",
		"session-id": "257-6417912-1788422",
		"session-token": "aKgjiosGVB1i4pr/2qc6Ova6WcrTbxX0F1ZGHi8BZYez0Teb5/Srlv7/oKQlY1QzOHCxvtpyfTJ7lnQonf5CQYt2QoRzHt4SmBkksyMzXNElVf1h8vXrj7Wk6rCXX7XozsINre4d3dYPYWF2lv81Zmi0L++99dHbYLAMS9TVs1DQgOgo9ChDc5sjRCQvY8Jg",
		"accept": "text/html",
		"ubid-acbde":"258-3360747-3990251",
		//"accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	}

	req, _ := http.NewRequest("GET", "http://www.amazon.de/gp/product/B00KY1U7GM", nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	//resp, err := http.Get("https://www.amazon.de/gp/product/B00KY1U7GM")
	//if err != nil{
	//	fmt.Print(err)
	//}

	dataBytes, err := io.ReadAll(resp.Body)
	fmt.Print(dataBytes)

	doc, err := html.Parse(resp.Body)
	if err != nil{
		fmt.Print(err)
	}

	fmt.Print(doc)
}