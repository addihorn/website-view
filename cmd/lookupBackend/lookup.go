package lookup

import (
	"fmt"
	"website-graph/internal/diagram"
	"website-graph/internal/website"
)

var maxDepth int
var ignoreList []string
var sitemap map[string]*website.Page

type Sitemap struct {
	sitemap map[string]*website.Page
}

var sitemapList map[string]*Sitemap

func newSitemap(baseUrl string) *Sitemap {

	newSitemap := &Sitemap{sitemap: make(map[string]*website.Page)}
	newSitemap.sitemap[baseUrl] = website.NewPage(baseUrl)
	return newSitemap
}

func StartUp() {
	sitemapList = make(map[string]*Sitemap)
}

func DoSearch(startUrl string, initialUrl string) string {

	if sitemapList[initialUrl] == nil {
		sitemapList[initialUrl] = newSitemap(startUrl)
	}

	maxDepth = 1

	startPage := sitemapList[initialUrl].sitemap[startUrl]

	todaysMem := make(map[string]*website.Page)
	enumeratePage(startPage, todaysMem, sitemapList[initialUrl].sitemap, 0)

	fmt.Println(todaysMem)

	return diagram.GenerateGraphViz(todaysMem)

}

func enumeratePage(page *website.Page, memory map[string]*website.Page, sitemap map[string]*website.Page, depthFromSource int) {

	if depthFromSource > maxDepth && memory[page.Path] == nil {
		website.CheckForAvailability(page)
		memory[page.Path] = page
		return
	}

	if memory[page.Path] != nil {
		//we already visited and enumerated this page
		return
	}

	website.ReadOutgoingLinks(page, sitemap)
	memory[page.Path] = page
	for _, link := range page.GetLinkList() {
		enumeratePage(link, memory, sitemap, depthFromSource+1)
	}

}
