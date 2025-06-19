package nex

type PRUDPConnectionSocket struct {
	remote *PRUDPConnection
}

func (cs *PRUDPConnectionSocket) Send(data []byte) {
	var messagePacket PRUDPPacketInterface

	if cs.remote.DefaultPRUDPVersion == 0 {
		messagePacket, _ = NewPRUDPPacketV0(cs.remote.endpoint.Server, cs.remote, nil)
	} else {
		messagePacket, _ = NewPRUDPPacketV1(cs.remote.endpoint.Server, cs.remote, nil)
	}

	messagePacket.SetPayload(data)

	cs.remote.endpoint.Server.Send(messagePacket)
}

func (cs *PRUDPConnectionSocket) Recv() []byte {
	return <-cs.remote.PacketChannel
}
