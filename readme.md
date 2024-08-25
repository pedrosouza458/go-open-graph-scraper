
##  Golang Open Graph Scraper

This is a very simples Golang scraping that uses website OG (Open Graph) to get some informations like website name, page title, description and images, feel free to contribute forking this repo, the main goal is to use this package and my future project (that will also be open-source) and to study Golang, cause i know that packages like that already exists, but i wanna make my own.

  
  <p>
  <a href="https://github.com/pedrosouza458/go-open-graph-scraper/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/pedrosouza458/go-open-graph-scraper" alt="GitHub License" />
  </a>
  <a href="https://github.com/pedrosouza458/go-open-graph-scraper/commits/main/">
    <img src="https://img.shields.io/github/last-commit/pedrosouza458/go-open-graph-scraper/main" alt="GitHub Last Commit" />
  </a>
</p>

###  Run this code locally

  

```curl

git clone https://github.com/pedrosouza458/go-open-graph-scraper

```

```go

go mod tidy

```

```go

go run .

```

Now just type the website url in terminal, some sites will not display some informations.

  

##  Example:

  

###  Input

```

https://www.tabnews.com.br/marlon/ajuda-sujestao-de-conteudos-de-golang

```

  

###  Output

```

Logo: https://i.imgur.com/O5OOa3s.png

Website name: tabnews

Image:https://www.tabnews.com.br/api/v1/contents/marlon/ajuda-sujestao-de-conteudos-de-golang/thumbnail

Page Name:[Ajuda] Sujestão de conteúdos de Golang · marlon

Page Description: Olá TabNews, estou querendo aprender Golang, vejo que é uma linguagem muito boa para trabalhar com DevOps e considerada muito rápida e fácil. Gostaria de pedir ajuda aos universitários co...

```