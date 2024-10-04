package structural

import "fmt"

type Client struct{}

func (c *Client) InsertLightingCableIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
