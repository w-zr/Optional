package optional

import (
	"errors"
)

// Optional is a wrapper for representing 'optional'(or 'nullable')
// objects who may not contain a valid value.
type Optional[T any] struct {
	storage *T
}

// Convert converts a value of type T into Optional[T].
func Convert[T any](val T) Optional[T] {
	opt := Optional[T]{}
	opt.Assign(val)
	return opt
}

// Assign assigns from a T.
func (o *Optional[T]) Assign(v T) {
	if !o.Initialized() {
		o.storage = new(T)
	}
	*o.storage = v
}

// Get returns the value if it is valid. Otherwise,
// returns the default value of T and an error.
func (o Optional[T]) Get() (T, error) {
	if o.Initialized() {
		return *o.storage, nil
	}
	return *new(T), errors.New("value is invalid")
}

// MustGet returns the value if it is valid. Otherwise,
// it panics.
func (o Optional[T]) MustGet() T {
	if o.Initialized() {
		return *o.storage
	}
	panic("value is invalid")
}

// GetOr returns the value if it is valid. Otherwise,
// it returns val.
func (o Optional[T]) GetOr(val T) T {
	if o.Initialized() {
		return *o.storage
	}
	return val
}

// Reset cleans the Optional.
func (o *Optional[T]) Reset() {
	o.storage = nil
}

// Initialized returns true if the value is valid.
func (o Optional[T]) Initialized() bool {
	return o.storage != nil
}

// Map applies function typed T->T to the value.
// TODO: support function typed T->U after parameterized methods is enabled.
// see: https://github.com/golang/proposal/blob/master/design/43651-type-parameters.md#no-parameterized-methods
func (o Optional[T]) Map(f func(T) T) Optional[T] {
	if o.Initialized() {
		return Convert(f(o.MustGet()))
	}
	return Optional[T]{}
}

// FlatMap applies function typed T->Optional[T] to the value.
// TODO: support function typed T->Optional[U] after parameterized methods is enabled.
// see: https://github.com/golang/proposal/blob/master/design/43651-type-parameters.md#no-parameterized-methods
func (o Optional[T]) FlatMap(f func(T) Optional[T]) Optional[T] {
	if o.Initialized() {
		return f(o.MustGet())
	}
	return Optional[T]{}
}
