package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
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

func renderTemplate(path string, data map[string]string) string {
	tpl, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	html := string(tpl)

	for k, v := range data {
		html = strings.ReplaceAll(html, "{{"+k+"}}", v)
	}

	return html
}

func renderWithBase(title, inner string) string {
	base := renderTemplate("templates/base.html", map[string]string{
		"TITLE":   title,
		"CONTENT": inner,
	})
	return base
}

/*
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
*/

func createPost(title string, md string) string {
	htmlContent := markdownToHTML(md)

	postInner := renderTemplate("templates/post.html", map[string]string{
		"TITLE":   title,
		"DATE":    "",
		"CONTENT": htmlContent,
	})

	full := renderWithBase(title, postInner)

	slug := slugify(title)
	outputDir := filepath.Join(publicDir, "posts", slug)
	createDirIfNotExist(outputDir)

	err := os.WriteFile(filepath.Join(outputDir, "index.html"), []byte(full), 0644)
	if err != nil {
		log.Fatal(err)
	}

	return slug
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
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM, // enables tables, strikethrough, task lists, etc.
		),
	)
	if err := md.Convert([]byte(content), &buf); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func readMarkdownPosts(dir string) []string {
	var posts []string

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
			return err
		}

		if info.IsDir() || filepath.Ext(info.Name()) != ".md" {
			return nil
		}

		content, _ := os.ReadFile(path)

		title := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
		slug := createPost(title, string(content))

		posts = append(posts, slug)
		return nil
	})

	return posts
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

func copyFile(src, dst string) {
	data, err := os.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func generateHome(posts []string) {
	indexInner := renderTemplate("templates/index.html", map[string]string{})

	var list strings.Builder
	for _, slug := range posts {
		list.WriteString(`<li><a href="/posts/` + slug + `/">` + slug + `</a></li>`)
	}

	indexInner = strings.Replace(indexInner, "{{BLOG_POSTS}}", list.String(), 1)

	final := renderWithBase("Home", indexInner)

	err := os.WriteFile(filepath.Join(publicDir, "index.html"), []byte(final), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func createPage(title string, md string) {
	htmlContent := markdownToHTML(md)

	pageInner := renderTemplate("templates/page.html", map[string]string{
		"TITLE":   title,
		"CONTENT": htmlContent,
	})

	full := renderWithBase(title, pageInner)

	slug := slugify(title)
	outputDir := filepath.Join(publicDir, slug)
	createDirIfNotExist(outputDir)

	err := os.WriteFile(filepath.Join(outputDir, "index.html"), []byte(full), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func readPages(dir string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
			return err
		}

		if info.IsDir() || filepath.Ext(info.Name()) != ".md" {
			return nil
		}

		content, _ := os.ReadFile(path)

		title := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
		createPage(title, string(content))

		return nil
	})
}

func main() {
	prepareOutputDirs()

	posts := readMarkdownPosts(postsDir)

	readPages(pagesDir)

	copyDir(notesDir, filepath.Join(publicDir, "notes"))
	copyDir(staticDir, publicDir)
	copyFile("CNAME", filepath.Join(publicDir, "CNAME"))

	generateHome(posts)
}
