// Package webContentParser 网页内容解析器
// HTML网页内容解析。jquery的方式解析内容，选择器使用方式参考：http://www.w3school.com.cn/cssref/css_selectors.asp。
package webContentParser

import (
    "github.com/PuerkitoBio/goquery"
    "strings"
)

// E网页内容解析器 用于解析 HTML 文本的工具类型，封装 goquery.Document。
type E网页内容解析器 struct {
    Document *goquery.Document
}

// New网页内容解析器 创建并返回一个 goquery.Document，用于解析提供的 HTML 文本。
func New网页内容解析器(html string) *goquery.Document {
    obj := new(E网页内容解析器).E解析(html).Document
    return obj
}

// E解析 解析传入的 HTML 文本，初始化并返回解析器自身。
func (this *E网页内容解析器) E解析(html string) *E网页内容解析器 {
    this.Document, _ = goquery.NewDocumentFromReader(strings.NewReader(html))
    return this
}
