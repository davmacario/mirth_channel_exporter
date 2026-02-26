package exporter

import "encoding/xml"

type DatabaseTasks struct {
	XMLName xml.Name            `xml:"map"`
	Entries []DatabaseTaskEntry `xml:"entry"`
}

type DatabaseTaskEntry struct {
	Id                  string   `xml:"id"`
	Status              string   `xml:"status"`
	Name                string   `xml:"name"`
	Description         string   `xml:"description"`
	ConfirmationMessage string   `xml:"confirmationMessage"`
	AffectedChannels    []string `xml:"affectedChannels"`
	StartDateTime       string   `xml:"startDateTime"`
}
