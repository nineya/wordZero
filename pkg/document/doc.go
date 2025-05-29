/*
Package document 提供了用于创建、编辑和操作 Microsoft Word 文档的 Golang 库。

WordZero 专注于现代的 Office Open XML (OOXML) 格式（.docx 文件），
提供了简单易用的 API 来创建和修改 Word 文档。

# 主要功能

## 基础功能
- 创建新的 Word 文档
- 打开和解析现有的 .docx 文件
- 添加和格式化文本内容
- 设置段落样式和对齐方式
- 配置字体、颜色和文本格式
- 设置行间距和段落间距
- 错误处理和日志记录

## 高级功能 ✨ 新增
- **页眉页脚**: 支持默认、首页、偶数页不同的页眉页脚设置
- **目录生成**: 基于标题样式自动生成目录，支持超链接和页码
- **脚注尾注**: 完整的脚注和尾注功能，支持多种编号格式
- **列表编号**: 无序列表和有序列表，支持多级嵌套
- **页面设置**: 页面尺寸、方向、边距等完整页面属性设置
- **表格功能**: 强大的表格创建、格式化和样式设置功能
- **样式系统**: 18种预定义样式和自定义样式支持

# 快速开始

创建一个简单的文档：

	doc := document.New()
	doc.AddParagraph("Hello, World!")
	err := doc.Save("hello.docx")

创建带格式的文档：

	doc := document.New()

	// 添加格式化标题
	titleFormat := &document.TextFormat{
		Bold:      true,
		FontSize:  18,
		FontColor: "FF0000", // 红色
		FontName:  "微软雅黑",
	}
	title := doc.AddFormattedParagraph("文档标题", titleFormat)
	title.SetAlignment(document.AlignCenter)

	// 添加正文段落
	para := doc.AddParagraph("这是正文内容...")
	para.SetSpacing(&document.SpacingConfig{
		LineSpacing:     1.5, // 1.5倍行距
		BeforePara:      12,  // 段前12磅
		AfterPara:       6,   // 段后6磅
		FirstLineIndent: 24,  // 首行缩进24磅
	})

	err := doc.Save("formatted.docx")

打开现有文档：

	doc, err := document.Open("existing.docx")
	if err != nil {
		log.Fatal(err)
	}

	// 读取段落内容
	for i, para := range doc.Body.Paragraphs {
		fmt.Printf("段落 %d: ", i+1)
		for _, run := range para.Runs {
			fmt.Print(run.Text.Content)
		}
		fmt.Println()
	}

# 高级功能示例

## 页眉页脚功能

	// 添加页眉
	doc.AddHeader(document.HeaderFooterTypeDefault, "这是页眉")

	// 添加带页码的页脚
	doc.AddFooterWithPageNumber(document.HeaderFooterTypeDefault, "第", true)

	// 设置首页不同
	doc.SetDifferentFirstPage(true)

## 目录生成

	// 添加带书签的标题
	doc.AddHeadingWithBookmark("第一章 概述", 1, "chapter1")
	doc.AddHeadingWithBookmark("1.1 背景", 2, "section1_1")

	// 生成目录
	tocConfig := document.DefaultTOCConfig()
	tocConfig.Title = "目录"
	tocConfig.MaxLevel = 3
	doc.GenerateTOC(tocConfig)

## 脚注和尾注

	// 添加脚注
	doc.AddFootnote("这是正文内容", "这是脚注内容")

	// 添加尾注
	doc.AddEndnote("更多说明", "这是尾注内容")

	// 自定义脚注配置
	footnoteConfig := &document.FootnoteConfig{
		NumberFormat: document.FootnoteFormatLowerRoman,
		StartNumber:  1,
		RestartEach:  document.FootnoteRestartEachPage,
		Position:     document.FootnotePositionPageBottom,
	}
	doc.SetFootnoteConfig(footnoteConfig)

## 列表功能

	// 无序列表
	doc.AddBulletList("列表项1", 0, document.BulletTypeDot)
	doc.AddBulletList("子项目", 1, document.BulletTypeCircle)

	// 有序列表
	doc.AddNumberedList("第一项", 0, document.ListTypeDecimal)
	doc.AddNumberedList("第二项", 0, document.ListTypeDecimal)

	// 多级列表
	items := []document.ListItem{
		{Text: "一级项目", Level: 0, Type: document.ListTypeDecimal},
		{Text: "二级项目", Level: 1, Type: document.ListTypeLowerLetter},
		{Text: "三级项目", Level: 2, Type: document.ListTypeLowerRoman},
	}
	doc.CreateMultiLevelList(items)

## 页面设置

	// 设置页面为A4横向
	doc.SetPageOrientation(document.OrientationLandscape)

	// 设置页面边距（毫米）
	doc.SetPageMargins(25, 25, 25, 25)

	// 完整页面设置
	pageSettings := &document.PageSettings{
		Size:           document.PageSizeLetter,
		Orientation:    document.OrientationPortrait,
		MarginTop:      30,
		MarginRight:    20,
		MarginBottom:   30,
		MarginLeft:     20,
		HeaderDistance: 15,
		FooterDistance: 15,
		GutterWidth:    0,
	}
	doc.SetPageSettings(pageSettings)

## 表格功能

	// 创建表格
	table := doc.CreateTable(&document.TableConfig{
		Rows:  3,
		Cols:  3,
		Width: 5000,
	})

	// 设置单元格内容
	table.SetCellText(0, 0, "标题")

	// 设置表格样式
	table.ApplyTableStyle(&document.TableStyleConfig{
		HeaderRow:    true,
		FirstColumn:  true,
		BandedRows:   true,
		BandedCols:   false,
	})

# 错误处理

库提供了统一的错误处理机制：

	doc, err := document.Open("nonexistent.docx")
	if err != nil {
		var docErr *document.DocumentError
		if errors.As(err, &docErr) {
			Errorf("文档操作失败 - 操作: %s, 错误: %v", docErr.Operation, docErr.Cause)
			fmt.Printf("操作: %s, 错误: %v\n", docErr.Operation, docErr.Cause)
		}
	}

# 日志记录

可以配置日志级别来控制输出：

	// 设置为调试模式
	document.SetGlobalLevel(document.LogLevelDebug)

	// 只显示错误
	document.SetGlobalLevel(document.LogLevelError)

# 文本格式

TextFormat 结构体支持多种文本格式选项：

	format := &document.TextFormat{
		Bold:      true,           // 粗体
		Italic:    true,           // 斜体
		FontSize:  14,             // 字体大小（磅）
		FontColor: "0000FF",       // 字体颜色（十六进制）
		FontName:  "Times New Roman", // 字体名称
	}

# 段落对齐

支持四种对齐方式：

	para.SetAlignment(document.AlignLeft)     // 左对齐
	para.SetAlignment(document.AlignCenter)   // 居中对齐
	para.SetAlignment(document.AlignRight)    // 右对齐
	para.SetAlignment(document.AlignJustify)  // 两端对齐

# 间距配置

可以精确控制段落间距：

	config := &document.SpacingConfig{
		LineSpacing:     1.5, // 行间距（倍数）
		BeforePara:      12,  // 段前间距（磅）
		AfterPara:       6,   // 段后间距（磅）
		FirstLineIndent: 24,  // 首行缩进（磅）
	}
	para.SetSpacing(config)

# 注意事项

- 字体大小以磅为单位，内部会自动转换为 Word 的半磅单位
- 颜色值使用十六进制格式，不需要 # 前缀
- 间距值以磅为单位，内部会转换为 TWIPs（1磅=20TWIPs）
- 所有文本内容都使用 UTF-8 编码
- 页眉页脚类型包括：Default（默认）、First（首页）、Even（偶数页）
- 脚注和尾注会自动编号，支持多种编号格式和重启规则
- 列表支持多级嵌套，最多支持9级缩进
- 目录功能需要先添加带书签的标题，然后调用生成目录方法

更多详细信息和示例，请参阅各个类型和函数的文档。
*/
package document
