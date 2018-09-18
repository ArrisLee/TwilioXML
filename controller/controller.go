package controller

import (
	"encoding/xml"
	"net/http"

	"github.com/labstack/echo"
)

//GetXML func
//pass text and retrieve XML file - TwiML to make the call
func GetXML(c echo.Context) error {
	type TwiML struct {
		XMLName xml.Name `xml:"Response"`
		Say     string   `xml:"Say"`
	}
	text := c.QueryParam("text")
	if text == "" {
		text = "Hello, this is the call from, GOGO, you have a new order!"
	}
	twiml := TwiML{Say: text}
	return c.XML(http.StatusOK, twiml)
}
