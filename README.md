# go-css-parser

[![Build Status](https://travis-ci.org/napsy/go-css-parser.svg?branch=master)](https://travis-ci.org/napsy/go-css-parser)
[![Coverage Status](http://codecov.io/github/napsy/go-css-parser/coverage.svg?branch=master)](http://codecov.io/github/vendor/package?branch=master)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/vendor/package/blob/master/LICENSE.md)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/napsy/go-css-parser)



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

stylesheet, err := css.Unmarshal([]byte(ex1))
if err != nil {
	panic(err)
}

fmt.Printf("Defined rules:\n")

for k, _ := range stylesheet {
	fmt.Printf("- rule %q\n", k)
}
```
