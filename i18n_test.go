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
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}

	expected = "core program was not found"
	actual = l.T("core program was not found")
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}

	expected = "not found core module: i18n-core"
	actual = l.T("not found core module: %s", "i18n-core")
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}

	expected = "current goroutine num: 3"
	actual = l.T("current goroutine num: %d", 3)
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}

	// 切换语言
	l.ToLang("zh")
	expected = "zh"
	actual = l.Lang()
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}

	expected = "未找到核心程序"
	actual = l.T("core program was not found")
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}

	expected = "未找到核心模块：i18n-core"
	actual = l.T("not found core module: %s", "i18n-core")
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}

	expected = "当前协程数：3"
	actual = l.T("current goroutine num: %d", 3)
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}
}

func TestI18nGroup(t *testing.T) {
	l := NewI18n().LoadFile("TestData/lang.json")
	if err := l.Error(); err != nil {
		t.Fatal(err)
	}

	_, err := l.GetGroup("test")
	expected := "unknown group: test"
	actual := err.Error()
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}

	m, err := l.GetGroup("test_zero")
	if err != nil {
		t.Fatal(err)
	}

	expected = "core program was not found"
	actual = m.GT("core program was not found")
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}

	expected = "not found core module: i18n-core"
	actual = m.GT("not found core module: %s", "i18n-core")
	if expected != actual {
		t.Fatalf("expected: %s\nactual: %s\nerror:expected != actual", expected, actual)
	}
}
