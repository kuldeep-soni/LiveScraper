package main

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"strings"
)

func AboutMovie(doc *html.Node) (*html.Node, error){
	var body, abtMovie *html.Node
	var getBody, getMovieInfo func(*html.Node)
	getBody = func(node *html.Node){
		if node.Type == html.ElementNode && node.Data == "body" {
			body = node
			return
		}
	}
	getMovieInfo = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			abtMovie = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			getMovieInfo(child)
		}
	}
	getBody(doc)
	if body == nil {
		return nil, errors.New("Missing <body> in the node tree")
	}
	getMovieInfo(body)
	if abtMovie == nil {
		return nil, errors.New("Missing movie meta in the node tree")
	}

	return abtMovie, nil
}

func Body(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if body != nil {
		return body, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func main22() {
	dataBytes, err := ioutil.ReadFile("B00KY1U7GM.txt")
	if err != nil{
		fmt.Print(err)
	}

	doc, _ := html.Parse(strings.NewReader(string(dataBytes)))
	bn, err := Body(doc)
	if err != nil {
		return
	}
	body := renderNode(bn)
	fmt.Println(body)
}

func renderDoc(doc *html.Node) {

}