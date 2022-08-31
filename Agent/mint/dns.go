package mint

import (
	"context"
	"github.com/dnsdao/nft.pass/config"
	"github.com/dnsdao/nft.pass/database/ldb"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
	"time"
)

type DNSAgent struct {
	quit chan struct{}
}

var (
	CcurrentBlkNumber uint64
	McurrentBlkNumber uint64
	lastDnsLoopTime   int64
	dnsAgent          *DNSAgent
	dnsAgentOnce      sync.Once
)

func GetDNSAgent() *DNSAgent {
	if dnsAgent == nil {
		dnsAgentOnce.Do(func() {
			dnsAgent = &DNSAgent{
				quit: make(chan struct{}, 8),
			}
		})
	}
	return dnsAgent
}

func currentBlk(confT string) (uint64, error) {
	var (
		cli *ethclient.Client
		err error
	)

	cli, err = ethclient.Dial(config.GetRConf().GetRPCEndPoint(confT))
	if err != nil {
		return 0, err
	}

	defer cli.Close()

	return cli.BlockNumber(context.TODO())
}

func loopForNewName() {
	var err error
	db := ldb.GetLdb()
	cdbblk := db.GetLatestBlkNum("CBlkNum")
	CcurrentBlkNumber, err = currentBlk("c_conf")
	if err != nil {
		return
	}
	BatchNewColdBootClient("c_conf", cdbblk, CcurrentBlkNumber)
	//defer db.CloseLdb()
	db.SaveLatestBlkNum("CBlkNum", CcurrentBlkNumber)

	mdbblk := db.GetLatestBlkNum("MBlkNum")
	McurrentBlkNumber, err = currentBlk("m_conf")
	if err != nil {
		return
	}
	BatchNewColdBootClient("m_conf", mdbblk, McurrentBlkNumber)
	db.SaveLatestBlkNum("MBlkNum", McurrentBlkNumber)
}

func (ag *DNSAgent) DNSLoopEvent() {

	for {
		now := time.Now().Unix()

		if now-lastDnsLoopTime > 300 {
			lastDnsLoopTime = now
			loopForNewName()
		} else {
			time.Sleep(time.Second)
		}

		select {
		case <-ag.quit:
			return
		default:
			//continue
		}
	}
}

func (ag *DNSAgent) Quit() {
	ag.quit <- struct{}{}
}
