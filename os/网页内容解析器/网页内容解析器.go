package 网页内容解析器

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type E网页内容解析器 struct {
	Document *goquery.Document
}

func New网页内容解析器(html string) *goquery.Document {
	obj := new(E网页内容解析器).E解析(html).Document
	return obj
}

func (this *E网页内容解析器) E解析(html string) *E网页内容解析器 {
	this.Document, _ = goquery.NewDocumentFromReader(strings.NewReader(html))
	return this
}
