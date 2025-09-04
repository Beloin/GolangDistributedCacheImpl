package cacher

type HashMapCacher struct{}

func (cacher *HashMapCacher) GetString(s string) (bool, string) {
	return false, ""
}

func (cacher *HashMapCacher) SetString(key string, value string) bool {
	return false
}
