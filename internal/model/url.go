package urlModel

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"sync"
)

type Url struct {
	OldURL       string `json:"url"`
	ShortenedURL string `json:"shortened_url"`
}

var (
	urlStore = map[string]string{}
	mu       sync.RWMutex
)

func ShortenUrl(url string) Url {
	mu.Lock()
	defer mu.Unlock()
	short := hashTheUrl(url)
	urlStore[short] = url
	return Url{
		OldURL:       url,
		ShortenedURL: short,
	}
}

func GetOldUrl(newUrl string) (string, error) {
	mu.RLock()
	defer mu.RUnlock()
	originalUrl, ok := urlStore[newUrl]
	if !ok {
		return "", errors.New("Not found")
	}
	return originalUrl, nil
}

func hashTheUrl(url string) string {
	hashed := sha256.Sum256([]byte(url))

	return hex.EncodeToString(hashed[:][:5])
}
