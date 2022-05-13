resource "exasol_database" "mydb" {
  name              = "dbName"
  region            = "us-east-2"
  size              = "XS"
  autostop          = 15
  main_cluster_name = "Main"
  initial_sql       = ["Initial sql"]
}
