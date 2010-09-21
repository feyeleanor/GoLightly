cd test
make
make install

cd ../storage
make
make install

cd ../go
gotest -benchmarks="Benchmark"

cd ../vm
gotest -benchmarks="Benchmark"