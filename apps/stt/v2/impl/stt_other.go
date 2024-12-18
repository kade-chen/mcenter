package impl

import (
	"fmt"

	"github.com/kade-chen/mcenter/apps/stt"
)

// Receives an project ID and endpoint and returns the parent string
// EXAMPLE projects/%s/locations/%s
func (s *speechToTextV2) parentString(projectID string, endpoint stt.ENDPOINT) string {
	return fmt.Sprintf("projects/%s/locations/%s", projectID, stt.ENDPOINTS_PARENT[endpoint])
}
