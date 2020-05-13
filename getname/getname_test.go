package getname

import (
	"testing"
)

func TestName(t *testing.T) {
	n := Name()
	if n != "Hanii" {
		t.Fatalf(`invalid name. want "Hanii" got %v`, n)
	}
}
