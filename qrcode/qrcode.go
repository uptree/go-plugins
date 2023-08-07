package qrcode

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/nfnt/resize"
	"github.com/tuotoo/qrcode"
)

// DataURISchemePng png
const DataURISchemePng = "data:image/png;base64,"

// CreateImage 生成二维码
func CreateImage(content string, width, height int) (image.Image, error) {
	if content == "" {
		return nil, errors.New("no content")
	}
	qrCode, err := qr.Encode(content, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}
	return barcode.Scale(qrCode, width, height)
}

// Create 生成二维码
func Create(content string, width, height int) (string, error) {
	qrCode, err := CreateImage(content, width, height)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = png.Encode(buf, qrCode); err != nil {
		return "", err
	}
	return DataURISchemePng + base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// CreateFile 生成二维码 文件
func CreateFile(content, fileName string, width, height int) error {
	qrCode, err := CreateImage(content, width, height)
	if err != nil {
		return err
	}
	file, _ := os.Create(fileName)
	defer file.Close()
	return png.Encode(file, qrCode)
}

// Parse 解析二维码
func Parse(fileBytes []byte) (string, error) {
	matrix, err := qrcode.Decode(bytes.NewReader(fileBytes))
	if err != nil {
		return "", err
	}
	return matrix.Content, nil
}

// ParseFile 解析二维码 文件
func ParseFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	matrix, err := qrcode.Decode(f)
	if err != nil {
		return "", err
	}
	return matrix.Content, nil
}

// CreateWithLogo 生成带logo的二维码
func CreateWithLogo(content string, size int, logo []byte, percent int) (string, error) {

	logoImage, _, err := image.Decode(bytes.NewBuffer(logo))
	if err != nil || logoImage == nil {
		return "", errors.New("decode logo file error")
	}
	return CreateWithLogoImage(content, size, logoImage, percent)
}

// CreateWithLogoImage 根据image生成带logo的二维码
func CreateWithLogoImage(content string, size int, logo image.Image, percent int) (string, error) {

	qrCode, err := CreateImage(content, size, size)
	if err != nil {
		return "", err
	}
	var img image.Image
	if logo != nil {
		if percent <= 0 || percent >= 100 {
			percent = 20
		}
		logoSize := float64(size) * float64(percent) / 100

		img, err = addLogo(qrCode, logo, int(logoSize))
	}

	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)

	return DataURISchemePng + base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// addLogo 添加Logo
func addLogo(srcImage image.Image, logo image.Image, size int) (image.Image, error) {

	logoImage, err := resizeLogo(logo, uint(size))
	if err != nil {
		return nil, err
	}

	offset := image.Pt((srcImage.Bounds().Dx()-logoImage.Bounds().Dx())/2,
		(srcImage.Bounds().Dy()-logoImage.Bounds().Dy())/2)
	b := srcImage.Bounds()
	m := image.NewNRGBA(b)
	draw.Draw(m, b, srcImage, image.Point{}, draw.Src)
	draw.Draw(m, logoImage.Bounds().Add(offset), logoImage, image.Point{}, draw.Over)

	return m, nil
}

// resizeLogo 缩放Logo
func resizeLogo(logo image.Image, size uint) (image.Image, error) {

	img := resize.Resize(size, size, logo, resize.Lanczos3)
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}
	return img, nil
}
