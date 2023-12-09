package core

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Image struct {
	ImgID      string `gorm:"primaryKey"`
	ImgName    string
	ImgContent []byte
}

var (
	db     *gorm.DB
	ErrMap = map[int]error{
		400: errors.New("图片已存在"),
		401: errors.New("创建图片失败"),
		402: errors.New("存储图片失败"),
		403: errors.New("查询图片失败"),
		404: errors.New("图片不存在"),
		405: errors.New("删除图片失败"),
	}
)

func init() {
	initDB()
}

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("images.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = db.AutoMigrate(&Image{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}
}

func md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func QueryByName(imgName string) (*Image, error) {
	imgID := md5Hash(imgName)

	img, ok := Query(imgID)
	if ok {
		return nil, ErrMap[403]
	}

	return img, nil
}

func Query(imgID string) (img *Image, _ bool) {
	img = new(Image)

	isExist := func(err error) bool {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		} else if err != nil {
			log.Error(err)
		}
		return true
	}

	err := db.First(
		img,
		"img_id = ?", imgID,
	).Error

	return img, isExist(err)
}

func GetAllImages() ([]Image, error) {
	var images []Image
	if err := db.Find(&images).Error; err != nil {
		return nil, ErrMap[403]
	}
	return images, nil
}

func CreateImage(imgName string, imgContent []byte) (string, error) {
	imgID := md5Hash(imgName)
	newImage := Image{
		ImgID:      imgID,
		ImgName:    imgName,
		ImgContent: imgContent,
	}

	if _, ok := Query(imgID); ok {
		return "", ErrMap[400]
	}

	if err := db.Create(&newImage).Error; err != nil {
		return "", ErrMap[401]
	}

	return imgID, nil
}

func UpdateImage(imgID string, newContent []byte) error {
	img, ok := Query(imgID)
	if !ok {
		return ErrMap[404]
	}

	img.ImgContent = newContent
	if err := db.Save(&img).Error; err != nil {
		return ErrMap[404]
	}

	return nil
}

func DeleteImage(imgID string) error {
	if err := db.Delete(&Image{}, "img_id = ?", imgID).Error; err != nil {
		return ErrMap[405]
	}
	return nil
}

func main() {

	CreateImage("百度", []byte{})

	imgs, err := GetAllImages()
	if err != nil {
		log.Error(err)
	}
	fmt.Println(imgs)
}
