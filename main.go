package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

    files, err := os.ReadDir("posts")
    if err != nil {
        panic(err)
    }

    fmt.Println("Generating site...")

    for _, file := range files {
        if filepath.Ext(file.Name()) == ".md" {
            processFile(file.Name(), outputDir)
        }
    }
    fmt.Println("Done! Site generated in /public")
}

func processFile(fileName, outputDir string) {
    // 1. Open Input (Markdown)
    inputPath := filepath.Join("posts", fileName)
    f, err := os.Open(inputPath)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // 2. Parse
    doc := ParseMarkdownReader(f)

    // 3. Create Output File (HTML)
    htmlName := strings.Replace(fileName, ".md", ".html", 1)
    outputPath := filepath.Join(outputDir, "posts", htmlName)
    
    outFile, err := os.Create(outputPath)
    if err != nil {
        panic(err)
    }
    defer outFile.Close()

    // 4. Write Content
    writer := bufio.NewWriter(outFile)
    
    // Write Header
    writer.WriteString("<!doctype html><html><head><title>Post</title></head><body>")
    writer.WriteString("<nav><a href='/'>&larr; Back to Home</a></nav><hr>")

    // Write Body
    for _, el := range doc {
        writer.WriteString(el.Render() + "\n")
    }

    // Write Footer
    writer.WriteString("</body></html>")
    writer.Flush()
}
