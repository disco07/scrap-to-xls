package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xuri/excelize/v2"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	Name, Title, Company string
}

func scrap(path string) ([]*Data, error) {
	var data []*Data

	page := 0
	for page < 211 {
		res, err := http.Get(path + fmt.Sprintf("=%v", page))
		if err != nil {
			return nil, errors.New("unsupported protocol scheme")
		}
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, err
		}
		doc.Find(".card-content").Each(func(i int, s *goquery.Selection) {
			var d Data
			var name string
			var title string
			var company string

			s.Find(".truncate").Each(func(i2 int, s2 *goquery.Selection) {
				switch i2 {
				case 0:
					name = s2.Text()
					d.Name = name
				case 1:
					job := strings.SplitN(strings.Split(s.Text(), "\n")[2], "-", 2)
					if len(job) > 1 {
						title = job[0]
						company = job[1]
					}
					d.Title = title
					d.Company = company
				}
			})
			data = append(data, &d)
		})
		page += 52
	}
	return data, nil
}

func excelizeData(data []*Data) error {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Name")
	f.SetCellValue("Sheet1", "B1", "Title")
	f.SetCellValue("Sheet1", "C1", "Company")
	for i, d := range data {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", i+2), d.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", i+2), d.Title)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", i+2), d.Company)
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs("export_dataframe.xlsx"); err != nil {
		return err
	}

	return nil
}

func main() {
	site := "https://a3f.fr/fr/annuaire_temp.php?all&depart_annuaire"
	data, err := scrap(site)
	if err != nil {
		log.Fatal(err)
	}

	err = excelizeData(data)
	if err != nil {
		log.Fatal(err)
	}
}
