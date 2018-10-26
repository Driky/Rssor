package atom

import (
	"encoding/xml"
	"github.com/antchfx/htmlquery"
)

const (
	// A generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>`
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Xmlns   string   `xml:"xmlns,attr"`
	Version string   `xml:"version,attr"`
	Rel     string   `xml:"rel,attr"`
	Chn     Channel  `xml:"channel"`
}

type Channel struct {
	Titre string `xml:"title"`
	Itm   []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

func GetTrashscanlationsLastChapters() RSS {
	var items []Item
	doc, _ := htmlquery.LoadURL("https://trashscanlations.com/")
	for _, itemDetail := range htmlquery.Find(doc, "//div[@id='loop-content']//div[@class='item-summary']") {
		titleLink := htmlquery.FindOne(itemDetail, "//div[@class='post-title font-title']/h5/a")
		numberAndLink := htmlquery.FindOne(itemDetail, "//div[@class='list-chapter']/div[@class='chapter-item'][1]/span[@class='chapter font-meta']/a")
		if titleLink == nil || numberAndLink == nil {
			continue
		}

		title := htmlquery.InnerText(titleLink)
		number := htmlquery.InnerText(numberAndLink)
		link := htmlquery.SelectAttr(numberAndLink, "href")
		items = append(items, Item{title + number, link})
	}

	return RSS{
		Xmlns:   "http://91.121.84.94:8081/TrashscanlationsRSS",
		Version: "1",
		Rel:     "Self",
		Chn: Channel{
			Titre: "Trashscanlations",
			Itm:   items}}
}
