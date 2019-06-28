package todo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"school/database"
)

type TodoHandler struct{}
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func (h *TodoHandler) GetListTodosHandler(c *gin.Context) {
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := database.ConnDB()
	if err != nil {
		log.Fatal("Open error: ", err.Error)
		fmt.Println("1 get list error ")
	}

	defer db.Close()
	fmt.Println("2.1 get list error ")
	//stmt ; err := db.Prepare("SELECT id, title ,status FROM todos")
	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("Prepare SQL error", err.Error)
	}

	rows, err := stmt.Query()
	fmt.Println("3 get list error ")
	if err != nil {
		log.Fatal("SQL error", err.Error)
	}

	todos := []Todo{}
	fmt.Println("4 get list error ")
	for rows.Next() {
		t := Todo{}
		fmt.Println("5 get list error ")
		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, t)
	}
	c.JSON(200, todos)
}

func (h *TodoHandler) GetTodosByIdHandler(c *gin.Context) {
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := database.ConnDB()
	if err != nil {
		log.Fatal(err.Error)
		fmt.Println("2 get error ")
	}
	defer db.Close()
	//stmt, _ := db.Prepare("SELECT id, title ,status from todos where id = $1")
	query := `SELECT id,title,status FROM todos WHERE id = $1;`
	if err != nil {
		log.Fatal("Sql error")
	}
	id := c.Param("id")
	fmt.Println("get --Pid-", id)
	row := db.QueryRow(query, id)
	t := Todo{}
	fmt.Println("List b4 scan ", (t))
	if err := row.Scan(&t.ID, &t.Title, &t.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errr": err.Error()})
		return
	}
	fmt.Println("After  scan ", (t))
	c.JSON(200, t)
}

func (h *TodoHandler) PostTodosHandler(c *gin.Context) {
	t := Todo{}
	fmt.Println("1) post ", (t))
	if err := c.ShouldBindJSON(&t); err != nil {
		//	c.JSON(http.StatusBadRequest, err.Error)
		c.JSON(http.StatusBadRequest, gin.H{"Error post ": err.Error()})
		return
	}
	fmt.Println("2) post ", (t))

	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := database.ConnDB()
	if err != nil {
		log.Fatal("Open error", err.Error)
	}
	defer db.Close()

	//stmt, err := db.Prepare("INSERT into todos ( Title ,Status) VALUES ($1 ,$2) RETURNING id")

	query := `INSERT INTO todos (title,status) VALUES ($1,$2) RETURNING id;`
	var id int
	row := db.QueryRow(query, t.Title, t.Status)
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("scan id fail ", id)
		return
	}

	t.ID = id
	fmt.Println("4) post ", (t))
	c.JSON(201, t)
}

func (h *TodoHandler) PutUpdateTodoHandler(c *gin.Context) {
	fmt.Println("1 put error ")
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := database.ConnDB()
	if err != nil {
		log.Fatal("Open Error ", err.Error)
		fmt.Println("4 put error ")
	}
	defer db.Close()
	//query := `UPDATE todos SET title = $2, status = S3 WHERE id = $1 `
	//row := db.QueryRow(query, Pid, t.Title, t.Status)

	stmt, err := db.Prepare("UPDATE todos SET title=$2, status=$3 WHERE id=$1;")
	if err != nil {
		log.Fatal("SQL eror", err.Error)
	}
	fmt.Println("1. put ")
	Pid := c.Param("id")
	fmt.Println("2. put Pid", Pid)
	t := Todo{}
	fmt.Println("3. put t", t)
	if err := (c.ShouldBindJSON(&t)); err != nil {
		c.JSON(http.StatusBadRequest, err.Error)
		return
	}

	t.ID, err = strconv.Atoi(Pid)
	fmt.Println("4. put t.ID ", t.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if _, err := stmt.Exec(Pid, t.Title, t.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Exec error": err.Error()})
		return
	}
	fmt.Println("Update successed ", t.ID, t.Title, t.Status)

	c.JSON(http.StatusOK, t)

}

func (h *TodoHandler) DeleteTodosHandler(c *gin.Context) {
	//db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	db, err := database.ConnDB()
	if err != nil {
		log.Fatal("Can not open database", err.Error)
	}
	defer db.Close()
	//todos := []Todo{}
	query := `DELETE FROM todos WHERE id=$1`
	var id int
	db.QueryRow(query, id)
	fmt.Println("Record Deleted ")
	c.JSON(200, gin.H{"status": "success"})
}
