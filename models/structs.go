package models

import "encoding/xml"

// RenderModel encapsulates the data used to render the content in a go template
type RenderModel struct {
	Class  string
	Object string
	ID     string
	Name   string
}

// Issue encapsulates an error when validating a glade file
type Issue struct {
	Description string
}

// Glade represents a glade xml file
type Glade struct {
	XMLName  xml.Name `xml:"interface"`
	Text     string   `xml:",chardata"`
	Requires struct {
		Text    string `xml:",chardata"`
		Lib     string `xml:"lib,attr"`
		Version string `xml:"version,attr"`
	} `xml:"requires"`
	Object []Object `xml:"object"`
}

// Object represents a GTK object
type Object struct {
	Text     string     `xml:",chardata"`
	Class    string     `xml:"class,attr"`
	ID       string     `xml:"id,attr"`
	Property []Property `xml:"property"`
	Child    []Child    `xml:"child"`
}

// Child is a child of a GTK object
type Child struct {
	XMLName xml.Name `xml:"child"`
	Text    string   `xml:",chardata"`
	Object  []Object `xml:"object"`
	Packing Packing  `xml:"packing"`
}

// Packing is the positioning of a GTK object
type Packing struct {
	Text     string     `xml:",chardata"`
	Property []Property `xml:"property"`
}

// Property of a GTK object
type Property []struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}
