package graphql

// Schema TODO
var Schema string = `
    scalar Time

    schema {
        query: Query
        mutation: Mutation
    }

    type Query {
        data(from: Time, to: Time): [Datum]
        datum(id: ID!): Datum
    }

    type Mutation {
        createDatum(input: CreateDatumInput!): Datum
        updateDatum(input: UpdateDatumInput!): Datum
        deleteDatum(id: ID!): Boolean!
    }

    input CreateDatumInput {
        timestamp: Time!
        x: Int
        y: Int
        z: Int
        temperature: Int
    }

    input UpdateDatumInput {
        id: ID!
        x: Int
        y: Int
        z: Int
        temperature: Int
    }

    type Datum {
        id: ID!
        timestamp: Time!
        x: Int
        y: Int
        z: Int
        temperature: Int
    }
`
