package shortener

type Repository interface {
	Save(shortCode string, longUrl string) error
	Get(shortCode string) (string, error)
}