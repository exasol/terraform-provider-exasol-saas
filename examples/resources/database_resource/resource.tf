resource "exasol_database" "mydb" {
  name   = "dbName"
  region = "us-east-2"
  size   = "XS"
}
