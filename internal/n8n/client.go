package n8n

import (
	"fmt"
	"net/http"
	"os"
)

func TriggerDemo(id string) error {
	url := os.Getenv("N8N_DEMO_URL")
	if url == "" {
		return fmt.Errorf("missing N8N_DEMO_URL")
	}

	resp, err := http.Post(fmt.Sprintf("%s/%s", url, id), "application/json", nil)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("n8n error: %d", resp.StatusCode)
	}

	return nil
}
