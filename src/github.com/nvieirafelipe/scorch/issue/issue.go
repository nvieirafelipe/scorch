package issue

import (
  "log"
  "time"
  "github.com/nvieirafelipe/go-github/github"
)

type Issue struct {
  Number     int               `json:"number"`
  State      string            `json:"state"`
  Title      string            `json:"title"`
  CreatedAt  *time.Time         `json:"created_at"`
  ClosedAt   *time.Time         `json:"closed_at"`
}

func IssuesFromMilestone(githubClient *github.Client, milestone string, organization string, repository string) []Issue {
  githubOpenIssues := issuesByRepo(githubClient, organization, repository, milestone,"open")
  githubClosedIssues := issuesByRepo(githubClient, organization, repository, milestone, "closed")
  return append(issuesFromGithub(githubOpenIssues), issuesFromGithub(githubClosedIssues)...)
}

func issuesByRepo(githubClient *github.Client, organization string, repository string, milestone string, state string) []github.Issue {
  opts := &github.IssueListByRepoOptions{Milestone: milestone, State: state}
  githubIssues, _, err := githubClient.Issues.ListByRepo(organization, repository, opts)
  if err != nil {
    log.Println("error: %v\n\n", err)
  }
  return githubIssues
}

func issuesFromGithub(githubIssues []github.Issue) []Issue {
  issues := make([]Issue, 0)
  for _, issue := range githubIssues {
    issues = append(issues, Issue{Number: *issue.Number, State: *issue.State, Title: *issue.Title, CreatedAt: issue.CreatedAt, ClosedAt: issue.ClosedAt})
  }
  return issues
}
