package result

type Option[T any] struct {
	value *T
}

func (opt Option[T]) IsNone() bool {
	return opt.value == nil
}
func (opt Option[T]) IsSome() bool {
	return opt.value != nil
}
func (opt Option[T]) IsSomeAnd(f func(value T) bool) bool {
	if opt.IsNone() {
		return false
	}
	return f(*opt.value)
}
func (opt Option[T]) OkOr(err error) Result[T] {
	if opt.IsSome() {
		return Ok(*opt.value)
	}
	return Err[T](err)
}
func (opt Option[T]) OkOrElse(err func() error) Result[T] {
	if opt.IsSome() {
		return Ok(*opt.value)
	}
	return Err[T](err())
}

func (opt Option[T]) Expect(msg string) T {
	if opt.IsSome() {
		return *opt.value
	}
	panic(msg)
}
func (opt Option[T]) Filter(predicate func(value T) bool) Option[T] {
	if opt.IsSome() {
		if predicate(*opt.value) {
			return Some(*opt.value)
		}
	}
	return None[T]()
}

// func (opt *Option[T]) Flatten() {
// }

func (opt Option[T]) Unwrap() T {
	if opt.IsSome() {
		return *opt.value
	}
	panic("called `Option.unwrap()` on a `None` value")
}

func (opt Option[T]) UnwrapOr(value T) T {
	if opt.IsSome() {
		return *opt.value
	}
	return value
}

func (opt Option[T]) UnwrapOrDefault() T {
	var value T
	return value
}

func (opt Option[T]) UnwrapOrElse(f func() T) T {
	if opt.IsSome() {
		return *opt.value
	}
	return f()
}

func (opt Option[T]) Or(value Option[T]) Option[T] {
	if opt.IsSome() {
		return Some(*opt.value)
	}
	return value
}

func (opt Option[T]) OrElse(f func() Option[T]) Option[T] {
	if opt.IsSome() {
		return Some(*opt.value)
	}
	return f()
}

func (opt Option[T]) Cloned() Option[T] {
	return Option[T]{value: opt.value}
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value: &value,
	}
}

func None[T any]() Option[T] {
	return Option[T]{}
}
