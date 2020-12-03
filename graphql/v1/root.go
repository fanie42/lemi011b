package graphql

import (
    "context"
    "net/http"

    "github.com/fanie42/simple/domain"
    "github.com/gorilla/mux"
    "github.com/nats-io/stan.go"

    "github.com/graph-gophers/graphql-go"
    gql "github.com/graph-gophers/graphql-go"
)

// GraphQL TODO
type GraphQL struct {
    dataRepo iDataRepository
}

// New TODO
func New(dataRepository iDataRepository) *GraphQL {
    return &GraphQL{
        dataRepo: dataRepository,
    }
}

type dataInput struct {
    From *graphql.Time
    To   *graphql.Time
}

// Data TODO
func (gql *GraphQL) Data(ctx context.Context, args dataInput) (*[]*Datum, error) {
    from := args.From.Time
    to := args.To.Time

    data, err := gql.dataRepo.GetByTime(ctx, from, to)

    resolvers := make(*[]*Datum, len(data))
    for i, datum := range data {
        resolvers[i] = &Datum{}
    }

    return data, err
}

type createDatumInput struct {
    Timestamp   gql.Time
    X           *int32
    Y           *int32
    Z           *int32
    Temperature *int32
}

// CreateDatum TODO
func (r *RootResolver) CreateDatum(ctx context.Context, args struct{ Input createDatumInput }) (*datumResolver, error) {
    datum := args.Input.adapt()

    err := r.data.Create(ctx, datum)
    if err != nil {
        return nil, err
    }

    resolver := datumResolver{
        datum: datum,
    }

    return &resolver, err
}

type updateDatumInput struct {
    ID          gql.ID `json:"id"`
    X           *int32 `json:"x"`
    Y           *int32 `json:"y"`
    Z           *int32 `json:"z"`
    Temperature *int32 `json:"temperature"`
}

func (input updateDatumInput) adapt() (*domain.Datum, error) {
    datum, err := domain.NewDatum(string(input.ID))
    if err != nil {
        return nil, err
    }

    datum.X = input.X
    datum.Y = input.Y
    datum.Z = input.Z
    datum.Temperature = input.Temperature

    return datum, nil
}

// UpdateDatum TODO
func (r *RootResolver) UpdateDatum(ctx context.Context, args struct{ Input updateDatumInput }) (*datumResolver, error) {
    datum, err := args.Input.adapt()
    if err != nil {
        return nil, err
    }

    err = r.data.Update(ctx, datum)
    if err != nil {
        return nil, err
    }

    resolver := datumResolver{
        datum: datum,
    }

    return &resolver, nil
}

// DeleteDatum TODO
func (r *RootResolver) DeleteDatum(ctx context.Context, args struct{ ID gql.ID }) (bool, error) {
    datum, err := domain.NewDatum(string(args.ID))
    if err != nil {
        return false, err
    }

    ok, err := r.data.Delete(ctx, datum.Timestamp)

    return ok, err
}

// HTTP TODO
func (r *GraphQL) HTTP() http.Handler {
    router := mux.NewRouter()

    router.HandleFunc("/graphql", r)

    return router
}

// ServeNATS TODO
func (r *GraphQL) NATS(subject string, conn *stan.Conn) {

}
