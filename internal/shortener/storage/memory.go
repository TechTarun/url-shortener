package storage

type MemoryStore struct {
	mapping map[string]string
}

func NewInMemoryStore() *MemoryStore {
	return &MemoryStore{make(map[string]string)}
}

func (m *MemoryStore) Save(shortCode string, longUrl string) error {
	m.mapping[shortCode] = longUrl
	return nil
}

func (m *MemoryStore) Get(shortCode string) (string, error) {
	longUrl := m.mapping[shortCode]
	return longUrl, nil
}
