package referees

import (
	"net/http"
	"testing"
)

func TestConsume(t *testing.T) {
	refs := NewRefereesReader(&http.Client{})
	refs.ReadReferees()
}
