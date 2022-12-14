# golang-clean-arch-reference

# Generate mocks

docker run -v "$PWD":/src -w /src vektra/mockery --all --inpackage

sudo chown -R $USER:$USER *

# List Customers

curl localhost:8080/customers

# Get Customer

curl localhost:8080/customer/512f08c4-8627-4a5d-b012-d803dc6650f0

# Post Customer

curl -X POST localhost:8080/customer -d '{"name":"another"}' -H 'Content-Type: application/json'
