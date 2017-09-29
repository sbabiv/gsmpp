package decoders

import (
	"net"
	"encoding/binary"
	"github.com/sbabiv/gsmpp/smpp/pdu"
)

func DecodeHeader(conn net.Conn) (*pdu.Header, error) {
	b := make([]byte, pdu.HeaderLength, pdu.HeaderLength)
	n, err := conn.Read(b)

	if n == 0 || err != nil {
		return nil, err
	}

	return &pdu.Header{
		Length:   binary.BigEndian.Uint32(b[0:4]),
		Id:       binary.BigEndian.Uint32(b[4:8]),
		Status:   binary.BigEndian.Uint32(b[8:12]),
		Sequence: binary.BigEndian.Uint32(b[12:16]),
	}, nil
}

func DecodeBindResp(h *pdu.Header, conn net.Conn) (*pdu.UnitRsp, error) {
	return pdu.D(h, conn,
		pdu.FieldNames{
			pdu.SystemId,
		})
}

func DecodeSubmitSmResp(h *pdu.Header, conn net.Conn) (*pdu.UnitRsp, error) {
	return pdu.D(h, conn,
		pdu.FieldNames{
			pdu.MessageId,
		})
}

func DecodeDeliverSm(h *pdu.Header, conn net.Conn)(*pdu.UnitRsp, error) {
	return pdu.D(h, conn,
	pdu.FieldNames{
		pdu.ServiceType,
		pdu.SourceAddrTon,
		pdu.SourceAddrNpi,
		pdu.SourceAddr,
		pdu.DestAddrTon,
		pdu.DestAddrNpi,
		pdu.DestinationAddr,
		pdu.EsmClass,
		pdu.ProtocolId,
		pdu.PriorityFlag,
		pdu.ScheduleDeliveryTime,
		pdu.ValidityPeriod,
		pdu.RegisteredDelivery,
		pdu.ReplaceIfPresentFlag,
		pdu.DataCoding,
		pdu.SmDefaultMsgId,
		pdu.SmLength,
		pdu.ShortMessage,
	})
}