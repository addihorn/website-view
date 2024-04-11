package website

import "net/url"

type Page struct {
	Path         string
	isSameServer bool
	Available    bool
	links        map[string]*Page
	images       map[string]*Image
	Url          *url.URL
}

type Image struct {
	Path string
	Url  *url.URL
}

func NewPage(URL string) *Page {
	uri, _ := url.Parse(URL)
	return &Page{Path: URL,
		isSameServer: true,
		Available:    false,
		links:        make(map[string]*Page),
		images:       make(map[string]*Image),
		Url:          uri}
}

func (page *Page) AddLink(URL string) *Page {
	page.links[URL] = NewPage(URL)
	return page.links[URL]
}

func (page *Page) LinkPage(ref *Page) {
	page.links[ref.Path] = ref
}

func (page *Page) GetLinkList() map[string]*Page {
	return page.links
}

func NewImage(URL string) *Image {
	uri, _ := url.Parse(URL)
	return &Image{Path: URL, Url: uri}
}

func (page *Page) AddImageLink(URL string) {
	page.images[URL] = NewImage(URL)
}

func (page *Page) LinkImage(ref *Image) {
	page.images[ref.Path] = ref
}

func (page *Page) GetImageList() map[string]*Image {
	return page.images
}
