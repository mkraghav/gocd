package main


import ("net/http"
"encoding/json"


)

func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/ping",func (w http.ResponseWriter, r *http.Request)  {
		if r.Method==http.MethodGet{
			data:= structs.Response{
				code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}

		// fmt.Println("Request Received")
		// w.Write([]byte("Hello World"))
		// fmt.Println(r.Method)
		// w.Write([]byte("Get going"))
		
	})
	// mux.HandleFunc("/getgoing/go",func (w http.ResponseWriter, r *http.Request)  {
	// 	fmt.Println("Request Received")
	// 	w.Write([]byte("Hello World go"))
	// 	fmt.Println(r.Method)
	
		
	// })
	http.ListenAndServe("localhost:3000",mux)
}