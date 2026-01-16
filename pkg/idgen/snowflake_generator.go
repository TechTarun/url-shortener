package idgen

import (
	"errors"
	"sync"
	"time"
)

const (
	nodeIdBits     = 2
	sequenceIdBits = 8
	maxNodeId      = -1 ^ (-1 << nodeIdBits)
	maxSequenceId  = -1 ^ (-1 << sequenceIdBits)

	nodeIdShift    = sequenceIdBits
	timestampShift = sequenceIdBits + nodeIdBits

	// Custom epoch: January 1, 2026 00:00:00 UTC
	// This reduces the timestamp value, allowing shorter Base62 codes
	customEpoch = 1767225600000
)

type Snowflake struct {
	length   int
	lastTS   int64
	nodeId   int64
	sequence int64
	mu       sync.Mutex
}

func NewSnowflakeGenerator(length int, nodeId int64) (*Snowflake, error) {
	if nodeId > maxNodeId {
		return nil, errors.New("Node Id is invalid")
	}

	if length <= 0 {
		return nil, errors.New("Short code length should be > 0")
	}

	return &Snowflake{length: length, nodeId: nodeId}, nil
}

// This method will generate short codes without conflict since every API call is sequential.
// Pros: Unique short code everytime, thread safe
// Cons: Short code length will vary starting from single digit to an endless number of digits
func (s *Snowflake) GenerateShortCode(longUrl string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UnixMilli() - customEpoch

	if now == s.lastTS {
		s.sequence = (s.sequence + 1) & maxSequenceId

		if s.sequence == 0 {
			for now <= s.lastTS {
				now = time.Now().UnixMilli() - customEpoch
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTS = now

	uniqueId := (s.lastTS << timestampShift) | (s.nodeId << nodeIdShift) | s.sequence
	base62_code := EncodeToBase62(uint64(uniqueId))
	short_code, error := PadShortCode(base62_code, s.length)
	if error != nil {
		return "", error
	}
	return short_code, nil
}
