// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package lcp

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service lcp.
type RPCService interface {
	LcpDefaultNsGet(ctx context.Context, in *LcpDefaultNsGet) (*LcpDefaultNsGetReply, error)
	LcpDefaultNsSet(ctx context.Context, in *LcpDefaultNsSet) (*LcpDefaultNsSetReply, error)
	LcpItfPairAddDel(ctx context.Context, in *LcpItfPairAddDel) (*LcpItfPairAddDelReply, error)
	LcpItfPairAddDelV2(ctx context.Context, in *LcpItfPairAddDelV2) (*LcpItfPairAddDelV2Reply, error)
	LcpItfPairGet(ctx context.Context, in *LcpItfPairGet) (RPCService_LcpItfPairGetClient, error)
	LcpItfPairReplaceBegin(ctx context.Context, in *LcpItfPairReplaceBegin) (*LcpItfPairReplaceBeginReply, error)
	LcpItfPairReplaceEnd(ctx context.Context, in *LcpItfPairReplaceEnd) (*LcpItfPairReplaceEndReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) LcpDefaultNsGet(ctx context.Context, in *LcpDefaultNsGet) (*LcpDefaultNsGetReply, error) {
	out := new(LcpDefaultNsGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) LcpDefaultNsSet(ctx context.Context, in *LcpDefaultNsSet) (*LcpDefaultNsSetReply, error) {
	out := new(LcpDefaultNsSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LcpItfPairAddDel(ctx context.Context, in *LcpItfPairAddDel) (*LcpItfPairAddDelReply, error) {
	out := new(LcpItfPairAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LcpItfPairAddDelV2(ctx context.Context, in *LcpItfPairAddDelV2) (*LcpItfPairAddDelV2Reply, error) {
	out := new(LcpItfPairAddDelV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LcpItfPairGet(ctx context.Context, in *LcpItfPairGet) (RPCService_LcpItfPairGetClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_LcpItfPairGetClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_LcpItfPairGetClient interface {
	Recv() (*LcpItfPairDetails, *LcpItfPairGetReply, error)
	api.Stream
}

type serviceClient_LcpItfPairGetClient struct {
	api.Stream
}

func (c *serviceClient_LcpItfPairGetClient) Recv() (*LcpItfPairDetails, *LcpItfPairGetReply, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, nil, err
	}
	switch m := msg.(type) {
	case *LcpItfPairDetails:
		return m, nil, nil
	case *LcpItfPairGetReply:
		if err := api.RetvalToVPPApiError(m.Retval); err != nil {
			c.Stream.Close()
			return nil, m, err
		}
		err = c.Stream.Close()
		if err != nil {
			return nil, m, err
		}
		return nil, m, io.EOF
	default:
		return nil, nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) LcpItfPairReplaceBegin(ctx context.Context, in *LcpItfPairReplaceBegin) (*LcpItfPairReplaceBeginReply, error) {
	out := new(LcpItfPairReplaceBeginReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LcpItfPairReplaceEnd(ctx context.Context, in *LcpItfPairReplaceEnd) (*LcpItfPairReplaceEndReply, error) {
	out := new(LcpItfPairReplaceEndReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
