package css

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSimple(t *testing.T) {
	ex1 := `rule {
		style1: value1;
		style2: value2;
}`

	cases := []struct {
		name     string
		CSS      string
		expected map[Rule]map[string]string
	}{
		{"Single rule (simple)", ex1, map[Rule]map[string]string{
			"rule": {
				"style1": "value1",
				"style2": "value2",
			},
		}},
	}

	for _, tt := range cases {
		t.Run("GoodCSS", func(t *testing.T) {
			css, err := Unmarshal([]byte(tt.CSS))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.expected, css, "Expected CSS to be equal")
		})
	}
}

func TestParseError(t *testing.T) {
	ex1 := `{
		style1: value1;
}`

	ex2 := `rule {
		style1: value1;
		style2:;
}`

	ex3 := `rule {
	   		style1: value1
	   		style2: value2;
}`

	ex4 := `}
rule {
		style1: value1;
		style2:;
}`
	_ = ex3

	cases := []struct {
		name string
		CSS  string
	}{
		{"Missing rule", ex1},
		{"Missing style", ex2},
		// TODO: this hsould not crash
		//{"Statement Missing Semicolon", ex3},
		{"BlockEndsWithoutBeginning", ex4},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Unmarshal([]byte(tt.CSS)); err == nil {
				t.Fatal("Should return error!")
			}
		})
	}
}

func TestParseHarder(t *testing.T) {
	t.Run("MultiRules", func(t *testing.T) {
		ex1 := `rule1 {
		style1: value1;
		style2: value2;
}
rule2 {
	style3: value3;
}`
		css, err := Unmarshal([]byte(ex1))
		if err != nil {
			t.Fatal(err)
		}
		if len(css) != 2 {
			t.Fatalf("expected 2 rules, got %d", len(css))
		}

		if _, ok := css["rule1"]; !ok {
			t.Fatal("missing rule 'rule1'")
		}
		if _, ok := css["rule2"]; !ok {
			t.Fatal("missing rule 'rule2'")
		}

		if len(css["rule1"]) != 2 {
			t.Fatalf("expected 2 styles for rule 'rule1', got %d", len(css["rule1"]))
		}

		if len(css["rule2"]) != 1 {
			t.Fatalf("expected 1 style for rule 'rule2', got %d", len(css["rule2"]))
		}
	})
	t.Run("PropertyWithSpace", func(t *testing.T) {
		ex1 := `body {
		font-family: 'Zil', serif;
}`
		css, err := Unmarshal([]byte(ex1))
		if err != nil {
			t.Fatal(err)
		}

		if css["body"]["font-family"] != "'Zil', serif" {
			t.Fatalf("invalid rule 'font-family', got %q", css["body"]["font-family"])
		}
	})
	t.Run("MergedRules", func(t *testing.T) {
		ex1 := `rule1 {
		style1: value1;
		style2: value2;
}
rule1 {
	style1: value3;
}`
		css, err := Unmarshal([]byte(ex1))
		if err != nil {
			t.Fatal(err)
		}
		if len(css) != 1 {
			t.Fatalf("there should be only one rule, got %d", len(css))
		}
		if len(css["rule1"]) != 2 {
			t.Fatalf("there should be two styles for rule 'rule1', got %d", len(css["rule1"]))
		}
		if css["rule1"]["style1"] != "value3" {
			t.Fatalf("value of 'style1' should be 'value3' but got '%v'", css["rule1"]["style1"])
		}
	})
	t.Run("RealWorldCSS", func(t *testing.T) {
		ex1 := `body {
    background-image: url("gradient_bg.png");
    background-repeat: repeat-x;
}`
		_, err := Unmarshal([]byte(ex1))
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestParseSelectors(t *testing.T) {
	ex1 := `.rule {
		style1: value1;
		style2: value2;
}
#rule1 sad asd {
	style3: value3;
	style4: value4;
}`

	css, err := Unmarshal([]byte(ex1))
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := css[".rule"]; !ok {
		t.Fatal("Missing '.rule' rule")
	}
	if _, ok := css["#rule1"]; !ok {
		t.Fatal("Missing '.rule' rule")
	}
}

func TestParseSelectorGroup(t *testing.T) {
	ex1 := `.rule1 #rule2 rule3 {
		style1: value1;
		style2: value2;
}`

	css, err := Unmarshal([]byte(ex1))
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := css[".rule1"]; !ok {
		t.Fatal("Missing '.rule1' rule")
	}
	if _, ok := css["#rule2"]; !ok {
		t.Fatal("Missing '#rule2' rule")
	}
	if _, ok := css["rule3"]; !ok {
		t.Fatal("Missing '.rule3' rule")
	}
}

func BenchmarkParser(b *testing.B) {
	ex1 := ""
	for i := 0; i < 100; i++ {
		ex1 += fmt.Sprintf(`block%d {
	style%d: value%d;
}`, i, i, i)
	}
	styleSheet := []byte(ex1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Unmarshal(styleSheet)
		if err != nil {
			b.Fatal(err)
		}
	}
}
