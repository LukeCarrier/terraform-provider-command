resource "command_command" "greeting" {
  name = "C:\\Windows\\System32\\cmd.exe"

  arguments = [
    "/c",
    "echo",
    "Hello, ${var.name}",
  ]
}
