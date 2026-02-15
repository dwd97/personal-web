package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/chasefleming/elem-go"
	"github.com/yuin/goldmark"
)

const (
	contentDir   = "content"
	postsDir     = "content/posts"
	pagesDir     = "content/pages"
	notesDir     = "content/notes"
	templatesDir = "templates"
	staticDir    = "static"
	publicDir    = "public"
	indexTpl     = "templates/index.html"
)

func prepareOutputDirs() {
	createDirIfNotExist(publicDir)
	createDirIfNotExist(filepath.Join(publicDir, "posts"))
	createDirIfNotExist(filepath.Join(publicDir, "notes"))
}

func layout(title string, content elem.Node) string {
	htmlPage := elem.Html(nil,
		elem.Head(nil, elem.Title(nil, elem.Text(title)),
			elem.Body(nil, elem.Header(nil, elem.H1(nil, elem.Text(title))), elem.Main(nil, content), elem.Footer(nil, elem.Text("Footer content here")))))
	return htmlPage.Render()
}

func createHTMLPage(title, content string) string {
	htmlOutput := layout(title, elem.Raw(content))

	filename := slugify(title) + ".html"
	filePath := filepath.Join(publicDir, "posts", filename)

	err := os.WriteFile(filePath, []byte(htmlOutput), 0644)
	if err != nil {
		log.Fatal(err)
	}

	return filename
}

func slugify(s string) string {
	s = strings.ToLower(s)

	// replace spaces with dash
	s = strings.ReplaceAll(s, " ", "-")

	// remove accents (č → c, ř → r, etc.)
	t := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) || r == '-' {
			t = append(t, r)
		}
	}
	s = string(t)

	// collapse multiple dashes
	re := regexp.MustCompile(`-+`)
	s = re.ReplaceAllString(s, "-")

	return strings.Trim(s, "-")
}

func markdownToHTML(content string) string {
	var buf bytes.Buffer
	md := goldmark.New()
	if err := md.Convert([]byte(content), &buf); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func readMarkdownPosts(dir string) []string {
	var posts []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			htmlContent := markdownToHTML(string(content))
			title := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			postFilename := createHTMLPage(title, htmlContent)

			posts = append(posts, postFilename)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return posts
}

func generatePostListHTML(postFilenames []string) string {
	var items strings.Builder

	for _, filename := range postFilenames {
		items.WriteString(`<li><a href="public/` + filename + `">` + filename + `</a></li>` + "\n")
	}

	return items.String()
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755) // or 0700 if you need it to be private
		if err != nil {
			log.Fatal(err)
		}
	}
}

func copyDir(src, dst string) {
	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)

		if info.IsDir() {
			createDirIfNotExist(target)
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(target, data, 0644)
	})
}

func generateHome(posts []string) {
	templateBytes, err := os.ReadFile(indexTpl)
	if err != nil {
		log.Fatal(err)
	}

	var items strings.Builder

	for _, p := range posts {
		items.WriteString(`<li><a href="posts/` + p + `">` + p + `</a></li>\n`)
	}

	finalHTML := strings.Replace(
		string(templateBytes),
		"{{BLOG_POSTS}}",
		items.String(),
		1,
	)

	err = os.WriteFile(filepath.Join(publicDir, "index.html"), []byte(finalHTML), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	prepareOutputDirs()

	posts := readMarkdownPosts(postsDir)

	copyDir(notesDir, filepath.Join(publicDir, "notes"))
	copyDir(staticDir, publicDir)

	generateHome(posts)
}
