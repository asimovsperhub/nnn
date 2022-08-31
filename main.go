package main

import (
	"fmt"
	"github.com/dnsdao/nft.pass/Agent/mint"
	"github.com/dnsdao/nft.pass/cmd"
	"github.com/dnsdao/nft.pass/cmd/cmdclient"
	"github.com/dnsdao/nft.pass/cmd/cmdserver"
	"github.com/dnsdao/nft.pass/config"
	"github.com/dnsdao/nft.pass/database/ldb"
	"github.com/dnsdao/nft.pass/server/httpserver"
	"github.com/kprc/nbsnetwork/tools"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"
)

const (
	PidFileName = ".pid"
	Version     = "0.0.1"
)

var param = struct {
	version bool
	passwd  string
}{}

var rootCmd = &cobra.Command{
	Use: "dnsr",

	Short: "mint resolve node",

	Long: `usage description::TODO::`,

	Run: mainRun,
}

func init() {

	flags := rootCmd.Flags()

	flags.BoolVarP(&param.version, "version",
		"v", false, "dnsr -v")

	rootCmd.AddCommand(cmdclient.ShowCmd)
	rootCmd.AddCommand(cmd.InitCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func mainRun(_ *cobra.Command, _ []string) {

	if b := tools.FileExists(config.ResolveHome()); !b {
		fmt.Println("please initial mint resolve server first!")
		return
	}

	if param.version {
		fmt.Println(Version)
		return
	}

	//go server.GetServerInstance().StartDaemon()

	go cmdserver.StartCmdService()
	go mint.GetDNSAgent().DNSLoopEvent()
	go httpserver.StartWebDaemon()

	waitShutdownSignal()
}

func waitShutdownSignal() {

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>mint resolve start at pid(%s)<<<<<<<<<<\n", pid)

	pidfile := path.Join(config.ResolveHome(), PidFileName)

	if err := ioutil.WriteFile(pidfile, []byte(pid), 0644); err != nil {
		fmt.Println("failed to write running pid", err)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1,
		syscall.SIGUSR2)

	sig := <-sigCh

	//server.GetServerInstance().StopDaemon()

	cmdserver.StopCmdService()

	ldb.GetLdb().CloseLdb()
	mint.GetDNSAgent().Quit()
	httpserver.StopWebDaemon()

	fmt.Printf("\n>>>>>>>>>>process finished(%s)<<<<<<<<<<\n", sig)
}
