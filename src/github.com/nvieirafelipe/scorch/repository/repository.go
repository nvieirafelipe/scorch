package repository

import (
  "net/http"
  "log"
  "os"
  "io"
  "encoding/json"
  "code.google.com/p/goauth2/oauth"
  "github.com/nvieirafelipe/go-github/github"
  "github.com/nvieirafelipe/scorch/milestone"
)

type RepositoryCollection struct {
  Repositories  []Repository      `json:"repositories"`
}

type Repository struct {
  Name          string            `json:"name"`
  Description   string            `json:"description"`
}

func RepositoriesFromGithub(githubRepos []github.Repository) RepositoryCollection {
  repos := make([]Repository, 0)
  for _, repo := range githubRepos {
    repos = append(repos, Repository{Name: *repo.Name, Description: *repo.Description})
  }
  return RepositoryCollection{Repositories: repos}
}

func WorkLeftVSTime(w http.ResponseWriter, req *http.Request) {
  repository := req.URL.Query().Get(":repository_name")
  organization := req.URL.Query().Get("organization")

  githubMilestones, _, err := githubClient().Milestones.List(organization, repository, nil)
  milestones := milestone.MilestonesFromGithub(githubMilestones, githubClient(), organization, repository)
  json, err := json.Marshal(milestones)

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
