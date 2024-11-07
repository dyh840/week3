package main

import (
	"io/ioutil"
	"log"
	"muxi-backend/tool/getDecryptedPaper"
	"muxi-backend/tool/savePaper"
	"net/http"
)

func main() {
	paperURL := "http://121.43.151.190:8000/paper"
	secretURL := "http://121.43.151.190:8000/secret"

	// 获取加密论文
	resp, err := http.Get(paperURL)
	if err != nil {
		log.Fatalf("Failed to get paper: %v", err)
	}
	defer resp.Body.Close()
	encryptedPaper, _ := ioutil.ReadAll(resp.Body)

	// 获取解密密钥
	resp, err = http.Get(secretURL)
	if err != nil {
		log.Fatalf("Failed to get secret key: %v", err)
	}
	defer resp.Body.Close()
	secretKey, _ := ioutil.ReadAll(resp.Body)

	// 解密论文
	decryptedPaper := getDecryptedPaper.GetDecryptedPaper(string(encryptedPaper), string(secretKey))

	// 保存解密后的论文
	savePath := "../paper/Academician Sun's papers.txt"
	savePaper.SavePaper(savePath, decryptedPaper)
}
