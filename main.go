package main

import (
  "github.com/mitchellh/packer/packer/plugin"

  "github.com/jgkim/packer-post-processor-docker-dockerfile/dockerfile"
)

func main() {
  server, err := plugin.Server()
  if err != nil {
    panic(err)
  }
  server.RegisterPostProcessor(new(dockerfile.PostProcessor))
  server.Serve()
}
