package maybe

type Maybe[T any] struct {
	value      T
	is_present bool
}

func (m Maybe[T]) Bind(bindf func(value T) (T, bool)) Maybe[T] {

	if m.is_present {
		return TupleToMaybe(bindf(m.value))
	}

	return None[T]()
}

func Some[T any](value T) Maybe[T] {
	return Maybe[T] {
		is_present: true,
		value: value,
	}
}

func None[T any]() Maybe[T] {
	return Maybe[T] {
		is_present: false,
	}
}

func TupleToMaybe[T any](value T, ok bool) Maybe[T] {
	if ok {
		return Some(value)
	}
	return None[T]()
}
