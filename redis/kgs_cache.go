package keystore

type KGScache interface {
	Set([]string)
	Get() []string
}
