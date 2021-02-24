package controller

import (
	"encoding/xml"
	"net/http"

	"github.com/labstack/echo"
)

//TwiXML func
//pass text and retrieve XML to make the call
func TwiXML(c echo.Context) error {
	//validate input, if nil, set defualt value
	text := c.QueryParam("text")
	shortID := c.QueryParam("shortID")
	if text == "" {
		text = "Hello, you have a new order, order number is"
	}
	if shortID == "" {
		shortID = "A A B B C C"
	}
	//compose XML structs
	type prosody struct {
		XMLName xml.Name `xml:"prosody"`
		ShortID string   `xml:",chardata"`
		Rate    string   `xml:"rate,attr"`
	}
	type say struct {
		Text    string  `xml:",chardata"`
		Prosody prosody `xml:",innerxml"`
	}
	type TwiML struct {
		XMLName xml.Name `xml:"Response"`
		Say     say      `xml:"Say"`
	}
	//assign input value to xml structs
	pr := prosody{}
	//set shortID
	pr.ShortID = shortID
	//set voice speed to 48%
	pr.Rate = "50%"
	sy := say{}
	//set basic text
	sy.Text = text
	sy.Prosody = pr
	twiml := TwiML{}
	twiml.Say = sy
	//retrun xml file
	return c.XMLPretty(http.StatusOK, twiml, " ")
}
