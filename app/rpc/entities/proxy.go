package entities

import (
	"github.com/RudyChow/proxy/app/io"
	"github.com/RudyChow/proxy/app/models"
)

type ProxyServer struct {}


type ProxyList struct {
	Data []models.ProxyShortcut
}

func (this *ProxyServer)GetBestProxy(req EmptyReq,res *models.ProxyShortcut) error {

	shortcut := io.Handler.GetBestUsefulProxyPool()

	res.Addr = shortcut.Addr
	res.Score = shortcut.Score


	return nil
}


func (this *ProxyServer)GetUsefulProxyList(req int,res *ProxyList) error {

	res.Data = io.Handler.GetShortcutFromUsefulProxyPool(int64(req))

	return nil
}