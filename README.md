# I18n

golang 简单的多语言

## example

### `lang.json`

```json
[
  {
    "id": "core program was not found",
    "langs": {
      "en": "core program was not found",
      "zh": "未找到核心程序"
    }
  },
  {
    "id": "not found core module: %s",
    "langs": {
      "en": "not found core module: %s",
      "zh": "未找到核心模块：%s"
    }
  },
  {
    "id": "current goroutine num: %d",
    "langs": {
      "en": "current goroutine num: %d",
      "zh": "当前协程数：%d"
    }
  }
]
```

### `main.go`

```go
package main

import (
	"fmt"

	"github.com/wsshow/I18n"
)

func main() {
	l := I18n.NewI18n().LoadFile("./lang.json")

	s := l.T("core program was not found")
	fmt.Println(s)

	s = l.T("not found core module: %s", "i18n-core")
	fmt.Println(s)

	s = l.T("current goroutine num: %d", 3)
	fmt.Println(s)

	l.ToLang("zh")

	s = l.T("core program was not found")
	fmt.Println(s)

	s = l.T("not found core module: %s", "i18n-core")
	fmt.Println(s)

	s = l.T("current goroutine num: %d", 3)
	fmt.Println(s)
}
```
