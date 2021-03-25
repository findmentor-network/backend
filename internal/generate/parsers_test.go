package generate

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestParseInterns(t *testing.T) {
	s, _ := DummyHTTP()
	res, err := GetData(s.URL)
	b := Bundle{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	ParseInterns(&b, &wg, res.GetInterns())
	assert.Equal(t, "adilahmetsargin@gmail.com", b.Interns[0].Email)
	assert.Equal(t, nil, err)
	assert.Equal(t, 30, len(b.Interns))
}

func TestParseJobs(t *testing.T) {
	s, _ := DummyHTTP()
	res, err := GetData(s.URL)
	b := Bundle{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	ParseJobs(&b, &wg, res.GetJobs())
	assert.Equal(t, "Malwation", b.Jobs[0].CompanyName)
	assert.Equal(t, nil, err)
	assert.Equal(t, 7, len(b.Jobs))
}

func TestParseMentorships(t *testing.T) {
	s, _ := DummyHTTP()
	res, err := GetData(s.URL)
	b := Bundle{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	ParseMentorships(&b, &wg, res.GetMentorships())
	assert.Equal(t, "https://github.com/findmentor-network/find-mentor", b.Mentorships[0].Mentorship)
	assert.Equal(t, nil, err)
	assert.Equal(t, 13, len(b.Mentorships))
}

func TestMapPerson(t *testing.T) {
	s, _ := DummyHTTP()
	res, err := GetData(s.URL)
	b := Bundle{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	ParsePeople(&b, &wg, res.GetPeople())
	assert.Equal(t, "Cagatay Cali", b.People[0].Name)
	assert.Equal(t, nil, err)
	assert.Equal(t, 675, len(b.People))
}
