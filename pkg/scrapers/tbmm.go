package scrapers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	log "github.com/Sirupsen/logrus"
	"github.com/eferhatg/hayri-irdal/pkg/models"
	"github.com/eferhatg/hayri-irdal/pkg/utils"
)

func ScrapeTBMMDeputyList(deputyListPage string) []models.Deputy {
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

//CrawlDeputyProfilePage crawl and find out deputy profile details
func ScrapeDeputyProfilePage(d models.Deputy, deputyProfileRoot string) models.Deputy {
	logr := log.WithFields(log.Fields{"action": "Crawling deputy profile"})
	doc := utils.GetDocument(deputyProfileRoot + "?p_donem=" + strconv.Itoa(d.TermNo) + "&p_sicil=" + strconv.Itoa(d.DeputyNo))
	logr.Info(d.Name, d.Surname)
	if strings.Split(doc.Find("div#mv_ili").Text(), " ")[0] == "" {
		logr.Debug("Deputy not found", d)
		d.IsActive = false
		return d

	}

	doc.Find("div#iletisim_bilgi table tbody tr").EachWithBreak(func(i int, tr *goquery.Selection) bool {
		tdArr := tr.Find("td")
		if tdArr.Length() > 1 && tdArr.Find("span").Length() != 0 {
			fc := tdArr.Get(0).FirstChild.Data
			sc := tdArr.Get(1).FirstChild.Data
			if fc != "span" || sc != "span" {
				return false
			}
			key := tdArr.Get(0).FirstChild.FirstChild.Data
			val := tdArr.Get(1).FirstChild.NextSibling.Data
			if key == "Adres" {
				d.Address = strings.TrimSpace(val)
			} else if key == "Telefon" {
				d.Tel = strings.TrimSpace(val)
			} else if key == "Faks" {
				d.Fax = strings.TrimSpace(val)
			} else if key == "E-Posta" {
				d.Email = strings.TrimSpace(tdArr.Get(1).LastChild.FirstChild.Data)
			} else if key == "Web" {
				d.Web = strings.TrimSpace(tdArr.Get(1).LastChild.FirstChild.Data)
			} else {

				fmt.Printf("%s\n", strings.TrimSpace(tdArr.Get(1).FirstChild.Data))

				tdArr.Last().Find("a").Each(func(k int, a *goquery.Selection) {
					href, ext := a.Attr("href")
					if ext {
						if strings.Contains(href, "facebook.com") {
							d.FbLink = href
						}
						if strings.Contains(href, "twitter.com") {
							d.TwLink = href
						}
					}
				})

			}

		}
		return true

	})

	d.Region = strings.Split(doc.Find("div#mv_ili").Text(), " ")[2]

	d.ProfilePic, _ = doc.Find("div#fotograf_alani img").Attr("src")

	divGorev := doc.Find("div#mv_gorev")
	position := ""
	for index := 0; index < divGorev.Children().Length(); index++ {
		pos := strings.TrimSpace(divGorev.Children().Get(index).NextSibling.Data)
		if len(pos) > 0 {
			position += pos + "||"
			fmt.Printf("%+v\n", pos)
		}

	}
	d.Position = strings.TrimSuffix(position, "||")

	divCv := doc.Find("div#ozgecmis_icerik")

	cv := strings.TrimSpace(divCv.Text())
	d.BirthDay = parseBirthDate(cv)
	d.BirthPlace = parseBirtPlace(cv)
	d.Profession = parseProffesion(cv)

	d.Cv = cv

	logr.Debug("Deputy profile crawled", d)
	return d
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

func parseBirthDate(cv string) time.Time {
	m := make(map[string]string)
	m["Ocak"] = "01"
	m["Şubat"] = "02"
	m["Mart"] = "03"
	m["Nisan"] = "04"
	m["Mayıs"] = "05"
	m["Haziran"] = "06"
	m["Temmuz"] = "07"
	m["Ağustos"] = "08"
	m["Eylül"] = "09"
	m["Ekim"] = "10"
	m["Kasım"] = "11"
	m["Aralık"] = "12"

	delim := [8]string{"'de", "'da", "'te", "'ta", "?de", "?da", "?te", "?ta"}
	tc := time.Time{}
	for _, d := range delim {
		s := strings.Split(cv, d)
		dp := strings.Split(s[0], " ")
		if len(dp) == 3 {
			dpi, err := strconv.Atoi(dp[0])
			if err != nil {
				break
			}
			bd := dp[2] + "-" + m[dp[1]] + "-" + fmt.Sprintf("%02d", dpi)
			fmt.Println(bd)
			tc, _ = time.Parse("2006-01-02", bd)
			break
		}
	}
	return tc
}

func parseBirtPlace(cv string) string {

	pIx := -1

	places := [81]string{"Adana", "Adıyaman", "Afyon", "Ağrı", "Amasya", "Ankara", "Antalya", "Artvin",
		"Aydın", "Balıkesir", "Bilecik", "Bingöl", "Bitlis", "Bolu", "Burdur", "Bursa", "Çanakkale",
		"Çankırı", "Çorum", "Denizli", "Diyarbakır", "Edirne", "Elazığ", "Erzincan", "Erzurum", "Eskişehir",
		"Gaziantep", "Giresun", "Gümüşhane", "Hakkari", "Hatay", "Isparta", "Mersin", "İstanbul", "İzmir",
		"Kars", "Kastamonu", "Kayseri", "Kırklareli", "Kırşehir", "Kocaeli", "Konya", "Kütahya", "Malatya",
		"Manisa", "Kahramanmaraş", "Mardin", "Muğla", "Muş", "Nevşehir", "Niğde", "Ordu", "Rize", "Sakarya",
		"Samsun", "Siirt", "Sinop", "Sivas", "Tekirdağ", "Tokat", "Trabzon", "Tunceli", "Şanlıurfa", "Uşak",
		"Van", "Yozgat", "Zonguldak", "Aksaray", "Bayburt", "Karaman", "Kırıkkale", "Batman", "Şırnak",
		"Bartın", "Ardahan", "Iğdır", "Yalova", "Karabük", "Kilis", "Osmaniye", "Düzce"}
	sIx := 2000000000
	for index, p := range places {

		x := strings.Index(cv, p)

		if x != -1 && x < sIx {
			sIx = x
			pIx = index
		}
	}
	if pIx == -1 {

		return ""
	}
	return places[pIx]
}

func parseProffesion(cv string) string {

	pLine := strings.Split(cv, "\n")
	if len(pLine) > 1 {
		pBlock := strings.Split(pLine[1], ";")[0]
		if len(pBlock) > 220 {
			return pBlock[:220]
		}
		return pBlock
	}
	return ""
}

func parseUniversity(cv string) string {

	pLine := strings.Split(cv, "\n")
	if len(pLine) > 1 {
		pBlock := strings.Split(pLine[1], ";")[1]
		pUniList := strings.Split(pBlock, "Üniver")
		if len(pUniList) > 1 {
			return pBlock[:220]
		}
		return pUniList[0]

	}
	return ""
}
