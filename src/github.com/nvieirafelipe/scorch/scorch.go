package main

import (
  "net/http"
  "os"
  "log"
  "strings"
  "github.com/bmizerany/pat"
  "github.com/nvieirafelipe/scorch/organization"
  "github.com/nvieirafelipe/scorch/repository"
)

func main() {
  routes := pat.New()
  routes.Get("/organizations/:organization_name/repositories", http.HandlerFunc(organization.Repositories))
  routes.Get("/repositories/:repository_name/workleft-vs-time", http.HandlerFunc(repository.WorkLeftVSTime))
  routes.Get("/public/", http.HandlerFunc(public))
  routes.Get("/", http.HandlerFunc(home))
  routes.Get("/:page", http.HandlerFunc(pages))

  http.Handle("/", routes)
  log.Println("listening...")

  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

func public(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, r.URL.Path[1:])
}

func home(w http.ResponseWriter, r *http.Request) {
  log.Println("home")
  http.ServeFile(w, r, r.URL.Path[1:]+"public/index.html")
}

func pages(w http.ResponseWriter, r *http.Request) {
  page := r.URL.Query().Get(":page")
  page = strings.Split(page, ".")[0]
  log.Println("pages %v", page)
  http.ServeFile(w, r, "public/"+page+".html")
}