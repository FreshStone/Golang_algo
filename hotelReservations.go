package main

import (
		"log"
		"fmt"
		"net/http"
		"os"
		"os/signal"
		"strconv"
)

var hotel [][][]int

func init(){
	hotel = make([][][]int, 3)
}

/*
api's usage-
getroom - http://localhost:8080/getroom?in=in_time&out=out_time
bookroom- http://localhost:8080/bookroom?in=in_time&out=out_time
*/

func main(){
	close := make(chan os.Signal, 1)
	signal.Notify(close, os.Interrupt)
	go func(){
		<-close
		log.Println("shutting down server on 8080")
		os.Exit(1)
	}()
	http.HandleFunc("/getroom", get_a_room)
	http.HandleFunc("/bookroom",book_a_room)
	log.Println("Listening on port 8080....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/* form parsing
				r.ParseForm()  // parse arguments, you have to call this by yourself
    fmt.Println(r.Form)  // print form information in server side
    //form is a map
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }

/* logging into a file
f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
if err != nil {
    log.Fatalf("error opening file: %v", err)
}
defer f.Close()

log.SetOutput(f)
log.Println("This is a test log entry")
*/

func get_a_room(w http.ResponseWriter, r *http.Request){
	log.Println("searching room")
	r.ParseForm()
	//log.Println(r.Form)
	if len(r.Form) != 2{
		log.Println("Bad room search request")
  fmt.Fprintf(w, "Bad request")
  return
	}
	var s, e int
	s, _ = strconv.Atoi(r.Form["in"][0])
	e, _ = strconv.Atoi(r.Form["out"][0])
	if s >= e{
				log.Println("Bad room search request")
    fmt.Fprintf(w, "Bad request...in_time must be less than out_time\n")
    return
 }

 room, _ := getroomUtil(s,e)
 if room == -1{
 	log.Println("Room not available")
 	fmt.Fprintf(w, "Room not available for in_time = %d and out_time = %d", s,e)
 	return
 }
 log.Println("Room number", room+1, "available")
	fmt.Fprintf(w, "Hello, there..Room number %d available for in_time = %d and out_time = %d \n", room+1, s,e)
}

func book_a_room(w http.ResponseWriter, r *http.Request){
	log.Println("reservation request")
	r.ParseForm()
	//log.Println(r.Form)
	if len(r.Form) != 2{
		log.Println("Bad room search request")
  fmt.Fprintf(w, "Bad request")
  return
	}
	var s, e int
	s, _ = strconv.Atoi(r.Form["in"][0])
	e, _ = strconv.Atoi(r.Form["out"][0])
	if s >= e{
				log.Println("Bad equest")
    fmt.Fprintf(w, "Bad request...in_time must be less than out_time\n")
    return
  }

  room, index := getroomUtil(s,e)
  if room == -1{
  	 log.Println("Room not available")
    fmt.Fprintf(w, "Room not available for in_time = %d and out_time = %d\n", s, e)
    return
  }
  //insert {s,e} in hotel[room] (insertion sort)
  hotel[room] = append(hotel[room], []int{s,e})
  for i := len(hotel[room])-1; i > index; i--{
    hotel[room][i][0], hotel[room][i-1][0] = hotel[room][i-1][0], hotel[room][i][0]
    hotel[room][i][1], hotel[room][i-1][1] = hotel[room][i-1][1], hotel[room][i][1]
  }
 log.Println("Room number", room+1, "booked")
	fmt.Fprintf(w, "Hello, there..Booked room %d for in_time = %d and out_time = %d \n", room+1, s, e)
}

func getroomUtil(s, e int)(int, int){
  var room, mid, low, high int
  var skip_room bool
  for ; room < len(hotel); room++{
    low, high = 0, len(hotel[room])-1
    skip_room = false
    for low <= high{
      mid = (low+high)/2
      if e < hotel[room][mid][1]{
        if e <= hotel[room][mid][0]{
          high = mid - 1
        }else{
          skip_room = true
          break
        }
      }else if e > hotel[room][mid][1] {
        if s < hotel[room][mid][1]{
          skip_room = true
          break
        }else{
          low = mid + 1
        }
      }else{
        skip_room = true
        break
      }
    }

    if !skip_room{
      return room, low 
    }
  }
  return -1, -1
}