package repository

import (
  "github.com/google/go-github/github"
)

type RepositoryCollection struct {
  Repositories   []Repository  `json:"repositories"`
}

type Repository struct {
    Name         string  `json:"name"`
    Description  string  `json:"description"`
}

func RepositoriesFromGithub(githubRepos []github.Repository) RepositoryCollection {
  repos := make([]Repository, 0)
  for _, repo := range githubRepos {
    repos = append(repos, Repository{Name: *repo.Name, Description: *repo.Description})
  }
  return RepositoryCollection{Repositories: repos}
}