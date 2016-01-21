# go-containership.cloud

## About

### Description
A Go API client for ContainerShip Cloud

### Author
ContainerShip Developers - developers@containership.io

## Usage

### Importing
`import github.com/containership/go-containership.cloud/client`

### Examples

Instantiating a new ContainerShip Cloud Client is as easy as supplying your ContainerShip Cloud Organization and API Key
```go
ContainerShipCloudClient := client.NewContainerShipCloudClient("my-org", "my-api-key")
```

Get an application by supplying the cluster ID and application ID to the `GetApplication` function
```go
response, err := ContainerShipCloudClient.GetApplication("my-cluster-id", "my-application")

if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response.Body)
}
```

Creating an application by supplying cluster ID, application ID, and the application schema as a map
```go
application := map[string]interface{} {
    "id": "application-id",
    "image": "library/image:tag",
    "cpus": 0.25,
    "memory": 512,
}

response, err := ContainerShipCloudClient.CreateApplication("my-cluster-id", "my-cool-application", application)

if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response.StatusCode)
}
```

Similarly, you can update an application by supplying the same parameters to the `UpdateApplication` function
```go
application := map[string]interface{} {
    "cpus": 0.1,
    "memory": 128,
}

response, err := ContainerShipCloudClient.UpdateApplication("my-cluster-id", "my-cool-application", application)

if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response.StatusCode)
}
```

Providing cluster ID and application ID to the `DeleteApplication` function will delete it from your cluster
```go
response, err := ContainerShipCloudClient.DeleteApplication("my-cluster-id", "my-cool-application")

if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(response.StatusCode)
}
```
