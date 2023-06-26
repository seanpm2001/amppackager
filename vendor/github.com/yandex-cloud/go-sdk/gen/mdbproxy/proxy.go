// Code generated by sdkgen. DO NOT EDIT.

//nolint
package mdbproxy

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	mdbproxy "github.com/yandex-cloud/go-genproto/yandex/cloud/serverless/mdbproxy/v1"
)

//revive:disable

// ProxyServiceClient is a mdbproxy.ProxyServiceClient with
// lazy GRPC connection initialization.
type ProxyServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements mdbproxy.ProxyServiceClient
func (c *ProxyServiceClient) Create(ctx context.Context, in *mdbproxy.CreateProxyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mdbproxy.NewProxyServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements mdbproxy.ProxyServiceClient
func (c *ProxyServiceClient) Delete(ctx context.Context, in *mdbproxy.DeleteProxyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mdbproxy.NewProxyServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements mdbproxy.ProxyServiceClient
func (c *ProxyServiceClient) Get(ctx context.Context, in *mdbproxy.GetProxyRequest, opts ...grpc.CallOption) (*mdbproxy.Proxy, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mdbproxy.NewProxyServiceClient(conn).Get(ctx, in, opts...)
}

// List implements mdbproxy.ProxyServiceClient
func (c *ProxyServiceClient) List(ctx context.Context, in *mdbproxy.ListProxyRequest, opts ...grpc.CallOption) (*mdbproxy.ListProxyResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mdbproxy.NewProxyServiceClient(conn).List(ctx, in, opts...)
}

type ProxyIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ProxyServiceClient
	request *mdbproxy.ListProxyRequest

	items []*mdbproxy.Proxy
}

func (c *ProxyServiceClient) ProxyIterator(ctx context.Context, req *mdbproxy.ListProxyRequest, opts ...grpc.CallOption) *ProxyIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ProxyIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ProxyIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Proxies
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ProxyIterator) Take(size int64) ([]*mdbproxy.Proxy, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*mdbproxy.Proxy

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ProxyIterator) TakeAll() ([]*mdbproxy.Proxy, error) {
	return it.Take(0)
}

func (it *ProxyIterator) Value() *mdbproxy.Proxy {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ProxyIterator) Error() error {
	return it.err
}

// ListAccessBindings implements mdbproxy.ProxyServiceClient
func (c *ProxyServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mdbproxy.NewProxyServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type ProxyAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ProxyServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *ProxyServiceClient) ProxyAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *ProxyAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ProxyAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ProxyAccessBindingsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListAccessBindings(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.AccessBindings
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ProxyAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*access.AccessBinding

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ProxyAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *ProxyAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ProxyAccessBindingsIterator) Error() error {
	return it.err
}

// ListOperations implements mdbproxy.ProxyServiceClient
func (c *ProxyServiceClient) ListOperations(ctx context.Context, in *mdbproxy.ListProxyOperationsRequest, opts ...grpc.CallOption) (*mdbproxy.ListProxyOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mdbproxy.NewProxyServiceClient(conn).ListOperations(ctx, in, opts...)
}

type ProxyOperationsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *ProxyServiceClient
	request *mdbproxy.ListProxyOperationsRequest

	items []*operation.Operation
}

func (c *ProxyServiceClient) ProxyOperationsIterator(ctx context.Context, req *mdbproxy.ListProxyOperationsRequest, opts ...grpc.CallOption) *ProxyOperationsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &ProxyOperationsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *ProxyOperationsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListOperations(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Operations
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ProxyOperationsIterator) Take(size int64) ([]*operation.Operation, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*operation.Operation

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *ProxyOperationsIterator) TakeAll() ([]*operation.Operation, error) {
	return it.Take(0)
}

func (it *ProxyOperationsIterator) Value() *operation.Operation {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ProxyOperationsIterator) Error() error {
	return it.err
}

// SetAccessBindings implements mdbproxy.ProxyServiceClient
func (c *ProxyServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mdbproxy.NewProxyServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Update implements mdbproxy.ProxyServiceClient
func (c *ProxyServiceClient) Update(ctx context.Context, in *mdbproxy.UpdateProxyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mdbproxy.NewProxyServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements mdbproxy.ProxyServiceClient
func (c *ProxyServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return mdbproxy.NewProxyServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}