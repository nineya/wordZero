const fs = require('fs-extra');
const path = require('path');
const { Document, Packer, Paragraph, TextRun, Table, TableRow, TableCell, WidthType, AlignmentType } = require('docx');

// 确保输出目录存在
const outputDir = '../results/javascript';
fs.ensureDirSync(outputDir);

// 性能测试工具类
class PerformanceTester {
    constructor() {
        this.results = [];
    }

    async runTest(name, testFunction, iterations = 10) {
        console.log(`\n开始测试: ${name}`);
        const times = [];
        
        for (let i = 0; i < iterations; i++) {
            const start = process.hrtime.bigint();
            await testFunction(i);
            const end = process.hrtime.bigint();
            
            const duration = Number(end - start) / 1000000; // 转换为毫秒
            times.push(duration);
            
            if (i % Math.ceil(iterations / 10) === 0) {
                console.log(`  进度: ${i + 1}/${iterations}`);
            }
        }
        
        const avgTime = times.reduce((a, b) => a + b, 0) / times.length;
        const minTime = Math.min(...times);
        const maxTime = Math.max(...times);
        
        const result = {
            name,
            avgTime: avgTime.toFixed(2),
            minTime: minTime.toFixed(2),
            maxTime: maxTime.toFixed(2),
            iterations
        };
        
        this.results.push(result);
        console.log(`  平均耗时: ${result.avgTime}ms`);
        console.log(`  最小耗时: ${result.minTime}ms`);
        console.log(`  最大耗时: ${result.maxTime}ms`);
        
        return result;
    }

    generateReport() {
        console.log('\n=== JavaScript 性能测试报告 ===');
        console.table(this.results);
        
        // 保存详细报告
        const reportPath = path.join(outputDir, 'performance_report.json');
        fs.writeJsonSync(reportPath, {
            timestamp: new Date().toISOString(),
            platform: 'JavaScript (Node.js)',
            nodeVersion: process.version,
            results: this.results
        }, { spaces: 2 });
        
        console.log(`\n详细报告已保存到: ${reportPath}`);
    }
}

// 基础文档创建测试
async function testBasicDocumentCreation(index) {
    const doc = new Document({
        sections: [{
            properties: {},
            children: [
                new Paragraph({
                    children: [
                        new TextRun("这是一个基础性能测试文档")
                    ]
                }),
                new Paragraph({
                    children: [
                        new TextRun("测试内容包括基本的文本添加功能")
                    ]
                })
            ]
        }]
    });

    const buffer = await Packer.toBuffer(doc);
    const filename = path.join(outputDir, `basic_doc_${index}.docx`);
    await fs.writeFile(filename, buffer);
}

// 复杂格式化测试
async function testComplexFormatting(index) {
    const doc = new Document({
        sections: [{
            properties: {},
            children: [
                new Paragraph({
                    children: [
                        new TextRun({
                            text: "性能测试报告",
                            bold: true,
                            size: 32,
                            color: "2E74B5"
                        })
                    ],
                    heading: "Heading1"
                }),
                new Paragraph({
                    children: [
                        new TextRun({
                            text: "测试概述",
                            bold: true,
                            size: 26,
                            color: "2E74B5"
                        })
                    ],
                    heading: "Heading2"
                }),
                new Paragraph({
                    children: [
                        new TextRun({
                            text: "粗体文本",
                            bold: true
                        }),
                        new TextRun(" "),
                        new TextRun({
                            text: "斜体文本",
                            italics: true
                        }),
                        new TextRun(" "),
                        new TextRun({
                            text: "彩色文本",
                            color: "FF0000"
                        })
                    ]
                })
            ]
        }]
    });

    // 添加多个格式化段落
    for (let j = 0; j < 10; j++) {
        doc.addSection({
            children: [
                new Paragraph({
                    children: [
                        new TextRun(`这是第${j + 1}个段落，包含复杂格式化`)
                    ],
                    alignment: AlignmentType.CENTER
                })
            ]
        });
    }

    const buffer = await Packer.toBuffer(doc);
    const filename = path.join(outputDir, `complex_formatting_${index}.docx`);
    await fs.writeFile(filename, buffer);
}

// 表格操作测试
async function testTableOperations(index) {
    const rows = [];
    
    // 创建10行5列的表格
    for (let row = 0; row < 10; row++) {
        const cells = [];
        for (let col = 0; col < 5; col++) {
            cells.push(new TableCell({
                children: [
                    new Paragraph({
                        children: [
                            new TextRun(`R${row + 1}C${col + 1}`)
                        ]
                    })
                ]
            }));
        }
        rows.push(new TableRow({ children: cells }));
    }

    const table = new Table({
        rows: rows,
        width: {
            size: 100,
            type: WidthType.PERCENTAGE
        }
    });

    const doc = new Document({
        sections: [{
            properties: {},
            children: [
                new Paragraph({
                    children: [
                        new TextRun({
                            text: "表格性能测试",
                            bold: true,
                            size: 32
                        })
                    ],
                    heading: "Heading1"
                }),
                table
            ]
        }]
    });

    const buffer = await Packer.toBuffer(doc);
    const filename = path.join(outputDir, `table_operations_${index}.docx`);
    await fs.writeFile(filename, buffer);
}

// 大表格测试
async function testLargeTable(index) {
    const rows = [];
    
    // 创建100行10列的大表格
    for (let row = 0; row < 100; row++) {
        const cells = [];
        for (let col = 0; col < 10; col++) {
            cells.push(new TableCell({
                children: [
                    new Paragraph({
                        children: [
                            new TextRun(`数据_${row + 1}_${col + 1}`)
                        ]
                    })
                ]
            }));
        }
        rows.push(new TableRow({ children: cells }));
    }

    const table = new Table({
        rows: rows,
        width: {
            size: 100,
            type: WidthType.PERCENTAGE
        }
    });

    const doc = new Document({
        sections: [{
            properties: {},
            children: [
                new Paragraph({
                    children: [
                        new TextRun({
                            text: "大表格性能测试",
                            bold: true,
                            size: 32
                        })
                    ],
                    heading: "Heading1"
                }),
                table
            ]
        }]
    });

    const buffer = await Packer.toBuffer(doc);
    const filename = path.join(outputDir, `large_table_${index}.docx`);
    await fs.writeFile(filename, buffer);
}

// 大型文档测试
async function testLargeDocument(index) {
    const children = [];
    
    children.push(new Paragraph({
        children: [
            new TextRun({
                text: "大型文档性能测试",
                bold: true,
                size: 32
            })
        ],
        heading: "Heading1"
    }));

    // 添加1000个段落
    for (let j = 0; j < 1000; j++) {
        if (j % 10 === 0) {
            // 每10个段落添加一个标题
            children.push(new Paragraph({
                children: [
                    new TextRun({
                        text: `章节 ${Math.floor(j / 10) + 1}`,
                        bold: true,
                        size: 26
                    })
                ],
                heading: "Heading2"
            }));
        }
        
        children.push(new Paragraph({
            children: [
                new TextRun(`这是第${j + 1}个段落。Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`)
            ]
        }));
    }

    // 添加一个中等大小的表格
    const tableRows = [];
    for (let row = 0; row < 20; row++) {
        const cells = [];
        for (let col = 0; col < 8; col++) {
            cells.push(new TableCell({
                children: [
                    new Paragraph({
                        children: [
                            new TextRun(`表格数据${row + 1}-${col + 1}`)
                        ]
                    })
                ]
            }));
        }
        tableRows.push(new TableRow({ children: cells }));
    }

    children.push(new Table({
        rows: tableRows,
        width: {
            size: 100,
            type: WidthType.PERCENTAGE
        }
    }));

    const doc = new Document({
        sections: [{
            properties: {},
            children: children
        }]
    });

    const buffer = await Packer.toBuffer(doc);
    const filename = path.join(outputDir, `large_document_${index}.docx`);
    await fs.writeFile(filename, buffer);
}

// 内存使用测试
async function testMemoryUsage(index) {
    const initialMemory = process.memoryUsage();
    
    const children = [];
    
    // 创建复杂内容
    for (let j = 0; j < 100; j++) {
        children.push(new Paragraph({
            children: [
                new TextRun(`段落${j + 1}: 这是一个测试段落，用于测试内存使用情况`)
            ]
        }));
    }

    // 添加表格
    const tableRows = [];
    for (let row = 0; row < 50; row++) {
        const cells = [];
        for (let col = 0; col < 6; col++) {
            cells.push(new TableCell({
                children: [
                    new Paragraph({
                        children: [
                            new TextRun(`单元格${row + 1}-${col + 1}`)
                        ]
                    })
                ]
            }));
        }
        tableRows.push(new TableRow({ children: cells }));
    }

    children.push(new Table({
        rows: tableRows,
        width: {
            size: 100,
            type: WidthType.PERCENTAGE
        }
    }));

    const doc = new Document({
        sections: [{
            properties: {},
            children: children
        }]
    });

    const buffer = await Packer.toBuffer(doc);
    const filename = path.join(outputDir, `memory_test_${index}.docx`);
    await fs.writeFile(filename, buffer);

    const finalMemory = process.memoryUsage();
    if (index === 0) {
        console.log(`  内存使用: ${Math.round((finalMemory.heapUsed - initialMemory.heapUsed) / 1024)}KB`);
    }
}

// 主测试函数
async function runBenchmarks() {
    console.log('开始 JavaScript (docx库) 性能基准测试...');
    console.log(`Node.js 版本: ${process.version}`);
    console.log(`输出目录: ${outputDir}`);
    
    const tester = new PerformanceTester();
    
    try {
        // 统一的测试配置，与其他语言保持一致
        const testIterations = {
            basic: 50,       // 基础文档创建
            complex: 30,     // 复杂格式化
            table: 20,       // 表格操作
            largeTable: 10,  // 大表格处理
            largeDoc: 5,     // 大型文档
            memory: 10,      // 内存使用测试
        };

        await tester.runTest('基础文档创建', testBasicDocumentCreation, testIterations.basic);
        await tester.runTest('复杂格式化', testComplexFormatting, testIterations.complex);
        await tester.runTest('表格操作', testTableOperations, testIterations.table);
        await tester.runTest('大表格处理', testLargeTable, testIterations.largeTable);
        await tester.runTest('大型文档', testLargeDocument, testIterations.largeDoc);
        await tester.runTest('内存使用测试', testMemoryUsage, testIterations.memory);
        
        tester.generateReport();
        
    } catch (error) {
        console.error('测试过程中发生错误:', error);
        process.exit(1);
    }
}

// 如果直接运行此文件，则执行基准测试
if (require.main === module) {
    runBenchmarks()
        .then(() => {
            console.log('\n所有测试完成！');
            process.exit(0);
        })
        .catch((error) => {
            console.error('基准测试失败:', error);
            process.exit(1);
        });
}

module.exports = {
    runBenchmarks,
    PerformanceTester,
    testBasicDocumentCreation,
    testComplexFormatting,
    testTableOperations,
    testLargeTable,
    testLargeDocument,
    testMemoryUsage
}; 