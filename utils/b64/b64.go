package b64

import (
	"encoding/base64"
	"log"
	"os"
)

func Img2base64(imgPath string) string {
	imgData, err := os.ReadFile(imgPath)
	if err != nil {
		log.Fatalln(err)
	}
	base64Img := base64.StdEncoding.EncodeToString(imgData)
	return base64Img
}

func WriteImgFileByBase64(value string, imgPath string) {
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		log.Fatalln(err)
	}
	err = os.WriteFile(imgPath, data, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

//func main() {
//	imgPath := "utils/icon/logo.png"
//	imgBase64 := Img2base64(imgPath)
//	fmt.Println(imgBase64)
//}
