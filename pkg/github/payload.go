package github

type Payload struct {
	Action string

	Repository struct {
		Name     string
		FullName string `json:"full_name"`
		HtmlUrl  string `json:"html_url"`
	}

	Url    string
	SShUrl string `json:"ssh_url"`

	// Push
	Compare string
	Commits []Commit

	// Issue
	Issue struct {
		Id      string
		HtmlUrl string `json:"html_url"`
		Title   string
		Body    string
	}
	Labels    []Label
	State     string
	Assignee  User
	Assignees []User

	// Pull request
	PullRequest struct {
		Number             int
		State              string
		Title              string
		Labels             []Label
		HtmlUrl            string `json:"html_url"`
		IssueUrl           string `json:"issue_url"`
		Assignee           User
		Assignees          []User
		RequestedReviewers []User `json:"requested_reviewers"`
	} `json:"pull_request"`
}

type User struct {
	Login     string
	Type      string
	AvatarUrl string `json:"avatar_url"`
	HtmlUrl   string `json:"html_url"`
}

type Commit struct {
	Message   string
	Timestamp string
	Url       string
	Author    struct {
		Name     string
		Email    string
		Username string
	}
	Added    []string
	Removed  []string
	Modified []string
}

type Label struct {
	Name  string
	Color string
}
