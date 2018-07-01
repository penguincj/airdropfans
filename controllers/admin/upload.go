package admin

import (
	//"fmt"

	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nfnt/resize"
	//"qiniupkg.com/api.v7/kodo"
)

var (
	imgdir   = "./static/img/upload/"
	abimgdir = "/static/img/upload/"
)

type UploadHandle struct {
	baseController
}

//上传页面
func (this *UploadHandle) UploadPage() {
	obj := this.GetString("obj")
	this.Data["obj"] = obj
	this.TplName = "admin/upload.html"

}

//upload picture handel
func (this *UploadHandle) UploadSave() {
	if this.IsPost() {
		var err error
		file, fileHead, err := this.GetFile("uploadfile")
		if err != nil {
			this.Ctx.WriteString("post err:" + err.Error())
			return
		}
		defer file.Close()
		obj := this.GetString("obj")
		filename := fileHead.Filename
		split_part := strings.Split(filename, ".")
		ext := "." + strings.ToLower(split_part[1])

		//save location file
		filename = this.GetGUID() + ext
		tempfilename := "temp_" + filename
		data, err := ioutil.ReadAll(file)
		err = ioutil.WriteFile(imgdir+tempfilename, data, 0777)

		imgUrl := abimgdir + tempfilename
		fmt.Println("cj imgUrl ", imgUrl)

		if err != nil {
			this.Ctx.WriteString("save to file err:" + err.Error())
			return
		}
		if len(obj) == 0 {
			obj = "photo"
		}

		//this.Ctx.WriteString("<a href='../upload/add'>重新上传..</a><script>window.parent.setphoto('" + obj + "','" + imgdir + tempfilename + "');</script>") //输出文件地址
		this.Ctx.WriteString("<a href='../upload/add'>重新上传..</a><script>window.parent.setphoto('" + obj + "','" + imgUrl + "');</script>") //输出文件地址
	} else {
		this.Ctx.WriteString("请求错误...")
		return
	}
}

//添加水印和略缩
func WaterMark(tempfilepath string, newfilepath string, ext string) bool {
	imgb, _ := os.Open(tempfilepath)
	img, _ := jpeg.Decode(imgb)
	defer imgb.Close()

	img = Resize(img)

	wmb, _ := os.Open("static/img/mark.png")
	watermark, _ := png.Decode(wmb)
	defer wmb.Close()

	//把水印写到右下角，并向0坐标各偏移10个像素
	offset := image.Pt(img.Bounds().Dx()/2-watermark.Bounds().Dx()/2, img.Bounds().Dy()-watermark.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

	//生成新图片，并设置图片质量..
	out, err := os.Create(newfilepath)
	defer out.Close()

	jpeg.Encode(out, m, &jpeg.Options{100})
	os.Remove(tempfilepath)
	if err != nil {
		return false
	} else {
		return true
	}
}

func Resize(img image.Image) (m image.Image) {
	var maximage uint
	maximage = 600
	if img.Bounds().Dx() > 600 {
		return resize.Resize(maximage, 0, img, resize.MitchellNetravali)
	} else {
		return img
	}
}
