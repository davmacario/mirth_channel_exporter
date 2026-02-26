package exporter

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

const systemStatsAPI = "/system/stats"

/* System stats XML example:
 *
 * <com.mirth.connect.model.SystemStats>
 *   <timestamp>
 *     <time>1772032297737</time>
 *     <timezone>Etc/UTC</timezone>
 *   </timestamp>
 *   <cpuUsagePct>0.04252184090624086</cpuUsagePct>
 *   <allocatedMemoryBytes>200278016</allocatedMemoryBytes>
 *   <freeMemoryBytes>111327256</freeMemoryBytes>
 *   <maxMemoryBytes>268435456</maxMemoryBytes>
 *   <diskFreeBytes>18790436864</diskFreeBytes>
 *   <diskTotalBytes>31526391808</diskTotalBytes>
 * </com.mirth.connect.model.SystemStats>
 */
type SystemStats struct {
	CPUUsagePct          float64              `xml:"cpuUsagePct"`
	AllocatedMemoryBytes int64                `xml:"allocatedMemoryBytes"`
	FreeMemoryBytes      int64                `xml:"freeMemoryBytes"`
	MaxMemoryBytes       int64                `xml:"maxMemoryBytes"`
	DiskFreeBytes        int64                `xml:"diskFreeBytes"`
	DiskTotalBytes       int64                `xml:"diskTotalBytes"`
	Timestamp            SystemStatsTimestamp `xml:"timestamp"`
}

type SystemStatsTimestamp struct {
	Time     int64  `xml:"time"`
	Timezone string `xml:"timezone"`
}

// Retrieve system metrics from Mirth and send them to Prometheus
func (e *Exporter) hitMirthStatsAPI(ch chan<- prometheus.Metric) error {
	resp, err := e.makeMirthAPIRequest(context.Background(), systemStatsAPI)
	if err != nil {
		return fmt.Errorf("failed to perform HTTP request for system statistics: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK status code %d from system statistics API: %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body for system statistics: %w", err)
	}

	var systemStats SystemStats
	if err := xml.Unmarshal(body, &systemStats); err != nil {
		return fmt.Errorf("failed to unmarshal XML for system statistics: %w", err)
	}

	ch <- prometheus.MustNewConstMetric(descs["cpu_usage_pct"], prometheus.GaugeValue, systemStats.CPUUsagePct, "system")
	ch <- prometheus.MustNewConstMetric(descs["allocated_memory_bytes"], prometheus.GaugeValue, float64(systemStats.AllocatedMemoryBytes), "system")
	ch <- prometheus.MustNewConstMetric(descs["free_memory_bytes"], prometheus.GaugeValue, float64(systemStats.FreeMemoryBytes), "system")
	ch <- prometheus.MustNewConstMetric(descs["disk_total_bytes"], prometheus.GaugeValue, float64(systemStats.DiskTotalBytes), "system")
	ch <- prometheus.MustNewConstMetric(descs["disk_free_bytes"], prometheus.GaugeValue, float64(systemStats.DiskFreeBytes), "system")

	return nil
}
