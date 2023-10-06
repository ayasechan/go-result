package result

import (
	"github.com/morikuni/failure"
)

type Result[T any] struct {
	value T
	err   error
}

func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic(r.err)
	}
	return r.value
}

func (r Result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic("called `called `Result.UnwrapErr()` on an `Ok` value")
	}
	return r.err
}

func (r Result[T]) UnwrapErrUnchecked() error {
	return r.err
}

func (r Result[T]) UnwrapOr(value T) T {
	if r.IsOk() {
		return r.value
	}
	return value
}

func (r Result[T]) UnwrapOrDefault() T {
	// not check ok
	return r.value
}

func (r Result[T]) UnwrapOrElse(op func(err error) T) T {
	if r.IsOk() {
		return r.value
	}
	return op(r.err)
}

func (r Result[T]) UnwrapUnchecked() T {
	return r.value
}

func (r Result[T]) Expect(msg string) T {
	if r.IsOk() {
		return r.value
	}
	panic(msg)
}

func (r Result[T]) ExpectErr(msg string) error {
	if r.IsErr() {
		return r.err
	}
	panic(msg)
}

func (r Result[T]) Inspect(f func(value T)) Result[T] {
	if r.IsOk() {
		f(r.value)
	}
	return r
}

func (r Result[T]) InspectErr(f func(err error)) Result[T] {
	if r.IsErr() {
		f(r.err)
	}
	return r
}

func (r Result[T]) Err() Option[error] {
	if r.IsOk() {
		return None[error]()
	}
	return Some[error](r.err)
}
func (r Result[T]) Ok() Option[T] {
	if r.IsOk() {
		return Some[T](r.value)
	}
	return None[T]()
}

func (r Result[T]) Or(value Result[T]) Result[T] {
	if r.IsOk() {
		return Ok(r.value)
	}
	return value
}
func (r Result[T]) OrElse(op func(err error) Result[T]) Result[T] {
	if r.IsOk() {
		return Ok(r.value)
	}
	return op(r.err)
}
func (r Result[T]) Cloned() Result[T] {
	return Result[T]{value: r.value, err: r.err}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}
func (r Result[T]) IsOkAnd(f func(value T) bool) bool {
	if r.IsErr() {
		return false
	}
	return f(r.value)
}
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) IsErrAnd(f func(err error) bool) bool {
	if r.IsOk() {
		return false
	}
	return f(r.err)
}
func (r Result[T]) MapErr(op func(err error) error) Result[T] {
	if r.IsOk() {
		return Ok(r.value)
	}
	return Err[T](op(r.err))
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: value,
		err:   nil,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		err: failure.Wrap(err),
	}
}

func AsResult[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok[T](value)
}
