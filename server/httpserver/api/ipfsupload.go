package api

import (
	"bytes"
	shell "github.com/ipfs/go-ipfs-api"
	"log"
)

func UploadIPFS(raw []byte) (string, error) {
	sh := shell.NewShell("localhost:5001")
	reader := bytes.NewReader(raw)
	// https://github.com/ipfs/go-ipfs-api/blob/master/add.go
	fileHash, err := sh.Add(reader)
	if err != nil {
		return "", err
	}
	log.Println(fileHash)
	return fileHash, nil
}

//func WriteHash(writeJson string, cont interface{}) {
//	// os.O_TRUNC 覆盖  os.O_APPEND 追加
//	if distFile, err := os.OpenFile(writeJson, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); err != nil {
//		log.Println("create file failed", err)
//	} else {
//		enc := json.NewEncoder(distFile)
//		if err1 := enc.Encode(cont); err1 != nil {
//			log.Println("write hash.json failed", err1)
//		} else {
//			log.Println("write hash.json successful")
//		}
//	}
//}
