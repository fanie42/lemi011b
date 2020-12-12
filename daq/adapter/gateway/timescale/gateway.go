package timescale

import (
    "context"

    "github.com/jackc/pgx/v4/pgxpool"
)

// Gateway TODO
type Gateway struct {
    db  *pgxpool.Pool
    log iLogger
}

// New TODO
func New(database *pgxpool.Pool, logger iLogger) *Gateway {
    return &Gateway{
        db:  database,
        log: logger,
    }
}

// Create TODO
func (gw *Gateway) Create(
    ctx context.Context,
    datum *lemi011b.Datum,
) error {
    q := `INSERT INTO data (time, sensor_id, x, y, z, temperature)
        VALUES($1, $2, $3, $4, $5, $6)
    `

    tag, err := gw.db.Exec(
        ctx,
        q,
        datum.Timestamp,
        datum.SensorID,
        datum.X,
        datum.Y,
        datum.Z,
        datum.Temperature,
    )
    if err != nil {
        gw.log.Error("failed to create datum", err)

        return err
    }
    gw.log.Info("successfully created datum")

    return nil
}
