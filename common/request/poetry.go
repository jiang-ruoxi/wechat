package request

type MakePDF struct {
	ImgList []ImgListIndex `json:"img_list"`
}

type ImgListIndex struct {
	Index string `json:"index"`
	Img   string `json:"img"`
}

type CompressPicture struct {
	ImgList []ImgListIndex `json:"img_list"`
}
