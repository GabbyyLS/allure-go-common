package fake_test

import (
	"testing"
	"github.com/GabbyyLS/allure-go-common/fake"
)

func TestFake(t *testing.T) {
	f := new(fake.Fake)
	if len(f.Name) != 0 {
		t.Errorf("Expected empty name, but - %s ",f.Name)
	}

	f.ChangeName("NewName")
	if f.Name != "NewName" {
		t.Errorf("Expected after change name 'NewName', but - %s ",f.Name)
	}
}
