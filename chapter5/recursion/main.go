package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type Node struct {
	Type                  NodeType
	Data                  string
	Attr                  []Attribute
	FirstChild, NextChild *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	outline(nil, doc)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}

func CountElement(counts map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	CountElement(counts, n.FirstChild)
	CountElement(counts, n.NextSibling)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
