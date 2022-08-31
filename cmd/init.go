package cmd

import (
	"fmt"
	"github.com/dnsdao/nft.pass/config"
	"github.com/dnsdao/nft.pass/database/ldb"
	"github.com/kprc/nbsnetwork/tools"
	"github.com/spf13/cobra"
	"os"
)

var InitCmd = &cobra.Command{
	Use: "init",

	Short: "netdev init  node",

	Long: `usage description::TODO::`,

	Run: initNetDev,
}

func initNetDev(cmd *cobra.Command, args []string) {
	if b := tools.FileExists(config.ResolveHome()); !b {
		os.MkdirAll(config.ResolveHome(), 0755)
	} else {
		fmt.Println("old node config in the user home dir!!!!")
		return
	}

	config.InitConfig()

	fmt.Println("initial node success!")
	db := ldb.GetLdb()
	HashMap := config.ReadIpfsHashServer().IpfsHash
	Ipfs := &ldb.Ipfs{
		Upload: HashMap,
	}
	db.SaveIpfsc(Ipfs)
	db.SaveIpfsm(Ipfs)
	fmt.Println("initial ipfs data success!")

}
