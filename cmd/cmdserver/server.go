package cmdserver

import (
	"github.com/dnsdao/nft.pass/cmd/pbs"
	"github.com/dnsdao/nft.pass/config"
	"github.com/kprc/nbsnetwork/tools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

var (
	gListener net.Listener
)

func StartCmdService() {

	if b := tools.FileExists(config.CmdSockFile()); b {
		os.RemoveAll(config.CmdSockFile())
	}

	gListener, err := net.Listen("unix", config.CmdSockFile())
	if err != nil {
		panic(err)
	}

	cmdServer := grpc.NewServer()

	pbs.RegisterCmdServiceServer(cmdServer, cmdServerInstance)

	reflection.Register(cmdServer)

	if err := cmdServer.Serve(gListener); err != nil {
		panic(err)
	}
}

func StopCmdService() {
	gListener.Close()
}
