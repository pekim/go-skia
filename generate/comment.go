package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

func parsedCommentToGoComment(comment clang.Comment) string {
	var b strings.Builder

	for i := uint32(0); i < comment.NumChildren(); i++ {
		child := comment.Child(i)

		switch child.Kind() {
		case clang.Comment_Paragraph, clang.Comment_VerbatimLine:
			for i := uint32(0); i < child.NumChildren(); i++ {
				child := child.Child(i)
				switch child.Kind() {
				case clang.Comment_Text:
					b.WriteString(strings.TrimSpace(child.TextComment_getText()))
					b.WriteString("\n")

				case clang.Comment_VerbatimLine:
					b.WriteString(strings.TrimSpace(child.TextComment_getText()))
					b.WriteString("\n")

				case clang.Comment_InlineCommand:
					// ignore commands

				default:
					fmt.Printf("unhandled comment kind, %s\n", child.Kind().Spelling())
				}
			}

		default:
			fmt.Printf("unhandled comment kind, %s\n", child.Kind().Spelling())
		}

		if child.Kind() == clang.Comment_Paragraph {
			b.WriteString("\n")
		}
	}

	text := strings.TrimSpace(b.String())
	if len(text) > 0 {
		return "/*\n" + text + "*/"
	}
	return ""
}
