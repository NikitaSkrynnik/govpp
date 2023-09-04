// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package acl

import (
	"context"
	"fmt"
	"io"

	memclnt "github.com/networkservicemesh/govpp/binapi/memclnt"
	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service acl.
type RPCService interface {
	ACLAddReplace(ctx context.Context, in *ACLAddReplace) (*ACLAddReplaceReply, error)
	ACLDel(ctx context.Context, in *ACLDel) (*ACLDelReply, error)
	ACLDump(ctx context.Context, in *ACLDump) (RPCService_ACLDumpClient, error)
	ACLInterfaceAddDel(ctx context.Context, in *ACLInterfaceAddDel) (*ACLInterfaceAddDelReply, error)
	ACLInterfaceEtypeWhitelistDump(ctx context.Context, in *ACLInterfaceEtypeWhitelistDump) (RPCService_ACLInterfaceEtypeWhitelistDumpClient, error)
	ACLInterfaceListDump(ctx context.Context, in *ACLInterfaceListDump) (RPCService_ACLInterfaceListDumpClient, error)
	ACLInterfaceSetACLList(ctx context.Context, in *ACLInterfaceSetACLList) (*ACLInterfaceSetACLListReply, error)
	ACLInterfaceSetEtypeWhitelist(ctx context.Context, in *ACLInterfaceSetEtypeWhitelist) (*ACLInterfaceSetEtypeWhitelistReply, error)
	ACLPluginControlPing(ctx context.Context, in *ACLPluginControlPing) (*ACLPluginControlPingReply, error)
	ACLPluginGetConnTableMaxEntries(ctx context.Context, in *ACLPluginGetConnTableMaxEntries) (*ACLPluginGetConnTableMaxEntriesReply, error)
	ACLPluginGetVersion(ctx context.Context, in *ACLPluginGetVersion) (*ACLPluginGetVersionReply, error)
	ACLPluginUseHashLookupGet(ctx context.Context, in *ACLPluginUseHashLookupGet) (*ACLPluginUseHashLookupGetReply, error)
	ACLPluginUseHashLookupSet(ctx context.Context, in *ACLPluginUseHashLookupSet) (*ACLPluginUseHashLookupSetReply, error)
	ACLStatsIntfCountersEnable(ctx context.Context, in *ACLStatsIntfCountersEnable) (*ACLStatsIntfCountersEnableReply, error)
	MacipACLAdd(ctx context.Context, in *MacipACLAdd) (*MacipACLAddReply, error)
	MacipACLAddReplace(ctx context.Context, in *MacipACLAddReplace) (*MacipACLAddReplaceReply, error)
	MacipACLDel(ctx context.Context, in *MacipACLDel) (*MacipACLDelReply, error)
	MacipACLDump(ctx context.Context, in *MacipACLDump) (RPCService_MacipACLDumpClient, error)
	MacipACLInterfaceAddDel(ctx context.Context, in *MacipACLInterfaceAddDel) (*MacipACLInterfaceAddDelReply, error)
	MacipACLInterfaceGet(ctx context.Context, in *MacipACLInterfaceGet) (*MacipACLInterfaceGetReply, error)
	MacipACLInterfaceListDump(ctx context.Context, in *MacipACLInterfaceListDump) (RPCService_MacipACLInterfaceListDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) ACLAddReplace(ctx context.Context, in *ACLAddReplace) (*ACLAddReplaceReply, error) {
	out := new(ACLAddReplaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ACLDel(ctx context.Context, in *ACLDel) (*ACLDelReply, error) {
	out := new(ACLDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ACLDump(ctx context.Context, in *ACLDump) (RPCService_ACLDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_ACLDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_ACLDumpClient interface {
	Recv() (*ACLDetails, error)
	api.Stream
}

type serviceClient_ACLDumpClient struct {
	api.Stream
}

func (c *serviceClient_ACLDumpClient) Recv() (*ACLDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *ACLDetails:
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

func (c *serviceClient) ACLInterfaceAddDel(ctx context.Context, in *ACLInterfaceAddDel) (*ACLInterfaceAddDelReply, error) {
	out := new(ACLInterfaceAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ACLInterfaceEtypeWhitelistDump(ctx context.Context, in *ACLInterfaceEtypeWhitelistDump) (RPCService_ACLInterfaceEtypeWhitelistDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_ACLInterfaceEtypeWhitelistDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_ACLInterfaceEtypeWhitelistDumpClient interface {
	Recv() (*ACLInterfaceEtypeWhitelistDetails, error)
	api.Stream
}

type serviceClient_ACLInterfaceEtypeWhitelistDumpClient struct {
	api.Stream
}

func (c *serviceClient_ACLInterfaceEtypeWhitelistDumpClient) Recv() (*ACLInterfaceEtypeWhitelistDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *ACLInterfaceEtypeWhitelistDetails:
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

func (c *serviceClient) ACLInterfaceListDump(ctx context.Context, in *ACLInterfaceListDump) (RPCService_ACLInterfaceListDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_ACLInterfaceListDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_ACLInterfaceListDumpClient interface {
	Recv() (*ACLInterfaceListDetails, error)
	api.Stream
}

type serviceClient_ACLInterfaceListDumpClient struct {
	api.Stream
}

func (c *serviceClient_ACLInterfaceListDumpClient) Recv() (*ACLInterfaceListDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *ACLInterfaceListDetails:
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

func (c *serviceClient) ACLInterfaceSetACLList(ctx context.Context, in *ACLInterfaceSetACLList) (*ACLInterfaceSetACLListReply, error) {
	out := new(ACLInterfaceSetACLListReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ACLInterfaceSetEtypeWhitelist(ctx context.Context, in *ACLInterfaceSetEtypeWhitelist) (*ACLInterfaceSetEtypeWhitelistReply, error) {
	out := new(ACLInterfaceSetEtypeWhitelistReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ACLPluginControlPing(ctx context.Context, in *ACLPluginControlPing) (*ACLPluginControlPingReply, error) {
	out := new(ACLPluginControlPingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ACLPluginGetConnTableMaxEntries(ctx context.Context, in *ACLPluginGetConnTableMaxEntries) (*ACLPluginGetConnTableMaxEntriesReply, error) {
	out := new(ACLPluginGetConnTableMaxEntriesReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ACLPluginGetVersion(ctx context.Context, in *ACLPluginGetVersion) (*ACLPluginGetVersionReply, error) {
	out := new(ACLPluginGetVersionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ACLPluginUseHashLookupGet(ctx context.Context, in *ACLPluginUseHashLookupGet) (*ACLPluginUseHashLookupGetReply, error) {
	out := new(ACLPluginUseHashLookupGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ACLPluginUseHashLookupSet(ctx context.Context, in *ACLPluginUseHashLookupSet) (*ACLPluginUseHashLookupSetReply, error) {
	out := new(ACLPluginUseHashLookupSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ACLStatsIntfCountersEnable(ctx context.Context, in *ACLStatsIntfCountersEnable) (*ACLStatsIntfCountersEnableReply, error) {
	out := new(ACLStatsIntfCountersEnableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MacipACLAdd(ctx context.Context, in *MacipACLAdd) (*MacipACLAddReply, error) {
	out := new(MacipACLAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MacipACLAddReplace(ctx context.Context, in *MacipACLAddReplace) (*MacipACLAddReplaceReply, error) {
	out := new(MacipACLAddReplaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MacipACLDel(ctx context.Context, in *MacipACLDel) (*MacipACLDelReply, error) {
	out := new(MacipACLDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MacipACLDump(ctx context.Context, in *MacipACLDump) (RPCService_MacipACLDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_MacipACLDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_MacipACLDumpClient interface {
	Recv() (*MacipACLDetails, error)
	api.Stream
}

type serviceClient_MacipACLDumpClient struct {
	api.Stream
}

func (c *serviceClient_MacipACLDumpClient) Recv() (*MacipACLDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *MacipACLDetails:
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

func (c *serviceClient) MacipACLInterfaceAddDel(ctx context.Context, in *MacipACLInterfaceAddDel) (*MacipACLInterfaceAddDelReply, error) {
	out := new(MacipACLInterfaceAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MacipACLInterfaceGet(ctx context.Context, in *MacipACLInterfaceGet) (*MacipACLInterfaceGetReply, error) {
	out := new(MacipACLInterfaceGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MacipACLInterfaceListDump(ctx context.Context, in *MacipACLInterfaceListDump) (RPCService_MacipACLInterfaceListDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_MacipACLInterfaceListDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_MacipACLInterfaceListDumpClient interface {
	Recv() (*MacipACLInterfaceListDetails, error)
	api.Stream
}

type serviceClient_MacipACLInterfaceListDumpClient struct {
	api.Stream
}

func (c *serviceClient_MacipACLInterfaceListDumpClient) Recv() (*MacipACLInterfaceListDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *MacipACLInterfaceListDetails:
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
