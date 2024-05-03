// O Time é um tipo que encapsula o time.Time e implementa a interface Marshaler e Unmarshaler do pacote encoding/json.
//
// A diferença do Time pro time.Time é que o Time aceita valores nulos, transformando-os em null ao serem serializados em JSON e deserializando-os para nil ao serem desserializados.
package time

import "time"

type Time struct {
	time.Time
}

// NewTimeNow cria um novo Time com o tempo atual.
func NewTimeNow() *Time {
	return &Time{Time: time.Now()}
}

// UnmarshalJSON implementa a interface Unmarshaler do pacote encoding/json.
func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t.Time = time.Time{}
		return nil
	}

	return t.Time.UnmarshalJSON(data)
}

// MarshalJSON implementa a interface Marshaler do pacote encoding/json.
func (t *Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	} else {
		return t.Time.MarshalJSON()
	}
}
