resource "command_command" "greeting" {
  name = "/bin/echo"

  arguments = [
    "-n",
    "Hello,",
    var.name,
  ]
}
