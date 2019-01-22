package simpleStruct

import "testing"

func TestSick(t *testing.T) {

	tests := []struct {
		name string
		args Suck
		wantName string
        wantKilled int
	}{
        {"just one", Suck{"abe fuckzo", 1000}, "ABE FUCKZO", 2000},
        {"yet another one", Suck{"Kim Jongilue", 30}, "KIM JONGILUE", 60},
        {"sucked one", Suck{"sushitaro", 40000}, "SUSHITARO", 80000},
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
