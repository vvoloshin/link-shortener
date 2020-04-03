package crypto

import "testing"

func TestGenerateBase58Str(t *testing.T) {
	s := GenerateBase58Str()
	if len(s) < 8 {
		t.Error("Not enough random literals")
	}
}
