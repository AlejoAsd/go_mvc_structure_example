package mvc

// RI is the base Repository interface
type RI interface {
	Create(value interface{}) error
	Get(out interface{}, where ...interface{}) error
	Update(value interface{}) error
	Delete(value interface{}) error
}

// R is the base Repository implementation
type R struct{
	tx *DB
}

func NewR(tx *DB) RI {
	return &R{
		tx: tx,
	}
}

// Create creates an entry in the storage
func (r *R) Create(value interface{}) error {
	var err error
	r.tx.Create(value).Error(err)

	return err
}

// Get gets an entry from the storage
func (r *R) Get(out interface{}, where ...interface{}) error {
	var err error
	r.tx.First(out, where...).Error(err)

	return err
}

// Update updates an entry in the storage
func (r *R) Update(value interface{}) error {
	var err error
	r.tx.Save(value).Error(err)

	return err
}

// Delete deletes an entry in the storage
func (r *R) Delete(value interface{}) error {
	var err error
	r.tx.Delete(value).Error(err)

	return err
}
