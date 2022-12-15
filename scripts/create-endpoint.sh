#!/bin/sh

if [ $# -ne 4 ]
then
    echo "Missing or too many command line arguments."
    echo "Usage: sh create-endpoint.sh [http verb] [endpoint name] [endpoint path] [expected json response]"
    exit 1
fi

# usage:
# sh create-endpoint.sh [http verb] [endpoint name] [expected json response]
# sh create-endpoint.sh "POST" "test" "{}"

# add router line to handleRequests()
function router_handle_func() {
    awk -v verb="$1" -v name="$2" -v path="$3" 'NR==30{print "     router.HandleFunc(\"/" path "\", " name ").Methods(\"" verb "\")"}1' main.go >> main-temp.go
    rm main.go && mv main-temp.go main.go
}

# add function
function create_handler() {
    echo "func $2(w http.ResponseWriter, r *http.Request) { // $1" >> main.go
    echo "     w.Header().Set(\"Content-Type\", \"application/json\")" >> main.go
    echo "     var res interface{}" >> main.go
    echo "     res, err := endpoints.GetResponse(\"$4\")" >> main.go
    echo "     if err != nil {" >> main.go
    echo "          return" >> main.go
    echo "     }" >> main.go
    echo "     json.NewEncoder(w).Encode(res)" >> main.go
    echo "}" >> main.go
}

router_handle_func $1 $2 $3 $4 &&
create_handler $1 $2 $3 $4