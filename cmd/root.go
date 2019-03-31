package cmd

import (
	"github.com/RudyChow/proxy/app/http"
	"github.com/RudyChow/proxy/app/rpc"
	"github.com/RudyChow/proxy/app/schedules"
	"github.com/RudyChow/proxy/config"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "proxy",
	Short: "proxypool",
	Long:  `blahblahblah`,
	Run: func(cmd *cobra.Command, args []string) {
		go schedules.Run()

		//开启http服务
		if(config.Conf.Http.Enable){
			go http.StartHttpServer()
		}


		//开启rpc服务
		if(config.Conf.Rpc.Enable){
			go rpc.StartRpcServer()
		}


		select{}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}
