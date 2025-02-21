package database

import (
    "database/sql"
    "log"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
    var err error
    connStr := "user=postgres password=yourpassword dbname=task_manager sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
}

func GetTasks(c *gin.Context) {
    rows, err := db.Query("SELECT id, title, status FROM tasks")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var tasks []map[string]interface{}
    for rows.Next() {
        var id int
        var title, status string
        rows.Scan(&id, &title, &status)
        tasks = append(tasks, map[string]interface{}{
            "id":     id,
            "title":  title,
            "status": status,
        })
    }
    c.JSON(200, tasks)
}

func CreateTask(c *gin.Context) {
    var task struct {
        Title string `json:"title"`
    }
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    _, err := db.Exec("INSERT INTO tasks (title) VALUES ($1)", task.Title)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Task created"})
}
