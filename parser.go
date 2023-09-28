package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type ContentInfo struct {
	url   string
	title string
	links []string
}

func (c ContentInfo) toStr() string {
	builder := strings.Builder{}
	builder.WriteString(c.url)
	builder.WriteString("\n")
	builder.WriteString(c.title)
	builder.WriteString("\n")
	if len(c.links) > 8 {
		for i := 0; i < 8; i++ {
			builder.WriteString("\t")
			builder.WriteString(c.links[i])
			builder.WriteString("\n")
		}
		builder.WriteString("and ")
		builder.WriteString(fmt.Sprint(len(c.links) - 8))
		builder.WriteString(" more\n")
	} else {
		for _, l := range c.links {
			builder.WriteString("\t")
			builder.WriteString(l)
			builder.WriteString("\n")
		}
	}
	builder.WriteString("\n\n")
	return builder.String()
}

func getContentInfo(content string, url string) ContentInfo {
	node, err := html.Parse(strings.NewReader(content))
	if err != nil {
		fmt.Println(err)
		return ContentInfo{}
	}
	return extractInfo(node, url)
}

func extractInfo(node *html.Node, url string) ContentInfo {
	var info ContentInfo
	var traverse func(*html.Node)
	info.url = url

	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			info.title = n.FirstChild.Data
		}
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, att := range n.Attr {
				if att.Key == "href" {
					info.links = append(info.links, att.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(node)
	return info
}
