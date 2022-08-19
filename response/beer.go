package response

import "fmt"

type Beer struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d *Beer) Info() string {
	return fmt.Sprintf("Beer: %s \nDesc: %s\n", d.Name, d.Description)
}
