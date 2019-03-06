package cmd

import (
	"github.com/RudyChow/proxy/app/http"
	"github.com/RudyChow/proxy/app/utils/spiders"
	"github.com/RudyChow/proxy/app/utils/filters"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "proxy",
	Short: "proxypool",
	Long:  `blahblahblah`,
	Run: func(cmd *cobra.Command, args []string) {
		//更新可用的代理ip池
		go filters.UpdateUsefulProxy()

		//开启http服务
		go http.StartHttpServer()

		//开始采集数据
		spiders.StartCrawl()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
