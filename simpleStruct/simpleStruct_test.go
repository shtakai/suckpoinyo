package simpleStruct

import (
	"reflect"
	"testing"
)

func TestSick(t *testing.T) {

	tests := []struct {
		name string
		args Suck
		want Suck
	}{
		{"just one", Suck{"abe fuckzo", 1000}, Suck{"ABE FUCKZO", 2000}},
		{"yet another one", Suck{"Kim Jongilue", 30}, Suck{"KIM JONGILUE", 60}},
		{"sucked one", Suck{"sushitaro", 40000}, Suck{"SUSHITARO", 80000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suck := tt.args
			got := Sick(&suck)

			if reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("name Sick() = %v, want:%v\n", got, tt)
			}
		})
	}
}

func TestSwapSick(t *testing.T) {
	tests := []struct {
		name  string
		arg1 Suck
		arg2 Suck
		want1  Suck
		want2 Suck
	}{
        {"simple swap", Suck{"name1", 1}, Suck{"name2", 2}, Suck{"name2",2}, Suck{"name1",1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := SwapSick(tt.arg1, tt.arg2)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SwapSick() got1 = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("SwapSick() got2 = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}
