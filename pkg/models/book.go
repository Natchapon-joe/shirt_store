package models

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

type Product struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Gender string `json:"gender"`
	Style  string `json:"style"`
	Size   string `json:"size"`
	Price  string `json:"price"`
}
