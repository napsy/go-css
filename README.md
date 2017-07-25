# go-css-parser

This parser understands simple CSS and comes with a basic CSS syntax checker.


```
go get github.com/napsy/go-css-parser
```

Example usage:

```go

import css "github.com/napsy/go-css-parser"

ex1 := `rule {
	style1: value1;
	style2: value2;
}`

css, err := css.Unmarshal(strings.NewReader(ex1))
if err != nil {
	panic(err)
}

fmt.Printf("Defined rules:\n")

for k, _ := range css {
	fmt.Printf("- rule %q\n", k)
}
```
