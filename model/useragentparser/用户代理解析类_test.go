package useragentparser

import (
	"testing"
)

func TestE解析(t *testing.T) {
	ua := E解析("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	if ua.Name != "Chrome" {
		t.Errorf("E名称 expected Chrome, got %s", ua.Name)
	}

	if ua.OS != "Windows" {
		t.Errorf("E系统名称 expected Windows, got %s", ua.OS)
	}

	if !ua.E为Windows系统() {
		t.Error("E为Windows系统 expected true, got false")
	}

	if !ua.E为谷歌浏览器() {
		t.Error("E为谷歌浏览器 expected true, got false")
	}

	if ua.E取系统名称() != EWindows系统 {
		t.Errorf("E取系统名称 expected %s, got %s", EWindows系统, ua.E取系统名称())
	}

	if ua.E取浏览器名称() != E谷歌浏览器 {
		t.Errorf("E取浏览器名称 expected %s, got %s", E谷歌浏览器, ua.E取浏览器名称())
	}

	if !ua.E是否解析成功() {
		t.Error("E是否解析成功 expected true, got false")
	}
}

func TestE解析_移动设备(t *testing.T) {
	ua := E解析("Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.1 Mobile/15E148 Safari/604.1")

	if !ua.E为IOS系统() {
		t.Error("E为IOS系统 expected true, got false")
	}

	if !ua.E为Safari浏览器() {
		t.Error("E为Safari浏览器 expected true, got false")
	}

	if ua.E取系统名称() != EIOS系统 {
		t.Errorf("E取系统名称 expected %s, got %s", EIOS系统, ua.E取系统名称())
	}

	if ua.E取浏览器名称() != ESafari浏览器 {
		t.Errorf("E取浏览器名称 expected %s, got %s", ESafari浏览器, ua.E取浏览器名称())
	}

	if !ua.Mobile {
		t.Error("E为移动设备 expected true, got false")
	}
}

func TestE解析_机器人(t *testing.T) {
	ua := E解析("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	if !ua.E为谷歌机器人() {
		t.Error("E为谷歌机器人 expected true, got false")
	}

	if ua.E取浏览器名称() != E谷歌机器人 {
		t.Errorf("E取浏览器名称 expected %s, got %s", E谷歌机器人, ua.E取浏览器名称())
	}

	if !ua.Bot {
		t.Error("E为机器人设备 expected true, got false")
	}
}

func TestE版本号(t *testing.T) {
	ua := E解析("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	if ua.E取版本号() != "91.0.4472" {
		t.Errorf("E取版本号 expected 91.0.4472, got %s", ua.E取版本号())
	}

	if ua.E取版本号2() != "91.0" {
		t.Errorf("E取版本号2 expected 91.0, got %s", ua.E取版本号2())
	}

	if ua.E取系统版本号() != "10.0.0" {
		t.Errorf("E取系统版本号 expected 10.0.0, got %s", ua.E取系统版本号())
	}

	if ua.E取系统版本号2() != "10.0" {
		t.Errorf("E取系统版本号2 expected 10.0, got %s", ua.E取系统版本号2())
	}
}
