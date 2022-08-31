package mint

import (
	"context"
	"fmt"
	"github.com/dnsdao/nft.pass/config"
	"github.com/dnsdao/nft.pass/contract/dnscontract"
	"github.com/dnsdao/nft.pass/database/ldb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

const svg = "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<!-- Generator: Adobe Illustrator 26.0.1, SVG Export Plug-In . SVG Version: 6.00 Build 0)  -->\n<svg version=\"1.1\" id=\"图层_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" x=\"0px\" y=\"0px\"\n\t viewBox=\"0 0 1200 1200\" style=\"enable-background:new 0 0 1200 1200;\" xml:space=\"preserve\">\n<style type=\"text/css\">\n\t.st0{fill:#FFFFFF;}\n\t.st1{font-family:'PingFang-SC-Heavy';}\n\t.st2{font-size:81.3892px;}\n</style>\n<image style=\"overflow:visible;\" width=\"1200\" height=\"1200\" xlink:href=\"%s\" >\n</image>\n<text transform=\"matrix(1 0 0 1 58.6783 129.0048)\" class=\"st0 st1 st2\">%s</text>\n</svg>"

func NewColdBootClient(confT string) (*dnscontract.ColdBoot, *ethclient.Client, error) {
	cli, err := ethclient.Dial(config.GetRConf().GetRPCEndPoint(confT))
	log.Println(fmt.Sprintf("%s client conf :%s", confT, config.GetRConf().GetRPCEndPoint(confT)))
	if err != nil {
		cli.Close()
		log.Println("NewColdBootClient", err)
		return nil, nil, err
	}

	coldBoot, err := dnscontract.NewColdBoot(common.HexToAddress(config.GetRConf().GetContractConf(confT)), cli)
	log.Println(fmt.Sprintf("%s ContractConf :%s", confT, config.GetRConf().GetContractConf(confT)))
	if err != nil {
		log.Println("NewColdBoot", err)
		return nil, nil, err
	}

	return coldBoot, cli, nil
}
func BatchNewColdBootClient(confT string, start, end uint64) {
	ColorList := []string{"noColor", "color", "gold", "green"}
	coldBoot, cli, err := NewColdBootClient(confT)
	if err != nil {
		log.Println("BatchNewColdBootClient", err)
		return
	}
	defer cli.Close()
	b, e := start, end
	log.Println("BatchNewColdBootClient start block", b)
	op := &bind.FilterOpts{
		Start:   b,
		End:     &e,
		Context: context.TODO(),
	}
	var iter *dnscontract.ColdBootEVMintCardIterator
	iter, err = coldBoot.FilterEVMintCard(op)
	if err != nil {
		log.Println("BatchNewColdBootClient", err)
		return
	}
	db := ldb.GetLdb()
	defer iter.Close()
	for iter.Next() {
		ev := iter.Event
		log.Println(ev.User, ev.CardId.String(), ColorList[ev.Color], ev.Color)
		cardId := ev.CardId.String()
		cn := strings.Split(confT, "_")
		err1 := db.SaveMintCard(cn[0]+cardId, ev.User.String()+"_"+ColorList[ev.Color])
		if err1 != nil {
			log.Println("BatchNewColdBootClient ", "save to db error")
			continue
		}
		label := "NO."
		switch len(cardId) {
		case 1:
			label += "000" + cardId
		case 2:
			label += "00" + cardId
		case 3:
			label += "0" + cardId
		case 4:
			label += cardId
		}
		if ev.Color > 0 {
			//path := "./passcard/" + cn[0] + ev.User.String() + "_" + cardId + ".svg"
			//info := fmt.Sprintf(svg, "http://31.12.95.78:8080/images/"+ColorList[ev.Color]+".gif", label)
			//writeFile(info, path)
			// 存储 钱包地址_cardId
			//CreatePassImage("./image/"+ColorList[ev.Color]+".png", "./passcard/"+cn[0]+ev.User.String()+"_"+cardId+".png", label)
		}
	}
}
