# golang-clean-arch-reference

# Generate mocks

docker run -v "$PWD":/src -w /src vektra/mockery --all --inpackage

sudo chown -R $USER:$USER *

# List Customers

curl localhost:8080/customers

# Get Customer

curl localhost:8080/customer/68352972-d1a6-4f8e-8e42-cb83c5a07348

# Post Customer

curl -X POST localhost:8080/customer -d '{"name":"my customer another"}' -H 'Content-Type: application/json'

# Put Customer

curl -X PUT localhost:8080/customer/68352972-d1a6-4f8e-8e42-cb83c5a07348 -d '{"name":"new changing name"}' -H 'Content-Type: application/json'
