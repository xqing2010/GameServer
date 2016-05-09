package Network

import (
    "GameServer/Common/PBProto"
    "github.com/golang/protobuf/proto"
    "bytes"
    "encoding/binary"
    "fmt"
)

//ProtoID Protocol id type
type ProtoID uint32
const (
    LoginID = iota;
)

const (
    IDEnd = 4;
    TypeEnd = 8;
)

type ProtoType uint32

const (
    PT_Login = iota;
    PT_ServerCenter;
    PT_GameSever;
    PT_GlobalServer;
)

//ProtocolInterface protocol interface
type ProtocolInterface interface {
    Marshal() ([]byte, error)
    UnMarshal(data []byte) (int, error)
}

//Protocol protocol 
type Protocol struct {
    ID ProtoID;
    PType ProtoType;
    Packet interface{}
}

var packetMap map[ProtoID]interface{}

func init()  {
    packetMap = make(map[ProtoID]interface{})
    packetMap[LoginID] = new(PBProto.Login);   
}

//NewPBPacket create PBMessage by id  the only way to create Packet!!!
func NewPBPacket(id ProtoID) interface{} {
    packet, ok := packetMap[id];
    if !ok {
        return nil;
    }
    ms, _ := packet.(proto.Message);
    
    return proto.Clone(ms)
}

//NewProtocol create Protocol
func NewProtocol(id ProtoID) *Protocol {
    protocol := new(Protocol);
    protocol.ID = id;
    protocol.Packet = NewPBPacket(id);
    
    return protocol;
}

//Marshal convert the protocol to bytes
func (protocol *Protocol)Marshal() ([]byte, error) {
    buff := new(bytes.Buffer);
    binary.Write(buff, binary.BigEndian, protocol.ID);
    binary.Write(buff, binary.BigEndian, protocol.PType)
    ms, ok := protocol.Packet.(proto.Message);
    if !ok {
		return nil, fmt.Errorf("Protocol error not valid protobuff");
	}
    data, err := proto.Marshal(ms)
    if nil != err {
        return nil, fmt.Errorf("Packet Marshal Error");
    }
    buff.Write(data);
    return buff.Bytes(), nil
}

//UnMarshal Protocol's Unmarshal UnSerialize protocol from bytes
func (protocol *Protocol) UnMarshal(data []byte) (int, error) {
    if(len(data) < TypeEnd) {
        //不完整的数据 待下次再读
        return 0, fmt.Errorf("incomplete data"); 
    }
    
    idSplit := data[:IDEnd]
    ptypeSplit := data[IDEnd:TypeEnd]
    packSplit := data[TypeEnd:]
    
    protocol.ID = ProtoID(binary.LittleEndian.Uint32(idSplit))
    protocol.PType = ProtoType(binary.LittleEndian.Uint32(ptypeSplit))
    
    protocol.Packet = NewPBPacket(protocol.ID)
    
    if nil == protocol.Packet {
        return -1, fmt.Errorf("Invalid packetId, need Close client??? id = %d, data = %v", protocol.ID, data)
    }
    
    ms, _ := protocol.Packet.(proto.Message);
    err := proto.Unmarshal(packSplit, ms)
    
    if nil != err {
        return 0, fmt.Errorf("PBMessage Unmarshal Error!!! incomplete packet or need close client")
    }
    
    msLen := proto.Size(ms)

    packetLen := msLen + TypeEnd
    return packetLen, nil
}

