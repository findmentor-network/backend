package person

type Person struct {
	RegisteredAt     string         `json:"registered_at"`
	Name             string         `json:"name"`
	TwitterHandle    string         `json:"twitter_handle"`
	Github           string         `json:"github"`
	Linkedin         string         `json:"linkedin"`
	Interests        string         `json:"interests"`
	Goals            string         `json:"goals"`
	Mentor           string         `json:"mentor"`
	Slug             string         `json:"slug"`
	Avatar           string         `json:"avatar"`
	DisplayInterests string         `json:"displayInterests"`
	IsHireable       bool           `json:"isHireable"`
	Mentorships      []string       `json:"mentorships"`
	Contributions    []Contribution `json:"contributors"`
}
type Contribution struct {
	string        `json:""`
	Mentor        string `json:"mentor"`
	ProjectAdress string `json:"project_adress"`
	Goal          string `json:"goal"`
	Slug          string `json:"slug"`
	Contributors  []struct {
		Username      string `json:"username"`
		GithubAddress string `json:"github_address"`
		Avatar        string `json:"avatar"`
		FmnURL        string `json:"fmn_url"`
	} `json:"contributions"`
}
