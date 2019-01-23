package oop

import "fmt"


//1.封装
type Person struct{
	Name string
	Sex int32
	Address string
}

func NewPerson

func (p Person)Publish(){
	fmt.Println("method : Publish")
}

func (p Person)Set(name string,sex int32,address string){
	p.Name = name
	p.Sex = sex
	p.Address = address
}

func (p Person)Get(){
	fmt.Printf("%v who is %v,lives in %v",p.Name,p.Sex,p.Address)
}

//2.继承
type ChinaPerson struct {
	Person
	area string
}

