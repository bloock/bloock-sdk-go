<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# availability

```go
import "github.com/bloock/bloock-sdk-go/v2/entity/availability"
```

## Index

- [type HostedLoader](#HostedLoader)
  - [func NewHostedLoader\(id string\) HostedLoader](#NewHostedLoader)
  - [func \(e HostedLoader\) ToProto\(\) \*proto.Loader](#HostedLoader.ToProto)
- [type HostedPublisher](#HostedPublisher)
  - [func NewHostedPublisher\(\) HostedPublisher](#NewHostedPublisher)
  - [func \(e HostedPublisher\) ToProto\(\) \*proto.Publisher](#HostedPublisher.ToProto)
- [type IpfsLoader](#IpfsLoader)
  - [func NewIpfsLoader\(hash string\) IpfsLoader](#NewIpfsLoader)
  - [func NewIpnsLoader\(hash string\) IpfsLoader](#NewIpnsLoader)
  - [func \(e IpfsLoader\) ToProto\(\) \*proto.Loader](#IpfsLoader.ToProto)
- [type IpfsPublisher](#IpfsPublisher)
  - [func NewIpfsPublisher\(\) IpfsPublisher](#NewIpfsPublisher)
  - [func \(e IpfsPublisher\) ToProto\(\) \*proto.Publisher](#IpfsPublisher.ToProto)
- [type IpnsKey](#IpnsKey)
  - [func NewIpnsKeyWithManagedKey\(key key.ManagedKey\) IpnsKey](#NewIpnsKeyWithManagedKey)
  - [func \(s IpnsKey\) ToProto\(\) \*proto.IpnsKey](#IpnsKey.ToProto)
- [type IpnsLoader](#IpnsLoader)
  - [func \(e IpnsLoader\) ToProto\(\) \*proto.Loader](#IpnsLoader.ToProto)
- [type IpnsPublisher](#IpnsPublisher)
  - [func NewIpnsPublisher\(ipnsKey IpnsKey\) IpnsPublisher](#NewIpnsPublisher)
  - [func \(e IpnsPublisher\) ToProto\(\) \*proto.Publisher](#IpnsPublisher.ToProto)
- [type Loader](#Loader)
- [type LoaderArgs](#LoaderArgs)
  - [func \(e LoaderArgs\) ToProto\(\) \*proto.LoaderArgs](#LoaderArgs.ToProto)
- [type Publisher](#Publisher)
- [type PublisherArgs](#PublisherArgs)
  - [func \(e PublisherArgs\) ToProto\(\) \*proto.PublisherArgs](#PublisherArgs.ToProto)


###### HostedLoader {#HostedLoader}
## type HostedLoader

HostedLoader represents a loader for hosted data availability.

```go
type HostedLoader struct {
    Type proto.DataAvailabilityType
    Args LoaderArgs
}
```

###### NewHostedLoader {#NewHostedLoader}
### func NewHostedLoader

```go
func NewHostedLoader(id string) HostedLoader
```

NewHostedLoader creates a HostedLoader instance with the provided identifier \(ex: c137fded\-cb04\-4c6e\-9415\-1e7baf48b659\).

###### HostedLoader.ToProto {#HostedLoader.ToProto}
### func \(HostedLoader\) ToProto

```go
func (e HostedLoader) ToProto() *proto.Loader
```



###### HostedPublisher {#HostedPublisher}
## type HostedPublisher

HostedPublisher represents a publisher for hosted data availability.

```go
type HostedPublisher struct {
    Type proto.DataAvailabilityType
    Args PublisherArgs
}
```

###### NewHostedPublisher {#NewHostedPublisher}
### func NewHostedPublisher

```go
func NewHostedPublisher() HostedPublisher
```

NewHostedPublisher creates a HostedPublisher instance with default publisher arguments.

###### HostedPublisher.ToProto {#HostedPublisher.ToProto}
### func \(HostedPublisher\) ToProto

```go
func (e HostedPublisher) ToProto() *proto.Publisher
```



###### IpfsLoader {#IpfsLoader}
## type IpfsLoader

IpfsLoader represents a loader for IPFS data availability.

```go
type IpfsLoader struct {
    Type proto.DataAvailabilityType
    Args LoaderArgs
}
```

###### NewIpfsLoader {#NewIpfsLoader}
### func NewIpfsLoader

```go
func NewIpfsLoader(hash string) IpfsLoader
```

NewIpfsLoader creates an IpfsLoader instance with the provided IPFS hash.

###### NewIpnsLoader {#NewIpnsLoader}
### func NewIpnsLoader

```go
func NewIpnsLoader(hash string) IpfsLoader
```

NewIpnsLoader creates an IpnsLoader instance with the provided IPNS hash.

###### IpfsLoader.ToProto {#IpfsLoader.ToProto}
### func \(IpfsLoader\) ToProto

```go
func (e IpfsLoader) ToProto() *proto.Loader
```



###### IpfsPublisher {#IpfsPublisher}
## type IpfsPublisher

IpfsPublisher represents a publisher for IPFS data availability.

```go
type IpfsPublisher struct {
    Type proto.DataAvailabilityType
    Args PublisherArgs
}
```

###### NewIpfsPublisher {#NewIpfsPublisher}
### func NewIpfsPublisher

```go
func NewIpfsPublisher() IpfsPublisher
```

IpfsPublisher represents a publisher for IPFS data availability.

###### IpfsPublisher.ToProto {#IpfsPublisher.ToProto}
### func \(IpfsPublisher\) ToProto

```go
func (e IpfsPublisher) ToProto() *proto.Publisher
```



###### IpnsKey {#IpnsKey}
## type IpnsKey

IpnsKey represents an object with various key types.

```go
type IpnsKey struct {
    ManagedKey         *key.ManagedKey
    ManagedCertificate *key.ManagedCertificate
}
```

###### NewIpnsKeyWithManagedKey {#NewIpnsKeyWithManagedKey}
### func NewIpnsKeyWithManagedKey

```go
func NewIpnsKeyWithManagedKey(key key.ManagedKey) IpnsKey
```

NewIpnsKeyWithManagedKey creates an IpnsKey instance with a managed key.

###### IpnsKey.ToProto {#IpnsKey.ToProto}
### func \(IpnsKey\) ToProto

```go
func (s IpnsKey) ToProto() *proto.IpnsKey
```



###### IpnsLoader {#IpnsLoader}
## type IpnsLoader

IpnsLoader represents a loader for IPNS data availability.

```go
type IpnsLoader struct {
    Type proto.DataAvailabilityType
    Args LoaderArgs
}
```

###### IpnsLoader.ToProto {#IpnsLoader.ToProto}
### func \(IpnsLoader\) ToProto

```go
func (e IpnsLoader) ToProto() *proto.Loader
```



###### IpnsPublisher {#IpnsPublisher}
## type IpnsPublisher

IpnsPublisher represents a publisher for IPNS data availability.

```go
type IpnsPublisher struct {
    Type proto.DataAvailabilityType
    Args PublisherArgs
}
```

###### NewIpnsPublisher {#NewIpnsPublisher}
### func NewIpnsPublisher

```go
func NewIpnsPublisher(ipnsKey IpnsKey) IpnsPublisher
```

NewIpnsPublisher represents a publisher for IPNS data availability.

###### IpnsPublisher.ToProto {#IpnsPublisher.ToProto}
### func \(IpnsPublisher\) ToProto

```go
func (e IpnsPublisher) ToProto() *proto.Publisher
```



###### Loader {#Loader}
## type Loader



```go
type Loader interface {
    ToProto() *proto.Loader
}
```

###### LoaderArgs {#LoaderArgs}
## type LoaderArgs

LoaderArgs represents the arguments for a data loader.

```go
type LoaderArgs struct {
    // Id is a unique identifier associated with the loader.
    Id string
}
```

###### LoaderArgs.ToProto {#LoaderArgs.ToProto}
### func \(LoaderArgs\) ToProto

```go
func (e LoaderArgs) ToProto() *proto.LoaderArgs
```



###### Publisher {#Publisher}
## type Publisher



```go
type Publisher interface {
    ToProto() *proto.Publisher
}
```

###### PublisherArgs {#PublisherArgs}
## type PublisherArgs

PublisherArgs represents the arguments for a data publisher.

```go
type PublisherArgs struct {
    // IpnsKey is a managed key or certificate object that will be used to create the IPNS record.
    IpnsKey IpnsKey
}
```

###### PublisherArgs.ToProto {#PublisherArgs.ToProto}
### func \(PublisherArgs\) ToProto

```go
func (e PublisherArgs) ToProto() *proto.PublisherArgs
```



Generated by [gomarkdoc](https://github.com/princjef/gomarkdoc)
