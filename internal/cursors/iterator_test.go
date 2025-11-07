package cursors_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/flussonic/go-flussonic/internal/cursors"
)

type MethodQuery struct {
	Cursor *string
}

func (q *MethodQuery) ToQueryString() (string, error) {
	values := url.Values{}
	if q.Cursor != nil {
		values.Set("cursor", *q.Cursor)
	}
	return values.Encode(), nil
}

func (q *MethodQuery) SetCursor(cursor *string) {
	q.Cursor = cursor
}

type MethodResponse struct {
	NextField *string
	Items     []Item
}

type Item struct {
	ID string
}

func (r MethodResponse) Collection() []Item {
	return r.Items
}

func (r MethodResponse) Next() *string {
	return r.NextField
}

func Method(ctx context.Context, opts *MethodQuery) (MethodResponse, error) {
	return MethodResponse{
		Items: []Item{
			{ID: "1"},
			{ID: "2"},
			{ID: "3"},
		},
		NextField: nil,
	}, nil
}

func TestIterator(t *testing.T) {
	ctx := context.Background()
	iter := cursors.Iterator(ctx, Method, &MethodQuery{})
	for item := range iter {
		fmt.Println(item)
	}
}
