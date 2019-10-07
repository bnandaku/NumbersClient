package NumbersClient

import "net/rpc"

type Client struct {
	DB *rpc.Client
}

type Args struct {

	number string
	value interface{}

}

type Response struct {
	Response interface{}
}

func Connect(addr string, port string) (*Client, error) {
	cl, err := rpc.DialHTTP("tcp", addr + ":" + port)
	if err != nil {
		return nil, err
	}
	client := &Client{
		DB:cl,
	}
	return client, err
}


func (c *Client) Insert(value interface{}, number string) bool {
	args := &Args{
		number:number,
		value:value,
	}
	var response *Response
	err := c.DB.Call("DataBase.Insert", args, response)

	if err != nil {
		return false
	}
	ans := response.Response.(bool)
	return ans
}


func (c *Client) Update(value interface{}, number string) bool {
	args := &Args{
		number:number,
		value:value,
	}
	var response *Response
	err := c.DB.Call("DataBase.Update", args, response)

	if err != nil {
		return false
	}
	ans := response.Response.(bool)
	return ans
}


func (c *Client) Search(number string) interface{} {
	args := &Args{
		number:number,
	}
	var response *Response
	err := c.DB.Call("DataBase.Search", args, response)

	if err != nil {
		return false
	}

	return response.Response
}

func (c *Client) Delete(number string) bool {
	args := &Args{
		number:number,
	}
	var response *Response
	err := c.DB.Call("DataBase.Delete", args, response)

	if err != nil {
		return false
	}
	ans := response.Response.(bool)
	return ans
}