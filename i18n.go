package I18n

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type message struct {
	ID    string            `json:"id"`
	Langs map[string]string `json:"langs"`
}

type i18n struct {
	messages map[string]message
	mutex    sync.RWMutex
	lang     string
	err      error
}

func NewI18n() *i18n {
	return &i18n{
		lang:     "en",
		messages: make(map[string]message),
		err:      nil,
	}
}

func (i *i18n) LoadFile(filename string) *i18n {
	data, err := os.ReadFile(filename)
	if err != nil {
		i.err = err
		return i
	}
	var messages []message
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

func (i *i18n) Lang() string {
	return i.lang
}

func (i *i18n) ToLang(lang string) *i18n {
	if i.err != nil || i.lang == lang {
		return i
	}
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.lang = lang
	return i
}

func (i *i18n) T(id string, args ...interface{}) string {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	if msg, ok := i.messages[id]; ok {
		if langStr, ok := msg.Langs[i.lang]; ok {
			return fmt.Sprintf(langStr, args...)
		}
	}
	return id
}

func (i *i18n) Error() error {
	return i.err
}
