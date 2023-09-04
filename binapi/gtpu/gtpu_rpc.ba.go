// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package gtpu

import (
	"context"
	"fmt"
	"io"

	memclnt "github.com/networkservicemesh/govpp/binapi/memclnt"
	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service gtpu.
type RPCService interface {
	GtpuAddDelTunnel(ctx context.Context, in *GtpuAddDelTunnel) (*GtpuAddDelTunnelReply, error)
	GtpuOffloadRx(ctx context.Context, in *GtpuOffloadRx) (*GtpuOffloadRxReply, error)
	GtpuTunnelDump(ctx context.Context, in *GtpuTunnelDump) (RPCService_GtpuTunnelDumpClient, error)
	GtpuTunnelUpdateTteid(ctx context.Context, in *GtpuTunnelUpdateTteid) (*GtpuTunnelUpdateTteidReply, error)
	SwInterfaceSetGtpuBypass(ctx context.Context, in *SwInterfaceSetGtpuBypass) (*SwInterfaceSetGtpuBypassReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) GtpuAddDelTunnel(ctx context.Context, in *GtpuAddDelTunnel) (*GtpuAddDelTunnelReply, error) {
	out := new(GtpuAddDelTunnelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GtpuOffloadRx(ctx context.Context, in *GtpuOffloadRx) (*GtpuOffloadRxReply, error) {
	out := new(GtpuOffloadRxReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GtpuTunnelDump(ctx context.Context, in *GtpuTunnelDump) (RPCService_GtpuTunnelDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GtpuTunnelDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GtpuTunnelDumpClient interface {
	Recv() (*GtpuTunnelDetails, error)
	api.Stream
}

type serviceClient_GtpuTunnelDumpClient struct {
	api.Stream
}

func (c *serviceClient_GtpuTunnelDumpClient) Recv() (*GtpuTunnelDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GtpuTunnelDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) GtpuTunnelUpdateTteid(ctx context.Context, in *GtpuTunnelUpdateTteid) (*GtpuTunnelUpdateTteidReply, error) {
	out := new(GtpuTunnelUpdateTteidReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetGtpuBypass(ctx context.Context, in *SwInterfaceSetGtpuBypass) (*SwInterfaceSetGtpuBypassReply, error) {
	out := new(SwInterfaceSetGtpuBypassReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
