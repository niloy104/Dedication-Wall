package domain

type Dedication struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name,omitempty" db:"name"`
	Message   string `json:"message" db:"message"`
	SongURL   string `json:"song_url" db:"song_url"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
