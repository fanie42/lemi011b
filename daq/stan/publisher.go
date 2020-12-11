package stan

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/fanie42/lemi011b"
    "github.com/nats-io/nats.go"
    "github.com/nats-io/stan.go"
    s "github.com/nats-io/stan.go"
)

// STAN TODO
type STAN struct {
    subject string
    nc      *nats.Conn
    sc      s.Conn
    log     iLogger
}

// New TODO
func New(config *Config, logger iLogger, site string, sensor string, version uint8) *STAN {
    v := strconv.FormatUint(uint64(version), 10)

    clientID := site + sensor + v
    clusterID := config.ClusterID

    serverURLs := []string{}
    for _, server := range config.Servers {
        port := strconv.FormatUint(uint64(server.Port), 10)
        url := "nats://" + server.Host + ":" + port
        serverURLs = append(serverURLs, url)
    }

    nc, _ := nats.Connect(strings.Join(serverURLs, ","))
    sc, _ := s.Connect(clusterID, clientID, stan.NatsConn(nc))

    mySTAN := &STAN{
        subject: strings.Join([]string{"eda", site, sensor, v, "datum", "create"}, "."),
        sc:      sc,
        nc:      nc,
        log:     logger,
    }

    return mySTAN
}

// Publish TODO
func (s *STAN) Publish(datum *lemi011b.Datum) {
    data := fmt.Sprintf("%+v", datum)

    err := s.sc.Publish(s.subject, []byte(data))
    if err != nil {
        s.log.Error("failed to publish datum", err)
    }

    return
}

// Close TODO
func (s *STAN) Close() {
    s.nc.Close()
    s.sc.Close()
}
