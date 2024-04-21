package service

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/jung-kurt/gofpdf"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
	"wechat/common/request"
	"wechat/utils"
)

type PDF struct {
}

//ApiMakePDF 生成PDF
func (p *PDF) ApiMakePDF(req request.MakePDF) (result string, name string, total int, size string, err error) {

	imgList := req.ImgList
	fmt.Printf("%#v \n", imgList)

	// 定义排序函数
	sortByIndex := func(i, j int) bool {
		return imgList[i].Index < imgList[j].Index
	}
	// 使用sort.Slice进行排序
	sort.Slice(imgList, sortByIndex)

	imgPathList := make([]string, 0)
	for _, item := range imgList {
		var imgPath string
		imgPath = utils.ReplaceURLPart(item.Img, "https://oss.58haha.com", "/data/static")
		imgPathList = append(imgPathList, imgPath)
	}
	fmt.Printf("%#v \n", imgPathList)
	pdf, err := p.doMakePDF(imgPathList)
	if err != nil {
		return
	}
	pdfFile := "https://oss.58haha.com/pdf-img/" + pdf

	oPath := "/data/static/pdf-img/" + pdf
	fileInfo, _ := os.Stat(oPath)
	fileSize := fileInfo.Size()
	size = humanize.Bytes(uint64(fileSize))

	return pdfFile, pdf, len(imgPathList), size, nil
}

func (p *PDF) doMakePDF(imageFiles []string) (string, error) {
	// 创建一个新的PDF文档
	pdf := gofpdf.New("P", "mm", "A4", "")
	// 遍历图片文件列表
	for _, file := range imageFiles {
		imgFile, err := os.Open(file)
		defer imgFile.Close()

		if err != nil {
			panic(err)
		}
		// 添加新的页面
		pdf.AddPage()
		imgWidth, imgHeight := pdf.GetPageSize()
		//imgWidth := width - 20
		//imgHeight := imgWidth * float64(img.Bounds().Dy()) / float64(img.Bounds().Dy())
		// 添加图片到页面
		pdf.Image(file, 0, 0, imgWidth, imgHeight, false, "", 0, "")
	}

	// 保存PDF文件
	pdfPath := "/data/static/pdf-img/"
	pdfFileName := strconv.FormatInt(time.Now().UnixNano(), 10) + ".pdf"
	pdfFile := pdfPath + pdfFileName
	err := pdf.OutputFileAndClose(pdfFile)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	for _, file := range imageFiles {
		os.Remove(file)
	}
	fmt.Println("PDF文件已生成")
	return pdfFileName, nil
}

func (p *PDF) CompressPicture(req request.CompressPicture) (imgRList []request.ImgListIndex, err error) {
	imgList := req.ImgList
	fmt.Printf("%#v \n", imgList)

	// 定义排序函数
	sortByIndex := func(i, j int) bool {
		return imgList[i].Index < imgList[j].Index
	}
	// 使用sort.Slice进行排序
	sort.Slice(imgList, sortByIndex)

	imgPathList := make([]string, 0)
	for _, item := range imgList {
		var imgPath string
		imgPath = utils.ReplaceURLPart(item.Img, "https://oss.58haha.com", "/data/static")
		imgPathList = append(imgPathList, imgPath)
	}
	fmt.Printf("%#v \n", imgPathList)
	err = p.doCompressPicture(imgPathList)
	if err != nil {
		return
	}
	return imgList, nil
}
func (p *PDF) doCompressPicture(imagePathList []string) (err error) {
	for _, imagePath := range imagePathList {
		file, err := os.Open(imagePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// 读取文件的前512个字节
		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil {
			return err
		}
		file.Seek(0, 0)

		// 根据MIME类型检查
		contentType := http.DetectContentType(buffer)

		var img image.Image
		var format string
		switch contentType {
		case "image/jpeg":
			img, err = jpeg.Decode(file)
			format = "jpeg"
		case "image/png":
			img, err = png.Decode(file)
			format = "png"
		default:
			return nil
		}

		if err != nil {
			return err
		}

		// 关闭并重新打开文件进行压缩
		file.Close()
		file, err = os.OpenFile(imagePath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		var opt jpeg.Options
		opt.Quality = 75 // 设置压缩质量为75

		switch format {
		case "jpeg":
			err = jpeg.Encode(file, img, &opt)
		case "png":
			encoder := png.Encoder{CompressionLevel: png.BestCompression}
			err = encoder.Encode(file, img)
		}
	}

	return err
}
