package entities

import "time"

type Video struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url"`
	LocalPath   string    `json:"local_path,omitempty"`
	IsCached    bool      `json:"is_cached"`
	CacheExpiry time.Time `json:"cache_expiry,omitempty"`
	UserID      int       `json:"user_id"`
}

func (v *Video) GetID() int                { return v.ID }
func (v *Video) GetTitle() string          { return v.Title }
func (v *Video) GetDescription() string    { return v.Description }
func (v *Video) GetURL() string            { return v.URL }
func (v *Video) GetLocalPath() string      { return v.LocalPath }
func (v *Video) GetIsCached() bool         { return v.IsCached }
func (v *Video) GetCacheExpiry() time.Time { return v.CacheExpiry }
func (v *Video) GetUserID() int            { return v.UserID }

func (v *Video) SetID(id int)                 { v.ID = id }
func (v *Video) SetTitle(title string)        { v.Title = title }
func (v *Video) SetDescription(desc string)   { v.Description = desc }
func (v *Video) SetURL(url string)            { v.URL = url }
func (v *Video) SetLocalPath(path string)     { v.LocalPath = path }
func (v *Video) SetIsCached(cached bool)      { v.IsCached = cached }
func (v *Video) SetCacheExpiry(exp time.Time) { v.CacheExpiry = exp }
func (v *Video) SetUserID(userID int)         { v.UserID = userID }

func (v *Video) IsCacheValid() bool {
	if !v.IsCached {
		return false
	}
	return v.CacheExpiry.IsZero() || time.Now().Before(v.CacheExpiry)
}

func (v *Video) ClearCache() {
	v.LocalPath = ""
	v.IsCached = false
	v.CacheExpiry = time.Time{}
}
