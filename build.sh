#!/usr/bin/env bash

export APP_NAME="btc-service"
export CURRENCY='BTCUSD,ETHBTC'
export PORT=3000


go install
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi

echo "Restoring all vendor versions ..."
godep restore
echo "Done."


echo "Doing some cleaning ..."
go clean
echo "Done."

echo "Running goimport ..."
goimports -w=true .
echo "Done."

echo "Running go vet ..."
go vet ./internal/...
if [ $? != 0 ]; then
  exit
fi
echo "Done."

echo "Running go generate ..."
go generate ./internal/...
echo "Done."

echo "Running go format ..."
gofmt -w .
echo "Done."

echo "Running go build ..."
go build -race
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi
echo "Done."

	echo "## Starting service ##"
    ./btc-service -worker="web"
fi
