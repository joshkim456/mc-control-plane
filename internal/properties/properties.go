package properties

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Properties struct {
	values map[string]string
}

func (p Properties) Get(key string) (string, bool) {
	value, ok := p.values[key]
	return value, ok
}

func Parse(r io.Reader) (Properties, error) {
	scanner := bufio.NewScanner(r)
	lineNum := 1
	props := map[string]string{}

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 || line[0] == '#' {
			lineNum++
			continue
		}

		key, value, found := bytes.Cut(line, []byte("="))
		if !found {
			return Properties{}, fmt.Errorf("line %d: missing \"=\": %q", lineNum, line)
		}
		props[string(key)] = string(value)
		lineNum++
	}

	parseErr := scanner.Err()
	if parseErr != nil {
		return Properties{}, parseErr
	}

	return Properties{values: props}, nil
}
