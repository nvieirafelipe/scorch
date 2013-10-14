package milestone

import (
  "github.com/nvieirafelipe/go-github/github"
)

type MilestoneCollection struct {
  Milestones  []Milestone       `json:"milestones"`
}

type Milestone struct {
  Title       string            `json:"title"`
  URL         string            `json:"url"`
  CreatedAt   github.Timestamp  `json:"created_at"`
  DueOn       github.Timestamp  `json:"due_on"`
}

func MilestonesFromGithub(githubMilestones []github.Milestone) MilestoneCollection {
  milestones := make([]Milestone, 0)
  for _, milestone := range githubMilestones {
    milestones = append(milestones, Milestone{URL: *milestone.URL, Title: *milestone.Title, CreatedAt: *milestone.CreatedAt, DueOn: *milestone.DueOn})
  }
  return MilestoneCollection{Milestones: milestones}
}