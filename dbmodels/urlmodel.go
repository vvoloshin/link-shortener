package dbmodels

import "time"

type UrlModel struct {
	Hashed  string
	Url     string
	Created time.Time
}
