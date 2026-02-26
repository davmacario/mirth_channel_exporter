package exporter

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	channelIdNameAPI = "/channels/idsAndNames"
	channelStatsAPI  = "/channels/statistics"
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
	Received  int64    `xml:"received"`
	Sent      int64    `xml:"sent"`
	Error     int64    `xml:"error"`
	Filtered  int64    `xml:"filtered"`
	Queued    int64    `xml:"queued"`
}

// gatherMirthChannelStats fetches channel statistics and updates Prometheus metrics.
func (e *Exporter) gatherMirthChannelStats(channelIDNameMap map[string]string, ch chan<- prometheus.Metric) error {
	resp, err := e.makeMirthAPIRequest(context.Background(), channelStatsAPI)
	if err != nil {
		return fmt.Errorf("failed to perform HTTP request for channel statistics: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK status code %d from channel statistics API: %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body for channel statistics: %w", err)
	}

	var channelStatsList ChannelStatsList
	if err := xml.Unmarshal(body, &channelStatsList); err != nil {
		return fmt.Errorf("failed to unmarshal XML for channel statistics: %w", err)
	}

	for _, channel := range channelStatsList.Channels {
		channelName, ok := channelIDNameMap[channel.ChannelId]
		if !ok {
			log.Printf("WARNING: Channel ID '%s' not found in ID-name map. Skipping metrics for this channel.", channel.ChannelId)
			continue
		}

		// Helper function to parse and send metric
		sendMetric := func(desc *prometheus.Desc, value int64) {
			valueFloat := float64(value)
			ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, valueFloat, channelName)
		}

		sendMetric(descs["messages_received_total"], channel.Received)
		sendMetric(descs["messages_sent_total"], channel.Sent)
		sendMetric(descs["messages_errored_total"], channel.Error)
		sendMetric(descs["messages_filtered_total"], channel.Filtered)
		sendMetric(descs["messages_queued"], channel.Queued)
	}

	return nil
}

// loadChannelIDNameMap fetches channel IDs and names from Mirth and returns them as a map.
func (e *Exporter) loadChannelIDNameMap() (map[string]string, error) {
	channelIDNameMap := make(map[string]string)

	resp, err := e.makeMirthAPIRequest(context.Background(), channelIdNameAPI)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request for channel ID names: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK status code %d from channel ID names API: %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body for channel ID names: %w", err)
	}

	var channelIDNameMapXML ChannelIdNameMap
	if err := xml.Unmarshal(body, &channelIDNameMapXML); err != nil {
		return nil, fmt.Errorf("failed to unmarshal XML for channel ID names: %w", err)
	}

	for _, entry := range channelIDNameMapXML.Entries {
		if len(entry.Values) == 2 {
			channelIDNameMap[entry.Values[0]] = entry.Values[1]
		} else {
			log.Printf("WARNING: Unexpected number of values (%d) in channel ID name entry: %v", len(entry.Values), entry.Values)
		}
	}

	return channelIDNameMap, nil
}
