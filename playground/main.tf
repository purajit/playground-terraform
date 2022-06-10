data "number" "one" {
  number = 1
}

output "one_output" {
  value = data.number.one.number
}
