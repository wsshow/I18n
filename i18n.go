package I18n

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Message struct {
	ID    string            `json:"id"`
	Langs map[string]string `json:"langs"`
}

type I18n struct {
	messages map[string]Message
	mutex    sync.RWMutex
	lang     string
	err      error
}

func NewI18n() *I18n {
	return &I18n{
		lang:     "en",
		messages: make(map[string]Message),
		err:      nil,
	}
}

func (i *I18n) LoadFile(filename string) *I18n {
	data, err := os.ReadFile(filename)
	if err != nil {
		i.err = err
		return i
	}
	var messages []Message
	err = json.Unmarshal(data, &messages)
	if err != nil {
		i.err = err
		return i
	}
	i.mutex.Lock()
	defer i.mutex.Unlock()
	for _, msg := range messages {
		i.messages[msg.ID] = msg
	}
	return i
}

func (i *I18n) ToLang(lang string) *I18n {
	if i.err != nil || i.lang == lang {
		return i
	}
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.lang = lang
	return i
}

func (i *I18n) T(id string, args ...interface{}) string {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	if msg, ok := i.messages[id]; ok {
		if langStr, ok := msg.Langs[i.lang]; ok {
			return fmt.Sprintf(langStr, args...)
		}
	}
	return id
}

func (i *I18n) Error() error {
	return i.err
}
