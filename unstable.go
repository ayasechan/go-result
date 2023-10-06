package result

func WithRecover[T any](fn func() Result[T]) (ret Result[T]) {
	defer func() {
		switch v := recover().(type) {
		case nil:
		case error:
			ret = Err[T](v)
		default:
			panic(v)
		}
	}()
	ret = fn()
	return
}
