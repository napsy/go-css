# go-css-parser

This parser understands simple CSS and comes with a basic CSS syntax checker.

Example usage:

```go
ex1 := `rule {
	style1: value1;
	style2: value2;
}`

css, err := Unmarshal(strings.NewReader(ex1))
if err != nil {
	panic(err)
}

fmt.Printf("Defined rules:\n")

for k, _ := range css {
	fmt.Printf("- rule %q\n", k)
}
```
