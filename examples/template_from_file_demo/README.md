# 从现有DOCX模板文件生成新文档

本示例演示了WordZero的核心模板功能：**从现有的Word模板文件（.docx）加载模板，然后填充数据生成新的文档**。

## 🎯 功能特点

✅ **支持现有DOCX文件**: 可以直接使用已存在的Word文档作为模板  
✅ **完整模板语法**: 支持变量替换、条件语句、循环语句等  
✅ **保持格式**: 保留原文档的格式和样式  
✅ **动态内容**: 根据数据动态生成文档内容  

## 📝 工作流程

```mermaid
graph TD
    A["现有DOCX模板文件"] --> B["document.Open()"]
    B --> C["engine.LoadTemplateFromDocument()"]
    C --> D["准备模板数据"]
    D --> E["engine.RenderToDocument()"]
    E --> F["生成新的DOCX文档"]
```

### 1. 打开现有DOCX模板
```go
// 打开现有的Word模板文件
templateDoc, err := document.Open("path/to/template.docx")
if err != nil {
    log.Fatal(err)
}
```

### 2. 从文档创建模板
```go
// 创建模板引擎
engine := document.NewTemplateEngine()

// 从文档创建模板
template, err := engine.LoadTemplateFromDocument("template_name", templateDoc)
if err != nil {
    log.Fatal(err)
}
```

### 3. 准备数据并渲染
```go
// 创建模板数据
data := document.NewTemplateData()
data.SetVariable("name", "张三")
data.SetVariable("company", "WordZero科技")

// 设置列表数据
items := []interface{}{
    map[string]interface{}{"product": "产品A", "price": 100},
    map[string]interface{}{"product": "产品B", "price": 200},
}
data.SetList("items", items)

// 设置条件
data.SetCondition("showDiscount", true)

// 渲染生成新文档（推荐方法）
newDoc, err := engine.RenderTemplateToDocument("template_name", data)
if err != nil {
    log.Fatal(err)
}

// 保存新文档
err = newDoc.Save("output/generated_document.docx")
```

## 🔧 支持的模板语法

### 变量替换
在DOCX模板中使用`{{变量名}}`来定义占位符：
```
客户姓名：{{customerName}}
联系电话：{{phone}}
```

### 条件语句
使用`{{#if 条件}}...{{/if}}`语法：
```
{{#if isVip}}
🎖️ 尊贵的VIP客户
{{/if}}
```

### 循环语句
使用`{{#each 列表}}...{{/each}}`语法：
```
商品清单：
{{#each items}}
{{@index}}. {{name}} - {{price}}元
{{/each}}
```

### 循环上下文变量
- `{{this}}`: 当前项的值
- `{{@index}}`: 当前索引（从0开始）
- `{{@first}}`: 是否第一项
- `{{@last}}`: 是否最后一项

## 📊 示例演示

本示例包含两个完整的演示：

### 演示1：商业发票生成
- **模板**: 包含发票的完整结构和格式
- **数据**: 出票方、收票方、商品明细、费用计算
- **功能**: 条件显示（折扣、税费）、商品列表循环

### 演示2：项目报告生成  
- **模板**: 项目进度报告结构
- **数据**: 团队成员、里程碑、成就、风险
- **功能**: 复杂的嵌套循环和条件判断

## 🚀 运行示例

1. **确保Go环境**: 需要Go 1.16+

2. **运行演示**:
```bash
cd examples/template_from_file_demo
go run main.go
```

3. **查看结果**: 
   - 模板文件: `examples/output/invoice_template.docx`
   - 生成的发票: `examples/output/generated_invoice_*.docx`
   - 生成的报告: `examples/output/generated_project_report_*.docx`

## 💡 实际应用场景

### 1. 业务文档自动化
- **发票生成**: 从ERP系统数据自动生成发票
- **合同生成**: 基于客户信息和产品配置生成合同
- **报告生成**: 自动生成周报、月报、季报

### 2. 个性化文档
- **邮件模板**: 根据用户信息生成个性化邮件
- **证书生成**: 批量生成培训证书、奖状等
- **通知函**: 根据不同场景生成通知文档

### 3. 数据驱动文档
- **财务报表**: 从数据库数据生成财务报告
- **库存报告**: 自动生成库存状况报告
- **客户报告**: 为每个客户生成专属服务报告

## 🔍 高级用法

### 从外部文件加载模板
```go
// 如果你已经有现成的Word模板文件
templateDoc, err := document.Open("templates/my_template.docx")
if err != nil {
    log.Fatal(err)
}

engine := document.NewTemplateEngine()
template, err := engine.LoadTemplateFromDocument("my_template", templateDoc)
```

### 批量生成文档
```go
// 批量数据
customers := []CustomerData{
    {Name: "张三", Phone: "138-0000-0001"},
    {Name: "李四", Phone: "138-0000-0002"},
    // ... 更多客户
}

// 为每个客户生成文档
for i, customer := range customers {
    data := document.NewTemplateData()
    err := data.FromStruct(customer)
    if err != nil {
        continue
    }
    
    doc, err := engine.RenderTemplateToDocument("customer_template", data)
    if err != nil {
        continue
    }
    
    filename := fmt.Sprintf("output/customer_%d.docx", i)
    doc.Save(filename)
}
```

### 结构体数据绑定
```go
type Invoice struct {
    Number    string
    Date      string
    Customer  string
    Amount    float64
    IsPaid    bool
}

invoice := Invoice{
    Number:   "INV-001",
    Date:     "2024-12-01", 
    Customer: "张三",
    Amount:   1000.00,
    IsPaid:   false,
}

data := document.NewTemplateData()
err := data.FromStruct(invoice)  // 自动转换结构体字段为模板变量
```

## ⚠️ 注意事项

1. **模板语法**: 确保DOCX模板中的语法正确，括号要配对
2. **文件路径**: 确保模板文件路径正确且文件存在
3. **数据类型**: 注意数据类型的匹配，特别是条件判断
4. **文件权限**: 确保有读取模板文件和写入输出文件的权限

## 📚 相关文档

- [模板功能详细教程](../../wordZero.wiki/12-模板功能.md)
- [API参考文档](../../pkg/document/README.md)
- [更多示例](../template_demo/)

---

这个功能让WordZero能够与现有的Word模板工作流无缝集成，大大提高了文档自动化的灵活性和实用性！ 