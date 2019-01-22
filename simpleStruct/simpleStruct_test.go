package simpleStruct

import "testing"

func TestSick(t *testing.T) {
	tests := []struct {
		name string
	}{
        {"just one"},
        {"yet another one"},
        {"sucked one"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sick()
		})
	}
}
