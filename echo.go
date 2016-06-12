package main

import (

    "net/http"
)

func getServerPort() (string) {
    return "7893";
}

func EchoHandler(writer http.ResponseWriter, request *http.Request) {

    writer.Header().Set("Access-Control-Allow-Origin", "*")

    // allow pre-flight headers
    writer.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")


    request.Write(writer)
}

func main() {
    http.HandleFunc("/", EchoHandler)
    http.ListenAndServe(":" + getServerPort(), nil)
}
