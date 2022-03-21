#!/bin/sh

if [ $# -ne 3 ]
then
    echo "Missing or too many command line arguments."
    echo "Usage: sh create-endpoint.sh [http verb] [endpoint name] [expected json response]"
    exit 1
fi

# usage:
# sh create-endpoint.sh [http verb] [endpoint name] [expected json response]
# sh create-endpoint.sh "POST" "test" "{}"

# add router line to handleRequests()
function router_handle_func() {
    VERB="$1"
    NAME="$2"
    awk -v verb="$VERB" -v name="$NAME" 'NR==65{print "\trouter.HandleFunc(\"/" name "\", " name ").Methods(\"" verb "\")"}1' main.go >> main-temp.go
    rm main.go && mv main-temp.go main.go
}

# add function
function create_handler() {
    echo "func $2(w http.ResponseWriter, r *http.Request) { // $1" >> main.go
    echo "\tw.Header().Set(\"Content-Type\", \"application/json\")" >> main.go
    echo "\tvar res interface{}" >> main.go
    echo "\tresJson := \`$3\`" >> main.go
    echo "\tjson.Unmarshal([]byte(resJson), &res)" >> main.go
    echo "\tjson.NewEncoder(w).Encode(res)" >> main.go
    echo "}" >> main.go
}

router_handle_func $1 $2 &&
create_handler $1 $2 $3