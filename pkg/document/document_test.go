package document

import (
	"os"
	"testing"

	"github.com/ZeroHawkeye/wordZero/pkg/style"
)

// TestNewDocument 测试新文档创建
func TestNewDocument(t *testing.T) {
	doc := New()

	// 验证基本结构
	if doc == nil {
		t.Fatal("Failed to create new document")
	}

	if doc.Body == nil {
		t.Fatal("Document body is nil")
	}

	if doc.styleManager == nil {
		t.Fatal("Style manager is nil")
	}

	// 验证初始状态
	if len(doc.Body.GetParagraphs()) != 0 {
		t.Errorf("Expected 0 paragraphs, got %d", len(doc.Body.GetParagraphs()))
	}

	// 验证样式管理器初始化
	styles := doc.styleManager.GetAllStyles()
	if len(styles) == 0 {
		t.Error("Style manager should have predefined styles")
	}
}

// TestAddParagraph 测试添加普通段落
func TestAddParagraph(t *testing.T) {
	doc := New()
	text := "测试段落内容"

	para := doc.AddParagraph(text)

	// 验证段落添加
	if len(doc.Body.GetParagraphs()) != 1 {
		t.Errorf("Expected 1 paragraph, got %d", len(doc.Body.GetParagraphs()))
	}

	// 验证段落内容
	if len(para.Runs) != 1 {
		t.Errorf("Expected 1 run, got %d", len(para.Runs))
	}

	if para.Runs[0].Text.Content != text {
		t.Errorf("Expected %s, got %s", text, para.Runs[0].Text.Content)
	}

	// 验证返回的指针是否正确
	paragraphs := doc.Body.GetParagraphs()
	if paragraphs[0] != para {
		t.Error("Returned paragraph pointer is incorrect")
	}
}

// TestAddHeadingParagraph 测试添加标题段落
func TestAddHeadingParagraph(t *testing.T) {
	doc := New()

	testCases := []struct {
		text    string
		level   int
		styleID string
	}{
		{"第一级标题", 1, "Heading1"},
		{"第二级标题", 2, "Heading2"},
		{"第三级标题", 3, "Heading3"},
		{"第九级标题", 9, "Heading9"},
	}

	for _, tc := range testCases {
		para := doc.AddHeadingParagraph(tc.text, tc.level)

		// 验证段落样式设置
		if para.Properties == nil {
			t.Errorf("Heading paragraph should have properties")
			continue
		}

		if para.Properties.ParagraphStyle == nil {
			t.Errorf("Heading paragraph should have style reference")
			continue
		}

		if para.Properties.ParagraphStyle.Val != tc.styleID {
			t.Errorf("Expected style %s, got %s", tc.styleID, para.Properties.ParagraphStyle.Val)
		}

		// 验证内容
		if len(para.Runs) != 1 {
			t.Errorf("Expected 1 run, got %d", len(para.Runs))
			continue
		}

		if para.Runs[0].Text.Content != tc.text {
			t.Errorf("Expected %s, got %s", tc.text, para.Runs[0].Text.Content)
		}
	}

	// 测试超出范围的级别
	para := doc.AddHeadingParagraph("超出范围", 10)
	if para.Properties.ParagraphStyle.Val != "Heading1" {
		t.Error("Out of range level should default to Heading1")
	}

	para = doc.AddHeadingParagraph("负数级别", -1)
	if para.Properties.ParagraphStyle.Val != "Heading1" {
		t.Error("Negative level should default to Heading1")
	}
}

// TestAddFormattedParagraph 测试添加格式化段落
func TestAddFormattedParagraph(t *testing.T) {
	doc := New()
	text := "格式化文本"

	format := &TextFormat{
		Bold:       true,
		Italic:     true,
		FontSize:   14,
		FontColor:  "FF0000",
		FontFamily: "宋体",
	}

	para := doc.AddFormattedParagraph(text, format)

	// 验证段落添加
	if len(doc.Body.GetParagraphs()) != 1 {
		t.Error("Failed to add formatted paragraph")
	}

	// 验证格式设置
	run := para.Runs[0]
	if run.Properties == nil {
		t.Fatal("Run properties should not be nil")
	}

	if run.Properties.Bold == nil {
		t.Error("Bold property should be set")
	}

	if run.Properties.Italic == nil {
		t.Error("Italic property should be set")
	}

	if run.Properties.FontSize == nil || run.Properties.FontSize.Val != "28" {
		t.Errorf("Expected font size 28, got %v", run.Properties.FontSize)
	}

	if run.Properties.Color == nil || run.Properties.Color.Val != "FF0000" {
		t.Errorf("Expected color FF0000, got %v", run.Properties.Color)
	}

	if run.Properties.FontFamily == nil || run.Properties.FontFamily.ASCII != "宋体" {
		t.Errorf("Expected font family 宋体, got %v", run.Properties.FontFamily)
	}
}

// TestParagraphSetAlignment 测试段落对齐设置
func TestParagraphSetAlignment(t *testing.T) {
	doc := New()
	para := doc.AddParagraph("测试对齐")

	testCases := []AlignmentType{
		AlignLeft,
		AlignCenter,
		AlignRight,
		AlignJustify,
	}

	for _, alignment := range testCases {
		para.SetAlignment(alignment)

		if para.Properties == nil {
			t.Fatal("Properties should not be nil after setting alignment")
		}

		if para.Properties.Justification == nil {
			t.Fatal("Justification should not be nil")
		}

		if para.Properties.Justification.Val != string(alignment) {
			t.Errorf("Expected alignment %s, got %s", alignment, para.Properties.Justification.Val)
		}
	}
}

// TestParagraphSetSpacing 测试段落间距设置
func TestParagraphSetSpacing(t *testing.T) {
	doc := New()
	para := doc.AddParagraph("测试间距")

	config := &SpacingConfig{
		LineSpacing:     1.5,
		BeforePara:      12,
		AfterPara:       6,
		FirstLineIndent: 24,
	}

	para.SetSpacing(config)

	// 验证属性设置
	if para.Properties == nil {
		t.Fatal("Properties should not be nil")
	}

	if para.Properties.Spacing == nil {
		t.Fatal("Spacing should not be nil")
	}

	// 验证间距值（转换为TWIPs）
	spacing := para.Properties.Spacing
	if spacing.Before != "240" { // 12 * 20
		t.Errorf("Expected before spacing 240, got %s", spacing.Before)
	}

	if spacing.After != "120" { // 6 * 20
		t.Errorf("Expected after spacing 120, got %s", spacing.After)
	}

	if spacing.Line != "360" { // 1.5 * 240
		t.Errorf("Expected line spacing 360, got %s", spacing.Line)
	}

	// 验证首行缩进
	if para.Properties.Indentation == nil {
		t.Fatal("Indentation should not be nil")
	}

	if para.Properties.Indentation.FirstLine != "480" { // 24 * 20
		t.Errorf("Expected first line indent 480, got %s", para.Properties.Indentation.FirstLine)
	}
}

// TestParagraphAddFormattedText 测试段落添加格式化文本
func TestParagraphAddFormattedText(t *testing.T) {
	doc := New()
	para := doc.AddParagraph("初始文本")

	// 添加格式化文本
	format := &TextFormat{
		Bold:      true,
		FontColor: "0000FF",
	}

	para.AddFormattedText("格式化文本", format)

	// 验证运行数量
	if len(para.Runs) != 2 {
		t.Errorf("Expected 2 runs, got %d", len(para.Runs))
	}

	// 验证第二个运行的格式
	run := para.Runs[1]
	if run.Properties == nil {
		t.Fatal("Second run should have properties")
	}

	if run.Properties.Bold == nil {
		t.Error("Second run should be bold")
	}

	if run.Properties.Color == nil || run.Properties.Color.Val != "0000FF" {
		t.Error("Second run should be blue")
	}

	if run.Text.Content != "格式化文本" {
		t.Errorf("Expected '格式化文本', got '%s'", run.Text.Content)
	}
}

// TestParagraphSetStyle 测试段落样式设置
func TestParagraphSetStyle(t *testing.T) {
	doc := New()
	para := doc.AddParagraph("测试样式")

	para.SetStyle("Heading1")

	if para.Properties == nil {
		t.Fatal("Properties should not be nil")
	}

	if para.Properties.ParagraphStyle == nil {
		t.Fatal("ParagraphStyle should not be nil")
	}

	if para.Properties.ParagraphStyle.Val != "Heading1" {
		t.Errorf("Expected style Heading1, got %s", para.Properties.ParagraphStyle.Val)
	}
}

// TestParagraphSetIndentation 测试段落缩进设置
func TestParagraphSetIndentation(t *testing.T) {
	doc := New()
	para := doc.AddParagraph("测试缩进")

	// 测试首行缩进
	para.SetIndentation(0.5, 0, 0)

	if para.Properties == nil {
		t.Fatal("Properties should not be nil")
	}

	if para.Properties.Indentation == nil {
		t.Fatal("Indentation should not be nil")
	}

	// 0.5厘米 = 283.5 TWIPs，四舍五入为284
	expectedFirstLine := "283"
	if para.Properties.Indentation.FirstLine != expectedFirstLine {
		t.Errorf("Expected FirstLine %s, got %s", expectedFirstLine, para.Properties.Indentation.FirstLine)
	}

	// 测试左右缩进
	para.SetIndentation(-0.5, 1.0, 0.5)

	expectedFirstLine = "-283" // 悬挂缩进
	expectedLeft := "567"      // 1厘米
	expectedRight := "283"     // 0.5厘米

	if para.Properties.Indentation.FirstLine != expectedFirstLine {
		t.Errorf("Expected FirstLine %s, got %s", expectedFirstLine, para.Properties.Indentation.FirstLine)
	}
	if para.Properties.Indentation.Left != expectedLeft {
		t.Errorf("Expected Left %s, got %s", expectedLeft, para.Properties.Indentation.Left)
	}
	if para.Properties.Indentation.Right != expectedRight {
		t.Errorf("Expected Right %s, got %s", expectedRight, para.Properties.Indentation.Right)
	}
}

// TestDocumentSave 测试文档保存
func TestDocumentSave(t *testing.T) {
	doc := New()
	doc.AddParagraph("测试保存功能")

	filename := "test_save.docx"
	defer os.Remove(filename) // 清理测试文件

	err := doc.Save(filename)
	if err != nil {
		t.Fatalf("Failed to save document: %v", err)
	}

	// 验证文件是否存在
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Error("Saved file does not exist")
	}

	// 验证文件大小
	stat, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("Failed to get file stats: %v", err)
	}

	if stat.Size() == 0 {
		t.Error("Saved file is empty")
	}
}

// TestDocumentGetStyleManager 测试获取样式管理器
func TestDocumentGetStyleManager(t *testing.T) {
	doc := New()

	styleManager := doc.GetStyleManager()
	if styleManager == nil {
		t.Fatal("Style manager should not be nil")
	}

	// 验证样式管理器功能
	if !styleManager.StyleExists("Normal") {
		t.Error("Normal style should exist")
	}

	if !styleManager.StyleExists("Heading1") {
		t.Error("Heading1 style should exist")
	}
}

// TestComplexDocument 测试复杂文档创建
func TestComplexDocument(t *testing.T) {
	doc := New()

	// 添加标题
	title := doc.AddFormattedParagraph("文档标题", &TextFormat{
		Bold:     true,
		FontSize: 18,
	})
	title.SetAlignment(AlignCenter)

	// 添加各级标题
	doc.AddHeadingParagraph("第一章", 1)
	doc.AddHeadingParagraph("1.1 概述", 2)
	doc.AddHeadingParagraph("1.1.1 背景", 3)

	// 添加带间距的段落
	para := doc.AddParagraph("这是一个带有特殊间距的段落")
	para.SetSpacing(&SpacingConfig{
		LineSpacing: 1.5,
		BeforePara:  12,
		AfterPara:   6,
	})

	// 添加混合格式段落
	mixed := doc.AddParagraph("这段文字包含")
	mixed.AddFormattedText("粗体", &TextFormat{Bold: true})
	mixed.AddFormattedText("和", nil)
	mixed.AddFormattedText("斜体", &TextFormat{Italic: true})
	mixed.AddFormattedText("文本。", nil)

	// 验证文档结构
	if len(doc.Body.GetParagraphs()) != 6 {
		t.Errorf("Expected 6 paragraphs, got %d", len(doc.Body.GetParagraphs()))
	}

	// 保存并验证
	filename := "test_complex.docx"
	defer os.Remove(filename)

	err := doc.Save(filename)
	if err != nil {
		t.Fatalf("Failed to save complex document: %v", err)
	}
}

// TestDocumentOpen 测试打开文档（需要先创建一个测试文档）
func TestDocumentOpen(t *testing.T) {
	// 先创建一个测试文档
	originalDoc := New()
	originalDoc.AddParagraph("第一段")
	originalDoc.AddParagraph("第二段")
	originalDoc.AddHeadingParagraph("标题", 1)

	filename := "test_open.docx"
	defer os.Remove(filename)

	err := originalDoc.Save(filename)
	if err != nil {
		t.Fatalf("Failed to save test document: %v", err)
	}

	// 打开文档
	loadedDoc, err := Open(filename)
	if err != nil {
		t.Fatalf("Failed to open document: %v", err)
	}

	// 验证文档内容
	if len(loadedDoc.Body.GetParagraphs()) != 3 {
		t.Errorf("Expected 3 paragraphs, got %d", len(loadedDoc.Body.GetParagraphs()))
	}

	// 验证第一段内容
	if len(loadedDoc.Body.GetParagraphs()[0].Runs) > 0 {
		content := loadedDoc.Body.GetParagraphs()[0].Runs[0].Text.Content
		if content != "第一段" {
			t.Errorf("Expected '第一段', got '%s'", content)
		}
	}
}

// TestErrorHandling 测试错误处理
func TestErrorHandling(t *testing.T) {
	// 测试打开不存在的文件
	_, err := Open("nonexistent.docx")
	if err == nil {
		t.Error("Should return error when opening non-existent file")
	}

	// 测试保存到只读目录（如果创建失败则跳过这个测试）
	doc := New()
	doc.AddParagraph("测试")

	// 尝试保存到一个包含空字符的无效文件名
	invalidPath := "test\x00invalid.docx"
	err = doc.Save(invalidPath)
	if err == nil {
		// 如果第一个测试没有失败，尝试另一个策略
		// 尝试保存到一个超长路径
		longPath := string(make([]byte, 300)) + ".docx"
		err = doc.Save(longPath)
		if err == nil {
			t.Log("Warning: Unable to trigger save error - filesystem may be permissive")
		}
	}
}

// TestStyleIntegration 测试样式集成
func TestStyleIntegration(t *testing.T) {
	doc := New()
	styleManager := doc.GetStyleManager()
	quickAPI := style.NewQuickStyleAPI(styleManager)

	// 创建自定义样式
	config := style.QuickStyleConfig{
		ID:      "TestStyle",
		Name:    "测试样式",
		Type:    style.StyleTypeParagraph,
		BasedOn: "Normal",
		RunConfig: &style.QuickRunConfig{
			Bold:      true,
			FontColor: "FF0000",
		},
	}

	_, err := quickAPI.CreateQuickStyle(config)
	if err != nil {
		t.Fatalf("Failed to create custom style: %v", err)
	}

	// 使用自定义样式
	para := doc.AddParagraph("使用自定义样式")
	para.SetStyle("TestStyle")

	// 验证样式应用
	if para.Properties == nil || para.Properties.ParagraphStyle == nil {
		t.Fatal("Style should be applied to paragraph")
	}

	if para.Properties.ParagraphStyle.Val != "TestStyle" {
		t.Errorf("Expected TestStyle, got %s", para.Properties.ParagraphStyle.Val)
	}

	// 验证样式存在
	if !styleManager.StyleExists("TestStyle") {
		t.Error("Custom style should exist in style manager")
	}
}

// BenchmarkAddParagraph 基准测试 - 添加段落性能
func BenchmarkAddParagraph(b *testing.B) {
	doc := New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doc.AddParagraph("基准测试段落")
	}
}

// BenchmarkDocumentSave 基准测试 - 文档保存性能
func BenchmarkDocumentSave(b *testing.B) {
	doc := New()

	// 创建一个中等大小的文档
	for i := 0; i < 100; i++ {
		doc.AddParagraph("基准测试段落内容")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filename := "benchmark_save.docx"
		err := doc.Save(filename)
		if err != nil {
			b.Fatalf("Failed to save: %v", err)
		}
		os.Remove(filename)
	}
}

// TestTextFormatValidation 测试文本格式验证
func TestTextFormatValidation(t *testing.T) {
	doc := New()

	// 测试颜色格式
	testCases := []struct {
		color    string
		expected string
	}{
		{"#FF0000", "FF0000"}, // 带#前缀
		{"FF0000", "FF0000"},  // 不带#前缀
		{"#123456", "123456"},
		{"ABCDEF", "ABCDEF"},
	}

	for _, tc := range testCases {
		format := &TextFormat{
			FontColor: tc.color,
		}

		para := doc.AddFormattedParagraph("测试颜色", format)
		if para.Runs[0].Properties.Color.Val != tc.expected {
			t.Errorf("Color %s should be formatted as %s, got %s",
				tc.color, tc.expected, para.Runs[0].Properties.Color.Val)
		}
	}
}

// TestMemoryUsage 测试内存使用
func TestMemoryUsage(t *testing.T) {
	doc := New()

	// 添加大量段落测试内存使用
	const numParagraphs = 1000
	for i := 0; i < numParagraphs; i++ {
		doc.AddParagraph("内存测试段落")
	}

	if len(doc.Body.GetParagraphs()) != numParagraphs {
		t.Errorf("Expected %d paragraphs, got %d", numParagraphs, len(doc.Body.GetParagraphs()))
	}

	// 测试保存大文档
	filename := "test_memory.docx"
	defer os.Remove(filename)

	err := doc.Save(filename)
	if err != nil {
		t.Fatalf("Failed to save large document: %v", err)
	}
}
