package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/Mohanbarman/leetemplate/internal/leetcode"
	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./gltemplate <LEETCODE_URL>\n")
		os.Exit(1)
	}
	args := os.Args[1:]
	leetcodeUrlStr := args[0]

	url, err := url.ParseRequestURI(leetcodeUrlStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s\n", err)
		os.Exit(1)
	}

	re := regexp.MustCompile(`/problems/([^/]+)/`)
	match := re.FindStringSubmatch(url.Path)
	slug := ""

	if len(match) > 1 {
		slug = match[1]
		fmt.Println("Slug:", slug)
	} else {
		fmt.Fprintf(os.Stderr, "Error: Only leetcode url supported\n")
		os.Exit(1)
	}

	question, err := leetcode.GetQuestion(slug)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	z := html.NewTokenizer(strings.NewReader(question.Content))

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			fmt.Println("content html parse error")
			os.Exit(1)
		case html.TextToken:
			fmt.Printf("%s\n", string(z.Text()))
			break
		case html.StartTagToken:
			tn, _ := z.TagName()
			fmt.Printf("open tag: %s\n", tn)
			break
		case html.EndTagToken:
			tn, _ := z.TagName()
			fmt.Printf("close tag: %s\n", tn)
			break
		}

	}
}
