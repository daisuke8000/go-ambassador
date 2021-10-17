package controllers

import (
	"ambassador-backend/src/database"
	"ambassador-backend/src/models"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func Ambassadors(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_ambassador = true").Find(&users)
	return c.JSON(users)
}

func Rankings(c *fiber.Ctx) error {
	//var users []models.User
	//database.DB.Find(&users, models.User{
	//	IsAmbassador: true,
	//})
	rankings, err := database.Cache.ZRevRangeByScoreWithScores(context.Background(), "rankings", &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	if err != nil{
		return err
	}

	//Patern1 Arfabet sort
	result := make(map[string]float64)

	for _, ranking := range rankings {
		result[ranking.Member.(string)] = ranking.Score
	}

	//Patern2 int sort
	//result := make(map[int]interface{})
	//
	//for i, ranking := range rankings {
	//	result[i] = ranking.Score
	//}


	//var result []interface{}

	//for _, user := range users {
	//	ambassador := models.Ambassador(user)
	//	ambassador.CalculateRevenue(database.DB)
	//
	//	result = append(result, fiber.Map{
	//		user.Name(): ambassador.Revenue,
	//	})
	//}
	return c.JSON(result)
}