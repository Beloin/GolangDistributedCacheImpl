// Package cacher
package cacher

type Cacher interface {
	GetString(s string) (bool, string)
	SetString(key string, value string) bool
}
