// Code generated by "stringer -type=tokenType"; DO NOT EDIT.

package css

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[tokenFirstToken - -1]
	_ = x[tokenBlockStart-0]
	_ = x[tokenBlockEnd-1]
	_ = x[tokenRuleName-2]
	_ = x[tokenValue-3]
	_ = x[tokenSelector-4]
	_ = x[tokenStyleSeparator-5]
	_ = x[tokenStatementEnd-6]
	_ = x[tokenCommentStart-7]
	_ = x[tokenCommentEnd-8]
}

const _tokenType_name = "tokenFirstTokentokenBlockStarttokenBlockEndtokenRuleNametokenValuetokenSelectortokenStyleSeparatortokenStatementEndtokenCommentStarttokenCommentEnd"

var _tokenType_index = [...]uint8{0, 15, 30, 43, 56, 66, 79, 98, 115, 132, 147}

func (i tokenType) String() string {
	i -= -1
	if i < 0 || i >= tokenType(len(_tokenType_index)-1) {
		return "tokenType(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _tokenType_name[_tokenType_index[i]:_tokenType_index[i+1]]
}
