// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": device Resource Client
//
// Command:
// $ main

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// AddDevicePath computes a request path to the add action of device.
func AddDevicePath(expeditionID int) string {
	param0 := strconv.Itoa(expeditionID)

	return fmt.Sprintf("/expeditions/%s/sources/devices", param0)
}

// Add a Device source
func (c *Client) AddDevice(ctx context.Context, path string, payload *AddDeviceSourcePayload) (*http.Response, error) {
	req, err := c.NewAddDeviceRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddDeviceRequest create the request corresponding to the add action endpoint of the device resource.
func (c *Client) NewAddDeviceRequest(ctx context.Context, path string, payload *AddDeviceSourcePayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// GetIDDevicePath computes a request path to the get id action of device.
func GetIDDevicePath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/sources/devices/%s", param0)
}

// Get a Device source
func (c *Client) GetIDDevice(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetIDDeviceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetIDDeviceRequest create the request corresponding to the get id action endpoint of the device resource.
func (c *Client) NewGetIDDeviceRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ListDevicePath computes a request path to the list action of device.
func ListDevicePath(project string, expedition string) string {
	param0 := project
	param1 := expedition

	return fmt.Sprintf("/projects/@/%s/expeditions/@/%s/sources/devices", param0, param1)
}

// List an expedition's Device sources
func (c *Client) ListDevice(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListDeviceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListDeviceRequest create the request corresponding to the list action endpoint of the device resource.
func (c *Client) NewListDeviceRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// UpdateDevicePath computes a request path to the update action of device.
func UpdateDevicePath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/sources/devices/%s", param0)
}

// Update an Device source
func (c *Client) UpdateDevice(ctx context.Context, path string, payload *UpdateDeviceSourcePayload) (*http.Response, error) {
	req, err := c.NewUpdateDeviceRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateDeviceRequest create the request corresponding to the update action endpoint of the device resource.
func (c *Client) NewUpdateDeviceRequest(ctx context.Context, path string, payload *UpdateDeviceSourcePayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// UpdateLocationDevicePath computes a request path to the update location action of device.
func UpdateLocationDevicePath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/sources/devices/%s/location", param0)
}

// Update an Device source location
func (c *Client) UpdateLocationDevice(ctx context.Context, path string, payload *UpdateDeviceSourceLocationPayload) (*http.Response, error) {
	req, err := c.NewUpdateLocationDeviceRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateLocationDeviceRequest create the request corresponding to the update location action endpoint of the device resource.
func (c *Client) NewUpdateLocationDeviceRequest(ctx context.Context, path string, payload *UpdateDeviceSourceLocationPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// UpdateSchemaDevicePath computes a request path to the update schema action of device.
func UpdateSchemaDevicePath(id int) string {
	param0 := strconv.Itoa(id)

	return fmt.Sprintf("/sources/devices/%s/schemas", param0)
}

// Update an Device source schema
func (c *Client) UpdateSchemaDevice(ctx context.Context, path string, payload *UpdateDeviceSourceSchemaPayload) (*http.Response, error) {
	req, err := c.NewUpdateSchemaDeviceRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateSchemaDeviceRequest create the request corresponding to the update schema action endpoint of the device resource.
func (c *Client) NewUpdateSchemaDeviceRequest(ctx context.Context, path string, payload *UpdateDeviceSourceSchemaPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
