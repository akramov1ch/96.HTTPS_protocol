package models

type Item struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
}

var Items = make(map[int]Item)
