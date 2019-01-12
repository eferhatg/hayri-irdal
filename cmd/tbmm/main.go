package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/eferhatg/hayri-irdal/pkg/scrapers"
)

const (
	deputyListPage = "https://www.tbmm.gov.tr/develop/owa/milletvekillerimiz_sd.liste"
)

func main() {

	log.SetLevel(log.DebugLevel)
	deplist := scrapers.CrawlTBMMDeputyList(deputyListPage)
	log.Info(len(deplist))
}
