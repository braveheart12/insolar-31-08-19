/*
 * The Clear BSD License
 *
 * Copyright (c) 2019 Insolar Technologies
 *
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without modification, are permitted (subject to the limitations in the disclaimer below) provided that the following conditions are met:
 *
 *  Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 *  Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 *  Neither the name of Insolar Technologies nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
 *
 * NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package transport

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"

	consensus "github.com/insolar/insolar/consensus/packets"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/log"
	"github.com/insolar/insolar/network/transport/host"
	"github.com/insolar/insolar/network/transport/packet"
	"github.com/insolar/insolar/network/transport/relay"
	"github.com/insolar/insolar/network/utils"
	"github.com/pkg/errors"
)

const udpMaxPacketSize = 1400

type udpTransport struct {
	baseTransport
	serverConn net.PacketConn
}

type udpSerializer struct{}

func (b *udpSerializer) SerializePacket(q *packet.Packet) ([]byte, error) {
	data, ok := q.Data.(consensus.ConsensusPacket)
	if !ok {
		return nil, errors.New("could not convert packet to ConsensusPacket type")
	}
	header := &consensus.RoutingHeader{
		OriginID:   q.Sender.ShortID,
		TargetID:   q.Receiver.ShortID,
		PacketType: q.Type,
	}
	err := data.SetPacketHeader(header)
	if err != nil {
		return nil, errors.Wrap(err, "could not set routing information for ConsensusPacket")
	}
	return data.Serialize()
}

func (b *udpSerializer) DeserializePacket(conn io.Reader) (*packet.Packet, error) {
	data, err := consensus.ExtractPacket(conn)
	if err != nil {
		return nil, errors.Wrap(err, "could not convert network datagram to ConsensusPacket")
	}
	header, err := data.GetPacketHeader()
	if err != nil {
		return nil, errors.Wrap(err, "could not get routing information from ConsensusPacket")
	}
	p := &packet.Packet{}
	p.Sender = &host.Host{ShortID: header.OriginID}
	p.Receiver = &host.Host{ShortID: header.TargetID}
	p.Type = header.PacketType
	return p, nil
}

func newUDPTransport(conn net.PacketConn, proxy relay.Proxy, publicAddress string) (*udpTransport, error) {
	transport := &udpTransport{
		baseTransport: newBaseTransport(proxy, publicAddress),
		serverConn:    conn}
	transport.sendFunc = transport.send
	transport.serializer = &udpSerializer{}

	return transport, nil
}

func (t *udpTransport) send(recvAddress string, data []byte) error {
	log.Debug("Sending PURE_UDP request")
	if len(data) > udpMaxPacketSize {
		return errors.New(fmt.Sprintf("udpTransport.send: too big input data. Maximum: %d. Current: %d",
			udpMaxPacketSize, len(data)))
	}

	// TODO: may be try to send second time if error
	// TODO: skip resolving every time by caching result
	udpAddr, err := net.ResolveUDPAddr("udp", recvAddress)
	if err != nil {
		return errors.Wrap(err, "udpTransport.send")
	}

	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return errors.Wrap(err, "udpTransport.send")
	}
	defer utils.CloseVerbose(udpConn)

	log.Debug("udpTransport.send: len = ", len(data))
	_, err = udpConn.Write(data)
	return errors.Wrap(err, "Failed to write data")
}

// Start starts networking.
func (t *udpTransport) Listen(ctx context.Context, started chan struct{}) error {
	inslogger.FromContext(ctx).Info("Start UDP transport")
	started <- struct{}{}
	for {
		buf := make([]byte, udpMaxPacketSize)
		n, addr, err := t.serverConn.ReadFrom(buf)
		if err != nil {
			<-t.disconnectFinished
			return err
		}

		go t.handleAcceptedConnection(buf[:n], addr)
	}
}

// Stop stops networking.
func (t *udpTransport) Stop() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	log.Info("Stop UDP transport")
	t.prepareDisconnect()

	utils.CloseVerbose(t.serverConn)
}

func (t *udpTransport) handleAcceptedConnection(data []byte, addr net.Addr) {
	r := bytes.NewReader(data)
	msg, err := t.serializer.DeserializePacket(r)
	if err != nil {
		log.Error("[ handleAcceptedConnection ] ", err)
		return
	}
	log.Debug("[ handleAcceptedConnection ] Packet processed. size: ", len(data), ". Address: ", addr)

	go t.packetHandler.Handle(context.TODO(), msg)
}
