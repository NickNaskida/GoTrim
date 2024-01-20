package main

type UrlShortener struct {
	urls map[string]string
}

func NewUrlShortener() *UrlShortener {
	return &UrlShortener{
		urls: make(map[string]string),
	}
}

func (u *UrlShortener) Add(url string) string {
	panic("not implemented")
}

func (u *UrlShortener) Get(key string) string {
	panic("not implemented")
}

func (u *UrlShortener) Remove(key string) {
	panic("not implemented")
}
