package practices

import "errors"

type TagManager struct {
	Tags map[string]struct{}
}

func NewTagManager() *TagManager {
	m := make(map[string]struct{})
	return &TagManager{
		Tags: m,
	}
}

func (t *TagManager) AddTag(tag string) error {
	if _, ok := t.Tags[tag]; ok {
		return errors.New("ошибка, тег уже существует")
	}
	t.Tags[tag] = struct{}{}
	return nil
}

func (t *TagManager) RemoveTag(tag string) error {
	for key := range t.Tags {
		if key == tag {
			delete(t.Tags, key)
			return nil
		}
	}
	return errors.New("ошибка, тег не существует")
}

func (t TagManager) TagExists(tag string) bool {
	for key := range t.Tags {
		if key == tag {
			return true
		}
	}
	return false
}

func (t TagManager) ListTags() []string {
	var res []string
	for key := range t.Tags {
		res = append(res, key)
	}
	return res
}
