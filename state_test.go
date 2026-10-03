package main

import "testing"

func TestD20StaysWithinDieRange(t *testing.T) {
	for i := 0; i < 50; i++ {
		roll := d20()
		if roll < 1 || roll > 20 {
			t.Fatalf("d20() returned %d, want a value in [1, 20]", roll)
		}
	}
}

func TestResolveHitRoll(t *testing.T) {
	cases := []struct {
		name          string
		roll          int
		accuracy      int
		weaponMod     int
		defenderDodge int
		wantHit       bool
	}{
		{"high roll with good accuracy hits", 15, 5, 0, 4, true},
		{"low roll with no accuracy misses", 2, 0, 0, 4, false},
		{"weapon accuracy penalty can turn a hit into a miss", 15, 0, -3, 4, false},
		{"exact threshold does not count as a hit", 14, 0, 0, 4, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := resolveHitRoll(c.roll, c.accuracy, c.weaponMod, c.defenderDodge)
			if got != c.wantHit {
				t.Errorf("resolveHitRoll(%d, %d, %d, %d) = %v, want %v", c.roll, c.accuracy, c.weaponMod, c.defenderDodge, got, c.wantHit)
			}
		})
	}
}

func TestMitigate(t *testing.T) {
	got := mitigate(100, 0.5)
	if got != 50 {
		t.Errorf("mitigate(100, 0.5) = %d, want 50", got)
	}
}
