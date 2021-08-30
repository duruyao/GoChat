package key

import "testing"

func TestCreateKeys(t *testing.T) {
	if err := CreateKeys(); err != nil {
		t.Fatal(err)
	}
}
