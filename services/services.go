package services

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

var redirectDomain string = "http://localhost:8080/"

type UrlShortener struct {
	urls map[string]string
}

var urlService *UrlShortener

func GetUrlService() *UrlShortener {
	if urlService == nil {
		urlService = &UrlShortener{
			urls: make(map[string]string),
		}
	}
	return urlService
}

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func (u *UrlShortener) Create(url string) (string, error) {
	if !isValidURL(url) {
		return "", errors.New(fmt.Sprintf("invalid url '%s'", url))
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	key := make([]byte, 6)
	for i := range key {
		key[i] = charset[random.Intn(len(charset))]
	}

	u.urls[string(key)] = url
	return redirectDomain + string(key), nil
}

func (u *UrlShortener) Get(key string) (string, error) {
	url, exists := u.urls[key]
	fmt.Println(url)
	if !exists {
		return "", errors.New(fmt.Sprintf("url with key '%s' not found", key))
	}
	return url, nil
}

func (u *UrlShortener) GetAll() map[string]string {
	return u.urls
}

func (u *UrlShortener) Delete(key string) error {
	_, exists := u.urls[key]
	if !exists {
		return errors.New(fmt.Sprintf("url with key '%s' not found", key))
	}
	delete(u.urls, key)
	return nil
}
