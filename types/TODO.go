package types

import "fmt"

type TODO struct {
	ID int
	Text string
	Time int64
}

func (t TODO) ToString() string {
	return fmt.Sprintf("Text: %s, Time: %d", t.Text, t.Time)
}
