package client

import (
    "context"

    "github.com/fanie42/lemi011b"
)

// ALL THE BUSINESS LOGIC SHOULD GO HERE! THIS INCLUDES VALIDATION!!!

// DAQClientService TODO
type DAQClientService struct {
    log iLogger
    pub iPublisher
}

// NewDAQClientService TODO
func NewDAQClientService(
    logger iLogger,
    publisher iPublisher,
) *DAQClientService {
    return &DAQClientService{
        log:  logger,
        repo: repository,
        pub:  publisher,
    }
}

// Create TODO
func (svc *DAQClientService) Create(
    ctx context.Context,
    datum *lemi011b.Datum,
) error {
    err := svc.pub.Publish(ctx, datum)
    if err != nil {
        svc.log.Error("failed to publish datum", err)

        return err
    }
    svc.log.Info("successfully published datum")

    return nil
}
