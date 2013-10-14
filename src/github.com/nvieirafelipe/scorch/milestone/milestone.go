package milestone

import (
  "strconv"
  "github.com/nvieirafelipe/go-github/github"
  "github.com/nvieirafelipe/scorch/issue"
)

type MilestoneCollection struct {
  Milestones  []Milestone       `json:"milestones"`
}

type Milestone struct {
  Number      int               `json:"number"`
  Title       string            `json:"title"`
  URL         string            `json:"url"`
  CreatedAt   github.Timestamp  `json:"created_at"`
  DueOn       github.Timestamp  `json:"due_on"`
  Issues      []issue.Issue           `json:"issues"`
}

func MilestonesFromGithub(githubMilestones []github.Milestone, githubClient *github.Client, organization string, repository string) MilestoneCollection {
  milestones := make([]Milestone, 0)
  for _, milestone := range githubMilestones {
    issues := issue.IssuesFromMilestone(githubClient, strconv.Itoa(*milestone.Number), organization, repository)
    milestones = append(milestones, Milestone{Number: *milestone.Number, URL: *milestone.URL, Title: *milestone.Title, CreatedAt: *milestone.CreatedAt, DueOn: *milestone.DueOn, Issues: issues})
  }
  return MilestoneCollection{Milestones: milestones}
}
