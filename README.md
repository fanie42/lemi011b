# lemi011b
All services for Lemi011b magnetometer

## DAQ
The DAQ application is the client that runs on the instrument logger. It collects data form the serial port, timestamps it, formats it and then publishes a "DatumAcquired" event on the
event bus - in this case we're using NATS with STAN as an event store.

## FS
This application can run anywhere and needs an allocated directory for storing the data in file format. It listens on the event bus and writes data to file.

## MON
Can run anywhere. Listens to the event bus and stores events in an SQL table that is easy to query by monitoring services - Grafana. In this case, we use TimescaleDB.