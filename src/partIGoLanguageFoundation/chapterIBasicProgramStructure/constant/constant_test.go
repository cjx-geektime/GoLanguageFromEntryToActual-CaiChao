package constant

import "testing"

const (
	Monday = iota + 1
	Tuesday
	wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Open = 1 << iota
	Close
	Pending
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday, wednesday, Thursday, Friday, Saturday, Sunday)
	t.Log(Open, Close, Pending)
}
