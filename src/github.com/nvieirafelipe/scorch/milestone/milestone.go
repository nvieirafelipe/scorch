package milestone

import (
  "strconv"
  "time"
  "github.com/nvieirafelipe/go-github/github"
  "github.com/nvieirafelipe/scorch/issue"
)

type MilestoneCollection struct {
  Milestones  []Milestone       `json:"milestones"`
}

type Milestone struct {
  Number          int               `json:"number"`
  Title           string            `json:"title"`
  URL             string            `json:"url"`
  CreatedAt       github.Timestamp  `json:"created_at"`
  DueOn           github.Timestamp  `json:"due_on"`
  WorkLeftVSTime  []int             `json:"work_left_vs_time"`
}

func MilestonesFromGithub(githubMilestones []github.Milestone, githubClient *github.Client, organization string, repository string) MilestoneCollection {
  milestones := make([]Milestone, 0)

  for _, milestone := range githubMilestones {
    if (milestone.DueOn != nil) {
      issues := issue.IssuesFromMilestone(githubClient, strconv.Itoa(*milestone.Number), organization, repository)
      workLeftVsTime := workLeftVsTime(*milestone.CreatedAt, *milestone.DueOn, issues)
      milestones = append(milestones, Milestone{Number: *milestone.Number, URL: *milestone.URL, Title: *milestone.Title, CreatedAt: *milestone.CreatedAt, DueOn: *milestone.DueOn, WorkLeftVSTime: workLeftVsTime})
    }
  }

  return MilestoneCollection{Milestones: milestones}
}

func workLeftVsTime(createdAt github.Timestamp, dueOn github.Timestamp, issues []issue.Issue) []int {
  createdDate := time.Date(createdAt.Year(), createdAt.Month(), createdAt.Day(), 0, 0, 0, 0, time.UTC)
  dueOnDate := time.Date(dueOn.Year(), dueOn.Month(), dueOn.Day(), 0, 0, 0, 0, time.UTC)
  workLeftVSTime := make([]int, 0)
  for currentDate := createdDate;; {
    issuesLeft := issuesLeftAt(issues, currentDate)
    workLeftVSTime = append(workLeftVSTime, issuesLeft)
    currentDate = currentDate.AddDate(0, 0, 1)
    if dueOnDate.Before(currentDate) {
      break;
    }
  }
  return workLeftVSTime
}

func issuesLeftAt(issues []issue.Issue, date time.Time) int {
  issuesLeft := 0
  for _, issue := range issues {
    if (issue.ClosedAt == nil || issueClosedAfter(issue, date)) {
      issuesLeft += 1
    }
  }
  return issuesLeft
}

func issueClosedAfter(issue issue.Issue, date time.Time) bool {
  issueClosedDate := time.Date(issue.ClosedAt.Year(), issue.ClosedAt.Month(), issue.ClosedAt.Day(), 0, 0, 0, 0, time.UTC)
  return issueClosedDate.After(date)
}
