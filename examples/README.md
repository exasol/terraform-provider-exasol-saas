# Examples

This directory contains examples that are mostly used for documentation, but can also be run/tested manually via the Terraform CLI.

The document generation tool looks for files in the following locations by default. All other *.tf files besides the ones mentioned below are ignored by the documentation tool. This is useful for creating examples that can run and/or ar testable even if some parts are not relevant for the documentation.

* **provider/provider.tf** example file for the provider index page
* **resources/database_resource/resource.tf** example file for creating a database
* **resources/cluster_resource/resource.tf** example file for adding cluster to a database
* **resources/network_allow_list_resource/resource.tf** example file for adding cidr block to allow list
