package models

type Actor struct {
	ID           int64  `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type Repositore struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Commit struct {
	SHA      string `json:"sha"`
	Author   Author `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	URL      string `json:"url"`
}

type Payload struct {
	RepositoryID int64    `json:"repository_id"`
	PushID       int64    `json:"push_id"`
	Size         int64    `json:"size"`
	DistinctSize int64    `json:"distinct_size"`
	Ref          string   `json:"ref"`
	RefType      string   `json:"ref_type"`
	MasterBranch string   `json:"master_branch"`
	Head         string   `json:"head"`
	Before       string   `json:"before"`
	Commits      []Commit `json:"commits"`
}

type Event struct {
	ID        string     `json:"id"`
	Type      string     `json:"type"`
	Actor     Actor      `json:"actor"`
	Repo      Repositore `json:"repo"`
	Pay       Payload    `json:"payload"`
	Public    bool       `json:"public"`
	CreatedAt string     `json:"created_at"`
}
