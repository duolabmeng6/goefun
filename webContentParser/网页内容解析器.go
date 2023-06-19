// Package webContentParser 网页内容解析器
// HTML网页内容解析。jquery的方式解析内容，选择器使用方式参考：http://www.w3school.com.cn/cssref/css_selectors.asp。
package webContentParser

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
