package main

import (
  "net/http"
  "os"
  "log"
  "github.com/bmizerany/pat"
  "github.com/nvieirafelipe/scorch/organization"
  "github.com/nvieirafelipe/scorch/repository"
)

func main() {
  routes := pat.New()
  routes.Get("/organizations/:organization_name/repositories", http.HandlerFunc(organization.Repositories))
  routes.Get("/repositories/:repository_name/workleft-vs-time", http.HandlerFunc(repository.WorkLeftVSTime))

  http.Handle("/", routes)
  log.Println("listening...")

  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
