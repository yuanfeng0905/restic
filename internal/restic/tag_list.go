package restic

import (
	"fmt"
	"strings"
)

// TagList is a list of tags.
type TagList []string

// splitTagList splits a string into a list of tags. The tags in the string
// need to be separated by commas. Whitespace is stripped around the individual
// tags.
func splitTagList(s string) (l TagList) {
	for _, t := range strings.Split(s, ",") {
		l = append(l, strings.TrimSpace(t))
	}
	return l
}

func (l TagList) String() string {
	return "[" + strings.Join(l, ", ") + "]"
}

// Set updates the TagList's value.
func (l *TagList) Set(s string) error {
	*l = splitTagList(s)
	return nil
}

// Type returns a description of the type.
func (TagList) Type() string {
	return "TagList"
}

// TagLists consists of several TagList.
type TagLists []TagList

// splitTagLists splits a slice of strings into a slice of TagLists using
// SplitTagList.
func splitTagLists(s []string) (l TagLists) {
	l = make([]TagList, 0, len(s))
	for _, t := range s {
		l = append(l, splitTagList(t))
	}
	return l
}

func (l TagLists) String() string {
	return fmt.Sprintf("%v", []TagList(l))
}

// Set updates the TagList's value.
func (l *TagLists) Set(s string) error {
	*l = append(*l, splitTagList(s))
	return nil
}

// Type returns a description of the type.
func (TagLists) Type() string {
	return "TagLists"
}
