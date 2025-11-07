package cursors

import (
	"context"
	"iter"
)

type Nextable interface {
	Next() *string
}

type Collectionable[C any] interface {
	Collection() []C
}

type Iterable[C any] interface {
	Nextable
	Collectionable[C]
}

type IsQuery interface {
	ToQueryString() (string, error)
	SetCursor(cursor *string)
}

func Iterator[C any, T Iterable[C], Q IsQuery](
	ctx context.Context,
	request func(ctx context.Context, opts Q) (T, error),
	initialOpts Q,
) iter.Seq2[C, error] {
	return func(yield func(item C, err error) bool) {
		opts := initialOpts

		for {
			res, apiErr := request(ctx, opts)
			if apiErr != nil {
				yield(*new(C), apiErr)
				return
			}

			for _, item := range res.Collection() {
				if !yield(item, nil) {
					return
				}
			}

			nextCursor := res.Next()
			if nextCursor == nil {
				return
			}

			opts.SetCursor(nextCursor)
		}
	}
}
