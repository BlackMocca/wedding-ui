package models

import "time"

type Celebrate struct {
	Id            int64     `json:"id" db:"id" type:"int64"`
	CelebrateText string    `json:"celebrate_text" db:"celebrate_text" type:"string"`
	CelebrateFrom string    `json:"celebrate_from" db:"celebrate_from" type:"string"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

func NewCelebrate(celText string, celFrom string) *Celebrate {
	return &Celebrate{
		CelebrateText: celText,
		CelebrateFrom: celFrom,
	}
}
