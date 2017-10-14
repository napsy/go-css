# go-css

[![Build Status](https://travis-ci.org/napsy/go-css.svg?branch=master)](https://travis-ci.org/napsy/go-css)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/vendor/package/blob/master/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/napsy/go-css)



This parser understands simple CSS and comes with a basic CSS syntax checker.


```
go get github.com/napsy/go-css
```

Example usage:

```go

import "github.com/napsy/go-css"

ex1 := `rule {
	style1: value1;
	style2: value2;
}`

stylesheet, err := css.Unmarshal([]byte(ex1))
if err != nil {
	panic(err)
}

fmt.Printf("Defined rules:\n")

for k, _ := range stylesheet {
	fmt.Printf("- rule %q\n", k)
}
```

You can get a CSS verifiable property by calling ``CSSStyle``:

```go
style, err := css.CSSStyle("background-color", styleSheet["body"])
if err != nil {
	fmt.Printf("Error checking body background color: %v\n", err)
} else {
	fmt.Printf("Body background color is %v", style)
}
```

Most of the CSS properties are currently not implemented, but you can always write your own handler by writing a ``StyleHandler`` function and adding it to the ``StylesTable`` map.
