package css

import (
	"strings"
	"testing"
)

func TestParseSimple(t *testing.T) {
	ex1 := `rule {
		style1: value1;
		style2: value2;
}`
	ex2 := `{
		style1: value1;
}`
	ex3 := `rule {
		style1: value1
		style2: value2;
}`

	t.Run("GoodCSS", func(t *testing.T) {
		css, err := parse(buildList(strings.NewReader(ex1)))
		if err != nil {
			t.Fatal(err)
		}
		rule, ok := css["rule"]
		if !ok {
			t.Fatal("rule 'rule' doesn't exist")
		}

		if value, ok := rule["style1"]; !ok {
			t.Fatal("syle 'style1' doesn't exist")
		} else if value != "value1" {
			t.Fatalf("incorrect value for 'style1', got '%v', expected 'value1'", value)
		}

		if value, ok := rule["style2"]; !ok {
			t.Fatal("style 'style2' doesn't exist")
		} else if value != "value2" {
			t.Fatalf("incorrect value for 'style2', got '%v', expected 'value2'", value)
		}
	})

	t.Run("MissingRule", func(t *testing.T) {
		if _, err := parse(buildList(strings.NewReader(ex2))); err == nil {
			t.Fatal("should error out")
		}
	})

	t.Run("StatementMissingSemicolon", func(t *testing.T) {
		if _, err := parse(buildList(strings.NewReader(ex3))); err == nil {
			t.Fatal("should error out")
		}
	})

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
		css, err := parse(buildList(strings.NewReader(ex1)))
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
	})
}
