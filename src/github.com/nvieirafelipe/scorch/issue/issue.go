package issue

import (
  "log"
  "github.com/nvieirafelipe/go-github/github"
)

type IssueCollection struct {
  Issues    []*Issue       `json:"issues"`
}
func (issueCollection IssueCollection) Len() int {
   return len(issueCollection.Issues)
}

type Issue struct {
  Number     *int                  `json:"number"`
  State      *string               `json:"state"`
  Title      *string               `json:"title"`
  Milestone  *Milestone            `json:"title"`
}

func IssuesFromGithub(githubIssues []github.Issue) IssueCollection {
  issues := make([]Issue, 0)
  for _, issue := range githubIssues {
    milestone := Milestone{URL: issue.Milestone.URL, Title: issue.Milestone.Title, CreatedAt: issue.Milestone.CreatedAt, DueOn: issue.Milestone.DueOn}
    log.Println("milestone: %v", github.Stringify(milestone))
    issues = append(issues, &Issue{Number: issue.Number,
                                  State: issue.State,
                                  Title: issue.Title,
                                  Milestone: &milestone})
  }
  return IssueCollection{Issues: issues}
}
