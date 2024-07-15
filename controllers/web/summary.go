package web

import (
	"eiyaro-htmx-explorer/service/node"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetLastBlock(c *fiber.Ctx) error {

	block, _ := node.GetLastBlock()

	return c.Render(
		"templates/layout/misc/cards/block",
		fiber.Map{
			"id":      "card-lastblock",
			"getPath": "lastBlock",
			"value":   block,
		},
	)
}

func GetLastBlocks(c *fiber.Ctx) error {

	linkPrefix := c.Query("prefix", "en")
	lastBlock, err := node.GetLastBlock()

	if err != nil {
		fmt.Println("Failed to get last block")
		return nil
	}

	blocks, err := node.GetBlocks(int64(lastBlock), 10)

	if err != nil {
		fmt.Println("Failed to get last blocks")
		return nil
	}

	return c.Render(
		"templates/layout/misc/tables/blockrow",
		fiber.Map{
			"linkPrefix": linkPrefix,
			"blocks":     blocks,
		},
	)
}

func GetHashRate(c *fiber.Ctx) error {

	hashRate, _ := node.GetHashRate()

	return c.Render(
		"templates/layout/misc/cards/block",
		fiber.Map{
			"id":      "card-hashrate",
			"getPath": "hashRate",
			"value":   fmt.Sprintf("%.2f Mh/s", hashRate),
		},
	)
}

func GetDifficulty(c *fiber.Ctx) error {

	difficulty, _ := node.GetDifficulty()

	return c.Render(
		"templates/layout/misc/cards/block",
		fiber.Map{
			"id":      "card-difficulty",
			"getPath": "difficulty",
			"value":   fmt.Sprintf("%.3f G", difficulty),
		},
	)
}

func GetPendingCount(c *fiber.Ctx) error {

	pendingCount, _ := node.GetPendingXfers()

	return c.Render(
		"templates/layout/misc/cards/block",
		fiber.Map{
			"id":      "card-pendings",
			"getPath": "pendings",
			"value":   pendingCount,
		},
	)
}
