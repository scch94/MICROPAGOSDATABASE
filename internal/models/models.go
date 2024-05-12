package models

type Request interface {
	ToString() string
	Name() string
}
