package person

import (
	"testing"
)

func TestSetPersonName(t *testing.T) {
	var p = new(Person)
	p.NewPerson("Pete", "Developer")
	p.SetPersonName("Peter")
	var actualResult = p.GetName()
	if actualResult != "Peter" {
		t.Error("Expected result: Peter, actual result: ", actualResult)
	}
}
