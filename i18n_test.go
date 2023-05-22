package I18n

import (
	"testing"
)

func TestI18n(t *testing.T) {

	l := NewI18n().LoadFile("TestData/lang.json")
	if err := l.Error(); err != nil {
		t.Fatal(err)
	}

	// 默认语言
	expected := "en"
	actual := l.Lang()
	MustEqual(t, expected, actual)

	expected = "core program was not found"
	actual = l.T("core program was not found")
	MustEqual(t, expected, actual)

	expected = "not found core module: i18n-core"
	actual = l.T("not found core module: %s", "i18n-core")
	MustEqual(t, expected, actual)

	expected = "current goroutine num: 3"
	actual = l.T("current goroutine num: %d", 3)
	MustEqual(t, expected, actual)

	// 切换语言
	l.ToLang("zh")
	expected = "zh"
	actual = l.Lang()
	MustEqual(t, expected, actual)

	expected = "未找到核心程序"
	actual = l.T("core program was not found")
	MustEqual(t, expected, actual)

	expected = "未找到核心模块：i18n-core"
	actual = l.T("not found core module: %s", "i18n-core")
	MustEqual(t, expected, actual)

	expected = "当前协程数：3"
	actual = l.T("current goroutine num: %d", 3)
	MustEqual(t, expected, actual)
}

func TestI18nGroup(t *testing.T) {
	l := NewI18n().LoadFile("TestData/lang.json")
	if err := l.Error(); err != nil {
		t.Fatal(err)
	}

	l0 := l.ToGroup("test_zero")
	MustEqual(t, l0.Lang(), "en")

	l1 := l.ToGroup("test_one").ToLang("zh")
	MustEqual(t, l1.Lang(), "zh")

	MustEqual(t, l0.Lang(), "en")

	MustEqual(t, l0.T("not found core module: %s", "i18n-core"), "not found core module: i18n-core")
	MustEqual(t, l1.T("not found core module: %s", "i18n-core"), "未找到核心模块：i18n-core")
}

func MustEqual(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}
}
