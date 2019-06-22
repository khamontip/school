package main

import (
 "github.com/gin-gonic/gin"
 "database/sql"
 "net/http"
  "os"
  _"github.com/lib/pq"
)


type Todo struct {
	ID int
	Title string
	Status string
}
func getTodosHandler(c *gin.Context){
	db,_ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	stmt , _ := db.Prepare("SELECT id, title ,status FROM todos")
	rows , _ := stmt.Query()
	todos := []Todo{}
	for rows.Next(){
		t := Todo{}
		err := rows.Scan(&t.ID , &t.Title ,&t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError , gin.H{"error": err.Error()})
			return
		}
		todos = append(todos,t)
	}

	c.JSON(200,"Okay")
}
func main() {
	r := gin.Default()
	r.GET("/api/todos", getTodosHandler)

	r.Run(":1234")
}


// var students = map[int]Student{}

// func pingHandler(c *gin.Context) {
// 	response := gin.H{
// 		"message": "This is get",
// 	}
// 	c.JSON(http.StatusOK, response)

// }

// func pingPostHandler(c *gin.Context) {
// 	response := gin.H{
// 		"message": "This is post",
// 	}
// 	c.JSON(http.StatusOK, response)
// }
// func pingDelHandler(c *gin.Context) {
// 	response := gin.H{
// 		"message": "success",
// 	}
// 	c.JSON(http.StatusOK, response)
// }
// func getOneStudentHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	myid, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, nil)
// 		return
// 	}
// 	c.JSON(http.StatusOK, students[myid])

// }
// func getListStudentHandler(c *gin.Context) {
//      listS := []Student{}
//  	for _, list := range students{
// 		listS = append(listS, list)
// 	}
// 	  	c.JSON(200 , listS)
// }
// func delHandler(c *gin.Context) {
// 	id := c.Param("id")
// //	fmt.Println("id=", id)
// 	var tempstudents = map[int]Student{}
// 	for i := 0; i <= len(students); i++ {
// 		curStudent := students[i]
// 		if strconv.Itoa(curStudent.ID) == id {
// 			//del
// 		} else {
// 			tempstudents[curStudent.ID] = curStudent
// 		}
// 	}
// 	students = tempstudents
// 	fmt.Printf("2 After delete % #v/n", students)
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": "success",
// 	})
// }
// func putStudentHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	myid, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest , err)
// 		return
// 	}
// 	 newStudent := students[myid]
// 	 if err := c.ShouldBindJSON(&newStudent); err != nil{
// 	  	c.JSON(http.StatusBadRequest, err)
// 		return
// 		 }
//         students[myid] = newStudent	
// 	 	c.JSON(200,newStudent)
// }
// func postStudentHandler(c *gin.Context) {
// 	s := Student{}
// 	fmt.Println("brfore bind % #v\n", s)
// 	if err := c.ShouldBindJSON(&s); err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	id := len(students)
// 	id++
// 	s.ID = id
// 	students[id] = s
// 	c.JSON(201, s)
// 	// fmt.Printf("after bind % #v\n , s")
// }