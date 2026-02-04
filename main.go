package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type MDElement interface {
	Render() string
}

type Header struct {
	Level int
	Content string
}

func (h Header) Render() string {
	return fmt.Sprintf("<h%d>%s</h%d>", h.Level, h.Content, h.Level)
}

type Paragraph struct {
	Content string
}

func (p Paragraph) Render() string {
	return fmt.Sprintf("<p>%s</p>", p.Content)
}

func ParseMarkdownReader(r io.Reader) []MDElement {
	var elements []MDElement
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {continue}
		if level, content, isHeader := parseHeader(line); isHeader {
			elements = append(elements, Header{Level: level, Content: content})
		} else {
			elements = append(elements, Paragraph{Content: line})
		}
	}
	return elements
}

func parseHeader(line string) (int, string, bool) {
	level := 0

	runes := []rune(line)

	for _, r := range runes {
		if r == '#' {
			level++
		} else {
			break
		}
	}

	if level == 0 || level > 6 {
		return 0, "", false
	}

	content := strings.TrimSpace(line[level:])

	return level, content, true
}

func main() {
	outputDir := "public"

	os.RemoveAll(outputDir)
	if err := os.MkdirAll(outputDir + "/posts", 0755); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w,r,"index.html")
	})

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fileServer))

	mux.HandleFunc("GET /posts/{slug}", handlePost)

	fmt.Println("Server starting on the https://localhost:3030")
	err := http.ListenAndServe(":3030", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	path := "posts/" + slug + ".md"

	f,err := os.Open(path)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	defer f.Close()

	doc := ParseMarkdownReader(f)

	fmt.Fprint(w, "<!doctype html><html><head><title>Post</title></head><body>")

	fmt.Fprint(w, "<nav><a href='/'>&larr; Back to Home</a></nav><hr>")

	for _, el := range doc {
		fmt.Fprint(w, el.Render())
	}

	fmt.Fprint(w, "</body></html>")
}
