package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kprc/nbsnetwork/tools"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
)

const (
	resolveHome = ".passcard"
	confFile    = "conf.json"
	cmdSock     = ".cmd-sock"
)

var (
	resolveconf       *RConfig
	resolveconfOnceDo sync.Once
)

func defaultConf() *RConfig {
	return &RConfig{
		WConf: &WebSeverConf{
			ListenServer: ":30030",
		},
		Cconf: &ContractConf{
			ChainName:     "rinkeby",
			ChainID:       4,
			ColdBoot:      "0xD3F2Aa878014Ae12496368917719514e4605BA9F",
			RpcEndPoint:   "https://rinkeby.infura.io/v3/",
			WsRPCEndPoint: "wss://rinkeby.infura.io/ws/v3/",
			RpcEndPointID: "ced16671f5894b2796224e49062999ca",
		},
		Mconf: &ContractConf{
			ChainName:     "mainnet",
			ChainID:       1,
			ColdBoot:      "0x22871b977AAe43d44FE50dF03f632134c3e3e490",
			RpcEndPoint:   "https://mainnet.infura.io/v3/",
			WsRPCEndPoint: "wss://mainnet.infura.io/ws/v3/",
			RpcEndPointID: "ced16671f5894b2796224e49062999ca",
		},
		SConf: &SysConf{
			DBPath: "ldb",
		},
		IConf: &IpfsConf{
			HashPath: "./hash.json",
		},
		Mysql: &DatabaseConf{
			Host:   "172.81.40.90:3306",
			Driver: "mysql",
			DbName: "white",
			User:   "root",
			Passwd: "Asimov@0518",
		},
	}
}
func (rc *RConfig) GetMysqlDb() *DatabaseConf {
	return rc.Mysql
}
func (rc *RConfig) GetIpfsHash() string {
	return rc.IConf.HashPath
}

func (rc *RConfig) GetDBPath() string {
	return path.Join(ResolveHome(), rc.SConf.DBPath)
}

func (rc *RConfig) GetContractConf(confType string) string {
	switch confType {
	case "c_conf":
		return rc.Cconf.ColdBoot
	case "m_conf":
		return rc.Mconf.ColdBoot
	}
	return ""
}

func (rc *RConfig) GetRPCEndPoint(confType string) string {
	switch confType {
	case "c_conf":
		return rc.Cconf.RpcEndPoint + rc.Cconf.RpcEndPointID
	case "m_conf":
		return rc.Mconf.RpcEndPoint + rc.Mconf.RpcEndPointID
	}
	return ""
}

func (rc *RConfig) GetWsRPCENDPoint(confType string) string {
	switch confType {
	case "c_conf":
		return rc.Cconf.WsRPCEndPoint + rc.Cconf.RpcEndPointID
	case "m_conf":
		return rc.Mconf.WsRPCEndPoint + rc.Mconf.RpcEndPointID
	}
	return ""
}

func GetRConf() *RConfig {
	if resolveconf == nil {
		resolveconfOnceDo.Do(func() {
			resolveconf = defaultConf()
			resolveconf.load()
			resolveconf.compareAndSave()
		})
	}

	return resolveconf
}

func InitConfig() {
	nc := defaultConf()
	nc.save()
}

func (nc *RConfig) compareAndSave() {
	data, err := json.Marshal(*nc)
	if err != nil {
		return
	}

	ncdefault := defaultConf()

	var dataDefault []byte
	dataDefault, err = json.Marshal(ncdefault)
	if err != nil {
		return
	}

	if cp := bytes.Compare(data, dataDefault); cp != 0 {
		nc.save()
		return
	}
	var dataload []byte

	ncload := &RConfig{}
	ncload.load()

	dataload, err = json.Marshal(ncload)
	if err != nil {
		return
	}

	if cp := bytes.Compare(dataload, dataDefault); cp != 0 {
		nc.save()
		return
	}

}

func ResolveHome() string {
	h, _ := tools.Home()

	return path.Join(h, resolveHome)
}

func ConfFile() string {
	return path.Join(ResolveHome(), confFile)
}

func CmdSockFile() string {
	return path.Join(ResolveHome(), cmdSock)
}

type WebSeverConf struct {
	ListenServer string `json:"listen_server"`
}
type DatabaseConf struct {
	Host   string `json:"host"`
	Driver string `json:"driver"`
	DbName string `json:"db_name"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
}
type ContractConf struct {
	ChainID       int64
	ChainName     string
	RpcEndPoint   string
	WsRPCEndPoint string
	RpcEndPointID string
	ColdBoot      string
}

type SysConf struct {
	DBPath string
}

type IpfsConf struct {
	HashPath string
}
type RConfig struct {
	WConf *WebSeverConf `json:"w_conf"`
	Cconf *ContractConf `json:"c_conf"`
	Mconf *ContractConf `json:"m_conf"`
	SConf *SysConf      `json:"s_conf"`
	IConf *IpfsConf     `json:"i_conf"`
	Mysql *DatabaseConf `json:"mysql"`
}

func (nc *RConfig) load() error {
	cfile := ConfFile()

	fdata, err := tools.OpenAndReadAll(cfile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(fdata, nc); err != nil {
		return err
	}

	return nil
}

func (nc *RConfig) save() error {
	cfile := ConfFile()

	if j, err := json.MarshalIndent(*nc, "\t", " "); err != nil {
		return err
	} else {
		return tools.Save2File(j, cfile)
	}
}

type ReadIpfsHash struct {
	IpfsHash map[string]string
}

func Read(filepath string) []byte {
	f, err := os.Open(filepath)
	if err != nil {
		log.Println("read file fail", err)
		return nil
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("read to fd fail", err)
		return nil
	}

	return fd
}

func NewReadIpfsHash() *ReadIpfsHash {
	f := Read(GetRConf().GetIpfsHash())
	hashmap := map[string]string{}
	err1 := json.Unmarshal(f, &hashmap)
	if err1 != nil {
		fmt.Println(err1)
	}
	return &ReadIpfsHash{
		hashmap,
	}
}

//
var ReadIpfsHashInstance *ReadIpfsHash

//var listenServerMutex sync.Mutex
var listenServerInstanceOnceDo sync.Once

func ReadIpfsHashServer() *ReadIpfsHash {
	if ReadIpfsHashInstance == nil {
		//listenServerMutex.Lock()
		//defer listenServerMutex.Unlock()
		//if listenServerInstance == nil {
		//	listenServerInstance = NewListenServer()
		//}
		listenServerInstanceOnceDo.Do(func() {
			ReadIpfsHashInstance = NewReadIpfsHash()
		})
	}
	return ReadIpfsHashInstance
}
