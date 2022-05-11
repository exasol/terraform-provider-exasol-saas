resource "exasol_network_allow_list" "ip1" {
  name       = "companyNetwork"
  cidr_block = "127.0.0.1/32"
}