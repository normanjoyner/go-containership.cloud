package client

type ContainerShipCloudClient struct {
	Opts *ClientOpts
}

func NewContainerShipCloudClient(organization string, apiKey string) *ContainerShipCloudClient {
	client := &ContainerShipCloudClient{}
	clientOpts := &ClientOpts{
		Organization: organization,
		ApiKey:       apiKey,
	}
	client.Opts = clientOpts

	return client
}
