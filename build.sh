go build -o mux/mux --ldflags="-X 'main.version=$1'"
cp -r completion mux
tar -zcvf mux.tar.gz mux/*
shasum --algorithm 256 mux.tar.gz | cut -d " " -f 1
