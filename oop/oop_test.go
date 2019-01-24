package oop

import "testing"

func TestChinaPerson_Get(t *testing.T) {
	person := NewPerson("qqq", 1, "wad")
	person.Get()
	cperson := NewChinaPerson("wqy", 1, "qwewq", "www")
	cperson.Get()
	cperson.Person.Get()
}
