package api

import (
	_ "github.com/lib/pq"
)

type MessageDTO struct {
	Name    string `json:"name"`
	Content   string `json:"content"`
}

