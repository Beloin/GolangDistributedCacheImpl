// Package cacher
package cacher

type Cacher interface {
	GetString(s string) (string, bool)
	SetString(key string, value string) bool
}
