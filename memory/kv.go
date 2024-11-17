package memory

type KV interface {
	Set(key string, value interface{}) error
	Get(key string) (string, error)
	Delete(key string) error
}

var inst KV

func Init(store KV) {
	inst = store
}

func Set(key string, value interface{}) error {
	return inst.Set(key, value)
}

func Get(key string) (string, error) {
	return inst.Get(key)
}

func Delete(key string) error {
	return inst.Delete(key)
}
