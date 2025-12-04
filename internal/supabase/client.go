package supabase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Automation struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func ListAutomations() ([]Automation, error) {
	url := os.Getenv("SUPABASE_AUTOMATIONS_URL")
	if url == "" {
		return nil, fmt.Errorf("missing SUPABASE_AUTOMATIONS_URL")
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", os.Getenv("SUPABASE_ANON_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out []Automation
	err = json.NewDecoder(resp.Body).Decode(&out)
	return out, err
}
