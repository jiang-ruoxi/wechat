package request

type PoetryVideoReq struct {
	OpenId   string `json:"open_id" comment:"open_id"`
	PoetryId int    `json:"poetry_id" comment:"poetry_id"`
	Mp3      string `json:"mp3" comment:"mp3"`
}

type MakePDF struct {
	ImgList []ImgListIndex `json:"img_list"`
}

type ImgListIndex struct {
	Index string `json:"index"`
	Img   string `json:"img"`
}
