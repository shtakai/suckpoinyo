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

            if reflect.DeepEqual(tt.args, tt.want){
				t.Errorf("name Sick() = %v, want:%v\n", got, tt)
			}
		})
	}
}
