# Personal web

## file structure

```
personal-web/
├─ .github/                     # entrypoints
│   └─ workflows/
│       └─ deploy.yml       # for deplyoing using github actions
├─ cmd/                     # entrypoints
│   └─ blog/
│       └─ main.go
├─ internal/                # core logic (not importable externally)
│   ├─ markdown/            # wrapper around goldmark
│   ├─ parser/              # frontmatter parsing
│   ├─ renderer/            # HTML rendering
│   ├─ site/                # site builder (routing, pages)
│   └─ fs/                  # file loading utilities
├─ templates/               # HTML templates
│   ├─ base.html
│   ├─ post.html
│   ├─ page.html
│   └─ index.html
├─ content/                 # ALL source content
│   ├─ posts/
│   ├─ notes/               # folder for notes, categorized by semester. In Markdown, PDF and etc.
│   └─ pages/               # other necessary pages
│       ├─ about.md
│       └─ contact.md
├─ static/                 # copied as-is
│   ├─ css/
│   ├─ images/
│   └─ fonts/
├─ public/                 # GENERATED output (deploy this)
│   ├─ index.html
│   ├─ posts/
│   ├─ notes/
│   └─ CNAME
├─ go.mod
├─ go.sum
├─ .gitignore
├─ Documentation.md
└─ README.md
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
  - [ ] last updated date
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
