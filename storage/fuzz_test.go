package storage

import (
    "testing"
)

func FuzzParseEntity(f *testing.F) {
    f.Add("user-123")
    f.Fuzz(func(t *testing.T, data string) {
        _, _ = parseEntity(data) 
    })
}
