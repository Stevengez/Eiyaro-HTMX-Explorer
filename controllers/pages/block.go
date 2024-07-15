package pages

import (
	"eiyaro-htmx-explorer/models"
	"eiyaro-htmx-explorer/service/node"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBlock(c *fiber.Ctx) error {
	translation := c.Locals("lang").(models.Translation)

	blockNumber := c.Params("height", "0")
	blockHeight, err := strconv.ParseInt(blockNumber, 10, 64)

	if err != nil {
		fmt.Println("Failed to parse block height value", err)
		return err
	}

	blockData, err := node.GetBlock(blockHeight)

	if err != nil {
		fmt.Println("Failed to retrieve block data", err)
		return err
	}

	return c.Render(
		"templates/block",
		fiber.Map{
			"Title": translation.M("pages").S("index"),
			"lang":  translation,
			"block": blockData,
		},
		"templates/layout/global",
	)
}
