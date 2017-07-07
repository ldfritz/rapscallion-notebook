package encoding_test

import (
	"encoding/xml"
	"fmt"
	"log"
)

func Example_encodingXMLBasic() {
	var xmlData = `
<?xml version="1.0" encoding="UTF-8"?>
<letter>
  <title maxlength="10"> Quote Letter </title>
  <salutation limit="40">Dear Daniel,</salutation>
  <text>Thank you f or sending us the information on <emphasis>SDL Trados Studio 2009</emphasis>.
    We like your products and think they certainly represent the most powerful translation solution on the market.
    We especially like the <component translate="yes">XML Parser rules</component> options in the <component translate="no">XML</component> filter.
    It has helped us to set up support for our XML files in a flash.
    We have already downloaded the latest version from your Customer Center.</text>
   <title maxlength="40"> Quote Details </title>
    <text> We would like to order 50 licenses.
    Please send us a quote. Keep up the good work!</text>
  
  <greetings minlength="10">Yours sincerely,</greetings>
  <signature> Paul Smith</signature>
  <address translate="yes">Smith &amp; Company Ltd.</address>
  <address translate="no">Smithtown</address>
  <weblink>http://www.smith-company-ltd.com</weblink>
  <logo alt="Logo of Smith and Company Ltd." address="http://www.smith-company-ltd.com/logo.jpg"/>
</letter>
`
	type Letter struct {
		Title      string `xml:"title"`
		Salutation string `xml:"salutation"`
	}
	var obj Letter
	err := xml.Unmarshal([]byte(xmlData), &obj)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(obj.Title[1:6])
	fmt.Println(obj.Salutation[0:4])
	// Output:
	// Quote
	// Dear
}