# golang-clean-arch-reference

# Generate mocks

docker run -v "$PWD":/src -w /src vektra/mockery --all --inpackage

sudo chown -R $USER:$USER *

# List Customers

curl localhost:8080/customers

# Get Customer

curl localhost:8080/customer/1cacdb20-3a42-4d25-9b59-fe25f97310ab

# Post Customer

curl -X POST localhost:8080/customer -d '{"name":"my customer another"}' -H 'Content-Type: application/json'

# Put Customer

curl -X PUT localhost:8080/customer/1cacdb20-3a42-4d25-9b59-fe25f97310ab -d '{"name":"new changing name"}' -H 'Content-Type: application/json'

# Delete Customer

curl -X DELETE localhost:8080/customer/32efd101-bdba-442e-8f66-c873d7729bf2
