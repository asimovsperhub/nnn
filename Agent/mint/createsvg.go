package mint

import (
	"log"
	"os"
)

func writeFile(info string, path string) {

	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		// 创建文件失败处理

	} else {
		_, err = f.Write([]byte(info))
		if err != nil {
			// 写入失败处理
			log.Println("write svg err", err)
		}
	}
}
