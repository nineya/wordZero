# WordZero 性能基准测试

对Golang WordZero库、JavaScript docx库和Python python-docx库进行Word文档操作性能测试和对比分析。

## 测试配置统一化

为了确保公平对比，所有语言的测试都使用相同的迭代次数：

| 测试项目 | 迭代次数 | 说明 |
|---------|---------|------|
| 基础文档创建 | 50次 | 创建包含基本文本的文档 |
| 复杂格式化 | 30次 | 创建包含多种格式化的文档 |
| 表格操作 | 20次 | 创建10行5列的表格 |
| 大表格处理 | 10次 | 创建100行10列的大表格 |
| 大型文档 | 5次 | 创建包含1000个段落+表格的文档 |
| 内存使用测试 | 10次 | 测试内存使用情况 |

## 快速开始

### 运行完整的性能对比测试（推荐）

```bash
# 安装所有依赖
task setup

# 运行固定迭代次数的对比测试
task benchmark-all-fixed

# 生成对比报告
task compare
```

### 运行单独的语言测试

```bash
# Golang 固定迭代次数测试（推荐用于对比）
task benchmark-golang-fixed

# Golang 标准基准测试（迭代次数由Go框架决定）
task benchmark-golang

# JavaScript 测试
task benchmark-js

# Python 测试
task benchmark-python
```

## 测试类型说明

### 1. Golang 测试

**固定迭代次数测试（推荐）**:
- 命令: `task benchmark-golang-fixed`
- 运行: `TestFixedIterationsPerformance`
- 与其他语言使用相同的迭代次数，确保公平对比

**Go标准基准测试**:
- 命令: `task benchmark-golang`
- 运行: `BenchmarkXXX` 系列函数
- 迭代次数由Go测试框架自动决定（通常是数千次）

### 2. JavaScript 测试

- 使用 `docx` 库
- 固定迭代次数，与其他语言保持一致
- 输出JSON格式的性能报告

### 3. Python 测试

- 使用 `python-docx` 库
- 固定迭代次数，与其他语言保持一致
- 支持pytest-benchmark和自定义测试

## 输出文件

测试完成后，会在 `results/` 目录下生成：

```
results/
├── golang/
│   ├── performance_report.json          # 固定迭代测试报告
│   ├── fixed_benchmark_output.txt       # 固定迭代测试日志
│   ├── benchmark_output.txt             # Go标准基准测试日志
│   └── *.docx                          # 生成的测试文档
├── javascript/
│   ├── performance_report.json          # JavaScript测试报告
│   └── *.docx                          # 生成的测试文档
├── python/
│   ├── performance_report.json          # Python测试报告
│   └── *.docx                          # 生成的测试文档
├── charts/                             # 性能对比图表
├── detailed_comparison_report.md        # 详细对比报告
└── performance_comparison.json          # 对比数据
```

## 性能指标

每个测试都会记录：
- **平均耗时** (avgTime): 所有迭代的平均执行时间
- **最小耗时** (minTime): 最快的一次执行时间
- **最大耗时** (maxTime): 最慢的一次执行时间
- **迭代次数** (iterations): 测试运行的次数

## 环境要求

### Golang
- Go 1.19+
- WordZero库依赖

### JavaScript
- Node.js 16+
- docx 库

### Python
- Python 3.8+
- python-docx 库
- 自动创建虚拟环境

## 测试最佳实践

1. **使用固定迭代次数测试进行对比**: `task benchmark-all-fixed`
2. **确保系统负载一致**: 测试期间避免运行其他大型程序
3. **多次运行取平均值**: 可以多次运行测试以获得更稳定的结果
4. **关注相对性能**: 重点关注各语言间的相对性能差异

## 故障排除

### Python SSL证书问题
```bash
task setup-python-ssl-fix
```

### 清理生成的文件
```bash
task clean
```

### 完整重置环境
```bash
task clean-all
task setup
```

## 测试说明

- 所有测试都包含文档创建和文件保存操作
- 测试文件保存在各自的results目录中
- 内存测试会输出内存使用情况
- 每种语言使用其生态系统中最流行的Word操作库 