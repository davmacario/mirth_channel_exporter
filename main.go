package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davmacario/mirth_channel_exporter/exporter"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// DefaultHTTPClientTimeout is the default timeout for HTTP client requests.
	DefaultHTTPClientTimeout = 10 * time.Second
)

var (
	// Command-line flags
	listenAddress = flag.String("web.listen-address", ":9141", "Address to listen on for telemetry")
	metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics")
	configPath    = flag.String("config.file-path", "", "Path to environment file")
)

func main() {
	flag.Parse()

	// Load environment variables
	if *configPath != "" {
		log.Printf("Loading environment file from: %s", *configPath)
		if err := godotenv.Load(*configPath); err != nil {
			log.Fatalf("FATAL: Error loading environment file %s: %v", *configPath, err)
		}
	} else {
		if err := godotenv.Load(); err != nil {
			log.Println("WARNING: Error loading .env file, assuming environment variables are set or not needed.")
		}
	}

	mirthEndpoint := os.Getenv("MIRTH_BASE_API_URL")
	mirthUsername := os.Getenv("MIRTH_USERNAME")
	mirthPassword := os.Getenv("MIRTH_PASSWORD")

	if mirthEndpoint == "" {
		log.Fatal("FATAL: MIRTH_BASE_API_URL environment variable is not set.")
	}
	if mirthUsername == "" || mirthPassword == "" {
		log.Fatal("FATAL: MIRTH_USERNAME or MIRTH_PASSWORD environment variables are not set. Basic authentication is required.")
	}

	exporter := exporter.NewExporter(mirthEndpoint, mirthUsername, mirthPassword)
	prometheus.MustRegister(exporter)
	log.Printf("Mirth Exporter started. Using Mirth endpoint: %s", mirthEndpoint)
	log.Printf("Metrics exposed on %s%s", *listenAddress, *metricsPath)

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<html>
             <head><title>Mirth Channel Exporter</title></head>
             <body>
             <h1>Mirth Channel Exporter</h1>
             <p><a href='%s'>Metrics</a></p>
             </body>
             </html>`, *metricsPath)
	})

	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
