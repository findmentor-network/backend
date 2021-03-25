package generate

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
