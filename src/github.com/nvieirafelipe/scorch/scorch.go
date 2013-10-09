package main

import (
  "net/http"
  "os"
  "log"
  "github.com/bmizerany/pat"
  "github.com/nvieirafelipe/scorch/organization"
)

func main() {
  m := pat.New()
  m.Get("/organizations/:organization_name/repositories", http.HandlerFunc(organization.Repositories))

  http.Handle("/", m)
  log.Println("listening...")

  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}