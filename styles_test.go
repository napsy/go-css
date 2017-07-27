package css

import "testing"

func TestStyles(t *testing.T) {
	_, err := CSSStyle("background-color", "bla")
	if err == nil {
		t.Fatal("should report invalid color")
	}
	_, err = CSSStyle("background-color", "#aabbccdd")
	if err != nil {
		t.Fatalf("should be valid color, but got %v", err)
	}
}
