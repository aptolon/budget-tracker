package crypto_hasher

type Hasher interface {
	Hash(string) (string, error)
	Compare(string, string) error
}
