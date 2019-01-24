package oop

import "fmt"

//1.封装
type Person struct {
	Name    string
	Sex     int32
	Address string
}

func (p Person) Publish() {
	fmt.Println("method : Publish")
}

func NewPerson(name string, sex int32, address string) Person {
	return Person{
		Name:    name,
		Sex:     sex,
		Address: address,
	}
}

func (p Person) Set(name string, sex int32, address string) {
	p.Name = name
	p.Sex = sex
	p.Address = address
}

func (p Person) Get() {
	//fmt.Printf("%v who is %v,lives in %v", p.Name, p.Sex, p.Address)
	fmt.Println("Person")
}

//2.继承
type ChinaPerson struct {
	Person
	Area string
}

func NewChinaPerson(name string, sex int32, address string, area string) ChinaPerson {
	p := NewPerson(name, sex, area)
	return ChinaPerson{
		Person: p,
		Area:   area,
	}

}

func (cp ChinaPerson) Set(p Person, area string) {
	cp.Person = p
	cp.Area = area
}

func (cp ChinaPerson) Get() {
	//fmt.Printf("%v who is %v,lives in %v %v", cp.Name, cp.Sex, cp.Address,cp.area)
	fmt.Println("ChinaPerson")
}

func testjicheng(cp ChinaPerson) {
	cp.Publish()
}

//
