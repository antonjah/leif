package utils

import (
	"github.com/patrickmn/go-cache"
)

func NewCache() cache.Cache {
	return *cache.New(cache.NoExpiration, cache.NoExpiration)
}