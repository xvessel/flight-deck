
curl http://127.0.0.1:8080/components
echo ""

curl -XPOST http://127.0.0.1:8080/products -d \
    '{"Name":"test-product"}'
echo ""

curl http://127.0.0.1:8080/products
echo ""

curl -w %{http_code}  http://127.0.0.1:8080/designs -d \
'{
"ProductName":"test-product",
"Revision":"test-revision",
"ComponentRefers":[
{
    "Role":"test-role",
    "ComponentName":"example",
    "Input":{"CPU":"1","MEM":"1"},
    "PreRole":[]
}
]
}'
echo ""

curl  http://127.0.0.1:8080/products/test-product/designs 
echo ""


curl -w %{http_code} http://127.0.0.1:8080/products/test-product/instances -d \
'
{
    "Name":"test-instance",
    "ProductName":"test-product",
    "DesignRevision":"test-revision",
    "KubeClusterName":"test-cluster"
}
'
echo ""

curl http://127.0.0.1:8080/products/test-product/instances
