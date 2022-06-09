provider "number" {
}

data "number" "1" {
  provider = "number"
  number = 1
}

output "1_output" {
  value = data.number.1.number
}
