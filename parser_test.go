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

	ex2 := `rule1 {
		style1: value1;
		style2: value2;
}
rule2 {
	style3: value3;
}`

	ex3 := `body {
		font-family: 'Zil', serif;
}`

	ex4 := `rule1 {
		style1: value1;
		style2: value2;
}
rule1 {
	style1: value3;
}`

	ex5 := `body {
    background-image: url("gradient_bg.png");
    background-repeat: repeat-x;
}`

	ex6 := `rule1 descendant {
	style1:value1;
}`

	ex7 := `rule1 {
	/* this is a comment */
	style: value;
}`

	ex8 := `.rule1 #rule2 {
	style: value;
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
		{"Multiple rules", ex2, map[Rule]map[string]string{
			"rule1": {
				"style1": "value1",
				"style2": "value2",
			},
			"rule2": {
				"style3": "value3",
			},
		}},
		{"Property with spaces", ex3, map[Rule]map[string]string{
			"body": {
				"font-family": "'Zil', serif",
			},
		}},
		{"Merged rules", ex4, map[Rule]map[string]string{
			"rule1": {
				"style1": "value3",
				"style2": "value2",
			},
		}},
		{"Real world css", ex5, map[Rule]map[string]string{
			"body": {
				"background-image":  "url(\"gradient_bg.png\")",
				"background-repeat": "repeat-x",
			},
		}},
		{"Descendant selector", ex6, map[Rule]map[string]string{
			"rule1 descendant": {
				"style1": "value1",
			},
		}},
		{"Comment in rule", ex7, map[Rule]map[string]string{
			"rule1": {
				"style": "value",
			},
		}},
		{"Selector with descentant ID and Class", ex8, map[Rule]map[string]string{
			".rule1 #rule2": {
				"style": "value",
			},
		}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
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

	ex5 := `
body {
	style1:value1;
	*/
}`

	cases := []struct {
		name string
		CSS  string
	}{
		{"Missing rule", ex1},
		{"Missing style", ex2},
		{"Statement Missing Semicolon", ex3},
		{"BlockEndsWithoutBeginning", ex4},
		{"Unexpected end of comment", ex5},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Unmarshal([]byte(tt.CSS)); err == nil {
				t.Fatal("Should return error!")
			}
		})
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
	fmt.Println(css)

	if _, ok := css[".rule1"]; !ok {
		t.Fatal("Missing '.rule1' rule")
	}
	if _, ok := css["#rule2"]; !ok {
		t.Fatal("Missing '#rule2' rule")
	}
	/*
		if _, ok := css["rule3"]; !ok {
			t.Fatal("Missing '.rule3' rule")
		}
	*/
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
