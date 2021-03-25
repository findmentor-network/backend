package generate

const baseURL = "https://findmentor.network/peer/"

type (
	Bundle struct {
		Interns     []Intern
		Jobs        []Job
		People      []Person
		Mentorships []Mentorship
	}

	Intern struct {
		Timestamp string
		FmnUrl    string
		Email     string
	}

	Job struct {
		Timestamp      string
		CompanyName    string
		Position       string
		Address        string
		Description    string
		Location       string
		CompanyLogo    string
		Labels         string
		IsRemoteOption string
	}

	Person struct {
		RegisteredAt  string `json:"registeredAt"`
		Name          string `json:"name"`
		TwitterHandle string `json:"TwitterHandle"`
		GitHub        string `json:"GitHub"`
		LinkedIn      string `json:"LinkedIn"`
		Interests     string `json:"Interests"`
		Goals         string `json:"Goals"`
		Mentor        string `json:"Mentor"`
		Stackoverflow string `json:"Stackoverflow"`

		// processed data
		Slug        string
		Mentorships []Mentorship
	}

	Mentorship struct {
		Timestamp  string `json:"timestamp"`
		FmnUrl     string `json:"fmn_url"`
		Mentorship string `json:"mentorship"`
		Goal       string `json:"goal"`
	}
)

type RawDataBundle interface {
	GetPeople() [][]string
	GetMentorships() [][]string
	GetJobs() [][]string
	GetInterns() [][]string
}

func (r *Bundle) AggregateMentorships() {
	for i := 0; i < len(r.People); i++ {
		r.People[i].Mentorships = r.filterMentorshipsBySlug(r.People[i].Slug)
	}
}

func (r *Bundle) filterMentorshipsBySlug(slug string) []Mentorship {
	var filtered []Mentorship

	for _, mentorship := range r.Mentorships {
		lastPartOfURL := mentorship.FmnUrl[len(baseURL):]
		if slug == lastPartOfURL {
			filtered = append(filtered, mentorship)
		}
	}

	return filtered
}

type GoogleSheetResponse struct {
	ValueRanges []struct {
		Range  string     `json:"range"`
		Values [][]string `json:"values"`
	} `json:"valueRanges"`
}

func (r GoogleSheetResponse) GetPeople() [][]string {
	return r.ValueRanges[0].Values[7:]
}

func (r GoogleSheetResponse) GetMentorships() [][]string {
	return r.ValueRanges[1].Values[2:]
}

func (r GoogleSheetResponse) GetJobs() [][]string {
	return r.ValueRanges[2].Values[2:]
}

func (r GoogleSheetResponse) GetInterns() [][]string {
	return r.ValueRanges[3].Values[2:]
}
