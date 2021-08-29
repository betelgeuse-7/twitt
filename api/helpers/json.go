package helpers

import (
	"encoding/json"
	"io"
)

// wrapper for json.NewEncoder(...).Encode(...)
func JSON(w io.Writer, value interface{}) {
	json.NewEncoder(w).Encode(value)
}
