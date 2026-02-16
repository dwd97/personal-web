package main

import (
	"bytes"
	"encoding/json"
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

type NoteItem struct {
	Title string
	URL   string
}

type NoteSection struct {
	Name  string
	Items []NoteItem
}

type NotesConfig struct {
	Sections []NotesSection `json:"sections"`
}

type NotesSection struct {
	ID    string      `json:"id"`
	Title string      `json:"title"`
	Items []NotesItem `json:"items"`
}

type NotesItem struct {
	File  string `json:"file"`
	Title string `json:"title"`
	Type  string `json:"type"` // "md" | "pdf" | "link"
	URL   string `json:"url"`
}

func loadNotesConfig() NotesConfig {
	data, err := os.ReadFile("content/notes/notes.json")
	if err != nil {
		log.Fatal(err)
	}

	var cfg NotesConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
func processNotesFromConfig(cfg NotesConfig) []NoteSection {
	var sections []NoteSection

	for _, s := range cfg.Sections {
		sec := NoteSection{Name: s.Title}

		for _, item := range s.Items {
			src := filepath.Join(notesDir, s.ID, item.File)

			switch item.Type {
			case "md":
				content, _ := os.ReadFile(src)
				url := createNotePage(
					filepath.Join(s.ID, item.File),
					item.Title,
					string(content),
				)
				sec.Items = append(sec.Items, NoteItem{Title: item.Title + " (MD)", URL: url})

			case "pdf":
				dst := filepath.Join(publicDir, "notes", s.ID, item.File)
				createDirIfNotExist(filepath.Dir(dst))
				copyFile(src, dst)

				url := "/notes/" + s.ID + "/" + item.File
				sec.Items = append(sec.Items, NoteItem{Title: item.Title + " (PDF)", URL: url})
			case "link":
				sec.Items = append(sec.Items, NoteItem{Title: item.Title, URL: item.URL})
			}
		}

		sections = append(sections, sec)
	}

	return sections
}

func renderNotes(sections []NoteSection) string {
	var b strings.Builder

	b.WriteString("<h2>Zápisky MFF UK</h2>")

	for _, s := range sections {
		b.WriteString("<h3>" + s.Name + "</h3><ul>")
		for _, item := range s.Items {
			b.WriteString(`<li><a href="` + item.URL + `">` +
				item.Title + `</a></li>`)
		}
		b.WriteString("</ul>")
	}

	return b.String()
}

func createNotePage(relPath, title, md string) string {
	htmlContent := markdownToHTML(md)

	noteInner := renderTemplate("templates/page.html", map[string]string{
		"TITLE":   title,
		"CONTENT": htmlContent,
	})

	full := renderWithBase(title, noteInner)

	outDir := filepath.Join(publicDir, "notes",
		strings.TrimSuffix(relPath, filepath.Ext(relPath)))

	createDirIfNotExist(outDir)

	os.WriteFile(filepath.Join(outDir, "index.html"), []byte(full), 0644)

	return "/notes/" + filepath.ToSlash(
		strings.TrimSuffix(relPath, filepath.Ext(relPath))) + "/"
}

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
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
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

func generateHome(posts []string, notes []NoteSection) {
	indexInner := renderTemplate("templates/index.html", map[string]string{})

	var list strings.Builder
	for _, slug := range posts {
		list.WriteString(`<li><a href="/posts/` + slug + `/">` + slug + `</a></li>`)
	}

	indexInner = strings.Replace(indexInner, "{{BLOG_POSTS}}", list.String(), 1)
	indexInner = strings.Replace(indexInner, "{{NOTES}}", renderNotes(notes), 1)

	final := renderWithBase("Home", indexInner)

	os.WriteFile(filepath.Join(publicDir, "index.html"), []byte(final), 0644)
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

	cfg := loadNotesConfig()
	notes := processNotesFromConfig(cfg)

	copyDir(staticDir, publicDir)
	copyFile("CNAME", filepath.Join(publicDir, "CNAME"))

	generateHome(posts, notes)
}
