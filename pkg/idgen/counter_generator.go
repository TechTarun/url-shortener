package idgen

import (
	"sync/atomic"
)

type CounterGenerator struct {
	counter uint64
}

func NewBase62Generator() *CounterGenerator {
	return &CounterGenerator{}
}

// This method will generate short codes without conflict since every API call is sequential.
// Pros: Unique short code everytime, thread safe
// Cons: Short code length will vary starting from single digit to an endless number of digits
func (g *CounterGenerator) GenerateShortCode(longUrl string) (string, error) {
	id := atomic.AddUint64(&g.counter, 1)
	short_code := EncodeToBase62(id)
	return short_code, nil
}
