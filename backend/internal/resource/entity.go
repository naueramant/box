package resource

import "time"

type Resource struct {
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	Size      int       `json:"size"`
	Extension string    `json:"extension,omitempty"`
	Modified  time.Time `json:"modified"`
	IsDir     bool      `json:"isDir"`
}
