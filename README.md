# golang-clean-arch-reference

# Generate mocks

docker run -v "$PWD":/src -w /src vektra/mockery --all --inpackage

sudo chown -R $USER:$USER *

# List Customers

curl localhost:8080/customers

# Get Customer

curl localhost:8080/customer/0d48b63d-3222-4cd1-b1e4-ad184a1459de

# Post Customer

curl -X POST localhost:8080/customer -d '{"name":"another"}' -H 'Content-Type: application/json'
