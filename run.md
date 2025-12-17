cd seicho
go mod tidy
go build -o seicho
./seicho


### example usage

./seicho add "Exercise" -d "30 minutes of physical activity" -f daily
./seicho log 1 done
./seicho stats
./seicho zen

go install
export PATH="$PATH:$(go env GOPATH)/bin"
