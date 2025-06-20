package nex

import "github.com/PretendoNetwork/nex-go/v2/constants"

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
	messagePacket.SetType(constants.DataPacket)
	messagePacket.AddFlag(constants.PacketFlagNeedsAck)
	messagePacket.AddFlag(constants.PacketFlagHasSize)
	messagePacket.AddFlag(constants.PacketFlagReliable)
	messagePacket.SetSourceVirtualPortStreamType(cs.remote.StreamType)
	messagePacket.SetSourceVirtualPortStreamID(cs.remote.endpoint.StreamID)
	messagePacket.SetDestinationVirtualPortStreamType(cs.remote.StreamType)
	messagePacket.SetDestinationVirtualPortStreamID(cs.remote.StreamID)
	messagePacket.SetSubstreamID(0)

	cs.remote.endpoint.Server.Send(messagePacket)
}

func (cs *PRUDPConnectionSocket) Recv() *[]byte {
	for conn := range cs.remote.PacketChannel {
		return conn
	}

	return nil
}
