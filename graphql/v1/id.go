package graphql

import (
    "encoding/json"
    "fmt"
)

const api uint8 = 1

// Time is a custom GraphQL type to represent an instant in time. It has to be added to a schema
// via "scalar Time" since it is not a predeclared GraphQL type like "ID".
// type Time struct {
//     time.Time
// }

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (sansa.UID) ImplementsGraphQLType(name string) bool {
    return name == "ID"
}

// UnmarshalGraphQL is a custom unmarshaler for ID
//
// This function will be called whenever you use the
// time scalar as an input
func (id *sansa.UID) UnmarshalGraphQL(input interface{}) error {
    switch input := input.(type) {
    case string:
        return id.Decode(input)
    default:
        return fmt.Errorf("wrong type for ID: %T", input)
    }
}

// MarshalJSON is a custom marshaler for ID
//
// This function will be called whenever you
// query for fields that use the ID type
func (id sansa.UID) MarshalJSON() ([]byte, error) {
    s, err := id.Encode()
    if err != nil {
        return nil, err
    }

    return json.Marshal(map[string]string{"id": s})
}
