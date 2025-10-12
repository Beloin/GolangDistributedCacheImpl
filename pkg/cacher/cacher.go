// Package cacher
package cacher

type Cacher interface {
	GetString(s string) (string, bool)
	SetString(key string, value string) bool
	Paginate(from, limit int) map[string]string
}
