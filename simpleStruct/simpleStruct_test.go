package simpleStruct

import "testing"

func TestSick(t *testing.T) {

	tests := []struct {
		name string
		args Suck
		want string
	}{
        {"just one", Suck{"abe fuckzo", 1000}, "abe fuckzo"},
        {"yet another one", Suck{"Kim Jongilue", 30}, "Kim Jongilue"},
        {"sucked one", Suck{"sushitaro", 40000}, "sushitaro"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sick()
            if got.name != tt.want {
				t.Errorf("Sick() = %v, want %v", got, tt.want)
			}
		})
	}
}
