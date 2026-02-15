# Create a blog with Go and MD

I believe that everyone should at least try to have a blog or an online presence of some kind to express himself on the internet. It's the best place to network or share ideas. That's why I decided to create my personal website [davidkrejci.dev](https://www.davidkrejci.dev) and implement a blog using Go and MD. In this post I will just share how to create a simple blog.

With Katex support:

$E=mc^2$

$$
\int_a^b f(x)\,dx
$$

I will be using these libraries:

- [goldmark](https://github.com/yuin/goldmark)

> I will use Github Pages for deployment and Github Actions workflow for automatic website generation

For the file structure, create something similar to the file structure below. This will be the starting point.

```
personal-web/
├─ .github/
│   └─ workflows/
│       └─ deploy.yml       # for deplyoing using github actions
├─ cmd/
│   └─ blog/
│       └─ main.go          # main Go file for blog generation
├─ templates/               # HTML templates
│   ├─ base.html
│   ├─ post.html
│   ├─ page.html
│   └─ index.html
├─ content/                 # ALL source content
│   └─ posts/
├─ static/                  # will be copied over to public/
│   ├─ css/
│   ├─ images/
│   └─ fonts/
├─ public/                  # Generated output by main.go - will be done by Github Actions
│   ├─ index.html
│   ├─ posts/
│   ├─ notes/
│   └─ CNAME
├─ go.mod
├─ go.sum
├─ .gitignore
└─ README.md                # documentation
```

## Setup project

1. Create the github repository
   - create .gitignore file with `/public` as its content. This will be managed using Github Actions externally, no need to commit.
2. Setup DNS and custom domain
   - in your DNS settings at your domain provider add these records:
     | Type | Host | Answer |
     | ----------- | -------------- | ---- |
     | A | domainname.com | 185.199.108.153 |
     | A | domainname.com | 185.199.109.153 |
     | A | domainname.com | 185.199.110.153 |
     | A | domainname.com | 185.199.111.153 |
     | CNAME | www.domainname.com | gitUsername.github.io |
   - more: [github docs](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/managing-a-custom-domain-for-your-github-pages-site)
3. In Github pages select `Deploy from branch`

## Creating the template files
