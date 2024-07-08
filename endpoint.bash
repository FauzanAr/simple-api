#!/bin/bash

# Define endpoint URLs
endpoint1="http://127.0.0.1:5000/api/v1/login"
endpoint2="http://127.0.0.1:5000/api/v1/profile"
endpoint3="http://127.0.0.1:5000/api/v1/profile"

# Function to test POST endpoint
test_post_endpoint() {
    url=$1

    echo "Testing POST endpoint: $url"

    # Example POST request with JSON body
    curl -X POST \
         -H "Content-Type: application/json" \
         -d '{"username": "fauzan", "password": ""}' \
         $url

    echo -e "\n"
}

# Function to test GET endpoint
test_get_endpoint() {
    url=$1

    echo "Testing GET endpoint: $url"

    # Example GET request
    curl -X GET \
         -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJmYXV6YW4iLCJlbWFpbCI6ImZhdXJpbnRhMTIiLCJzdGF0dXMiOiJzdGF0dXMiLCJleHAiOjE3MjA0NDc3NDAsImlhdCI6MTcyMDQyOTc0MH0.G6iwb4IiDDIicO6lGzduwyFRVN31bjttS9bcCMu6R3I" \
         $url

    echo -e "\n"
}

# Function to test PUT endpoint
test_put_endpoint() {
    url=$1

    echo "Testing PUT endpoint: $url"

    # Example PUT request with JSON body
    curl -X PUT \
         -H "Content-Type: application/json" \
         -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJmYXV6YW4iLCJlbWFpbCI6ImZhdXJpbnRhMTIiLCJzdGF0dXMiOiJzdGF0dXMiLCJleHAiOjE3MjA0NDc3NDAsImlhdCI6MTcyMDQyOTc0MH0.G6iwb4IiDDIicO6lGzduwyFRVN31bjttS9bcCMu6R3I" \
         -d '{"username": "fauzan", "status": "status", "email": "email"}' \
         $url

    echo -e "\n"
}

# Test each endpoint with appropriate function
test_post_endpoint $endpoint1
test_get_endpoint $endpoint2
test_put_endpoint $endpoint3
