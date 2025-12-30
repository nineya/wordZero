# 段落格式自定义功能演示

本示例演示了WordZero库提供的全面段落格式自定义功能。

## 功能特性

### 分页控制
- **SetKeepWithNext** - 确保段落与下一段落保持在同一页
- **SetKeepLines** - 防止段落被分页拆分
- **SetPageBreakBefore** - 在段落前强制插入分页符

### 其他高级功能
- **SetWidowControl** - 孤行控制，提升排版质量
- **SetOutlineLevel** - 大纲级别设置，便于文档导航
- **SetParagraphFormat** - 一次性设置多个格式属性

## 运行示例

```bash
go run ./examples/paragraph_format_demo/main.go
```

运行后将生成 `examples/output/paragraph_format_demo.docx` 文件。

## 示例内容

演示文档包含以下场景：

1. **章节标题格式** - 使用居中对齐、段前分页、与下段保持、大纲级别
2. **正文段落格式** - 使用两端对齐、行间距、首行缩进、孤行控制
3. **二级标题格式** - 使用段前段后间距、与下段保持、大纲级别
4. **悬挂缩进列表** - 演示负首行缩进和左缩进组合
5. **引用块格式** - 使用左右缩进、行保持、居中对齐
6. **多级标题** - 演示不同大纲级别的设置

## API使用方式

### 方法1：单独设置各个属性

```go
title := doc.AddParagraph("第一章 概述")
title.SetAlignment(document.AlignCenter)
title.SetKeepWithNext(true)
title.SetPageBreakBefore(true)
title.SetOutlineLevel(0)
```

### 方法2：使用ParagraphFormatConfig一次性设置

```go
para := doc.AddParagraph("重要内容")
para.SetParagraphFormat(&document.ParagraphFormatConfig{
    Alignment:       document.AlignJustify,
    Style:           "Normal",
    LineSpacing:     1.5,
    BeforePara:      12,
    AfterPara:       6,
    FirstLineCm:     0.5,
    KeepWithNext:    true,
    KeepLines:       true,
    WidowControl:    true,
    OutlineLevel:    0,
})
```

## 应用场景

- **文档结构化** - 使用大纲级别创建清晰的文档层次结构
- **专业排版** - 使用分页控制确保标题和内容的关联性
- **内容保护** - 使用行保持防止重要段落被分页
- **章节管理** - 使用段前分页实现章节的页面独立性

## 相关文档

- [pkg/document API文档](../../pkg/document/README.md)
- [段落格式设置文档](../../pkg/document/doc.go)
