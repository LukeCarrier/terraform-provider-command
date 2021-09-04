terraform {
  required_version = "= 1.0.6"

  required_providers {
    command = {
      source = "LukeCarrier/command"
    }
  }
}

provider "command" {}
