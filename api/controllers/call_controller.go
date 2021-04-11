package controllers

import (
   "io/ioutil"
   "log"
   "net/http"

   "github.com/gin-gonic/gin"
)

func CreateCall(c *gin.Context) {

   url := "http://localhost:8080/api/campaign/"

   // Create a Bearer string by appending string access token
   var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJsdW1lbi1qd3QiLCJzdWIiOjIsImlhdCI6MTYxODEyMTk5NCwiZXhwIjoxNjE4MTI1NTk0fQ.dJ8_d1OvAMN5VDURwOFWyNTroPOt3SI2OMkYMHBb5UU"
   log.Printf(bearer)

   // Create a new request using http
   req, err := http.NewRequest("GET", url, nil)
   // resp, err := http.Get("http://localhost:8080/api/campaign/")
   req.Header.Add("Authorization", bearer)
   req.Header.Add("Accept", "application/json")

   // if err != nil {
   //    log.Fatalln(err)
   // }

   client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        log.Println("Error on response.\n[ERROR] -", err)
    }
    defer resp.Body.Close()

//We Read the response body on the line below.
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
//Convert the body to type string
   sb := string(body)
   // log.Printf(sb)

   // c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": sb})
   c.JSON(http.StatusOK, sb)
	// responses.JSON(w, http.StatusCreated, postCreated)

}

// func store(c *gin.Context) {
// 	var todos []todoModel
// 	var _todos []transformedTodo

// 	db.Find(&todos)

// 	if len(todos) <= 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}

// 	//transforms the todos for building a good response
// 	for _, item := range todos {
// 		completed := false
// 		if item.Completed == 1 {
// 			completed = true
// 		} else {
// 			completed = false
// 		}
// 		_todos = append(_todos, transformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
// 	}
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
// }