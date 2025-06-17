package goldmark_test

import (
	"testing"

	. "github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
)

func TestMarkdownLexerDebug(t *testing.T) {
	markdownContent := `# Heading 1
## Heading 2
This is a paragraph with **bold** and *italic* text.

### Heading 3

- List item 1
- List item 2

## Heading 4
> Blockquote

~~~go
fmt.Println("Hello, world!")
~~~
`
	source := []byte(markdownContent)

	markdown := New(
		WithExtensions(
			extension.Table,
			extension.GFM,
		),
	)

	//// 1. 打开文件
	//file, err := os.Open("README.md")
	//if err != nil {
	//	fmt.Println("打开文件失败:", err)
	//	return
	//}
	//defer file.Close() // 确保关闭文件句柄
	//
	//// 2. 创建带缓冲的读取器
	//reader1 := bufio.NewReader(file)
	//
	//// 3. 循环读取每一行
	//for {
	//	line, err := reader1.ReadString('\n') // 以换行符为分隔符
	//	if err != nil {
	//		// 4. 处理文件结束或错误
	//		if err == io.EOF {
	//			if line != "" { // 处理最后一行无换行符的情况
	//				fmt.Print(line)
	//			}
	//			break
	//		}
	//		fmt.Println("读取错误:", err)
	//		return
	//	}
	//	// 5. 处理每行内容（含结尾的\n）
	//	fmt.Print(line) // 直接输出或进行其他处理
	//}

	reader := text.NewReader(source)
	//ctx := parser.NewContext()
	//for {
	//	block := markdown.Parser().NextBlock(reader, parser.WithContext(ctx))
	//	block.Node.Dump(source, 0) // 直接输出或进行其他处理
	//}

	doc := markdown.Parser().Parse(reader)

	doc.Dump(source, 0)

	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		t.Logf("Node kind: %s", n.Kind())
		if n.Kind() == ast.KindText {
			t.Logf("  Text: %s", n.Text(source))
		}
		return ast.WalkContinue, nil
	})
}
