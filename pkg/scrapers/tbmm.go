package scrapers

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"github.com/eferhatg/hayri-irdal/pkg/models"
	"github.com/eferhatg/hayri-irdal/pkg/utils"
)

func CrawlTBMMDeputyList(deputyListPage string) []models.Deputy {
	logr := log.WithFields(log.Fields{"action": "Crawling deputy"})
	doc := utils.GetDocument(deputyListPage)

	deputies := []models.Deputy{}
	c := 0
	doc.Find(".grid_12 table tbody tr").Each(func(i int, tr *goquery.Selection) {

		bgcolor, exist := tr.Attr("bgcolor")
		if exist && bgcolor == "#FFFFFF" {

			tdName := tr.Find("td")
			name, surname := parseFullName(tdName.Get(0).FirstChild.FirstChild.Data)
			pLink, _ := tdName.Find("a").Attr("href")
			deputyNo, _ := strconv.Atoi(strings.Split(pLink, "p_sicil=")[1])
			termNo, _ := strconv.Atoi(strings.Split(pLink, "p_donem=")[1][:2])
			party := tdName.Get(2).FirstChild.Data
			deputy := models.Deputy{Name: name, Surname: surname, DeputyNo: deputyNo, TermNo: termNo, Party: party}

			logr.Debug("Deputy found", deputy)
			deputies = append(deputies, deputy)
			c++
		}
	})
	return deputies
}

func parseFullName(fullname string) (name string, surname string) {

	nameArr := strings.Split(fullname, " ")

	for i := 0; i < len(nameArr); i++ {
		part := nameArr[i]

		runes := []rune(part)

		lastChar := runes[len(runes)-1:][0]

		if unicode.IsUpper(lastChar) {
			surname += part + " "
		} else {
			name += part + " "
		}
	}
	return strings.TrimSuffix(name, " "), strings.TrimSuffix(surname, " ")
}
