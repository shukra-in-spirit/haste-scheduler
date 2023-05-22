#!/bin/bash

# Function to run the curl command
run_curl() {
    local name="$1"
    curl --location 'http://localhost:8080/schedule' \
        --header 'Content-Type: text/plain' \
        --data '{
            "message": "Hello",
            "url": "test",
            "period": 5,
            "name": "'"$name"'"
        }'
}

# Loop through 10 items
for ((i=1; i<=25; i++)); do
    # Generate the name for each run
    name="test_$i"
    
    # Run the curl command in the background
    run_curl "$name" &
done

# Wait for all background processes to finish
wait