package test

import (
	"os"
	"testing"

	"github.com/ZeroHawkeye/wordZero/pkg/document"
)

func TestCreateDocument(t *testing.T) {
	// 创建新文档
	doc := document.New()

	// 添加段落
	doc.AddParagraph("Hello, World!")
	doc.AddParagraph("这是一个使用 WordZero 创建的 Word 文档。")

	// 保存文档
	err := doc.Save("test_output/test_document.docx")
	if err != nil {
		t.Fatalf("Failed to save document: %v", err)
	}

	// 检查文件是否存在
	if _, err := os.Stat("test_output/test_document.docx"); os.IsNotExist(err) {
		t.Fatal("Document file was not created")
	}

	// 清理测试文件
	defer os.RemoveAll("test_output")

	t.Log("Document created successfully")
}

func TestOpenDocument(t *testing.T) {
	// 首先创建一个文档
	doc := document.New()
	doc.AddParagraph("Test paragraph")

	testFile := "test_output/test_open.docx"
	err := doc.Save(testFile)
	if err != nil {
		t.Fatalf("Failed to save document: %v", err)
	}

	// 打开文档
	openedDoc, err := document.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open document: %v", err)
	}

	// 验证内容
	paragraphs := openedDoc.Body.GetParagraphs()
	if len(paragraphs) != 1 {
		t.Fatalf("Expected 1 paragraph, got %d", len(paragraphs))
	}

	if paragraphs[0].Runs[0].Text.Content != "Test paragraph" {
		t.Fatalf("Paragraph content mismatch")
	}

	// 清理测试文件
	defer os.RemoveAll("test_output")

	t.Log("Document opened successfully")
}
