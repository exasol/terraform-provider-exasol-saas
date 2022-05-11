resource "exasol_cluster" "workercluster01" {
  database = exasol_database.mydb.id // Reference to database
  size     = "XS"
  name     = "WorkerClusterName"
}
