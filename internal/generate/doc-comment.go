package generate

import (
	"encoding/xml"
	"fmt"
	"strings"
	"unicode"

	"github.com/go-clang/clang-v15/clang"
)

// A Relax NG schema for the XML can be found in comment-xml-schema.rng file inside clang source tree.
type Comment struct {
	Name     string
	Abstract struct {
		Paras CommentParas `xml:"Para"`
	}
	Parameters struct {
		Parameter []struct {
			Name       string
			Discussion struct {
				Paras CommentParas `xml:"Para"`
			}
		}
	}
	ResultDiscussion struct {
		Paras CommentParas `xml:"Para"`
	}
	Discussion struct {
		Paras CommentParas `xml:"Para"`
	}
}

type CommentParas []string

func docComment(comment clang.Comment) string {
	xmlData := comment.FullComment_getAsXML()
	if len(xmlData) == 0 {
		return ""
	}

	var function Comment
	err := xml.Unmarshal([]byte(xmlData), &function)
	fatalOnError(err)

	var text strings.Builder

	// abstract
	for _, para := range function.Abstract.Paras {
		text.WriteString(strings.TrimSpace(para))
		text.WriteString("\n\n")
	}

	// discussion
	for _, para := range function.Discussion.Paras {
		text.WriteString(strings.TrimSpace(para))
		text.WriteString("\n\n")
	}

	// parameters
	if len(function.Parameters.Parameter) > 0 {
		text.WriteString("# parameters")
		text.WriteString("\n\n")

		for _, param := range function.Parameters.Parameter {
			text.WriteString("  - " + param.Name + " - ")
			for i, para := range param.Discussion.Paras {
				if i == 0 {
					text.WriteString(para)
				} else {
					text.WriteString("  " + para)
				}
				text.WriteString("\n")
			}
		}

		text.WriteString("\n\n")
	}

	// return
	if len(function.ResultDiscussion.Paras) > 0 {
		text.WriteString("# return")
		text.WriteString("\n\n")

		for _, para := range function.ResultDiscussion.Paras {
			text.WriteString("  - " + strings.TrimSpace(para))
			text.WriteString("\n\n")
		}
	}

	if text.Len() == 0 {
		return ""
	}
	return "/*\n" + strings.TrimSpace(text.String()) + "\n*/\n"
}

/*
addDocCommentLinks turn references to types in a doc comment in to links.
*/
func addDocCommentLinks(comment string, api api) string {
	if len(comment) == 0 {
		return comment
	}

	comment = strings.TrimPrefix(comment, "/*")
	comment = strings.TrimSuffix(comment, "*/\n")

	var newComment strings.Builder
	var word strings.Builder
	for _, r := range comment {
		if unicode.IsLetter(r) {
			word.WriteRune(r)
		} else {
			if word.Len() > 0 {
				possibleName := word.String()
				goName, ok := api.goNameForCppName(possibleName)
				if ok {
					newComment.WriteString(fmt.Sprintf("[%s]", goName))
				} else {
					newComment.WriteString(possibleName)
				}
				word.Reset()
			}

			newComment.WriteRune(r)
		}
	}

	return "/*" + newComment.String() + "*/\n"
}
