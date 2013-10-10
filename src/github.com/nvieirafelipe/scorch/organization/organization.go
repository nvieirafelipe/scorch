package organization

import (
  "net/http"
  "log"
  "os"
  "io"
  "code.google.com/p/goauth2/oauth"
  "github.com/google/go-github/github"
  "github.com/nvieirafelipe/scorch/repository"
)

func Repositories(w http.ResponseWriter, req *http.Request) {
  organization := req.URL.Query().Get(":organization_name")

  githubRepos, _, err := githubClient().Repositories.ListByOrg(organization, nil)
  repositories := repository.RepositoriesFromGithub(githubRepos)

  if err != nil {
    log.Println("error: %v\n\n", err)
  } else {
    io.WriteString(w, "repositories for "+req.URL.Query().Get(":organization_name")+"\n"+github.Stringify(repos))
  }
}

func githubClient() *github.Client {
  t := &oauth.Transport{
    Token: &oauth.Token{AccessToken: os.Getenv("GITHUB_KEY")},
  }

  return github.NewClient(t.Client())
}
