// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	context "context"

	archive "github.com/docker/docker/pkg/archive"

	io "io"

	mock "github.com/stretchr/testify/mock"

	types "github.com/docker/docker/api/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// ArchiveDirectory provides a mock function with given fields: srcPath, options
func (_m *Client) ArchiveDirectory(srcPath string, options *archive.TarOptions) (io.ReadCloser, error) {
	ret := _m.Called(srcPath, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string, *archive.TarOptions) io.ReadCloser); ok {
		r0 = rf(srcPath, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *archive.TarOptions) error); ok {
		r1 = rf(srcPath, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageBuild provides a mock function with given fields: ctx, buildContext, options
func (_m *Client) ImageBuild(ctx context.Context, buildContext io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	ret := _m.Called(ctx, buildContext, options)

	var r0 types.ImageBuildResponse
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, types.ImageBuildOptions) types.ImageBuildResponse); ok {
		r0 = rf(ctx, buildContext, options)
	} else {
		r0 = ret.Get(0).(types.ImageBuildResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, io.Reader, types.ImageBuildOptions) error); ok {
		r1 = rf(ctx, buildContext, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImagePush provides a mock function with given fields: ctx, image, options
func (_m *Client) ImagePush(ctx context.Context, image string, options types.ImagePushOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, image, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ImagePushOptions) io.ReadCloser); ok {
		r0 = rf(ctx, image, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ImagePushOptions) error); ok {
		r1 = rf(ctx, image, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NegotiateAPIVersion provides a mock function with given fields: ctx
func (_m *Client) NegotiateAPIVersion(ctx context.Context) {
	_m.Called(ctx)
}
