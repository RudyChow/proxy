package schedules

import (
	"github.com/RudyChow/proxy/app/utils/filters"
	"github.com/RudyChow/proxy/app/utils/spiders"
	"time"
)

func Run() {
	oneMin := time.NewTicker(time.Minute)

	for {
		select {
		case <-oneMin.C:
			go filters.UpdateUsefulProxy()
			go spiders.StartCrawl()
		}
	}
}
