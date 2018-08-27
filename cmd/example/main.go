package main

import (
	"log"

	"github.com/forfuncsake/smartswitch"
)

type sw struct {
	status bool
}

func (s *sw) Status() (bool, error) {
	log.Printf("status requested (%t)\n", s.status)
	return s.status, nil
}

func (s *sw) Set(status bool) error {
	log.Printf("status set to %t\n", status)
	s.status = status
	return nil
}

func main() {
	control := make(chan struct{})
	s := &sw{}

	c := smartswitch.NewController("tester", s)
	loc, err := c.Start()
	if err != nil {
		panic(err)
	}
	log.Printf("example switch service location: %s\n", loc)
	defer c.Stop()

	<-control
}
