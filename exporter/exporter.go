package exporter

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	// DefaultHTTPClientTimeout is the default timeout for HTTP client requests.
	DefaultHTTPClientTimeout = 10 * time.Second
)

var (
	// httpClient is a shared HTTP client with insecure TLS verification for Mirth.
	httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: DefaultHTTPClientTimeout,
	}
	// Prometheus descs
	descs = make(map[string]*prometheus.Desc)
)

type metricDef struct {
	name        string
	help        string
	extraLabels []string
	typeLabel   string
}

func init() {
	const namespace = "mirth"

	metrics := []metricDef{
		{"up", "Was the last Mirth query successful.", nil, "status"},
		{"messages_received_total", "How many messages have been received (per channel).", []string{"channel"}, "channels"},
		{"messages_filtered_total", "How many messages have been filtered (per channel).", []string{"channel"}, "channels"},
		{"messages_queued", "How many messages are currently queued (per channel).", []string{"channel"}, "channels"},
		{"messages_sent_total", "How many messages have been sent (per channel).", []string{"channel"}, "channels"},
		{"messages_errored_total", "How many messages have errored (per channel).", []string{"channel"}, "channels"},
		{"cpu_usage_pct", "CPU usage percentage.", nil, "system"},
		{"allocated_memory_bytes", "Allocated memory in bytes.", nil, "system"},
		{"free_memory_bytes", "Free memory in bytes.", nil, "system"},
		{"disk_total_bytes", "Total disk space in bytes.", nil, "system"},
		{"disk_free_bytes", "Free disk space in bytes.", nil, "system"},
		{"num_db_tasks", "Number of database tasks.", nil, "database"},
	}

	for _, m := range metrics {
		desc := prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", m.name),
			m.help,
			m.extraLabels,
			prometheus.Labels{"type": m.typeLabel},
		)
		descs[m.name] = desc
	}
}

type Exporter struct {
	mirthEndpoint string
	mirthUsername string
	mirthPassword string
	lastMetrics   *MetricsResponse
}

type MetricsResponse struct {
	systemStats  *SystemStats
	channelStats *ChannelStats
}

// NewExporter creates a new Exporter instance.
func NewExporter(mirthEndpoint, mirthUsername, mirthPassword string) *Exporter {
	return &Exporter{
		mirthEndpoint: mirthEndpoint,
		mirthUsername: mirthUsername,
		mirthPassword: mirthPassword,
	}
}

// Describe sends the metric descriptions to the provided channel.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(e, ch)
}

// Collect fetches the Mirth statistics and delivers them as Prometheus metrics.
// This is the main entry point for the Prometheus client to collect metrics.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	channelIDNameMap, err := e.loadChannelIDNameMap()
	if err != nil {
		log.Printf("ERROR: Failed to load channel ID to name map: %v", err)
		ch <- prometheus.MustNewConstMetric(descs["up"], prometheus.GaugeValue, 0)
		return
	}

	ch <- prometheus.MustNewConstMetric(descs["up"], prometheus.GaugeValue, 1) // Mirth API is accessible => Mirth is up

	if err := e.gatherMirthChannelStats(channelIDNameMap, ch); err != nil {
		log.Printf("ERROR: Failed to collect channel statistics: %v", err)
		// Optionally set 'up' to 0 if subsequent collection fails after initial success
		// ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 0)
	} else {
		log.Println("Successfully scraped Mirth endpoint (channels).")
	}

	// Collect system statistics
	if err := e.gatherSystemStats(ch); err != nil {
		log.Printf("ERROR: Failed to collect system statistics: %v", err)
	} else {
		log.Println("Successfully scraped Mirth endpoint (system).")
	}

	// Collect DB stats
	if err := e.gatherDatabaseTasks(ch); err != nil {
		log.Printf("ERROR: Failed to collect database task statistics: %v", err)
	} else {
		log.Println("Successfully scraped Mirth endpoint (database).")
	}
}

// Submit request to API endpoint specified in `path`
func (e *Exporter) makeMirthAPIRequest(ctx context.Context, path string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, e.mirthEndpoint+path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for channel ID names: %w", err)
	}
	req.SetBasicAuth(e.mirthUsername, e.mirthPassword)
	req.Header.Set("X-Requested-With", "OpenAPI")
	return httpClient.Do(req)
}
