// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/helper/footer.gohtml

package helper

import (
	"io"
	"strings"
)

// Footer generates templates/helper/footer.gohtml
func Footer() string {
	var _b strings.Builder
	RenderFooter(&_b)
	return _b.String()
}

// RenderFooter render templates/helper/footer.gohtml
func RenderFooter(_buffer io.StringWriter) {
	_buffer.WriteString("<div>copyright 2014</div>")

}
