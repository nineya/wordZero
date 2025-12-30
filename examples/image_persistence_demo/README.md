# Image Persistence Demo

This demo demonstrates the fix for the image disappearance issue when modifying and saving Word documents.

## Problem

Previously, when you:
1. Opened a Word document containing images
2. Modified the document (added text, paragraphs, etc.)
3. Saved it as a new file

The images would disappear from the new file.

## Solution

The issue has been fixed by:
- Parsing the `word/_rels/document.xml.rels` file when opening documents
- Preserving existing image relationships
- Updating the image ID counter to prevent conflicts

## Running the Demo

```bash
go run ./examples/image_persistence_demo/main.go
```

## What the Demo Does

1. **Creates an original document** with 2 images (red and blue)
2. **Opens the document** and verifies the images are loaded
3. **Modifies the document** by adding a new paragraph and a 3rd image (green)
4. **Saves the modified document** to a new file
5. **Verifies all 3 images** are present in the final document

## Output Files

The demo creates two files in `examples/output/`:
- `image_persistence_original.docx` - Original document with 2 images
- `image_persistence_modified.docx` - Modified document with all 3 images

You can open these files in Microsoft Word or any compatible word processor to verify that all images are correctly preserved.

## Expected Result

```
✓ 图片持久性测试成功！

演示结果：
  - 原始文档: examples/output/image_persistence_original.docx (包含2张图片)
  - 修改后的文档: examples/output/image_persistence_modified.docx (包含3张图片)

说明：
  1. 原有的2张图片（红色和蓝色）在打开和重新保存后没有丢失
  2. 新添加的1张图片（绿色）正确保存
  3. 所有图片在Word中可以正常显示
```

## Related Issue

This demo addresses the issue: "修改后图片消失了" (Images disappear after modification)

## Technical Details

The fix involves:
1. **parseDocumentRelationships()**: Parses existing image and resource relationships
2. **updateNextImageID()**: Updates the image ID counter based on existing images
3. **Enhanced Open() and OpenFromMemory()**: Both functions now preserve image relationships

See the commit message and PR description for more details.
