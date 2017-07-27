package css

import (
	"errors"
	"strconv"
	"strings"
)

func checkColor(color string) error {
	var errColor = errors.New("invalid color")
	if strings.HasPrefix(color, "#") {
		if len(color) == 0 || len(color) > 9 {
			return errColor
		}
		if _, err := strconv.ParseUint(color[1:], 16, 32); err != nil {
			return errColor
		}
		return nil
	}

	switch color {
	case
		"black",
		"silver",
		"gray",
		"white",
		"maroon",
		"red",
		"purple",
		"fuchsia",
		"green",
		"lime",
		"olive",
		"yellow",
		"navy",
		"blue",
		"teal",
		"aqua",
		"orange",
		"aliceblue",
		"antiquewhite",
		"aquamarine",
		"azure",
		"beige",
		"bisque",
		"blanchedalmond",
		"blueviolet",
		"brown",
		"burlywood",
		"cadetblue",
		"chartreuse",
		"chocolate",
		"coral",
		"cornflowerblue",
		"cornsilk",
		"crimson",
		"cyan",
		"darkblue",
		"darkcyan",
		"darkgoldenrod",
		"darkgray",
		"darkgreen ",
		"darkgrey",
		"darkkhaki",
		"darkmagenta",
		"darkolivegreen",
		"darkorange",
		"darkorchid",
		"darkred",
		"darksalmon",
		"darkseagreen",
		"darkslateblue",
		"darkslategray",
		"darkslategrey",
		"darkturquoise",
		"darkviolet",
		"deeppink	",
		"deepskyblue",
		"dimgray",
		"dimgrey",
		"dodgerblue",
		"firebrick",
		"floralwhite",
		"forestgreen",
		"gainsboro",
		"ghostwhite",
		"gold",
		"goldenrod",
		"greenyellow",
		"grey",
		"honeydew",
		"hotpink",
		"indianred",
		"indigo",
		"ivory",
		"khaki",
		"lavender",
		"lavnderblush",
		"lawgreen",
		"lemonchiffon",
		"lightblue",
		"lightcoral",
		"lightcyan",
		"lightgoldenrodyellow",
		"lightgray",
		"lightgreen",
		"lightgrey",
		"lightpink",
		"lightsalmon",
		"lightseagreen",
		"lightskyblue",
		"lightslategray",
		"lightslategrey",
		"lightsteelblue",
		"lightyellow",
		"limegreen",
		"linen",
		"magenta",
		"mediumaquamarine",
		"mediumblue",
		"mediumorchid",
		"mediumpurple",
		"mediumseagreen",
		"mediumslateblue",
		"mediumspringgreen",
		"mediumturquoise",
		"mediumvioletred",
		"midnightblue",
		"mintcream",
		"mistyrose",
		"moccasin",
		"navajowhite",
		"oldlace",
		"olivedrab",
		"orangered",
		"orchid",
		"palegoldenrod",
		"palegreen",
		"paleturquoise",
		"palevioletred",
		"papayawhip",
		"peachpuff",
		"peru",
		"pink",
		"plum",
		"powderblue",
		"rosybrown",
		"royalblue",
		"saddlebrown",
		"salmon",
		"sandybrown",
		"seagreen",
		"seashell",
		"sienna",
		"skyblue",
		"slateblue",
		"slategray",
		"slategrey",
		"snow",
		"springgreen",
		"steelblue",
		"tan",
		"thistle",
		"tomato",
		"turquoise",
		"violet",
		"wheat",
		"whitesmoke",
		"yellowgreen":
		return nil

	}
	return errColor
}

func background(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func backgroundAttachment(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func backgroundColor(value string) (Style, error) {
	if err := checkColor(value); err != nil {
		return Style{}, err
	}

	style := Style{
		Value: value,
	}
	return style, nil
}
func backgroundImage(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func backgroundPosition(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func backgroundRepeat(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func border(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderBottom(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderBottomColor(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderBottomStyle(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderBottomWidth(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderColor(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderLeft(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderLeftColor(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderLeftStyle(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderLeftWidth(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderRight(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderRightColor(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderRightStyle(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderRightWidth(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderStyle(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderTop(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderTopColor(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderTopStyle(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderTopWidth(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func borderWidth(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func clear(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func clip(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func color(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func cursor(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func display(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func filter(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func font(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func fontFamily(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func fontSize(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func fontVariant(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func fontWeight(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func height(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func left(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func letterSpacing(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func lineHeight(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func listStyle(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func listStyleImage(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func listStylePosition(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func listStyleType(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func margin(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func marginBottom(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func marginLeft(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func marginRight(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func marginTop(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func overflow(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func padding(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func paddingBottom(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func paddingLeft(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func paddingRight(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func paddingTop(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func pageBreakAfter(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func pageBreakBefore(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func position(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func float(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func textAlign(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func textDecoration(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func textDecorationBlink(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func textDecorationLineThrough(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func textDecorationNone(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func textDecorationOverline(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func textDecorationUnderline(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func textIndent(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func textTransform(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func top(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func verticalAlign(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func visibility(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func width(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
func zIndex(value string) (Style, error) {
	return Style{}, errors.New("not implemented")
}
