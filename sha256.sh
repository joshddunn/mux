curl -sL https://github.com/joshddunn/mux/archive/refs/tags/$1.tar.gz | shasum --algorithm 256 | cut -d " " -f 1
