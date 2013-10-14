package organization

import (
  "net/http"
  "log"
  "os"
  "io"
  "encoding/json"
  "code.google.com/p/goauth2/oauth"
  "github.com/nvieirafelipe/go-github/github"
  "github.com/nvieirafelipe/scorch/repository"
)

func Repositories(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  organization := req.URL.Query().Get(":organization_name")

  githubRepos, _, err := githubClient().Repositories.ListByOrg(organization, nil)
  repositories := repository.RepositoriesFromGithub(githubRepos)
  json, err := json.Marshal(repositories)

  if err != nil {
    log.Println("error: %v\n\n", err)
  } else {
    io.WriteString(w, string(json))
  }
}

func githubClient() *github.Client {
  t := &oauth.Transport{
    Token: &oauth.Token{AccessToken: os.Getenv("GITHUB_KEY")},
  }

  return github.NewClient(t.Client())
}
