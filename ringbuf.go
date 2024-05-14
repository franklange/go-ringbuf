package ringbuf

type Ringbuf[T any] struct {
	buf []T
	cap int
	idx int
}

type Iterator[T any] struct {
	buf *Ringbuf[T]
	idx int
	lap bool
}

func NewRingBuf[T any](n int) *Ringbuf[T] {
	return &Ringbuf[T]{
		buf: make([]T, 0, n),
		cap: n,
		idx: 0,
	}
}

func (r *Ringbuf[T]) full() bool {
	return r.Len() == r.Cap()
}

func (r *Ringbuf[T]) Len() int {
	return len(r.buf)
}

func (r *Ringbuf[T]) Cap() int {
	return cap(r.buf)
}

func (r *Ringbuf[T]) Raw() []T {
	return r.buf
}

func (r *Ringbuf[T]) Put(t ...T) {
	for _, v := range t {
		if !r.full() {
			r.buf = append(r.buf, v)
		} else {
			r.buf[r.idx] = v
		}
		r.idx = (r.idx + 1) % r.cap
	}
}

func (r *Ringbuf[T]) Iter() *Iterator[T] {
	res := &Iterator[T]{
		buf: r,
		lap: false,
	}
	if !r.full() {
		res.idx = 0
	} else {
		res.idx = r.idx
	}
	return res
}

func (iter *Iterator[T]) Next() bool {
	if iter.buf.Len() == 0 {
		return false
	}
	if !iter.buf.full() {
		return (iter.idx < iter.buf.idx)
	}
	return !iter.lap
}

func (iter *Iterator[T]) Get() *T {
	i := iter.idx
	iter.idx = (iter.idx + 1) % iter.buf.cap
	if iter.idx == iter.buf.idx {
		iter.lap = true
	}
	return &iter.buf.buf[i]
}
