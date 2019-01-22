package simpleInt

import "testing"

func TestSimpleInt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simply not delicious", args{1}, 1},
		{"more not delicious", args{2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleInt(tt.args.i); got != tt.want {
				t.Errorf("SimpleInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
