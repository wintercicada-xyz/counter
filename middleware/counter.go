package middleware

import (
	"counter/counter2"
	"fmt"

	"github.com/gin-gonic/gin"
)

const day = 1000 * 60 * 60 * 24

type Counter func() gin.HandlerFunc

func (c *Counter) Init(f func(m map[string]*uint64)) {
	counter2.Init()
	counter2.Flush2broker(day, f)
	*c = func() gin.HandlerFunc {
		return func(c *gin.Context) {
			key := c.Request.URL.Path
			counter2.Incr(key, 1)
			c.Next()
		}
	}
}

func Flush2brokerStderr(m map[string]*uint64) {
	for key, count := range m {
		fmt.Println("PMS request ", key, " count: ", *count)
		// Write to database
		// result, err := db.Exec("INSERT INTO pms VALUES (CURDATE(), ?, ?)", key, *count)
	}
}
