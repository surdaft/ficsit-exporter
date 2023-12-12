package exporter

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/benbjohnson/clock"
)

var timeRegex = regexp.MustCompile(`\d\d:\d\d:\d\d`)

var Clock = clock.New()

func parseTimeSeconds(timeStr string) *float64 {
	match := timeRegex.FindStringSubmatch(timeStr)
	if len(match) < 1 {
		return nil
	}
	parts := strings.Split(match[0], ":")
	duration := parts[0] + "h" + parts[1] + "m" + parts[2] + "s"
	t, _ := time.ParseDuration(duration)
	seconds := t.Seconds()
	return &seconds
}

func parseBool(b bool) float64 {
	if b {
		return 1
	} else {
		return 0
	}
}

func retrieveData(ctx context.Context, frmAddress string, details any) error {
	http.DefaultClient.Timeout = time.Second * 5
	req, _ := http.NewRequestWithContext(ctx, "GET", frmAddress, nil)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("error fetching statistics from FRM: %s\n", err)
		return err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&details)
	return err
}
