# Personal web

## file structure

```
personal-web/
├── .github/
│   └── workflows/
│       └── deploy.yml        # GitHub Actions config
├── cmd/
│   └── blog/
│       └── main.go           # Small entry point, few lines
├── internal/                 # Core application logic
│   ├── fs/
│   │   └── fs.go             # File system helpers (Copy, Write, Mkdir)
│   ├── markdown/
│   │   └── renderer.go       # Goldmark config (MathJax, Syntax Highlight)
│   ├── parser/
│   │   ├── frontmatter.go    # Metadata structs and parsing
│   │   └── utils.go          # Helpers (Slugify, Reading Time, Headings)
│   ├── renderer/
│   │   └── template.go       # HTML template management
│   └── site/
│       ├── site.go           # Main builder logic (The Orchestrator)
│       └── types.go          # Shared data models (Post, Note, Config)
├── templates/                # (Existing) HTML files
│   ├── base.html
│   ├── index.html
│   ├── page.html
│   └── post.html
├── content/                  # (Existing) Content source
│   ├── notes/
│   │   └── notes.json
│   ├── pages/                # (New) Move loose .md files here
│   │   ├── about.md
│   │   └── contact.md
│   └── posts/
├── static/                   # (Existing) Assets
│   ├── css/
│   │   └── style.css
│   ├── images/
│   ├── fonts/
│   └── CNAME                 # (Move here)
├── public/                   # (Generated) Gitignored
├── go.mod
├── go.sum
├── .gitignore
├── Documentation.md
└── README.md
```

## tasks

- [ ] blog
  - [ ] reading time estimate
  - [ ] sharing - creating the preview when sharing
    - [ ] rss feed
    - [ ] x, bluesky, linkedin, whatsapp, facebook, telegram, pinterest, mail, copy link
  - [ ] tags/categories
  - [ ] frontmatter yaml
  - [ ] seo
  - [ ] last updated date (maybe release date? separate?)
  - [ ] previous and next article
  - [ ] breadcrumbs
  - [ ] latex support
  - [ ] diagram support - from js code
  - [ ] code blocks - syntax highligting
  - [ ] table of contents - at the start or sticky
  - [ ] email newsletter
  - [ ] back to top button at the bottom
  - [ ] edit on github link
  - [ ] animation when opening blog post - title gets animated to the top
  - [ ] hearts, comments, views
- [ ] sitemap.xml
- [ ] ssg architecture - pre rendered html
- [ ] responsive
- [ ] dark/light mode
- [ ] seo meta tags
- [ ] canonical urls - dont have duplicates
- [ ] sitemap.xml
- [ ] robots.txt
- [ ] custom 404 page
- [ ] social links - github, x, bluesky, linkedin, mail
- [ ] image optimization
- [ ] sections of the website
  - [ ] Home
  - [ ] Posts (sorted by year)
  - [ ] About
  - [ ] Projects
