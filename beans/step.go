package beans

import (
	"time"
)

func NewStep(name string, start time.Time) *Step {
	test := new(Step)
	test.Name = name

	if !start.IsZero() {
		test.Start = start.Unix()
	} else {
		test.Start = time.Now().Unix()
	}

	return test
}

type Step struct {
	Parent      *Step         `xml:"-"`

	Status      string        `xml:"status,attr"`
	Start       int16         `xml:"start,attr"`
	Stop        int16         `xml:"stop,attr"`
	Name        string        `xml:"name"`
	Steps       []*Step       `xml:"steps"`
	Attachments []*Attachment `xml:"attachments"`
}

func (s *Step) End(status string, end time.Time) {
	if !end.IsZero() {
		s.Stop = end.Unix()
	} else {
		s.Stop = time.Now().Unix()
	}
	s.Status = status
}

func (s *Step) AddStep(step *Step) {
	if step != nil {
		s.Steps = append(s.Steps,step)
	}
}
