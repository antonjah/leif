package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func NewCache() cache.Cache {
	return *cache.New(1*time.Hour, 1*time.Hour)
}
