# Protobuf plugin for appevents

This `protoc` plugin is meant to be used together with https://github.com/jobilla/go-app-events.

It enables the use of a `jobilla.appevents.event_name` option on messages, which
will generate a `StringType` method for your message's generated Go struct,
making it compliant with the `Message` interface from the library.

This library only exists due to shortcomings in the design of the original
library, and is primarily used internally at Jobilla. We will likely deprecate
and archive this tool during 2023 as we migrate over to [CloudEvents](https://cloudevents.io).

## Installation

We recommend using [Buf](https://buf.build) to generate your protobuf code.
To use the tool, you will need the `jobilla/appevents/options.proto` from this library.
This is published on Buf, and you can import it from your `buf.yaml`:

```yaml
version: v1
deps:
  - buf.build/jobilla/appevents
```

If you're not using `buf`, you'll have to copy the proto file to your codebase. 

You'll also need the `protoc-gen-go-appevents` binary. This is published with every
release on this repository, and also in the `jobilla/protoc-gen-go-appevents` Docker
image. We encourage building your own Buf image and importing the binary:

```dockerfile
FROM bufbuild/buf

COPY --from=jobilla/protoc-gen-go-appevents /protoc-gen-go-appevents /usr/local/bin/protoc-gen-go-appevents
```

## Usage

You can declare the compatibility name of your message using the `jobilla.appevents.event_name` option:

```protobuf
syntax="proto3";

import "jobilla/appevents/options.proto";

message User {
  option (jobilla.appevents.event_name) = "user";

  string email = 1;
  string name = 2;
}
```
