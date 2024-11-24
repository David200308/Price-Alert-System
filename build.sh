#!/bin/bash

echo "Building Backend..."
cd Backend || exit 1
go build .
if [ $? -eq 0 ]; then
  echo "Backend built successfully."
else
  echo "Error building Backend."
  exit 1
fi
cd .. || exit 1

echo "Building Scheduler..."
cd Scheduler || exit 1
go build .
if [ $? -eq 0 ]; then
  echo "Scheduler built successfully."
else
  echo "Error building Scheduler."
  exit 1
fi

echo "Building daily_status_update/daily_update.go..."
go build ./daily_status_update/daily_update.go
if [ $? -eq 0 ]; then
  echo "daily_update built successfully."
else
  echo "Error building daily_update."
  exit 1
fi
cd .. || exit 1

echo "Built successfully!"
