package parsers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInterns(t *testing.T) {
	s, _ := DummyHTTP()
	res, err := getData(s.URL)
	interns := ParseInterns(res.getInterns())
	assert.Equal(t, "adilahmetsargin@gmail.com", interns[0].Email)
	assert.Equal(t, nil, err)
	assert.Equal(t, 30, len(interns))
}

func TestParseJobs(t *testing.T) {
	s, _ := DummyHTTP()
	res, err := getData(s.URL)
	jobs := ParseJobs(res.getJobs())
	assert.Equal(t, "Malwation", jobs[0].CompanyName)
	assert.Equal(t, nil, err)
	assert.Equal(t, 7, len(jobs))
}

func TestParseMentorships(t *testing.T) {
	s, _ := DummyHTTP()
	res, err := getData(s.URL)
	mentorships, err := ParseMentorships(res.getMentorships())
	assert.Equal(t, "https://github.com/findmentor-network/find-mentor", mentorships[0].Mentorship)
	assert.Equal(t, nil, err)
	assert.Equal(t, 13, len(mentorships))
}

func TestMapPerson(t *testing.T) {
	s, _ := DummyHTTP()
	res, err := getData(s.URL)
	people := ParsePeople(res.getPeople())
	assert.Equal(t, "Cagatay Cali", people[0].Name)
	assert.Equal(t, nil, err)
	assert.Equal(t, 675, len(people))
}
