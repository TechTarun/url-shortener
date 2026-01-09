package shortener

import (
	"errors"
	"url-shortener/pkg/idgen"
)

type Service struct {
	repo      Repository
	generator idgen.Generator
}

func NewService(repo Repository, generator idgen.Generator) *Service {
	return &Service{repo: repo, generator: generator}
}

func (s *Service) Shorten(longUrl string) (string, error) {
	if longUrl == "" {
		return "", errors.New("longUrl is empty")
	}

	shortCode, error := s.generator.GenerateShortCode(longUrl)

	if error != nil {
		return "", error
	}

	s.repo.Save(shortCode, longUrl)
	return shortCode, nil
}

func (s *Service) Resolve(shortUrl string) (string, error) {
	if shortUrl == "" {
		return "", errors.New("shortUrl is empty")
	}

	longUrl, ok := s.repo.Get(shortUrl)
	if ok != nil {
		return "", errors.New(ok.Error())
	}

	return longUrl, nil
}
