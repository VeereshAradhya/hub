// Code generated by goa v3.4.0, DO NOT EDIT.
//
// resource client
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package resource

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "resource" service client.
type Client struct {
	ListEndpoint         goa.Endpoint
	VersionsByIDEndpoint goa.Endpoint
}

// NewClient initializes a "resource" service client given the endpoints.
func NewClient(list, versionsByID goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:         list,
		VersionsByIDEndpoint: versionsByID,
	}
}

// List calls the "List" endpoint of the "resource" service.
func (c *Client) List(ctx context.Context) (res *Resources, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*Resources), nil
}

// VersionsByID calls the "VersionsByID" endpoint of the "resource" service.
func (c *Client) VersionsByID(ctx context.Context, p *VersionsByIDPayload) (res *ResourceVersions, err error) {
	var ires interface{}
	ires, err = c.VersionsByIDEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ResourceVersions), nil
}
