// Package document 提供Word文档域字段结构
package document

import (
	"encoding/xml"
	"fmt"
)

// FieldChar 域字符
type FieldChar struct {
	XMLName       xml.Name `xml:"w:fldChar"`
	FieldCharType string   `xml:"w:fldCharType,attr"`
}

// InstrText 域指令文本
type InstrText struct {
	XMLName xml.Name `xml:"w:instrText"`
	Space   string   `xml:"xml:space,attr,omitempty"`
	Content string   `xml:",chardata"`
}

// HyperlinkField 超链接域
type HyperlinkField struct {
	BeginChar    FieldChar
	InstrText    InstrText
	SeparateChar FieldChar
	EndChar      FieldChar
}

// CreateHyperlinkField 创建超链接域
func CreateHyperlinkField(anchor string) HyperlinkField {
	return HyperlinkField{
		BeginChar: FieldChar{
			FieldCharType: "begin",
		},
		InstrText: InstrText{
			Space:   "preserve",
			Content: fmt.Sprintf(" HYPERLINK \\l %s ", anchor),
		},
		SeparateChar: FieldChar{
			FieldCharType: "separate",
		},
		EndChar: FieldChar{
			FieldCharType: "end",
		},
	}
}

// PageRefField 页码引用域
type PageRefField struct {
	BeginChar    FieldChar
	InstrText    InstrText
	SeparateChar FieldChar
	EndChar      FieldChar
}

// CreatePageRefField 创建页码引用域
func CreatePageRefField(anchor string) PageRefField {
	return PageRefField{
		BeginChar: FieldChar{
			FieldCharType: "begin",
		},
		InstrText: InstrText{
			Space:   "preserve",
			Content: fmt.Sprintf(" PAGEREF %s \\h ", anchor),
		},
		SeparateChar: FieldChar{
			FieldCharType: "separate",
		},
		EndChar: FieldChar{
			FieldCharType: "end",
		},
	}
}
