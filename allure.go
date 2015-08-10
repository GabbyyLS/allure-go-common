package allure

import (
    "github.com/GabbyyLS/allure-go-common/beans"
    "time"
    "bytes"
    "errors"
//    "mime"
    "code.google.com/p/go-uuid/uuid"
    "io/ioutil"
    "path/filepath"
    "os"
    "encoding/xml"
)

//
type Allure struct {
	Suites    []*beans.Suite
	TargetDir string
}

//SetOptions()js -> New
func New(suites []*beans.Suite) *Allure {
	return &Allure{Suites: suites, TargetDir: "allure-results"}
}

//getCurrentSuite -> 0
func (a *Allure) GetCurrentSuite() *beans.Suite {
	return a.Suites[0]
}


func (a *Allure) StartSuite(name string,start time.Time) {
    a.Suites = append(a.Suites,beans.NewSuite(name,start))
}


func (a *Allure) EndSuite (end time.Time) {
    suite := a.GetCurrentSuite()
    suite.End(end)
    if(suite.HasTests()) {
        writeSuite(a.TargetDir,suite)
    }
    //remove first/current suite
    a.Suites = a.Suites[1:]
}

var currentState = map[*beans.Suite]*beans.TestCase{}
var currentStep = map[*beans.Suite]*beans.Step{}

func (a *Allure) StartCase (testName string, start time.Time) {
    var (
        test = beans.NewTestCase(testName, start)
        step = beans.NewStep(testName, start)
        suite = a.GetCurrentSuite()
    )

    currentState[suite] = test
    //strange logic((((
    currentStep[suite] = step

    suite.AddTest(test)
}

func (a *Allure) EndCase (status string, err error, end time.Time) {
    suite := a.GetCurrentSuite()
    test, ok := currentState[suite]
    if ok {
        test.End(status,err,end)
        currentState[suite]
    }
}

func (a *Allure) CreateStep(name string, stepFunc func() ) {
    status := `passed`
    a.StartStep(name,time.Now())
    // if test error
    stepFunc()
    //end
    a.EndStep(status,time.Now())
}

func (a *Allure) StartStep (stepName string, start time.Time) {
    var (
        step = beans.NewStep(stepName, start)
        suite = a.GetCurrentSuite()
    )

    step = currentStep[suite]
    step.Parent.AddStep(step)
    currentStep[suite] = step
}

func (a *Allure) EndStep (status string, end time.Time) {
    suite := a.GetCurrentSuite()
    currentStep[suite].End(status, end)
    currentStep[suite] = currentStep[suite].Parent
}

func (a *Allure) AddAttachment (attachmentName, buf bytes.Buffer, typ string) {
    mime,ext := getBufferInfo(buf,typ)
    name,_ := writeBuffer(a.TargetDir,buf,ext)
    currentState[a.GetCurrentSuite()].AddAttachment(beans.NewAttachment(attachmentName,mime,name,buf.Len()))
}

func (a *Allure) PendingCase(testName string, start time.Time) {
    a.StartCase(testName, start)
    a.EndCase("pending", errors.New("Test ignored"), start)
}

//utils
func getBufferInfo(buf bytes.Buffer, typ string) (string,string) {
//    exts,err := mime.ExtensionsByType(typ)
//    if err != nil {
//        mime.ParseMediaType()
//    }
    return "text/plain","txt"
}

func writeBuffer(pathDir string,buf bytes.Buffer,ext string) (string,error) {
    fileName := uuid.New()+`-attachment.`+ext
    err := ioutil.WriteFile(filepath.Join(pathDir,fileName),buf.Bytes(),os.O_CREATE|os.O_WRONLY)
    return fileName,err
}

func writeSuite(pathDir string,suite *beans.Suite) error {
    bytes, err := xml.Marshal(suite)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(filepath.Join(pathDir,uuid.New()+`-testsuite.xml`),bytes,os.O_CREATE|os.O_WRONLY)
}