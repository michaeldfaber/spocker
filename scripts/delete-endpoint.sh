#!/bin/sh

if [ $# -ne 2 ]
then
    echo "Missing or too many command line arguments."
    echo "Usage: sh delete-endpoint.sh [http verb] [endpoint name]"
    exit 1
fi

# usage:
# sh delete-endpoint.sh [http verb] [endpoint name]
# sh delete-endpoint.sh "POST" "test"

# delete router line from handleRequests()
function delete_router_handle_func() {
    router_handle=$(grep -n "router.HandleFunc(\"/$2\", $2).Methods(\"$1\")" main.go)
    if [ -z "$router_handle" ]
    then
        return
    fi
    
    $(grep -v "router.HandleFunc(\"/$2\", $2).Methods(\"$1\")" main.go >> main-temp.go)
    rm main.go && mv main-temp.go main.go
}

# delete function
function delete_handler() {
    handler=$(grep -n "func $2(w http.ResponseWriter, r \*http.Request) { // $1" main.go)
    if [ -z "$handler" ]
    then
        return
    fi

    handler_line_start=$(echo $handler | cut -d':' -f 1)
    handler_line_end=$(($handler_line_start+6))

    sed "$handler_line_start"','"$handler_line_end"'d' main.go >> main-temp.go
    rm main.go && mv main-temp.go main.go
}

# update json
function update_json() {
    endpoint=$(grep -n "{ \"httpVerb\": \"$1\", \"name\": \"$2\", \"path\": \"$2\"," endpoints.json)
    if [ -z "$endpoint" ]
    then
        return
    fi

    endpoint_line=$(echo $endpoint | cut -d':' -f 1)

    sed "$endpoint_line"'d' endpoints.json >> endpoints-temp.json
    rm endpoints.json && mv endpoints-temp.json endpoints.json
}

delete_router_handle_func $1 $2 &&
delete_handler $1 $2 &&
update_json $1 $2