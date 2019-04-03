package main

import (
	"flag"
	"os"

	"github.com/linkchain/common/util/log"
	"github.com/robfig/cron"
	"hum.tan/server/server"
	"hum.tan/server/server/client"
)

func main() {
	logLevel := flag.Int("loglevel", 3, "log level")

	//init log
	log.Root().SetHandler(
		log.LvlFilterHandler(log.Lvl(*logLevel),
			log.StreamHandler(os.Stdout, log.TerminalFormat(true))))

	height := client.GetDBLastBlock()
	client.SyncBlockByHeight(height + 1)

	go func() {
		c := cron.New()
		spec := "0 0/2 * * * ?"
		_ = c.AddFunc(spec, func() {
			height := client.GetDBLastBlock()
			client.SyncBlockByHeight(height + 1)
		})
		log.Info("Start sync block and transaction info")
		c.Start()
	}()

	s := server.NewServer()
	s.Start()
}
