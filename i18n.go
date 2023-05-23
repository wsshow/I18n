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

type messageGroup struct {
	GroupName string    `json:"groupname"`
	Messages  []message `json:"messages"`
}

type i18n struct {
	messages map[string]map[string]message
	mutex    sync.RWMutex
	langs    []string
	groups   []string
	group    string
	lang     string
	err      error
}

func NewI18n() *i18n {
	return &i18n{
		lang:     "en",
		messages: make(map[string]map[string]message),
		err:      nil,
	}
}

func (i *i18n) LoadFile(filename string) *i18n {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	data, err := os.ReadFile(filename)
	if err != nil {
		i.err = err
		return i
	}
	var messageGroups []messageGroup
	err = json.Unmarshal(data, &messageGroups)
	if err != nil {
		i.err = err
		return i
	}
	for _, group := range messageGroups {
		mg := make(map[string]message)
		for _, msg := range group.Messages {
			mg[msg.ID] = msg
			for lang := range msg.Langs {
				if !i.InLangs(lang) {
					i.langs = append(i.langs, lang)
				}
			}
		}
		i.messages[group.GroupName] = mg
		i.groups = append(i.groups, group.GroupName)
	}
	return i
}

func (i *i18n) InLangs(lang string) bool {
	for _, l := range i.langs {
		if l == lang {
			return true
		}
	}
	return false
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

func (i *i18n) ToGroup(group string) *i18n {
	if i.err != nil || i.group == group {
		return i
	}
	i.mutex.Lock()
	defer i.mutex.Unlock()
	n := &i18n{
		messages: i.messages,
		group:    group,
		lang:     i.lang,
		err:      i.err,
	}
	return n
}

func (i *i18n) GetLangs() []string {
	return i.langs
}

func (i *i18n) GetGroups() []string {
	return i.groups
}

func (i *i18n) T(id string, args ...interface{}) string {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	if len(i.group) > 0 {
		if mg, ok := i.messages[i.group]; ok {
			if msg, ok := mg[id]; ok {
				if langStr, ok := msg.Langs[i.lang]; ok {
					return fmt.Sprintf(langStr, args...)
				}
			}
		}
		return id
	}
	for _, mg := range i.messages {
		if msg, ok := mg[id]; ok {
			if langStr, ok := msg.Langs[i.lang]; ok {
				return fmt.Sprintf(langStr, args...)
			}
		}
	}
	return id
}

func (i *i18n) Error() error {
	return i.err
}
