package idgen

type Generator interface {
	GenerateShortCode(longUrl string) (string, error)
}