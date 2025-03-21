package utils

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleArray(c *gin.Context){
	var reqDetails struct{
		List	[]string	`json:"data" binding:"required"`
	}

	type response struct{
		Success	bool		`json:"is_success"`
		UserID	string		`json:"user_id"`
		Email	string		`json:"email"`
		Roll	string		`json:"roll_number"`
		Nums	[]string	`json:"numbers"`
		Alpha	[]string	`json:"alphabets"`	
		High    []string	`json:"highest_alphabet"`
	}

	err := c.BindJSON(&reqDetails)
	if err != nil {
		c.IndentedJSON(400, gin.H{"msg":"error bad request"})
		return
	}

	alphabets := []string{}
	highest := []string{}
	nums := []string{}

	for _, item := range reqDetails.List{
		if(item >= "a" && item <= "z" || item >= "A" && item <= "Z"){
			alphabets = append(alphabets, item)
		}else{
			x, err := strconv.Atoi(item)
			if err != nil{
				log.Printf("error: %v", err)
				c.IndentedJSON(500, gin.H{"msg":"internal function error"})
				return
			}
			nums = append(nums, fmt.Sprint(x))
		}
	}

	if len(alphabets) != 0{
		tempList := alphabets
		sort.Sort(sort.Reverse(sort.StringSlice(tempList)))
		s := tempList[0];

		highest = append(highest, s)
	}


	requestResponse := response{
		Success: true,
		UserID: "aaryaj",
		Roll: "12",
		Email: "aaryaj@example.com",
		Nums: nums,
		Alpha: alphabets,
		High: highest,
	}

	c.IndentedJSON(201, requestResponse)
}


func HandleOperation(c *gin.Context){
	c.IndentedJSON(200, gin.H{"operation_code":1})
}