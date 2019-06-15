// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": Tasks Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CheckTasksPath computes a request path to the check action of Tasks.
func CheckTasksPath() string {

	return fmt.Sprintf("/tasks/check")
}

// Run periodic checks
func (c *Client) CheckTasks(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCheckTasksRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCheckTasksRequest create the request corresponding to the check action endpoint of the Tasks resource.
func (c *Client) NewCheckTasksRequest(ctx context.Context, path string) (*http.Request, error) {
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

// FiveTasksPath computes a request path to the five action of Tasks.
func FiveTasksPath() string {

	return fmt.Sprintf("/tasks/five")
}

// Run periodic checks
func (c *Client) FiveTasks(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFiveTasksRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFiveTasksRequest create the request corresponding to the five action endpoint of the Tasks resource.
func (c *Client) NewFiveTasksRequest(ctx context.Context, path string) (*http.Request, error) {
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

// RefreshTasksPath computes a request path to the refresh action of Tasks.
func RefreshTasksPath(deviceID string, fileTypeID string) string {
	param0 := deviceID
	param1 := fileTypeID

	return fmt.Sprintf("/tasks/refresh/%s/%s", param0, param1)
}

// Refresh a device by ID
func (c *Client) RefreshTasks(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRefreshTasksRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRefreshTasksRequest create the request corresponding to the refresh action endpoint of the Tasks resource.
func (c *Client) NewRefreshTasksRequest(ctx context.Context, path string) (*http.Request, error) {
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
