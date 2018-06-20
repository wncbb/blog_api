package marshal_json

import "encoding/json"

type InnerJSON struct {
}

func (p *InnerJSON) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (p *InnerJSON) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func NewInnerJSON() *InnerJSON {
	return &InnerJSON{}
}
