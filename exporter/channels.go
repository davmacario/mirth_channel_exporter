package exporter

import (
	"encoding/xml"
)

// ChannelIdNameMap represents the XML structure for channel IDs and names.
type ChannelIdNameMap struct {
	XMLName xml.Name       `xml:"map"`
	Entries []ChannelEntry `xml:"entry"`
}

// ChannelEntry represents a single entry in the ChannelIdNameMap.
type ChannelEntry struct {
	XMLName xml.Name `xml:"entry"`
	Values  []string `xml:"string"`
}

// ChannelStatsList represents the XML structure for channel statistics.
type ChannelStatsList struct {
	XMLName  xml.Name       `xml:"list"`
	Channels []ChannelStats `xml:"channelStatistics"`
}

// ChannelStats represents a single channel's statistics.
type ChannelStats struct {
	XMLName   xml.Name `xml:"channelStatistics"`
	ServerId  string   `xml:"serverId"`
	ChannelId string   `xml:"channelId"`
	Received  string   `xml:"received"`
	Sent      string   `xml:"sent"`
	Error     string   `xml:"error"`
	Filtered  string   `xml:"filtered"`
	Queued    string   `xml:"queued"`
}
