package model

type Item interface {
	GetID() string
	GetType() string
	GetName() string
}