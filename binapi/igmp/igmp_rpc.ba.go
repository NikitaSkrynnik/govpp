// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package igmp

import (
	"context"
	"fmt"
	"io"

	memclnt "github.com/networkservicemesh/govpp/binapi/memclnt"
	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service igmp.
type RPCService interface {
	IgmpClearInterface(ctx context.Context, in *IgmpClearInterface) (*IgmpClearInterfaceReply, error)
	IgmpDump(ctx context.Context, in *IgmpDump) (RPCService_IgmpDumpClient, error)
	IgmpEnableDisable(ctx context.Context, in *IgmpEnableDisable) (*IgmpEnableDisableReply, error)
	IgmpGroupPrefixDump(ctx context.Context, in *IgmpGroupPrefixDump) (RPCService_IgmpGroupPrefixDumpClient, error)
	IgmpGroupPrefixSet(ctx context.Context, in *IgmpGroupPrefixSet) (*IgmpGroupPrefixSetReply, error)
	IgmpListen(ctx context.Context, in *IgmpListen) (*IgmpListenReply, error)
	IgmpProxyDeviceAddDel(ctx context.Context, in *IgmpProxyDeviceAddDel) (*IgmpProxyDeviceAddDelReply, error)
	IgmpProxyDeviceAddDelInterface(ctx context.Context, in *IgmpProxyDeviceAddDelInterface) (*IgmpProxyDeviceAddDelInterfaceReply, error)
	WantIgmpEvents(ctx context.Context, in *WantIgmpEvents) (*WantIgmpEventsReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) IgmpClearInterface(ctx context.Context, in *IgmpClearInterface) (*IgmpClearInterfaceReply, error) {
	out := new(IgmpClearInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IgmpDump(ctx context.Context, in *IgmpDump) (RPCService_IgmpDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IgmpDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IgmpDumpClient interface {
	Recv() (*IgmpDetails, error)
	api.Stream
}

type serviceClient_IgmpDumpClient struct {
	api.Stream
}

func (c *serviceClient_IgmpDumpClient) Recv() (*IgmpDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IgmpDetails:
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

func (c *serviceClient) IgmpEnableDisable(ctx context.Context, in *IgmpEnableDisable) (*IgmpEnableDisableReply, error) {
	out := new(IgmpEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IgmpGroupPrefixDump(ctx context.Context, in *IgmpGroupPrefixDump) (RPCService_IgmpGroupPrefixDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IgmpGroupPrefixDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IgmpGroupPrefixDumpClient interface {
	Recv() (*IgmpGroupPrefixDetails, error)
	api.Stream
}

type serviceClient_IgmpGroupPrefixDumpClient struct {
	api.Stream
}

func (c *serviceClient_IgmpGroupPrefixDumpClient) Recv() (*IgmpGroupPrefixDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IgmpGroupPrefixDetails:
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

func (c *serviceClient) IgmpGroupPrefixSet(ctx context.Context, in *IgmpGroupPrefixSet) (*IgmpGroupPrefixSetReply, error) {
	out := new(IgmpGroupPrefixSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IgmpListen(ctx context.Context, in *IgmpListen) (*IgmpListenReply, error) {
	out := new(IgmpListenReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IgmpProxyDeviceAddDel(ctx context.Context, in *IgmpProxyDeviceAddDel) (*IgmpProxyDeviceAddDelReply, error) {
	out := new(IgmpProxyDeviceAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IgmpProxyDeviceAddDelInterface(ctx context.Context, in *IgmpProxyDeviceAddDelInterface) (*IgmpProxyDeviceAddDelInterfaceReply, error) {
	out := new(IgmpProxyDeviceAddDelInterfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WantIgmpEvents(ctx context.Context, in *WantIgmpEvents) (*WantIgmpEventsReply, error) {
	out := new(WantIgmpEventsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
