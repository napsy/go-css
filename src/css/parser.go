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

type tokenType int

const (
	tokenBlockStart tokenType = iota
	tokenBlockEnd
	tokenRuleName
	tokenValue
	tokenStyleSeparator
	tokenStatementEnd
)

type Rule string

type tokenEntry struct {
	value string
	pos   scanner.Position
}

type tokenizer struct {
	s *scanner.Scanner
}

func (rule Rule) Type() string {
	if strings.HasPrefix(string(rule), ".") {
		return "class"
	}
	if strings.HasPrefix(string(rule), ".") {
		return "id"
	}
	return "tag"
}

func (e tokenEntry) typ() tokenType {
	return newTokenType(e.value)
}

func (t *tokenizer) next() (tokenEntry, error) {
	token := t.s.Scan()
	if token == scanner.EOF {
		return tokenEntry{}, errors.New("EOF")
	}
	value := t.s.TokenText()
	pos := t.s.Pos()
	return tokenEntry{
		value: value,
		pos:   pos,
	}, nil
}

func (t tokenType) String() string {
	switch t {
	case tokenBlockStart:
		return "BLOCK_START"
	case tokenBlockEnd:
		return "BLOCK_END"
	case tokenStyleSeparator:
		return "STYLE_SEPARATOR"
	case tokenStatementEnd:
		return "STATEMENT_END"
	}
	return "VALUE"
}

func newTokenType(typ string) tokenType {
	switch typ {
	case "{":
		return tokenBlockStart
	case "}":
		return tokenBlockEnd
	case ":":
		return tokenStyleSeparator
	case ";":
		return tokenStatementEnd
	}
	return tokenValue
}

func newTokenizer(r io.Reader) *tokenizer {
	s := &scanner.Scanner{}
	s.Init(r)
	return &tokenizer{
		s: s,
	}
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
		rule  string
		style string
		value string

		// Parsed styles.
		css    = make(map[Rule]map[string]string)
		styles = make(map[string]string)

		// Previous token for the state machine.
		prevToken = tokenType(-1)
	)

	for e := l.Front(); e != nil; e = l.Front() {
		token := e.Value.(tokenEntry)
		l.Remove(e)

		switch token.typ() {
		case tokenValue:
			switch prevToken {
			case -1, tokenBlockEnd:
				rule = token.value
			case tokenBlockStart, tokenStatementEnd:
				style = token.value
			case tokenStyleSeparator:
				value = token.value
			default:
				return css, fmt.Errorf("invalid syntax: line %d\n", token.pos.Line)
			}
		case tokenBlockStart:
			if prevToken != tokenValue {
				return css, fmt.Errorf("block without rule name: line %d\n", token.pos.Line)
			}
		case tokenStatementEnd:
			styles[style] = value
		case tokenBlockEnd:
			css[Rule(rule)] = styles
		}
		prevToken = token.typ()
	}
	return css, nil
}

func Unmarshal(b []byte) (map[Rule]map[string]string, error) {
	return parse(buildList(bytes.NewReader(b)))
}
