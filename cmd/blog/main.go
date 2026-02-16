package main

import (
	"bytes"
	"encoding/json"
	stdhtml "html"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	mathjax "github.com/litao91/goldmark-mathjax" // You'll need to go get this
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
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

type ImageLinkRenderer struct{}

func (r *ImageLinkRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindImage, r.renderImageWithLink)
}

func (r *ImageLinkRenderer) renderImageWithLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.Image)
	src := string(n.Destination)

	// Create the WebP path
	ext := filepath.Ext(src)
	webpSrc := strings.TrimSuffix(src, ext) + ".webp"

	w.WriteString(`<a href="` + src + `" class="img-expand-link">`)
	w.WriteString(`<picture>`)
	// The browser will try webp first, but fall back to the png/jpg immediately if it fails
	w.WriteString(`<source srcset="` + webpSrc + `" type="image/webp">`)
	w.WriteString(`<img src="` + src + `" alt="` + string(n.Title) + `" loading="lazy" onerror="this.parentElement.querySelector('source').remove();">`)
	w.WriteString(`</picture>`)
	w.WriteString(`</a>`)

	return ast.WalkSkipChildren, nil
}

type CodeBlockRenderer struct{}

func (r *CodeBlockRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
}

func (r *CodeBlockRenderer) renderFencedCodeBlock(
	w util.BufWriter,
	source []byte,
	node ast.Node,
	entering bool,
) (ast.WalkStatus, error) {

	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.FencedCodeBlock)

	language := string(n.Language(source))

	filename := ""
	if n.Info != nil {
		meta := string(n.Info.Value(source))
		fields := strings.Fields(meta)
		if len(fields) >= 2 {
			filename = fields[1]
		}
	}

	// outer card
	w.WriteString(`<div class="code-block">`)

	// header
	if filename != "" {
		w.WriteString(`<div class="code-header">`)
		w.WriteString(stdhtml.EscapeString(filename))
		w.WriteString(`</div>`)
	}

	// scroll wrapper (THIS is the key change)
	w.WriteString(`<div class="code-scroll">`)

	// code
	w.WriteString(`<pre><code class="language-`)
	w.WriteString(stdhtml.EscapeString(language))
	w.WriteString(`">`)

	lines := n.Lines()
	for i := 0; i < lines.Len(); i++ {
		line := lines.At(i)
		w.WriteString(stdhtml.EscapeString(string(line.Value(source))))
	}

	w.WriteString(`</code></pre>`)

	// close scroll + card
	w.WriteString(`</div></div>`)

	return ast.WalkSkipChildren, nil
}

func renderWithBase(title, inner string) string {
	base := renderTemplate("templates/base.html", map[string]string{
		"TITLE":   title,
		"CONTENT": inner,
		"YEAR":    strconv.Itoa(time.Now().Year()),
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

type Heading struct {
	Level int
	Text  string
	ID    string
}

func extractHeadings(md string) []Heading {
	var result []Heading
	lines := strings.Split(md, "\n")

	// store current heading path per level
	path := make(map[int]string)

	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if !strings.HasPrefix(trim, "#") {
			continue
		}

		level := strings.Count(trim, "#")
		text := strings.TrimSpace(trim[level:])
		slug := slugify(text)

		// update path at this level
		path[level] = slug

		// clear deeper levels
		for l := level + 1; l <= 6; l++ {
			delete(path, l)
		}

		// build hierarchical id
		var parts []string
		for l := 1; l <= level; l++ {
			if p, ok := path[l]; ok {
				parts = append(parts, p)
			}
		}
		id := strings.Join(parts, "-")

		result = append(result, Heading{
			Level: level,
			Text:  text,
			ID:    id,
		})
	}

	return result
}

func addHeadingIDs(html string, headings []Heading) string {
	re := regexp.MustCompile(`<h([1-6])>(.*?)</h[1-6]>`)
	i := 0

	return re.ReplaceAllStringFunc(html, func(match string) string {
		if i >= len(headings) {
			return match
		}

		sub := re.FindStringSubmatch(match)
		level := sub[1]
		text := sub[2]
		id := headings[i].ID
		i++

		return `<h` + level + ` id="` + id + `">` + text + `</h` + level + `>`
	})
}

func renderTOC(headings []Heading) string {
	if len(headings) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString(`<nav class="toc"><strong>Index</strong><ul>`)

	for _, h := range headings {

		// normalize depth so first heading has no indent
		baseLevel := headings[0].Level
		depth := h.Level - baseLevel
		if depth < 0 {
			depth = 0
		}

		indentPx := depth * 16

		b.WriteString(`<li style="margin-left:` +
			strconv.Itoa(indentPx) + `px">` +
			`<a href="#` + h.ID + `">` + h.Text + `</a></li>`)
	}

	b.WriteString(`</ul></nav>`)
	return b.String()
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

	b.WriteString("<h2>Notes MFF UK (in Czech)</h2>")

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
	headings := extractHeadings(md)

	htmlContent := markdownToHTML(md)
	htmlContent = addHeadingIDs(htmlContent, headings)

	toc := renderTOC(headings)
	htmlContent = toc + htmlContent

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
	headings := extractHeadings(md)

	htmlContent := markdownToHTML(md)
	htmlContent = addHeadingIDs(htmlContent, headings)

	toc := renderTOC(headings)
	htmlContent = toc + htmlContent

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
			extension.GFM,
			mathjax.MathJax, // Ujisti se, že používáš verzi, která podporuje $
		),
		goldmark.WithRendererOptions(
			renderer.WithNodeRenderers(
				util.Prioritized(&CodeBlockRenderer{}, 100),
				util.Prioritized(&ImageLinkRenderer{}, 99),
			),
		),
	)
	if err := md.Convert([]byte(content), &buf); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func optimizeStaticImages(root string) {
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Println("Skipping optimization: ffmpeg not found")
		return
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".jpg" || ext == ".png" || ext == ".jpeg" {
			// Generate both modern formats
			webpPath := strings.TrimSuffix(path, ext) + ".webp"

			// 2. Optimize to WebP (Will work in GitHub Actions)
			if _, err := os.Stat(webpPath); os.IsNotExist(err) {
				cmd := exec.Command("ffmpeg", "-y", "-i", path, "-q:v", "75", webpPath)
				_ = cmd.Run() // Might fail locally but that's okay!
			}
		}
		return nil
	})
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

	// Optimize images in the static/images folder before processing posts
	optimizeStaticImages("static/images")

	posts := readMarkdownPosts(postsDir)
	readPages(pagesDir)

	cfg := loadNotesConfig()
	notes := processNotesFromConfig(cfg)

	copyDir(staticDir, publicDir)
	copyFile("CNAME", filepath.Join(publicDir, "CNAME"))

	generateHome(posts, notes)
}
