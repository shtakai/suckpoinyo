package simpleStruct

import "testing"

func TestSick(t *testing.T) {

	tests := []struct {
		name string
		args Suck
		wantName string
        wantKilled int
	}{
        {"just one", Suck{"abe fuckzo", 1000}, "abe fuckzo", 1000},
        {"yet another one", Suck{"Kim Jongilue", 30}, "Kim Jongilue", 30},
        {"sucked one", Suck{"sushitaro", 40000}, "sushitaro", 40000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			suck := Suck{
				tt.args.name,
				tt.args.killed,
			}
			got := Sick(suck)

            if got.name != tt.wantName || got.killed != tt.wantKilled {
				t.Errorf("name Sick() = %v, want name:%v killed: %v\n", got, tt.wantName, tt.wantKilled)
			}
		})
	}
}
