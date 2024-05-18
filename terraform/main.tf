provider "docker" {}

resource "docker_image" "myapp" {
  name = "myapp"
  build {
    context = "."
  }
}

resource "docker_container" "myapp" {
  image = docker_image.myapp.latest
  name  = "myapp"
  ports {
    internal = 8080
    external = 8080
  }
}
