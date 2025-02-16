go build -o mux --ldflags="-X 'main.version=$1'"
shasum --algorithm 256 ./mux | cut -d " " -f 1
