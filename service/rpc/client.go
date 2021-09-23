package rpc

import (
	"context"
	"fmt"
	client3 "github.com/rpcxio/rpcx-etcd/client"
	client2 "github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

type Cli interface {
	Call(ctx context.Context, fn string, request, reply interface{}) error
	Broadcast(fn string, request, reply interface{}) error
	Run() error
	Close() error
}

type ClientOptions struct {
	client2.Option

	Addr        string
	Port        int
	Name        string
	EtcdServers []string
	Selector    client2.Selector
}

type BaseClient struct {
	cli     client2.XClient
	options *ClientOptions
	id      string
}

func NewBaseClient(options *ClientOptions) (*BaseClient, error) {
	ret := &BaseClient{
		options: options,
		id:      fmt.Sprintf("%s@%s:%d", "", "", 1),
	}
	etcd, err := client3.NewEtcdV3Discovery(BaseServicePath, options.Name, options.EtcdServers, false, nil)
	if err != nil {
		return nil, err
	}
	if options.SerializeType == protocol.SerializeNone {
		// using protobuffer serializer by default
		options.SerializeType = protocol.ProtoBuffer
	}
	ret.cli = client2.NewXClient(options.Name, client2.Failtry, client2.RoundRobin, etcd, options.Option)

	if options.Selector != nil {
		ret.cli.SetSelector(options.Selector)
	} else {
		// using round robbin selector by default
		ret.cli.SetSelector(NewServerSelector())
	}
	return ret, nil
}

func (c *BaseClient) Call2(fn string, arg interface{}, reply interface{}) error {
	return c.Call(context.Background(), fn, arg, reply)
}

func (c *BaseClient) Broadcast(fn string, request, reply interface{}) error {
	return c.cli.Broadcast(context.Background(), fn, request, reply)
}

func (c *BaseClient) Call(ctx context.Context, fn string, arg interface{}, reply interface{}) error {
	err := c.cli.Call(ctx, fn, arg, reply)
	return err
}

func (c *BaseClient) Run() error {
	return nil
}

func (c *BaseClient) Close() error {
	return c.cli.Close()
}
