package tenants

import "github.com/rackspace/rack/internal/github.com/rackspace/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("tenants")
}
