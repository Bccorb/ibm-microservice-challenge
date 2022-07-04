package cmd_test

import (
	"context"
	"errors"
	"testing"

	pb "github.com/bccorb/pkg/gts/globalTradeSystem"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Make a Mock Stream
func makeStreamMock() *StreamMock {
	return &StreamMock{
		ctx:            context.Background(),
		recvToServer:   make(chan pb.TradeRequest),
		sentFromServer: make(chan pb.TradeResponse),
	}
}

// Mock Stream stucture
type StreamMock struct {
	grpc.ServerStream
	ctx            context.Context
	recvToServer   chan pb.TradeRequest
	sentFromServer chan pb.TradeResponse
}

// Mock GlobalTradeSystem Stream Server Service
type GlobalTradeSystem_GetTradeListStreamServer interface {
	Send(pb.TradeResponse) error
	Recv() (pb.TradeRequest, error)
	grpc.ServerStream
}

// Mock Context
func (m *StreamMock) Context() context.Context {
	return m.ctx
}

// Mock Send service
func (m *StreamMock) Send(resp pb.TradeResponse) error {
	m.sentFromServer <- resp
	return nil
}

// Mock Recv service
func (m *StreamMock) Recv() (pb.TradeRequest, error) {
	req, more := <-m.recvToServer
	if !more {
		return pb.TradeRequest{}, errors.New("empty")
	}
	return req, nil
}

// Mock a client request
func (m *StreamMock) SendFromClient(req pb.TradeRequest) error {
	m.recvToServer <- req
	return nil
}

// Mock a client repsonse
func (m *StreamMock) RecvToClient() (pb.TradeResponse, error) {
	response, more := <-m.sentFromServer
	if !more {
		return pb.TradeResponse{}, errors.New("empty")
	}
	return response, nil
}

// Mock API structure
type Api struct{}

// Mock Register func for setting up a server
func RegisterGlobalTradeSystemServer() *Api {
	return &Api{}
}

// Mocked Global Trade System Stream with Request and Reponse
func (a Api) GlobalTradeSystemStream(stream GlobalTradeSystem_GetTradeListStreamServer) error {
	for {
		TradeRequest, err := stream.Recv()
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		err = stream.Send(pb.TradeResponse{OfferedPokemon: TradeRequest.RequestedPokemon, RequestedPokemon: TradeRequest.OfferedPokemon})

		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
}

// Test that given a Request with an offered pokemon and requested pokemon that the Server will
// acknowledge and return a response that the trade is avalible to make.
func TestGlobalTradeSystem(t *testing.T) {
	stream := createStream(t)
	err := stream.SendFromClient(pb.TradeRequest{OfferedPokemon: "Zubat", RequestedPokemon: "Snorlax"})
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	TradeResponse, err := stream.RecvToClient()
	if err != nil {
		t.Error(err.Error())
		return
	}

	requestedPokemon := TradeResponse.OfferedPokemon
	offeredPokemon := TradeResponse.RequestedPokemon

	if !(requestedPokemon == "Snorlax") || !(offeredPokemon == "Zubat") {
		t.Errorf("Expected to trade a Snoarlax for a Zubat, instead received %v for a %v", requestedPokemon, offeredPokemon)
	}
}

// Mock Stream
func createStream(t *testing.T) *StreamMock {
	stream := makeStreamMock()
	go func() {
		api := RegisterGlobalTradeSystemServer()
		err := api.GlobalTradeSystemStream(stream)
		if err != nil {
			t.Errorf(err.Error())
		}
		close(stream.sentFromServer)
		close(stream.recvToServer)
	}()
	return stream
}
