package api

import (
	"encoding/json"
	"fmt"
	"github.com/dnsdao/nft.pass/Agent/mint"
	"github.com/dnsdao/nft.pass/database/ldb"
	"github.com/gorilla/mux"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

type Api struct {
}

func NewApi() *Api {
	return &Api{}
}

type MysqlRes struct {
	Addr string `json:"addr"`
}
type Res struct {
	Name          string              `json:"name"`
	Image         string              `json:"image"`
	External_Url  string              `json:"external_url"`
	Animation_Url string              `json:"animation_url"`
	Description   string              `json:"description"`
	Attributes    []map[string]string `json:"attributes"`
}

func GetLabel(tokenId int) string {
	label := "NO."
	scardId := fmt.Sprint(tokenId)
	switch len(scardId) {
	case 1:
		label += "000" + scardId
	case 2:
		label += "00" + scardId
	case 3:
		label += "0" + scardId
	case 4:
		label += scardId
	}
	return label
}

func GetPassType(confT string, tokenId int) (color string) {
	tokenId_ := big.NewInt(int64(tokenId))
	ColorList := []string{"nocolor", "color", "gold", "green"}
	coldBoot, cli, err := mint.NewColdBootClient(confT)
	if err != nil {
		log.Println("NewColdBootClient", err)
		return ColorList[0]
	}
	defer cli.Close()
	cardId, err1 := coldBoot.PassCardType(nil, tokenId_)
	if err1 != nil {
		log.Println("GetPassType", err1)
		return ColorList[0]
	}
	log.Println("cardId", cardId)
	return ColorList[cardId]

}
func (a *Api) CTokenId(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}
	vars := mux.Vars(request)
	log.Println("CTokenId query", vars)
	stoken_id := vars["token_id"]
	tokenId, err := strconv.Atoi(stoken_id)
	if err != nil {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}
	token_id := int32(tokenId)
	msg := "not found tokenid"
	if token_id < 0 || token_id >= 9999 {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}
	db := ldb.GetLdb()
	//mdb := mysqldb.GetMdb()
	val, err := db.GetKey("c" + stoken_id)
	var respstr []byte
	label := GetLabel(tokenId)
	var color string
	Ipfs, _ := db.GetIpfsc()
	if err == nil {
		if len(val) > 0 {
			useraddr_color := strings.Split(val, "_")
			//useraddr := useraddr_color[0]
			//var c MysqlRes
			//merr := mdb.Mysql.QueryRow(fmt.Sprintf("SELECT addr FROM Address WHERE addr='%s' LIMIT 1;", useraddr)).Scan(&c.Addr)
			//if merr != nil {
			//	log.Println("CTokenId Mysql Select err", merr)
			//}
			color = useraddr_color[1]
			res := new(Res)
			res.Name = fmt.Sprintf("%s card #%s", color, label)
			//if c.Addr == "" {
			//	res.Name = fmt.Sprintf("*%s card #%s", color, label)
			//	white, _ := db.GetKey("notinwhite")
			//	if white != "" {
			//		white += stoken_id + ","
			//	} else {
			//		white = stoken_id + ","
			//	}
			//	db.SaveKey("notinwhite", white)
			//}
			// hash_value := common.HexToHash(val)
			res.Image = fmt.Sprintf("https://dnsdaonftpasscard.s3.ap-east-1.amazonaws.com/cards/%s_%s.gif", color, stoken_id)
			res.External_Url = Ipfs.Upload[fmt.Sprintf("%s_%s", color, stoken_id)]
			res.Animation_Url = fmt.Sprintf("m%s_%s.gif", color, stoken_id)
			res.Description = fmt.Sprintf("a %s card for udid free name", color)
			respstr, err = json.Marshal(res)
			if err == nil {
				msg = string(respstr)
			}
		}

	} else {
		color = GetPassType("c_conf", tokenId)
		if color != "nocolor" {
			res := new(Res)
			res.Name = fmt.Sprintf("%s card #%s", color, label)
			// hash_value := common.HexToHash(val)
			// https://dnsdaonftpasscard.s3.ap-east-1.amazonaws.com/cards/green_177.gif
			res.Image = fmt.Sprintf("https://dnsdaonftpasscard.s3.ap-east-1.amazonaws.com/cards/%s_%s.gif", color, stoken_id)
			res.External_Url = Ipfs.Upload[fmt.Sprintf("%s_%s", color, stoken_id)]
			res.Animation_Url = fmt.Sprintf("c%s_%s.gif", color, stoken_id)
			res.Description = fmt.Sprintf("a %s card for udid free name", color)
			respstr, err = json.Marshal(res)
			if err == nil {
				msg = string(respstr)
			}
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))
	//Ss3, _ := db.GetS3c()
	//filename := fmt.Sprintf("%s_%s.gif", color, stoken_id)
	//savepath := fmt.Sprintf("./cards/%s", filename)
	//if color != "green" && color != "nocolor" {
	//	if Ss3 == nil {
	//		//./原始图片
	//		CreateGif(label, color, savepath)
	//		raw := Read(savepath)
	//		s3res, errs3 := UploadFileToS3(raw, filename)
	//		if errs3 != nil {
	//			log.Println("S3 Upload failed", err)
	//		} else {
	//			Ss3 = &ldb.S3{
	//				Upload: map[string]string{fmt.Sprintf("%s_%s", color, stoken_id): "yes"},
	//			}
	//			db.SaveS3c(Ss3)
	//			log.Println("S3 upload success", s3res)
	//		}
	//
	//	} else {
	//		if Ss3.Upload[fmt.Sprintf("%s_%s", color, stoken_id)] != "yes" {
	//			//./原始图片
	//			CreateGif(label, color, savepath)
	//			raw := Read(savepath)
	//			s3res, errs3 := UploadFileToS3(raw, filename)
	//			if errs3 != nil {
	//				log.Println("S3 Upload failed", errs3)
	//			} else {
	//				Ss3.Upload[fmt.Sprintf("%s_%s", color, stoken_id)] = "yes"
	//				db.SaveS3c(Ss3)
	//				log.Println("S3 upload success", s3res)
	//			}
	//		}
	//
	//	}
	//}
	//if color != "nocolor" {
	//	if Ipfs == nil {
	//		//./原始图片
	//		CreateGif(label, color, savepath)
	//		raw := Read(savepath)
	//		Ipfsres, erripfs := UploadIPFS(raw)
	//		if erripfs != nil {
	//			log.Println("Ipfs Upload failed", erripfs)
	//		} else {
	//			Ipfs = &ldb.Ipfs{
	//				Upload: map[string]string{fmt.Sprintf("%s_%s", color, stoken_id): fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres)},
	//			}
	//			db.SaveIpfsc(Ipfs)
	//			log.Println("Ipfs upload success", fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres))
	//		}
	//	} else {
	//		if Ipfs.Upload[fmt.Sprintf("%s_%s", color, stoken_id)] == "" {
	//			//./原始图片
	//			CreateGif(label, color, savepath)
	//			raw := Read(savepath)
	//			Ipfsres, erripfs := UploadIPFS(raw)
	//			if erripfs != nil {
	//				log.Println("Ipfs Upload failed", erripfs)
	//			} else {
	//				Ipfs.Upload[fmt.Sprintf("%s_%s", color, stoken_id)] = fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres)
	//				db.SaveIpfsc(Ipfs)
	//				log.Println("Ipfs upload success", fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres))
	//			}
	//		}
	//	}
	//}
}

func (a *Api) MTokenId(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}
	vars := mux.Vars(request)
	log.Println("MTokenId query", vars)
	stoken_id := vars["token_id"]
	tokenId, err := strconv.Atoi(stoken_id)
	if err != nil {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}
	token_id := int32(tokenId)
	msg := "not found tokenid"
	if token_id < 0 || token_id >= 9999 {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}
	db := ldb.GetLdb()
	val, err := db.GetKey("m" + stoken_id)
	var respstr []byte
	label := GetLabel(tokenId)
	var color string
	Ipfs, _ := db.GetIpfsm()
	if err == nil {
		if len(val) > 0 {
			useraddr_color := strings.Split(val, "_")
			color = useraddr_color[1]
			res := new(Res)
			tokenName, _ := db.GetTokenName()
			if tokenName != nil && tokenName.TokenName[stoken_id] != "" {
				//res.Name = tokenName.TokenName[stoken_id]
				res.Attributes = []map[string]string{{"trait_type": "Minting Rights", "value": "Used"}}
			} else {
				// res.Name = fmt.Sprintf("%s card #%s", color, label)
				res.Attributes = []map[string]string{{"trait_type": "minting rights", "value": "Ready"}}
			}
			res.Name = fmt.Sprintf("%s card #%s", color, label)
			res.Image = fmt.Sprintf("https://dnsdaonftpasscard.s3.ap-east-1.amazonaws.com/cards/%s_%s.gif", color, stoken_id)
			res.External_Url = Ipfs.Upload[fmt.Sprintf("%s_%s", color, stoken_id)]
			res.Animation_Url = fmt.Sprintf("m%s_%s.gif", color, stoken_id)
			res.Description = fmt.Sprintf("a %s card for udid free name", color)
			respstr, err = json.Marshal(res)
			if err == nil {
				msg = string(respstr)
			}
		}

	} else {
		color = GetPassType("m_conf", tokenId)
		if color != "nocolor" {
			res := new(Res)
			res.Name = fmt.Sprintf("%s card #%s", color, label)
			// hash_value := common.HexToHash(val)
			res.Image = fmt.Sprintf("https://dnsdaonftpasscard.s3.ap-east-1.amazonaws.com/cards/%s_%s.gif", color, stoken_id)
			res.External_Url = Ipfs.Upload[fmt.Sprintf("%s_%s", color, stoken_id)]
			res.Animation_Url = fmt.Sprintf("m%s_%s.gif", color, stoken_id)
			res.Description = fmt.Sprintf("a %s card for udid free name", color)
			respstr, err = json.Marshal(res)
			if err == nil {
				msg = string(respstr)
			}
		}
	}
	if color != "green" && color != "nocolor" {
		card, _ := db.GetCard()
		if card != nil {
			card.Name[fmt.Sprintf("%s_%s.gif", color, stoken_id)] = "yes"
		} else {
			card = &ldb.Card{
				Name: map[string]string{fmt.Sprintf("%s_%s.gif", color, stoken_id): "yes"},
			}
		}
		log.Println("not green card", color)
		log.Println(card)
		err := db.SaveCard(card)
		if err != nil {
			log.Println("Save Card err", err)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))
	//Ss3, _ := db.GetS3m()
	//filename := fmt.Sprintf("%s_%s.gif", color, stoken_id)
	//savepath := fmt.Sprintf("./cards/%s", filename)
	//if color != "green" && color != "nocolor" {
	//	if Ss3 == nil {
	//		//./原始图片
	//		CreateGif(label, color, savepath)
	//		raw := Read(savepath)
	//		s3res, errs3 := UploadFileToS3(raw, filename)
	//		if errs3 != nil {
	//			log.Println("S3 Upload failed", err)
	//		} else {
	//			Ss3 = &ldb.S3{
	//				Upload: map[string]string{fmt.Sprintf("%s_%s", color, stoken_id): "yes"},
	//			}
	//			db.SaveS3m(Ss3)
	//			log.Println("s3 upload success", s3res)
	//		}
	//
	//	} else {
	//		if Ss3.Upload[fmt.Sprintf("%s_%s", color, stoken_id)] != "yes" {
	//			//./原始图片
	//			CreateGif(label, color, savepath)
	//			raw := Read(savepath)
	//			s3res, errs3 := UploadFileToS3(raw, filename)
	//			if errs3 != nil {
	//				log.Println("S3 Upload failed", errs3)
	//			} else {
	//				Ss3.Upload[fmt.Sprintf("%s_%s", color, stoken_id)] = "yes"
	//				db.SaveS3m(Ss3)
	//				log.Println("S3 upload success", s3res)
	//			}
	//		}
	//	}
	//	if Ipfs == nil {
	//		//./原始图片
	//		CreateGif(label, color, savepath)
	//		raw := Read(savepath)
	//		Ipfsres, erripfs := UploadIPFS(raw)
	//		if erripfs != nil {
	//			log.Println("Ipfs Upload failed", erripfs)
	//		} else {
	//			Ipfs = &ldb.Ipfs{
	//				Upload: map[string]string{fmt.Sprintf("%s_%s", color, stoken_id): fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres)},
	//			}
	//			db.SaveIpfsm(Ipfs)
	//			log.Println("s3 upload success", fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres))
	//		}
	//	} else {
	//		if Ipfs.Upload[fmt.Sprintf("%s_%s", color, stoken_id)] == "" {
	//			//./原始图片
	//			CreateGif(label, color, savepath)
	//			raw := Read(savepath)
	//			Ipfsres, erripfs := UploadIPFS(raw)
	//			if erripfs != nil {
	//				log.Println("Ipfs Upload failed", erripfs)
	//			} else {
	//				Ipfs.Upload[fmt.Sprintf("%s_%s", color, stoken_id)] = fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres)
	//				db.SaveIpfsm(Ipfs)
	//				log.Println("s3 upload success", fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres))
	//			}
	//		}
	//	}

	//}
	//if color != "nocolor" {
	//	if Ipfs == nil {
	//		//./原始图片
	//		CreateGif(label, color, savepath)
	//		raw := Read(savepath)
	//		Ipfsres, erripfs := UploadIPFS(raw)
	//		if erripfs != nil {
	//			log.Println("Ipfs Upload failed", erripfs)
	//		} else {
	//			Ipfs = &ldb.Ipfs{
	//				Upload: map[string]string{fmt.Sprintf("%s_%s", color, stoken_id): fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres)},
	//			}
	//			db.SaveIpfsm(Ipfs)
	//			log.Println("s3 upload success", fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres))
	//		}
	//	} else {
	//		if Ipfs.Upload[fmt.Sprintf("%s_%s", color, stoken_id)] == "" {
	//			//./原始图片
	//			CreateGif(label, color, savepath)
	//			raw := Read(savepath)
	//			Ipfsres, erripfs := UploadIPFS(raw)
	//			if erripfs != nil {
	//				log.Println("Ipfs Upload failed", erripfs)
	//			} else {
	//				Ipfs.Upload[fmt.Sprintf("%s_%s", color, stoken_id)] = fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres)
	//				db.SaveIpfsm(Ipfs)
	//				log.Println("s3 upload success", fmt.Sprintf("https://ipfs.io/ipfs/%s?filename=%s", Ipfsres, Ipfsres))
	//			}
	//		}
	//	}
	//}

}

func (a *Api) GetWhiteCard(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}
	data := map[string]map[string]string{}
	msg := "not found "
	db := ldb.GetLdb()
	white, _ := db.GetNotWhite()
	card, _ := db.GetCard()
	if white != nil {
		data["notinwhite"] = white.Name
	}
	if card != nil {
		data["card"] = card.Name
	}
	var respstr []byte
	log.Println(data)
	respstr, err := json.Marshal(data)
	if err == nil {
		msg = string(respstr)
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

type Respons struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (a *Api) GetTokenName(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}
	query := request.URL.Query()
	log.Println("GetTokenName query ", query)
	msg := "err"
	var (
		v  []string
		ok bool
	)
	name := ""
	tokenid := ""
	if v, ok = query["name"]; ok {
		name = v[0]
	} else {
		res := &Respons{
			Code:    "0",
			Message: "not found name",
		}
		respstr, err := json.Marshal(res)
		if err == nil {
			msg = string(respstr)
		}
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
		return
	}
	if v, ok = query["tokenid"]; ok {
		tokenid = v[0]
	} else {
		res := &Respons{
			Code:    "0",
			Message: "not found name",
		}
		respstr, err := json.Marshal(res)
		if err == nil {
			msg = string(respstr)
		}
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
		return
	}
	db := ldb.GetLdb()
	tokenName, _ := db.GetTokenName()
	if tokenName == nil {
		tokenName = &ldb.TokenIdName{TokenName: map[string]string{tokenid: name}}
	} else {
		tokenName.TokenName[tokenid] = name
	}
	err1 := db.SaveTokenName(tokenName)
	if err1 != nil {
		res := &Respons{
			Code:    "0",
			Message: "SaveTokenName err",
		}
		respstr, err := json.Marshal(res)
		if err == nil {
			msg = string(respstr)
		}
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
		return
	} else {
		log.Println(fmt.Sprintf("SaveTokenName %s:%s", tokenid, name))
	}

	var respstr []byte
	res := &Respons{
		Code:    "1",
		Message: tokenid + "_" + name,
	}
	respstr, err := json.Marshal(res)
	if err == nil {
		msg = string(respstr)
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}
