package website

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/antchfx/htmlquery"
)

func CheckForAvailability(page *Page) error {
	fmt.Fprintf(os.Stderr, "[FINE] Trying to read from page %v \n", page.Path)
	if !page.Available {
		_, err := htmlquery.LoadURL(page.Path)

		if err != nil {
			page.Available = false
			newErr := errors.New(fmt.Sprintf("Could not load page %v", page.Path))
			return errors.Join(newErr, err)
		}
		page.Available = true
	}

	return nil
}

func ReadOutgoingLinks(page *Page, sitemap map[string]*Page) error {

	fmt.Fprintf(os.Stderr, "[FINE] Trying to read from page %v \n", page.Path)
	site, err := htmlquery.LoadURL(page.Path)
	if err != nil {
		page.Available = false
		newErr := errors.New(fmt.Sprintf("Could not load page %v", page.Path))
		return errors.Join(newErr, err)
	}
	page.Available = true

	//do not scrap from external servers
	if !page.isSameServer {
		return errors.New("Tried to enumerate external server")
	}

	links := htmlquery.Find(site, "//a")

	for _, link := range links {

		linkUrlString := htmlquery.SelectAttr(link, "href")
		linkUrl, _ := url.Parse(linkUrlString)

		if linkUrl.Path == "" && linkUrl.Fragment != "" {
			//check next link, this is an internal anchor
			continue
		}

		if linkUrl.Host == "" {
			linkUrl.Host = page.Url.Host

		}
		if linkUrl.Scheme == "" {
			linkUrl.Scheme = page.Url.Scheme
		}
		var linkedPage *Page
		if sitemap[linkUrl.String()] != nil {
			linkedPage = sitemap[linkUrl.String()]
			page.LinkPage(linkedPage)
		} else {
			linkedPage = page.AddLink(linkUrl.String())
			sitemap[linkedPage.Path] = linkedPage
		}

		//check of both source and target server are same
		if linkedPage.Url.Host != page.Url.Host {
			linkedPage.isSameServer = false
			linkedPage.Available = true
		}
	}
	return nil

}
