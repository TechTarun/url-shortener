package idgen

import (
	"testing"
)

func TestSnowflakeGenerator_IsUnique(t *testing.T) {
	generator, error := NewSnowflakeGenerator(8, 1)
	if error != nil {
		t.Fatal(error)
	}
	seen := make(map[string]bool)
	longUrl := "https://www.google.com"

	for i := 0; i < 10000000; i++ {
		short_code, error := generator.GenerateShortCode(longUrl)
		if error != nil {
			t.Fatal(error)
		}
		if seen[short_code] {
			t.Fatalf("Duplicate found at i = %d", i)
		}
		seen[short_code] = true
	}
}
