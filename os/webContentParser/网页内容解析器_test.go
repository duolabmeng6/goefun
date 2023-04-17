package 网页内容解析器

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
	. "github.com/duolabmeng6/goefun/ecore"
	. "github.com/duolabmeng6/goefun/os/ehttp"
)

func TestNew网页内容解析器(t *testing.T) {
	ehttp := NewHttp()
	html, _ := ehttp.Get("https://news.baidu.com/guonei")
	d := New网页内容解析器(html)
	//json := NewJson()
	d.Find("#col_focus > div.l-left-col > div.b-left > ul:nth-child(1) > li > a").Each(func(i int, selection *goquery.Selection) {
		name := selection.Text()
		href, _ := selection.Attr("href")
		E调试输出(name, href)
		//json.SetArray("data", H{
		//	"name": name,
		//	"href": href,
		//})
	})
	//E调试输出(json.ToJson(true))
}
