package shortener

import (
	"errors"
	"log"
	"os"
	"strconv"
	"url-shortener/pkg/idgen"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	repo      Repository
	generator idgen.Generator
}

func NewService(repo Repository, generator idgen.Generator) *Service {
	return &Service{repo: repo, generator: generator}
}

func (s *Service) Shorten(longUrl string, customShortCode string) (string, error) {
	if longUrl == "" {
		return "", errors.New("longUrl is empty")
	}

	var shortCode string
	var err error

	if len(customShortCode) > 0 {
		ok, err := s.validShortCode(customShortCode)
		if !ok {
			return "", err
		}
		shortCode = customShortCode
	} else {
		shortCode, err = s.generator.GenerateShortCode(longUrl)
		if err != nil {
			return "", err
		}
	}

	err = s.repo.Save(shortCode, longUrl)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func (s *Service) Resolve(shortUrl string) (string, error) {
	if shortUrl == "" {
		return "", errors.New("shortUrl is empty")
	}

	longUrl, err := s.repo.Get(shortUrl)
	if err != nil {
		return "", err
	}

	return longUrl, nil
}

func (s *Service) validShortCode(customShortCode string) (bool, error) {
	allowed_short_code_length, _ := strconv.Atoi(os.Getenv("SHORT_CODE_LENGTH"))
	if len(customShortCode) > allowed_short_code_length {
		return false, errors.New("given short code exceeds allowed length")
	}
	existingUrl, err := s.repo.Get(customShortCode)
	if err != nil && err != redis.Nil {
		return false, err
	}
	if existingUrl != "" {
		return false, errors.New("given short code is already taken")
	}
	return true, nil
}
