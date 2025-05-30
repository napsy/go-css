package css

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"io"
	"strings"
	"text/scanner"
)

var InvalidCSSError = errors.New("invalid CSS")

//go:generate stringer -type=tokenType

type tokenType int

const (
	tokenFirstToken tokenType = iota - 1
	tokenBlockStart
	tokenBlockEnd
	tokenRuleName
	tokenValue
	tokenSelector
	tokenStyleSeparator
	tokenStatementEnd
	tokenCommentStart
	tokenCommentEnd
)

type tokenEntry struct {
	value string
	pos   scanner.Position
}

func newTokenType(typ string) tokenType {
	types := map[string]tokenType{
		"{":  tokenBlockStart,
		"}":  tokenBlockEnd,
		":":  tokenStyleSeparator,
		";":  tokenStatementEnd,
		".":  tokenSelector,
		"#":  tokenSelector,
		"/*": tokenCommentStart,
		"*/": tokenCommentEnd,
	}

	result, ok := types[typ]
	if ok {
		return result
	}

	return tokenValue
}

func (e tokenEntry) typ() tokenType {
	return newTokenType(e.value)
}

type tokenizer struct {
	s *scanner.Scanner
}

func newTokenizer(r io.Reader) *tokenizer {
	s := &scanner.Scanner{}
	s.Init(r)

	return &tokenizer{
		s: s,
	}
}

func (t *tokenizer) next() (tokenEntry, error) {
	token := t.s.Scan()
	if token == scanner.EOF {
		return tokenEntry{}, errors.New("EOF")
	}
	value := t.s.TokenText()
	pos := t.s.Pos()
	if newTokenType(value) == tokenStyleSeparator {
		t.s.IsIdentRune = func(ch rune, i int) bool { // property value can contain spaces
			if ch == -1 || ch == '\n' || ch == '\r' || ch == '\t' || ch == ':' || ch == ';' {
				return false
			}
			return true
		}
	} else {
		t.s.IsIdentRune = func(ch rune, i int) bool { // other tokens can't contain spaces
			if ch == -1 || ch == '.' || ch == '#' || ch == '\n' || ch == '\r' || ch == ' ' || ch == '\t' || ch == ':' || ch == ';' {
				return false
			}
			return true
		}
	}
	return tokenEntry{
		value: value,
		pos:   pos,
	}, nil
}

// Rule is a string type that represents a CSS rule.
type Rule string

// Type returns the rule type, which can be a class, id or a tag.
func (rule Rule) Type() string {
	if strings.HasPrefix(string(rule), ".") {
		return "class"
	}
	if strings.HasPrefix(string(rule), "#") {
		return "id"
	}
	return "tag"
}

func buildList(r io.Reader) *list.List {
	l := list.New()
	t := newTokenizer(r)
	for {
		token, err := t.next()
		if err != nil {
			break
		}
		l.PushBack(token)
	}
	return l
}

// TODO: rules can be comma separated
func parse(l *list.List) (map[Rule]map[string]string, error) {
	var (
		// Information about the current block that is parsed.
		rule     []string
		style    string
		value    string
		selector string

		isBlock   bool
		isValue   bool
		isComment bool

		// Parsed styles.
		css    = make(map[Rule]map[string]string)
		styles = make(map[string]string)

		// Previous token for the state machine.
		prevToken = tokenType(tokenFirstToken)
	)

	for e := l.Front(); e != nil; e = l.Front() {
		token := e.Value.(tokenEntry)
		l.Remove(e)

		// handle comment - we continue after this because we don't want to override prevToken
		switch token.typ() {
		case tokenCommentStart:
			isComment = true
			continue
		case tokenCommentEnd:
			// handle standalone endComment token
			if !isComment {
				return css, fmt.Errorf("line %d: unexpected end of comment: %w", token.pos.Line, InvalidCSSError)
			}

			isComment = false
			continue
		}

		if isComment { // skip everything regardless what it is if processing in comment mode
			continue
		}

		switch token.typ() {
		case tokenValue:
			switch prevToken {
			case tokenFirstToken, tokenBlockEnd:
				rule = append(rule, token.value)
			case tokenSelector:
				rule = append(rule, selector+token.value)
			case tokenBlockStart, tokenStatementEnd: // { or ;
				style = token.value
			case tokenStyleSeparator:
				if isValue { // multiple separators without ;
					return css, fmt.Errorf("line %d: multiple style names before value: %w", token.pos.Line, InvalidCSSError)
				}

				isValue = true
				value = token.value
			case tokenValue:
				if !isBlock { // descendant selector
					rule[len(rule)-1] += " " + token.value
				} else { // technically, this could mean we put multiple style values.
					if !isValue { // want to parse multiple style names? denied.
						return css, fmt.Errorf("line %d: expected only one name before value: %w", token.pos.Line, InvalidCSSError)
					}

					value += " " + token.value
				}
			default:
				return css, fmt.Errorf("line %d: invalid syntax: %w", token.pos.Line, InvalidCSSError)
			}
		case tokenSelector:
			selector = token.value
		case tokenBlockStart:
			if prevToken != tokenValue {
				return css, fmt.Errorf("line %d: block is missing rule identifier: %w", token.pos.Line, InvalidCSSError)
			}
			isBlock = true
			isValue = false
		case tokenStatementEnd:
			if prevToken != tokenValue || style == "" || value == "" {
				return css, fmt.Errorf("line %d: expected style before semicolon: %w", token.pos.Line, InvalidCSSError)
			}
			styles[style] = value
			isValue = false
		case tokenBlockEnd:
			if !isBlock {
				return css, fmt.Errorf("line %d: rule block ends without a beginning: %w", token.pos.Line, InvalidCSSError)
			}

			for i := range rule {
				r := Rule(rule[i])
				oldRule, ok := css[r]
				if ok {
					// merge rules
					for style, value := range oldRule {
						if _, ok := styles[style]; !ok {
							styles[style] = value
						}
					}
				}

				css[r] = styles

			}

			styles = map[string]string{}
			style, value = "", ""
			isBlock = false
			rule = make([]string, 0)
		}
		prevToken = token.typ()
	}

	return css, nil
}

// Unmarshal will take a byte slice, containing sylesheet rules and return
// a map of a rules map.
func Unmarshal(b []byte) (map[Rule]map[string]string, error) {
	return parse(buildList(bytes.NewReader(b)))
}

// CSSStyle returns an error-checked parsed style, or an error if the
// style is unknown. Most of the styles are not supported yet.
func CSSStyle(name string, styles map[string]string) (Style, error) {
	value := styles[name]
	styleFn, ok := StylesTable[name]
	if !ok {
		return Style{}, errors.New("unknown style")
	}
	return styleFn(value)
}
