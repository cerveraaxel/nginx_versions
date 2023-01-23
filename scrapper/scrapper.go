package scrapper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

const NGINX_URL = "http://nginx.org/en/download.html"

type NginxVersions struct {
	Mainline string
	Stable   string
	Legacy   string
}

func GetVersions(c *gin.Context) {

	coll := colly.NewCollector()

	var versions []string

	coll.OnHTML("#content > table > tbody > tr > td:nth-child(2) > a:first-of-type", func(e *colly.HTMLElement) {

		// Obtain the text of the attributes and create a slice of all of them
		v := e.Text
		versions = append(versions, v)

	})

	coll.Visit(NGINX_URL)

	Versions := NginxVersions{
		Mainline: versions[0],
		Stable:   versions[1],
		Legacy:   versions[2],
	}

	c.IndentedJSON(http.StatusOK, Versions)

}
