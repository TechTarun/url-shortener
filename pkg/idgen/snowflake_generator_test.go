package idgen

import (
	"sync"
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

func TestSnowflake_ConcurrentUniqueness(t *testing.T) {
	const (
		goroutines    = 50
		idsPerRoutine = 20000
		expectedTotal = goroutines * idsPerRoutine
	)

	gen, err := NewSnowflakeGenerator(8, 1)
	if err != nil {
		t.Fatalf("failed to create snowflake: %v", err)
	}

	var (
		wg      sync.WaitGroup
		seen    sync.Map
		longUrl string
	)

	longUrl = "https://www.google.com"

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < idsPerRoutine; j++ {
				id, _ := gen.GenerateShortCode(longUrl)

				// Check for duplicates
				if _, exists := seen.LoadOrStore(id, true); exists {
					t.Fatalf("duplicate ID detected: %s", id)
				}
			}
		}()
	}

	wg.Wait()

	// Count results
	count := 0
	seen.Range(func(_, _ any) bool {
		count++
		return true
	})

	if count != expectedTotal {
		t.Fatalf("expected %d ids, got %d", expectedTotal, count)
	}
}
