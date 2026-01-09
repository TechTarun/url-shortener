package idgen

import "testing"

// Steps:
// 1. Initiate generator
// 2. Take a hashmap seen to record all the short codes already seen
// 3. Run an infinite loop, start counter with 1 and keep generating short codes until duplicate is encountered
// 4. Return counter at which duplicate happens

func TestCounterGenerator_IsUnique(t *testing.T) {
	generator := NewBase62Generator()
	seen := make(map[string]bool)
	longUrl := "https://www.google.com"

	for i := 0; i < 50000000; i++ {
		short_code, _ := generator.GenerateShortCode(longUrl)
		if seen[short_code] {
			t.Fatalf("Duplicate found at i = %d", i)
		}
		seen[short_code] = true
	}
}
