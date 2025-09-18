package cacher

type HashMapCacher struct {
	data map[string]string
}

func (cacher *HashMapCacher) GetString(s string) (string, bool) {
	res, has := cacher.data[s]
	return res, has
}

func (cacher *HashMapCacher) SetString(key string, value string) bool {
	cacher.data[key] = value
	return true
}
