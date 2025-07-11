package parser

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type paragraphParser struct {
}

var defaultParagraphParser = &paragraphParser{}

// NewParagraphParser returns a new BlockParser that
// parses paragraphs.
func NewParagraphParser() BlockParser {
	return defaultParagraphParser
}

func (b *paragraphParser) Trigger() []byte {
	return nil
}

func (b *paragraphParser) Open(parent ast.Node, reader text.Reader, pc Context) (ast.Node, State) {
	_, segment := reader.PeekLine()
	segment = segment.TrimLeftSpace(reader)
	if segment.IsEmpty() {
		return nil, NoChildren
	}
	node := ast.NewParagraph()
	node.Lines().Append(segment)
	reader.AdvanceToEOL()
	return node, NoChildren
}

func (b *paragraphParser) Continue(block *Block, reader text.Reader, pc Context) State {
	node := block.Node
	line, segment := reader.PeekLine()
	if util.IsBlank(line) {
		return Close
	}
	if block.Length+segment.Len() > pc.SizeLimit() {
		newNode := ast.NewParagraph()
		newNode.Lines().Append(segment)
		reader.AdvanceToEOL()
		if node.Parent() != nil {
			node.Parent().AppendChild(node.Parent(), newNode)
		}
		be := Block{Node: newNode, Parser: block.Parser, Length: segment.Len(), Start: segment.Start, BufferWindow: reader.BufferWindow()}
		pc.SetOpenedBlocks(append(pc.OpenedBlocks(), be))
		return Close
	}
	node.Lines().Append(segment)
	block.setLength(block.Length + segment.Len())
	reader.AdvanceToEOL()
	return Continue | NoChildren
}

func (b *paragraphParser) Close(block *Block, reader text.Reader, pc Context) {
	node := block.Node
	lines := node.Lines()
	if lines.Len() != 0 {
		// trim leading spaces
		for i := 0; i < lines.Len(); i++ {
			l := lines.At(i)
			lines.Set(i, l.TrimLeftSpace(reader))
		}

		// trim trailing spaces
		length := lines.Len()
		lastLine := node.Lines().At(length - 1)
		node.Lines().Set(length-1, lastLine.TrimRightSpace(reader))
	}
	if lines.Len() == 0 {
		node.Parent().RemoveChild(node.Parent(), node)
		return
	}
}

func (b *paragraphParser) CanInterruptParagraph() bool {
	return false
}

func (b *paragraphParser) CanAcceptIndentedLine() bool {
	return false
}
