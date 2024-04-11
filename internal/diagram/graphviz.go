package diagram

import (
	"fmt"
	"os"
	"website-graph/internal/website"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func GenerateGraphViz(sitemap map[string]*website.Page) string {
	g := graph.New(graph.StringHash, graph.Directed())

	for _, page := range sitemap {
		attributes := make(map[string]string)

		attributes["style"] = "filled"
		attributes["URL"] = fmt.Sprintf("%s?remoteUrl=%s", "", page.Url.String())
		attributes["fillcolor"] = "green"

		if !page.Available {
			attributes["fillcolor"] = "red"
		}
		g.AddVertex(page.Path, graph.VertexAttributes(attributes))
	}

	for _, page := range sitemap {
		for _, link := range page.GetLinkList() {
			g.AddEdge(page.Path, link.Path)
		}
	}

	//g1, _ := graph.TransitiveReduction(g)

	file, _ := os.Create("my-graph.gv")
	//buf := new(strings.Builder)
	//_ = draw.DOT(g1, buf)
	//fmt.Println(buf.String())
	_ = draw.DOT(g, file)

	/*
		krokiResponse, _ := http.Post("https://kroki.io/graphviz/svg", "text/plain", strings.NewReader(buf.String()))

		b, _ := io.ReadAll(krokiResponse.Body)

		fmt.Println(string(b))

		return string(b)
	*/
	return "foo"
}
