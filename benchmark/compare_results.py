#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
性能测试结果对比分析工具

读取Golang、JavaScript和Python的性能测试结果，
生成对比图表和详细分析报告。
"""

import json
import os
import sys
from pathlib import Path
from typing import Dict, List, Any
import matplotlib.pyplot as plt
import pandas as pd
import seaborn as sns
from datetime import datetime

# 设置中文字体
plt.rcParams['font.sans-serif'] = ['SimHei', 'Arial Unicode MS', 'DejaVu Sans']
plt.rcParams['axes.unicode_minus'] = False

class PerformanceAnalyzer:
    """性能分析器"""
    
    def __init__(self, results_dir: str = "results"):
        # 确保找到正确的results目录
        current_dir = Path.cwd()
        if current_dir.name == "python":
            # 如果当前在 python 目录，向上一级找 results
            self.results_dir = current_dir.parent / results_dir
        else:
            self.results_dir = Path(results_dir)
        self.golang_results = {}
        self.javascript_results = {}
        self.python_results = {}
        
    def load_results(self):
        """加载所有测试结果"""
        try:
            # 加载Golang结果（可能需要从其他格式转换）
            golang_path = self.results_dir / "golang" / "performance_report.json"
            if golang_path.exists():
                with open(golang_path, 'r', encoding='utf-8') as f:
                    self.golang_results = json.load(f)
            
            # 加载JavaScript结果
            js_path = self.results_dir / "javascript" / "performance_report.json"
            if js_path.exists():
                with open(js_path, 'r', encoding='utf-8') as f:
                    self.javascript_results = json.load(f)
            
            # 加载Python结果
            python_path = self.results_dir / "python" / "performance_report.json"
            if python_path.exists():
                with open(python_path, 'r', encoding='utf-8') as f:
                    self.python_results = json.load(f)
                    
            print("测试结果加载完成")
            
        except Exception as e:
            print(f"加载测试结果时发生错误: {e}")
    
    def create_comparison_dataframe(self) -> pd.DataFrame:
        """创建对比数据框"""
        data = []
        
        # 处理各个平台的结果
        platforms = {
            'Golang': self.golang_results,
            'JavaScript': self.javascript_results,
            'Python': self.python_results
        }
        
        for platform, results in platforms.items():
            if results and 'results' in results:
                for result in results['results']:
                    data.append({
                        'Platform': platform,
                        'Test': result['name'],
                        'AvgTime': float(result['avgTime']),
                        'MinTime': float(result['minTime']),
                        'MaxTime': float(result['maxTime']),
                        'Iterations': result['iterations']
                    })
        
        return pd.DataFrame(data)
    
    def generate_comparison_charts(self, df: pd.DataFrame):
        """生成对比图表"""
        if df.empty:
            print("没有数据可供生成图表")
            return
        
        # 创建图表目录
        charts_dir = self.results_dir / "charts"
        charts_dir.mkdir(exist_ok=True)
        
        # 1. 平均时间对比条形图
        plt.figure(figsize=(15, 8))
        pivot_df = df.pivot(index='Test', columns='Platform', values='AvgTime')
        ax = pivot_df.plot(kind='bar', width=0.8)
        plt.title('各平台平均执行时间对比 (毫秒)', fontsize=16, fontweight='bold')
        plt.xlabel('测试项目', fontsize=12)
        plt.ylabel('平均执行时间 (ms)', fontsize=12)
        plt.xticks(rotation=45, ha='right')
        plt.legend(title='平台')
        plt.grid(axis='y', alpha=0.3)
        plt.tight_layout()
        plt.savefig(charts_dir / 'avg_time_comparison.png', dpi=300, bbox_inches='tight')
        plt.close()
        
        # 2. 性能比率图（以Golang为基准）
        if 'Golang' in pivot_df.columns:
            plt.figure(figsize=(15, 8))
            ratio_df = pivot_df.div(pivot_df['Golang'], axis=0)
            ratio_df.plot(kind='bar', width=0.8)
            plt.title('相对性能比率 (以Golang为基准=1.0)', fontsize=16, fontweight='bold')
            plt.xlabel('测试项目', fontsize=12)
            plt.ylabel('性能比率', fontsize=12)
            plt.xticks(rotation=45, ha='right')
            plt.legend(title='平台')
            plt.axhline(y=1.0, color='red', linestyle='--', alpha=0.7, label='Golang基准')
            plt.grid(axis='y', alpha=0.3)
            plt.tight_layout()
            plt.savefig(charts_dir / 'performance_ratio.png', dpi=300, bbox_inches='tight')
            plt.close()
        
        # 3. 热力图
        plt.figure(figsize=(12, 8))
        sns.heatmap(pivot_df, annot=True, fmt='.1f', cmap='YlOrRd', 
                    cbar_kws={'label': '执行时间 (ms)'})
        plt.title('性能热力图', fontsize=16, fontweight='bold')
        plt.xlabel('平台', fontsize=12)
        plt.ylabel('测试项目', fontsize=12)
        plt.tight_layout()
        plt.savefig(charts_dir / 'performance_heatmap.png', dpi=300, bbox_inches='tight')
        plt.close()
        
        # 4. 箱线图（显示性能分布）
        plt.figure(figsize=(15, 8))
        melted_df = df.melt(id_vars=['Platform', 'Test'], 
                           value_vars=['MinTime', 'AvgTime', 'MaxTime'],
                           var_name='Metric', value_name='Time')
        sns.boxplot(data=melted_df, x='Test', y='Time', hue='Platform')
        plt.title('各测试项目性能分布对比', fontsize=16, fontweight='bold')
        plt.xlabel('测试项目', fontsize=12)
        plt.ylabel('执行时间 (ms)', fontsize=12)
        plt.xticks(rotation=45, ha='right')
        plt.legend(title='平台')
        plt.grid(axis='y', alpha=0.3)
        plt.tight_layout()
        plt.savefig(charts_dir / 'performance_distribution.png', dpi=300, bbox_inches='tight')
        plt.close()
        
        print(f"图表已保存到: {charts_dir}")
    
    def generate_detailed_report(self, df: pd.DataFrame):
        """生成详细分析报告"""
        if df.empty:
            print("没有数据可供生成报告")
            return
        
        report_path = self.results_dir / "detailed_comparison_report.md"
        
        with open(report_path, 'w', encoding='utf-8') as f:
            f.write("# WordZero 跨语言性能对比分析报告\n\n")
            f.write(f"生成时间: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}\n\n")
            
            # 系统信息
            f.write("## 测试环境\n\n")
            for platform, results in [('Golang', self.golang_results), 
                                     ('JavaScript', self.javascript_results),
                                     ('Python', self.python_results)]:
                if results:
                    f.write(f"- **{platform}**: ")
                    if 'nodeVersion' in results:
                        f.write(f"Node.js {results['nodeVersion']}\n")
                    elif 'pythonVersion' in results:
                        f.write(f"Python {results['pythonVersion']}\n")
                    elif platform == 'Golang':
                        f.write("Go 1.19+\n")
                    else:
                        f.write("未知版本\n")
            
            f.write("\n## 性能对比摘要\n\n")
            
            # 计算平均性能
            platform_avg = df.groupby('Platform')['AvgTime'].mean()
            fastest_platform = platform_avg.idxmin()
            f.write(f"- **总体最快平台**: {fastest_platform} (平均 {platform_avg[fastest_platform]:.2f}ms)\n")
            
            # 各测试项目最快平台
            f.write("- **各测试项目最快平台**:\n")
            for test in df['Test'].unique():
                test_data = df[df['Test'] == test]
                fastest = test_data.loc[test_data['AvgTime'].idxmin()]
                f.write(f"  - {test}: {fastest['Platform']} ({fastest['AvgTime']:.2f}ms)\n")
            
            f.write("\n## 详细测试结果\n\n")
            
            # 创建详细表格
            pivot_df = df.pivot(index='Test', columns='Platform', values='AvgTime')
            f.write("### 平均执行时间对比 (毫秒)\n\n")
            f.write(pivot_df.to_markdown())
            f.write("\n\n")
            
            # 性能比率分析（以最快的平台为基准）
            f.write("### 相对性能分析\n\n")
            if 'Golang' in pivot_df.columns:
                ratio_df = pivot_df.div(pivot_df['Golang'], axis=0)
                f.write("以Golang为基准的性能比率:\n\n")
                f.write(ratio_df.to_markdown())
                f.write("\n\n")
            
            # 性能建议
            f.write("## 性能建议\n\n")
            f.write("### 各语言优势分析\n\n")
            
            for platform in df['Platform'].unique():
                platform_data = df[df['Platform'] == platform]
                best_tests = []
                for test in df['Test'].unique():
                    test_data = df[df['Test'] == test]
                    if test_data.loc[test_data['AvgTime'].idxmin(), 'Platform'] == platform:
                        best_tests.append(test)
                
                f.write(f"**{platform}**:\n")
                if best_tests:
                    f.write(f"- 最适合: {', '.join(best_tests)}\n")
                else:
                    f.write("- 在所有测试项目中都不是最快的\n")
                
                avg_time = platform_data['AvgTime'].mean()
                f.write(f"- 平均性能: {avg_time:.2f}ms\n\n")
            
            f.write("### 选型建议\n\n")
            f.write("- **高性能需求**: 选择Golang实现，内存占用小，执行速度快\n")
            f.write("- **快速开发**: JavaScript/Node.js生态丰富，开发效率高\n")
            f.write("- **数据处理**: Python有丰富的数据处理库，适合复杂文档操作\n")
            f.write("- **跨平台**: 三种语言都支持跨平台，根据团队技术栈选择\n\n")
        
        print(f"详细报告已保存到: {report_path}")
    
    def run_analysis(self):
        """运行完整分析"""
        print("开始性能对比分析...")
        
        # 加载数据
        self.load_results()
        
        # 创建对比数据框
        df = self.create_comparison_dataframe()
        
        if df.empty:
            print("未找到测试结果数据，请先运行各平台的性能测试")
            return
        
        print(f"成功加载 {len(df)} 条测试结果")
        
        # 生成图表
        try:
            self.generate_comparison_charts(df)
        except Exception as e:
            print(f"生成图表时发生错误: {e}")
        
        # 生成详细报告
        self.generate_detailed_report(df)
        
        # 打印简要对比结果
        print("\n=== 性能对比摘要 ===")
        pivot_df = df.pivot(index='Test', columns='Platform', values='AvgTime')
        print(pivot_df)
        
        print("\n分析完成！")


def main():
    """主函数"""
    analyzer = PerformanceAnalyzer()
    analyzer.run_analysis()


if __name__ == '__main__':
    main() 