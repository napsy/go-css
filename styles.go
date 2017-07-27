package css

import "fmt"

type UnitValue float64

type UnitType int

const (
	UnitPixels UnitType = iota
	UnitEm
	UnitRem
	UnitPercent
	UnitPt
	UnitAuto
	UnitNone
)

type Style struct {
	Value interface{}
}

func (style Style) Unit() UnitType {
	return UnitNone
}

func (style Style) String() string {
	return fmt.Sprintf("%v", style.Value)
}

type styleHandler func(value string) (Style, error)

// Common CSS styles
var stylesTable = map[string]styleHandler{
	"background":                    background,
	"background-attachment":         backgroundAttachment,
	"background-color":              backgroundColor,
	"background-image":              backgroundImage,
	"background-position":           backgroundPosition,
	"background-repeat":             backgroundRepeat,
	"border":                        border,
	"border-bottom":                 borderBottom,
	"border-bottom-color":           borderBottomColor,
	"border-bottom-style":           borderBottomStyle,
	"border-bottom-width":           borderBottomWidth,
	"border-color":                  borderColor,
	"border-left":                   borderLeft,
	"border-left-color":             borderLeftColor,
	"border-left-style":             borderLeftStyle,
	"border-left-width":             borderLeftWidth,
	"border-right":                  borderRight,
	"border-right-color":            borderRightColor,
	"border-right-style":            borderRightStyle,
	"border-right-width":            borderRightWidth,
	"border-style":                  borderStyle,
	"border-top":                    borderTop,
	"border-top-color":              borderTopColor,
	"border-top-style":              borderTopStyle,
	"border-top-width":              borderTopWidth,
	"border-width":                  borderWidth,
	"clear":                         clear,
	"clip":                          clip,
	"color":                         color,
	"cursor":                        cursor,
	"display":                       display,
	"filter":                        filter,
	"font":                          font,
	"font-family":                   fontFamily,
	"font-size":                     fontSize,
	"font-variant":                  fontVariant,
	"font-weight":                   fontWeight,
	"height":                        height,
	"left":                          left,
	"letter-spacing":                letterSpacing,
	"line-height":                   lineHeight,
	"list-style":                    listStyle,
	"list-style-image":              listStyleImage,
	"list-style-position":           listStylePosition,
	"list-style-type":               listStyleType,
	"margin":                        margin,
	"margin-bottom":                 marginBottom,
	"margin-left":                   marginLeft,
	"margin-right":                  marginRight,
	"margin-top":                    marginTop,
	"overflow":                      overflow,
	"padding":                       padding,
	"padding-bottom":                paddingBottom,
	"padding-left":                  paddingLeft,
	"padding-right":                 paddingRight,
	"padding-top":                   paddingTop,
	"page-break-after":              pageBreakAfter,
	"page-break-before":             pageBreakBefore,
	"position":                      position,
	"float":                         float,
	"text-align":                    textAlign,
	"text-decoration":               textDecoration,
	"text-decoration: blink":        textDecorationBlink,
	"text-decoration: line-through": textDecorationLineThrough,
	"text-decoration: none":         textDecorationNone,
	"text-decoration: overline":     textDecorationOverline,
	"text-decoration: underline":    textDecorationUnderline,
	"text-indent":                   textIndent,
	"text-transform":                textTransform,
	"top":                           top,
	"vertical-align":                verticalAlign,
	"visibility":                    visibility,
	"width":                         width,
	"z-index":                       zIndex,
}
