// Package style 预定义样式常量
package style

// 预定义样式ID常量
const (
	// StyleNormal 普通文本样式
	StyleNormal = "Normal"

	// 标题样式
	StyleHeading1 = "Heading1"
	StyleHeading2 = "Heading2"
	StyleHeading3 = "Heading3"
	StyleHeading4 = "Heading4"
	StyleHeading5 = "Heading5"
	StyleHeading6 = "Heading6"
	StyleHeading7 = "Heading7"
	StyleHeading8 = "Heading8"
	StyleHeading9 = "Heading9"

	// 文档标题样式
	StyleTitle    = "Title"    // 文档标题
	StyleSubtitle = "Subtitle" // 副标题

	// 字符样式
	StyleEmphasis = "Emphasis" // 强调（斜体）
	StyleStrong   = "Strong"   // 加粗
	StyleCodeChar = "CodeChar" // 代码字符

	// 段落样式
	StyleQuote         = "Quote"         // 引用样式
	StyleListParagraph = "ListParagraph" // 列表段落
	StyleCodeBlock     = "CodeBlock"     // 代码块
)

// GetPredefinedStyleNames 获取所有预定义样式名称映射
func GetPredefinedStyleNames() map[string]string {
	return map[string]string{
		StyleNormal:        "普通文本",
		StyleHeading1:      "标题 1",
		StyleHeading2:      "标题 2",
		StyleHeading3:      "标题 3",
		StyleHeading4:      "标题 4",
		StyleHeading5:      "标题 5",
		StyleHeading6:      "标题 6",
		StyleHeading7:      "标题 7",
		StyleHeading8:      "标题 8",
		StyleHeading9:      "标题 9",
		StyleTitle:         "文档标题",
		StyleSubtitle:      "副标题",
		StyleEmphasis:      "强调",
		StyleStrong:        "加粗",
		StyleCodeChar:      "代码字符",
		StyleQuote:         "引用",
		StyleListParagraph: "列表段落",
		StyleCodeBlock:     "代码块",
	}
}

// StyleConfig 样式配置帮助结构
type StyleConfig struct {
	StyleID     string
	Name        string
	Description string
	StyleType   StyleType
}

// GetPredefinedStyleConfigs 获取所有预定义样式配置
func GetPredefinedStyleConfigs() []StyleConfig {
	return []StyleConfig{
		{
			StyleID:     StyleNormal,
			Name:        "普通文本",
			Description: "默认的段落样式，使用Calibri字体，11磅字号",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleHeading1,
			Name:        "标题 1",
			Description: "一级标题，16磅蓝色粗体，段前12磅间距",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleHeading2,
			Name:        "标题 2",
			Description: "二级标题，13磅蓝色粗体，段前6磅间距",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleHeading3,
			Name:        "标题 3",
			Description: "三级标题，12磅蓝色粗体，段前6磅间距",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleHeading4,
			Name:        "标题 4",
			Description: "四级标题，12磅蓝色粗体，段前6磅间距",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleHeading5,
			Name:        "标题 5",
			Description: "五级标题，12磅蓝色粗体，段前6磅间距",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleHeading6,
			Name:        "标题 6",
			Description: "六级标题，12磅蓝色粗体，段前6磅间距",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleHeading7,
			Name:        "标题 7",
			Description: "七级标题，12磅蓝色粗体，段前6磅间距",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleHeading8,
			Name:        "标题 8",
			Description: "八级标题，12磅蓝色粗体，段前6磅间距",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleHeading9,
			Name:        "标题 9",
			Description: "九级标题，12磅蓝色粗体，段前6磅间距",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleTitle,
			Name:        "文档标题",
			Description: "文档标题样式",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleSubtitle,
			Name:        "副标题",
			Description: "副标题样式",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleEmphasis,
			Name:        "强调",
			Description: "斜体文本样式",
			StyleType:   StyleTypeCharacter,
		},
		{
			StyleID:     StyleStrong,
			Name:        "加粗",
			Description: "粗体文本样式",
			StyleType:   StyleTypeCharacter,
		},
		{
			StyleID:     StyleCodeChar,
			Name:        "代码字符",
			Description: "等宽字体，红色文本，适用于代码片段",
			StyleType:   StyleTypeCharacter,
		},
		{
			StyleID:     StyleQuote,
			Name:        "引用",
			Description: "引用段落样式，斜体灰色，左右各缩进0.5英寸",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleListParagraph,
			Name:        "列表段落",
			Description: "列表段落样式",
			StyleType:   StyleTypeParagraph,
		},
		{
			StyleID:     StyleCodeBlock,
			Name:        "代码块",
			Description: "代码块样式",
			StyleType:   StyleTypeParagraph,
		},
	}
}
