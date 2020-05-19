package main

type Brackets struct {
	start string
	end   string
}

type BracketMap struct {
	brackets []Brackets
}

func (b BracketMap) isOpenB(s string) bool {
	for _, v := range b.brackets {
		if v.start == s {
			return true
		}
	}
	return false
}

func (b BracketMap) isCloseB(s string) bool {
	for _, v := range b.brackets {
		if v.end == s {
			return true
		}
	}
	return false
}

func (b BracketMap) getOpenB(end string) string {
	for _, v := range b.brackets {
		if v.end == end {
			return v.start
		}
	}
	return ""
}

func NewBracketMap() *BracketMap {
	return &BracketMap{
		[]Brackets{
			{start: "(", end: ")"},
			{start: "[", end: "]"},
			{start: "{", end: "}"},
		},
	}
}

func isValid(str string) bool {
	var stack []string
	bm := NewBracketMap()

	for _, s := range str {
		chr := string(s)
		if bm.isOpenB(chr) {
			stack = append(stack, chr)
		}
		if bm.isCloseB(chr) {
			n := len(stack) - 1
			if n < 0 || bm.getOpenB(chr) != stack[n] {
				return false
			}
			stack = stack[:n]
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
