package discogs

type Discogs struct {
	secret   string
	key      string
	callback string
}

func DiscogsWithAuth(key, secret, callback string) *Discogs {
	return &Discogs{
		key:      key,
		secret:   secret,
		callback: callback,
	}
}
