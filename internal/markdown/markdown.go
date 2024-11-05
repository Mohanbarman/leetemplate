package markdown

type TokenType int

const (
	H1     TokenType = iota + 1 // Headings 1
	H2                          // Headings 2
	H3                          // Headings 3
	H4                          // Headings 4
	H5                          // Headings 5
	H6                          // Headings 6
	P                           // paragraph
	Br                          // line break
	Strong                      // Bold
	Em                          // Italic
)
