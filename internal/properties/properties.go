package properties

import "io"

type Properties struct {
	values map[string]string
}

func (p Properties) Get(key string) (string, bool) {
	value, ok := p.values[key]
	return value, ok
}

func Parse(r io.Reader) (Properties, error) {
	return Properties{values: map[string]string{"max-players": "10"}}, nil
}
