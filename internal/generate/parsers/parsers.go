package parsers

import (
	"github.com/gosimple/slug"
	"sync"
)

func ParseBundle(raw RawDataBundle) (b Bundle) {
	ch := sync.WaitGroup{}
	ch.Add(4)

	go ParseInterns(&b, &ch, raw.GetInterns())
	go ParseJobs(&b, &ch, raw.GetJobs())
	go ParsePeople(&b, &ch, raw.GetPeople())
	go ParseMentorships(&b, &ch, raw.GetMentorships())

	ch.Wait()
	return
}

func ParseInterns(b *Bundle, wg *sync.WaitGroup, rawInterns [][]string) {
	for _, intern := range rawInterns {
		b.Interns = append(b.Interns, Intern{
			Timestamp: intern[0],
			FmnUrl:    intern[1],
			Email:     intern[2],
		})
	}
	wg.Done()
}

func ParseJobs(b *Bundle, wg *sync.WaitGroup, rawJobs [][]string) {
	for _, job := range rawJobs {
		b.Jobs = append(b.Jobs, parseJob(job))
	}
	wg.Done()
}

func parseJob(fields []string) Job {
	j := Job{
		Timestamp:   fields[0],
		CompanyName: fields[1],
		Position:    fields[2],
		Address:     fields[3],
		Description: fields[4],
		Location:    fields[5],
		CompanyLogo: fields[6],
		Labels:      fields[7],
	}
	if len(fields) > 8 {
		j.IsRemoteOption = fields[8]
	}
	return j
}

func ParseMentorships(b *Bundle, wg *sync.WaitGroup, rawMentorships [][]string) {
	for _, mentorship := range rawMentorships {
		b.Mentorships = append(b.Mentorships, Mentorship{
			Timestamp:  mentorship[0],
			FmnUrl:     mentorship[1],
			Mentorship: mentorship[2],
			Goal:       mentorship[3],
		})
	}
	wg.Done()
}

func ParsePeople(b *Bundle, wg *sync.WaitGroup, rawPeople [][]string) {
	for _, person := range rawPeople {
		b.People = append(b.People, parsePerson(person))
	}
	wg.Done()
}

func parsePerson(fields []string) Person {
	p := Person{
		RegisteredAt:  fields[0],
		Name:          fields[1],
		TwitterHandle: fields[2],
		GitHub:        fields[3],
		LinkedIn:      fields[4],
		Interests:     fields[5],
		Goals:         fields[6],
		Mentor:        fields[7],
	}

	// Sorry for shitty google spreadsheet api, I'll be fix when I join google as developer.
	if len(fields) > 8 {
		p.Stackoverflow = fields[8]
	}

	p.Slug = slug.MakeLang(p.Name, "tr")
	return p
}
