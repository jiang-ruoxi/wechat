package model

type ChineseBookAlbum struct {
	Id       int    `json:"-"`
	BookId   string `json:"book_id"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	Position uint8  `json:"position"`
}

func (ChineseBookAlbum) TableName() string {
	return "s_album_picture"
}

type ChineseAlbumInfo struct {
	Id       int    `json:"id"`
	BookId   string `json:"book_id"`
	Mp3      string `json:"mp3"`
	Pic      string `json:"pic"`
	Position uint8  `json:"position"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Duration string `json:"duration"`
}

func (ChineseAlbumInfo) TableName() string {
	return "s_album_picture_info"
}

type ChineseBookNavName struct {
	Id         int    `json:"-"`
	CategoryId int    `json:"category_id"`
	Name       string `json:"name"`
	SSort      int    `json:"s_sort"`
	SType      int    `json:"s_type"`
}

func (ChineseBookNavName) TableName() string {
	return "s_book_name"
}

type ChineseBook struct {
	Id       int    `json:"-"`
	BookId   string `json:"book_id"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	Type     uint8  `json:"type"`
	Position uint8  `json:"position"`
}

func (ChineseBook) TableName() string {
	return "s_chinese_picture"
}

type ChineseBookInfo struct {
	Id       int    `json:"id"`
	BookId   string `json:"book_id"`
	Mp3      string `json:"mp3"`
	Pic      string `json:"pic"`
	Position uint8  `json:"position"`
}

func (ChineseBookInfo) TableName() string {
	return "s_chinese_picture_info"
}

type PoetryPicture struct {
	Id     int    `json:"-"`
	BookId int    `json:"book_id"`
	Title  string `json:"title"`
	Icon   string `json:"icon"`
	Author string `json:"author"`
	TypeId int    `json:"type_id"`
}

func (PoetryPicture) TableName() string {
	return "s_poetry_picture"
}

type Lexicon struct {
	Id         int    `json:"id"`
	Word       string `json:"word"`
	Mark       string `json:"mark"`
	Annotation string `json:"annotation"`
	Explain    string `json:"explain"`
	Type       int    `json:"-"`
	Status     int    `json:"-"`
}

func (Lexicon) TableName() string {
	return "s_lexicon"
}

type PoetryPictureInfo struct {
	Id       int    `json:"-"`
	BookId   int    `json:"book_id"`
	CN       string `json:"cn"`
	Pic      string `json:"pic"`
	Mp3      string `json:"mp3"`
	Position int    `json:"position"`
}

func (PoetryPictureInfo) TableName() string {
	return "s_poetry_picture_info"
}
