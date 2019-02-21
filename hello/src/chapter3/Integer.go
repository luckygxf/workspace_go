package chapter3

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}
