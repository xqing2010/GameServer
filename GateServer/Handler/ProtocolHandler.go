package Handler

import (
    "GameServer/Common/Network"
)

//ProtocolHandler protocol handler
func ProtocolHandler(session *Network.Session, protocol *Network.Protocol)  {
    session.SendPacket(protocol)
}