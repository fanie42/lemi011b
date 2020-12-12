package client

import (
    "bufio"
    "fmt"

    "github.com/tarm/serial"
)

// Serial TODO
type Serial struct {
    sp *serial.Port
}

// NewSerial TODO
func NewSerial(serialPort *serial.Port) (*Serial, error) {
    if serialPort != nil {
        return &Serial{
            sp: serialPort,
        }, nil
    }

    return nil, fmt.Errorf("serial port may not be nil")
}

// Control TODO
func (s *Serial) Control(svc *) {
    scanner := bufio.NewScanner(s.sp)
    for scanner.Scan() {
        b := scanner.Bytes()
        s
    }
}
