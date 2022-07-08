# check go test
if ! command -v gotest &> /dev/null
then
    go get -u github.com/rakyll/gotest
fi 
echo "=== go-test OK ==="

# check mockgen
if ! command -v mockgen &> /dev/null
then
    go install github.com/golang/mock/mockgen
fi
echo "=== mockgen OK ==="

# check dependency setup
make download

echo "=== Initialization Completed ==="