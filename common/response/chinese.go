package response

type ResponseChineseBook struct {
	Id        int    `json:"id"`
	BookId    string `json:"book_id"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Level      uint8  `json:"level"`
	Position  uint8  `json:"position"`
	BookCount string `json:"book_count"`
}

type ResponseBookInfoCount struct {
	BookId    string `json:"book_id"`
	BookCount string `json:"book_count"`
}
