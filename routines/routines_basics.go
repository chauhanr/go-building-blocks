package routines

func sendData(ch chan string, c []string) {
	for _, city := range c {
		ch <- city
	}
}

func getData(ch chan string, c *CityStruct) {
	var input string

	for {
		input = <-ch
		c.cities = append(c.cities, input)
	}
}
