package packet

import (
	"encoding/binary"
	"github.com/danielpaulus/quicktime_video_hack/usb/dict"
)

//Async Packet types
const (
	AsynPacketMagic uint32 = 0x6173796E
	FEED            uint32 = 0x66656564 //These contain CMSampleBufs which contain raw h264 Nalus
	TJMP            uint32 = 0x746A6D70
	SRAT            uint32 = 0x73726174
	SPRP            uint32 = 0x73707270
	TBAS            uint32 = 0x74626173
	RELS            uint32 = 0x72656C73
	HPD1            uint32 = 0x68706431 //hpd1 - 1dph | Maybe Hotplug Detection?
	HPA1            uint32 = 0x68706131 //hpa1 - 1aph | high performance addressing?
)

const (
	hpa1Header uint64 = 0x00000001198D57B0
	hpd1Header uint64 = 0x0000000000000001
)

//It seems like Need Packets are constant
var AsynNeedPacketBytes = asynNeedPacketBytes()

type AsyncPacket struct {
	Header                     uint64 //I don't know what the first 8 bytes are for currently
	HumanReadableTypeSpecifier uint32 //One of the packet types above
	Payload                    interface{}
}

func NewAsynHpd1Packet(stringKeyDict dict.StringKeyDict) []byte {
	return newAsynDictPacket(stringKeyDict, HPD1, hpd1Header)
}

func NewAsynHpa1Packet(stringKeyDict dict.StringKeyDict) []byte {
	return newAsynDictPacket(stringKeyDict, HPA1, hpa1Header)
}

func newAsynDictPacket(stringKeyDict dict.StringKeyDict, subtypeMarker uint32, asynTypeHeader uint64) []byte {
	serialize := dict.SerializeStringKeyDict(stringKeyDict)
	length := len(serialize) + 20
	header := make([]byte, 20)
	binary.LittleEndian.PutUint32(header, uint32(length))
	binary.LittleEndian.PutUint32(header[4:], AsynPacketMagic)
	binary.LittleEndian.PutUint64(header[8:], asynTypeHeader)
	binary.LittleEndian.PutUint32(header[16:], subtypeMarker)
	return append(header, serialize...)
}

func asynNeedPacketBytes() []byte {
	needPacketLength := 20
	packet := make([]byte, needPacketLength)
	binary.LittleEndian.PutUint32(packet, uint32(needPacketLength))
	binary.LittleEndian.PutUint32(packet, AsynPacketMagic)
	binary.LittleEndian.PutUint64(packet, 0x0000000104CBB860) //don't know what these mean but they seem constant
	binary.LittleEndian.PutUint32(packet, 0x6E656564)         //need - deen
	return packet
}