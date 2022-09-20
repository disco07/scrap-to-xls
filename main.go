package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	Name, Title, Company string
}

func scrap() ([]*Data, error) {
	var data []*Data

	page := 0
	for page < 211 {
		var d Data
		res, err := http.Get(fmt.Sprintf("https://a3f.fr/fr/annuaire_temp.php?all&depart_annuaire=%v#div-annuaire", page))
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, err
		}
		doc.Find(".card-content").Each(func(i int, s *goquery.Selection) {
			var name string
			var title string
			var company string

			s.Find(".truncate").Each(func(i int, s2 *goquery.Selection) {
				switch i {
				case 0:
					name = s2.Text()
				case 1:
					job := strings.SplitN(s.Text(), "-", 2)
					if len(job) > 1 {
						title = job[0]
						company = job[1]
					}
				}
				d.Name = name
				d.Title = title
				d.Company = company
			})
			data = append(data, &d)
		})
		page += 52
	}
	return data, nil
}

func excelize(data []*Data) error {

	return nil
}

func main() {
	data, err := scrap()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(data)
}
