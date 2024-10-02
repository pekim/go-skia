package generate

import (
	"slices"
	"strings"
)

type docComment struct {
	lines []string
}

func makeDocComment(rawComment string) string {
	// fmt.Println(rawComment)
	dc := docComment{
		lines: strings.Split(rawComment, "\n"),
	}
	dc.removeCommentDelimiters()
	dc.trimLinesWhitespace()
	// dc.orderedLists() TODO
	dc.unorderedLists()
	dc.params()
	dc.return_()

	return "/*" +
		strings.Join(dc.lines, "\n") +
		"*/"
}

func (dc *docComment) removeCommentDelimiters() {
	if len(dc.lines) == 0 {
		return
	}

	dc.lines[0] = strings.TrimPrefix(dc.lines[0], "/**")
	dc.lines[0] = strings.TrimPrefix(dc.lines[0], "/*")

	last := len(dc.lines) - 1
	dc.lines[last] = strings.TrimSuffix(dc.lines[last], "*/")
}

func (dc *docComment) trimLinesWhitespace() {
	for i, line := range dc.lines {
		dc.lines[i] = strings.TrimSpace(line)
	}
}

func (dc *docComment) unorderedLists() {
	for i, line := range dc.lines {
		if strings.HasPrefix(line, "- ") {
			dc.lines[i] = "  " + line
		}
	}
}

func (dc *docComment) params() {
	// Add params heading.
	insertedHeading := false
	for i, line := range dc.lines {
		if strings.HasPrefix(line, "@param ") {
			if !insertedHeading {
				dc.lines = slices.Insert(dc.lines, i, []string{"# params", ""}...)
				insertedHeading = true
			}
		}
	}

	// Make params a bullet list.
	for i, line := range dc.lines {
		if strings.HasPrefix(line, "@param ") {
			line = strings.TrimPrefix(line, "@param ")
			paramNameEndIndex := strings.Index(line, " ")
			paramName := line[:paramNameEndIndex]
			restOfLine := line[paramNameEndIndex:]
			line = paramName + " =>" + restOfLine
			dc.lines[i] = "  - " + line
		}
	}
}

func (dc *docComment) return_() {
	for i, line := range dc.lines {
		if strings.HasPrefix(line, "@return ") {
			line = strings.TrimPrefix(line, "@return ")
			line = strings.TrimSpace(line)
			line = "  - " + line
			dc.lines[i] = line

			dc.lines = slices.Insert(dc.lines, i, []string{"# return", ""}...)
			break
		}
	}
}

// type docComment struct {
// 	sections []string
// }

// func makeDocComment(comment clang.Comment) string {
// 	dc := docComment{}

// 	for i := range comment.NumChildren() {
// 		child := comment.Child(uint32(i))
// 		switch child.Kind() {
// 		case clang.Comment_Paragraph:
// 			dc.addParagraph(child)

// 		case clang.Comment_ParamCommand:
// 			// fmt.Println("  ", child.ParamCommandComment_getParamName())
// 			// for i := range child.NumChildren() {
// 			// 	paramChild := child.Child(i)
// 			// 	if paramChild.Kind() == clang.Comment_Paragraph {
// 			// 		for i := range paramChild.NumChildren() {
// 			// 			paraChild := paramChild.Child(i)
// 			// 			fmt.Println("    ", i, paraChild.Kind().Spelling())
// 			// 			fmt.Println("      ", paraChild.TextComment_getText())
// 			// 		}
// 			// 	}
// 			// }

// 		case clang.Comment_VerbatimLine:
// 			// fmt.Println(child.VerbatimLineComment_getText())

// 		case clang.Comment_BlockCommand:
// 			// fmt.Println(child.BlockCommandComment_getParagraph(), child.BlockCommandComment_getCommandName())
// 			// fmt.Println(i, "  block::", child.BlockCommandComment_getArgText(i))

// 		default:
// 			fatalf("Unsupported comment child of kind %s", child.Kind().Spelling())
// 		}
// 	}

// 	return "/*\n" + strings.Join(dc.sections, "\n\n") + "\n*/"
// }

// func (dc *docComment) addParagraph(para clang.Comment) {
// 	var text strings.Builder

// 	for i := range para.NumChildren() {
// 		child := para.Child(i)
// 		switch child.Kind() {
// 		case clang.Comment_Text:
// 			text.WriteString(child.TextComment_getText())

// 		case clang.Comment_InlineCommand:
// 			text.WriteString(fmt.Sprintf("@%s", child.InlineCommandComment_getCommandName()))

// 		default:
// 			fatalf("Unsupported comment paragraph child of kind %s", child.Kind().Spelling())
// 		}
// 	}

// 	text.WriteRune('\n')
// 	dc.sections = append(dc.sections, strings.TrimSpace(text.String()))
// }
