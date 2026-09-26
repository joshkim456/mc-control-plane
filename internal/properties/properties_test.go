package properties

import (
	"string"
	"testing"
)

func TestParseNormal(t *testing.T) {
	props, err := Prase(strings.NewReader("max_players=10"))
	if err != nil {
		t.Fataf("Parse error: %v", err)
	}

	got, ok := props.Get("max-players")
	if got != "10" || !ok {
		t.Errorf("Get = %q, %v; want %q, true", got, ok, "10")
	}
}