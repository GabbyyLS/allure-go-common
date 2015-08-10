package beans

import (
	"strings"
	"time"
)

//start new test case
func NewTestCase(name string, start time.Time) *TestCase {
	test := new(TestCase)
	test.Name = name

	if !start.IsZero() {
		test.Start = start.Unix()
	} else {
		test.Start = time.Now().Unix()
	}

	return test
}

type TestCase struct {
	Status      string       `xml:"status,attr"`
	Start       int16        `xml:"start,attr"`
	Stop        int16        `xml:"stop,attr"`
	Name        string       `xml:"name"`
	Steps       []*Step       `xml:"steps"`
	Labels      []*Label      `xml:"labels"`
	Attachments []*Attachment `xml:"attachments"`
	Desc        string       `xml:"description"`
	Failure     struct {
		Msg   string `xml:"message"`
		Trace string `xml:"stack-trace"`
	} `xml:"failure,omitempty"`
}

type Label struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func (t *TestCase) SetDescription(desc string) {
	t.Desc = desc
}

func (t *TestCase) AddLabel(label *Label) {
	t.Labels = append(t.Labels, label)
}

func (t *TestCase) AddStep(step *Step) {
	t.Steps = append(t.Steps, step)
}

func (t *TestCase) AddAttachment(attach *Attachment) {
	t.Attachments = append(t.Attachments, attach)
}

func (t *TestCase) End(status string, err error, end time.Time) {
	if !end.IsZero() {
		t.Stop = end.Unix()
	} else {
		t.Stop = time.Now().Unix()
	}
	t.Status = status
	if err != nil {
		msg := strings.Split("\n", err.Error())
		t.Failure.Msg = msg[0]
		t.Failure.Trace = strings.Join(msg[1:], "\n")
	}
}
