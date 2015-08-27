package membership

import (
	"log"
	"net/rpc"
	"time"
)

type Membership struct {
}

func (m *Membership) Start() {
	m.startRPCServer()

	t := time.NewTimer(time.Second)
	for {
		select {
		case <-t:
			m.doGossip()
		}
	}
}

func (m *Membership) startRPCServer() {
	rpc.Register(m)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", 8080)
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	go http.Serve(l, nil)
}

func (m *Membership) doGossip() {
	m.doSyn()
}

func (m *Membership) doSyn() {
	args := new(SynPacket)
	args.Id = "test"

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	client.Go("Membership.HandleSyn", &args, nil, nil)
}

func (m *Membership) doSynAck() {
	args := new(SynAckPacket)
	args.Id = "test"

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	client.Go("Membership.HandleSynAck", &args, nil, nil)
}

func (m *Membership) doAck() {
	args := new(AckPacket)
	args.Id = "test"

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	client.Go("Membership.HandleSyn", &args, nil, nil)
}

func (m *Membership) HandleSyn(s *SynPacket, e *Empty) error {
	fmt.Println("fuck")
	return nil
}

func (m *Membership) HandleSynAck(sa *SynAckPacket, e *Empty) error {
	return nil
}

func (m *Membership) HandleAck(a *AckPacket, e *Empty) error {
	return nil
}
