package result

import "github.com/morikuni/failure"

type Result[T any] struct {
	value T
	err   error
}

func (r *Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}

func (r *Result[T]) UnwrapOr(value T) T {
	if r.IsOk() {
		return r.value
	}
	return value
}
func (r *Result[T]) Expect(msg string) T {
	if r.IsOk() {
		return r.value
	}
	panic(msg)
}
func (r *Result[T]) ExpectErr(msg string) error {
	if r.IsErr() {
		return r.err
	}
	panic(msg)
}
func (r *Result[T]) Error() error {
	return r.err
}
func (r *Result[T]) IsOk() bool {
	return r.err == nil
}
func (r *Result[T]) IsErr() bool {
	return r.err != nil
}
func Ok[T any](value T) *Result[T] {
	return &Result[T]{
		value: value,
		err:   nil,
	}
}

func Err[T any](err error) *Result[T] {
	return &Result[T]{
		err: failure.Wrap(err),
	}
}
