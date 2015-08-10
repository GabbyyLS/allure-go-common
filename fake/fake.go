package fake

type Fake struct {
	Name string
}

func (f *Fake) ChangeName(newName string) {
	if newName != "fakename" {
		f.Name = newName
	}

}
