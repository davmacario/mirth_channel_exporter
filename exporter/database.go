package exporter

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	databaseTasksAPI = "/databaseTasks"
)

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

func (e *Exporter) gatherDatabaseTasks(ch chan<- prometheus.Metric) error {
	resp, err := e.makeMirthAPIRequest(context.Background(), databaseTasksAPI)
	if err != nil {
		return fmt.Errorf("failed to perform HTTP request for channel statistics: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK status code %d from database tasks API: %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body for database tasks: %w", err)
	}

	var databaseTasks DatabaseTasks
	if err := xml.Unmarshal(body, &databaseTasks); err != nil {
		return fmt.Errorf("failed to unmarshal XML for database tasks: %w", err)
	}

	// TODO: provider better metrics about DB tasks
	numDatabaseTasks := len(databaseTasks.Entries)
	ch <- prometheus.MustNewConstMetric(descs["num_db_tasks"], prometheus.GaugeValue, float64(numDatabaseTasks))
	return nil
}
