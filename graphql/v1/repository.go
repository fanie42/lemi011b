package graphql

import (
    "context"
    "time"
)

type iDataRepository interface {
    GetByID(ctx context.Context, id lemi011b.DatumID) (*lemi011b.Datum, error)
    GetByTime(ctx context.Context, from time.Time, to time.Time) (*[]*lemi011b.Datum, err)
    // Fetch(ctx context.Context, page *domain.Page) (*domain.DataConnection, error)
    // Fetch(page *domain.Page, filter *domain.Filter, order *domain.Order) (*domain.DataConnection, error)
    Create(ctx context.Context, datum *lemi011b.Datum) error
    Update(ctx context.Context, datum *lemi011b.Datum) error
    Delete(ctx context.Context, timestamp time.Time) (bool, error)
}
