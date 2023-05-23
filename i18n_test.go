package I18n

import (
	"bytes"
	"reflect"
	"testing"
)

func ObjEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

func MustEqual(t *testing.T, expected, actual interface{}) {
	if ObjEqual(expected, actual) {
		return
	}
	t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
}

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
	l := NewI18n().LoadFile("TestData/lang_0.json")
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

func TestMultiFile(t *testing.T) {
	l := NewI18n().LoadFile("TestData/lang_0.json").LoadFile("TestData/lang_1.json")
	MustEqual(t, l.T("not found core module: %s", "i18n-core"), "not found core module: i18n-core")
	MustEqual(t, l.T("this is information from other files"), "this is information from other files")
}

func TestGetLangs(t *testing.T) {
	l := NewI18n().LoadFile("TestData/lang_0.json")
	MustEqual(t, []string{"en", "zh"}, l.GetLangs())
}

func TestGetGroups(t *testing.T) {
	l := NewI18n().LoadFile("TestData/lang_0.json").LoadFile("TestData/lang_1.json")
	MustEqual(t, []string{"test_zero", "test_one", "test_two"}, l.GetGroups())
}
