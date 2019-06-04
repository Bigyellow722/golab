package oop

type db interface {
	Connect()
	SetState()
	Getstate()
}