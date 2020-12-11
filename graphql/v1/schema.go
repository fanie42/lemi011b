package graphql

// Schema TODO
var Schema string = `
    scalar Time

    extend enum InstrumentType {
        LEMI011B
    }

    type Datum {
        timestamp: Time!
        x: Int
        y: Int
        z: Int
        temperature: Int
    }

    extend type Query {
        data(from: Time, to: Time): [Datum]
        datum(timestamp: Time): Datum
    }

    extend type Mutation {
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

    
`
