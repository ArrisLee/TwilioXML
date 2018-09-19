package controller

import (
	"encoding/xml"
	"net/http"

	"github.com/labstack/echo"
)

//TwiXML func
//pass text and retrieve XML to make the call
func TwiXML(c echo.Context) error {
	type TwiML struct {
		XMLName xml.Name `xml:"Response"`
		Say     string   `xml:"Say"`
	}
	text := c.QueryParam("text")
	if text == "" {
		text = "Q,W,D,3,2,1"
	}
	twiml := TwiML{Say: text}
	return c.XMLPretty(http.StatusOK, twiml, " ")
}
