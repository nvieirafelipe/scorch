package repository

import (
  "net/http"
  "log"
  "os"
  "io"
  "encoding/json"
  "code.google.com/p/goauth2/oauth"
  "github.com/nvieirafelipe/go-github/github"
)

type RepositoryCollection struct {
  Repositories  []Repository  `json:"repositories"`
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
  milestones := MilestonesFromGithub(githubMilestones)
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

type MilestoneCollection struct {
  Milestones    []Milestone       `json:"milestones"`
}

type Milestone struct {
  Title         string            `json:"title"`
  URL           string            `json:"url"`
  CreatedAt     github.Timestamp  `json:"created_at"`
  DueOn         github.Timestamp  `json:"due_on"`
}

func MilestonesFromGithub(githubMilestones []github.Milestone) MilestoneCollection {
  milestones := make([]Milestone, 0)
  for _, milestone := range githubMilestones {
    milestones = append(milestones, Milestone{URL: *milestone.URL, Title: *milestone.Title, CreatedAt: *milestone.CreatedAt, DueOn: *milestone.DueOn})
  }
  return MilestoneCollection{Milestones: milestones}
}