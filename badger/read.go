package badger

// Read read
type Read struct{}

// NewRead new read
func NewRead() *Read {
	return &Read{}
}

// Get Get
func (r *Read) Get(k string) ([]byte, error) {
	<-Conn
	defer func() {
		Conn <- 1
	}()
	txn := Pool.DB.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get([]byte(k))
	if err != nil {
		return nil, err
	}

	var _ret []byte
	err = item.Value(func(v []byte) error {
		_ret = v
		return nil
	})
	if err != nil {
		return _ret, err
	}
	return _ret, nil
}
