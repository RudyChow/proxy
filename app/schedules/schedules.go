package schedules

import (
	"github.com/RudyChow/proxy/app/utils/filters"
	"github.com/RudyChow/proxy/app/utils/spiders"
	"time"
)

func Run() {
	oneMin := time.NewTicker(time.Minute)
	tenSec := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-oneMin.C:
		case <-tenSec.C:
			go spiders.StartCrawl()
			go filters.UpdateUsefulProxy()
		}
	}
}
