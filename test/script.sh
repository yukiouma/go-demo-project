echo "creating customer..."
curl localhost:8081/api/customer/v1/register -X POST -d '{"name": "yuki"}' --header "Content-Type: application/json"
echo
echo "creating book..."
curl localhost:8080/api/book/v1/new -X POST -d '{"name": "golang"}' --header "Content-Type: application/json"
echo
echo "finding book..."
curl localhost:8080/api/book/v1/find/1
echo
echo "finding customer..."
curl localhost:8081/api/customer/v1/find/1
echo
echo "saling book..."
curl localhost:8080/api/book/v1/sale -X POST -d '{"id": 1, "customerId": 1}' --header "Content-Type: application/json"
echo
echo "ending test..."

