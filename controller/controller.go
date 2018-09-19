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
		text = "Hello, this is a call from gogo, you have a new oder"
	}
	twiml := TwiML{Say: text}
	return c.XMLPretty(http.StatusOK, twiml, " ")
}
