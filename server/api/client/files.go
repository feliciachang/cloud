// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": files Resource Client
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

// CsvFilesPath computes a request path to the csv action of files.
func CsvFilesPath(fileID string) string {
	param0 := fileID

	return fmt.Sprintf("/files/%s/data.csv", param0)
}

// Export file
func (c *Client) CsvFiles(ctx context.Context, path string, dl *bool) (*http.Response, error) {
	req, err := c.NewCsvFilesRequest(ctx, path, dl)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCsvFilesRequest create the request corresponding to the csv action endpoint of the files resource.
func (c *Client) NewCsvFilesRequest(ctx context.Context, path string, dl *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if dl != nil {
		tmp159 := strconv.FormatBool(*dl)
		values.Set("dl", tmp159)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// DeviceInfoFilesPath computes a request path to the device info action of files.
func DeviceInfoFilesPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/devices/%s", param0)
}

// Device info
func (c *Client) DeviceInfoFiles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeviceInfoFilesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeviceInfoFilesRequest create the request corresponding to the device info action endpoint of the files resource.
func (c *Client) NewDeviceInfoFilesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// FileFilesPath computes a request path to the file action of files.
func FileFilesPath(fileID string) string {
	param0 := fileID

	return fmt.Sprintf("/files/%s", param0)
}

// File info
func (c *Client) FileFiles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFileFilesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFileFilesRequest create the request corresponding to the file action endpoint of the files resource.
func (c *Client) NewFileFilesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetDeviceLocationHistoryFilesPath computes a request path to the get device location history action of files.
func GetDeviceLocationHistoryFilesPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/devices/%s/locations", param0)
}

// device location history
func (c *Client) GetDeviceLocationHistoryFiles(ctx context.Context, path string, page *int) (*http.Response, error) {
	req, err := c.NewGetDeviceLocationHistoryFilesRequest(ctx, path, page)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetDeviceLocationHistoryFilesRequest create the request corresponding to the get device location history action endpoint of the files resource.
func (c *Client) NewGetDeviceLocationHistoryFilesRequest(ctx context.Context, path string, page *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if page != nil {
		tmp160 := strconv.Itoa(*page)
		values.Set("page", tmp160)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListDeviceDataFilesFilesPath computes a request path to the list device data files action of files.
func ListDeviceDataFilesFilesPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/devices/%s/files/data", param0)
}

// List device files
func (c *Client) ListDeviceDataFilesFiles(ctx context.Context, path string, page *int) (*http.Response, error) {
	req, err := c.NewListDeviceDataFilesFilesRequest(ctx, path, page)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListDeviceDataFilesFilesRequest create the request corresponding to the list device data files action endpoint of the files resource.
func (c *Client) NewListDeviceDataFilesFilesRequest(ctx context.Context, path string, page *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if page != nil {
		tmp161 := strconv.Itoa(*page)
		values.Set("page", tmp161)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListDeviceLogFilesFilesPath computes a request path to the list device log files action of files.
func ListDeviceLogFilesFilesPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/devices/%s/files/logs", param0)
}

// List device files
func (c *Client) ListDeviceLogFilesFiles(ctx context.Context, path string, page *int) (*http.Response, error) {
	req, err := c.NewListDeviceLogFilesFilesRequest(ctx, path, page)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListDeviceLogFilesFilesRequest create the request corresponding to the list device log files action endpoint of the files resource.
func (c *Client) NewListDeviceLogFilesFilesRequest(ctx context.Context, path string, page *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if page != nil {
		tmp162 := strconv.Itoa(*page)
		values.Set("page", tmp162)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListDevicesFilesPath computes a request path to the list devices action of files.
func ListDevicesFilesPath() string {

	return fmt.Sprintf("/files/devices")
}

// List devices
func (c *Client) ListDevicesFiles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListDevicesFilesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListDevicesFilesRequest create the request corresponding to the list devices action endpoint of the files resource.
func (c *Client) NewListDevicesFilesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RawFilesPath computes a request path to the raw action of files.
func RawFilesPath(fileID string) string {
	param0 := fileID

	return fmt.Sprintf("/files/%s/data.fkpb", param0)
}

// Export file
func (c *Client) RawFiles(ctx context.Context, path string, dl *bool) (*http.Response, error) {
	req, err := c.NewRawFilesRequest(ctx, path, dl)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRawFilesRequest create the request corresponding to the raw action endpoint of the files resource.
func (c *Client) NewRawFilesRequest(ctx context.Context, path string, dl *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if dl != nil {
		tmp163 := strconv.FormatBool(*dl)
		values.Set("dl", tmp163)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// StatusFilesPath computes a request path to the status action of files.
func StatusFilesPath() string {

	return fmt.Sprintf("/files/status")
}

// File backend status
func (c *Client) StatusFiles(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewStatusFilesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewStatusFilesRequest create the request corresponding to the status action endpoint of the files resource.
func (c *Client) NewStatusFilesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateDeviceInfoFilesPath computes a request path to the update device info action of files.
func UpdateDeviceInfoFilesPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/devices/%s", param0)
}

// Device info
func (c *Client) UpdateDeviceInfoFiles(ctx context.Context, path string, payload *UpdateDeviceInfoPayload) (*http.Response, error) {
	req, err := c.NewUpdateDeviceInfoFilesRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateDeviceInfoFilesRequest create the request corresponding to the update device info action endpoint of the files resource.
func (c *Client) NewUpdateDeviceInfoFilesRequest(ctx context.Context, path string, payload *UpdateDeviceInfoPayload) (*http.Request, error) {
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
	return req, nil
}
