package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func main() {

	file := "data.txt"                   //System path to the file we want to gather input data from.
	inFile, err := ioutil.ReadFile(file) //Let's store the input data from the file and check for any errors while opening/reading the file.
	if err != nil {
		log.Fatalln(err)
	}

	//Let's parse and format the string in-order to make the input comply with the Unmarshal method. This will enable storing the "strings" into an array of string.
	s2 := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.TrimLeft(strings.TrimRight(string(inFile), "``"), "``"), " ", ""), "\n", ""), "text", "")

	endpoint := "http://localhost:3456/submit"
	var items []string                 //This is our string array which will house all of the "unmarshalled" data.
	json.Unmarshal([]byte(s2), &items) //Yep, it's time to pull the data out and stuff it into our array.
	data := url.Values{}               //We will stuff all of the string elements of the array intot his before attempting to enconde and post at `submit` EP
	i := 0                             //This is our loop control counter.

	//Now let's feed the String array into data in-order to interatively POST to the endpoint.
	for i < len(items) {
		pItem1 := fmt.Sprintf("%v,", items[i]) //Payload item 1
		data.Set("importantData", pItem1)      //Let's start giving `data` some meat

		//Are we going beyond the bounds of `items`? If not, fill the second element of the payload with the next string. Otherwise, fill with `#`
		if i+1 < len(items) {
			pItem2 := fmt.Sprintf("%v,", items[i+1]) //Payload item 2 is the next string
			data.Add("importantData", pItem2)        //Stuff it in `data`
		} else {
			pItem2 := "#,"
			data.Add("importantData", pItem2)
		}
		//Same as above conditional
		if i+2 < len(items) {
			pItem3 := fmt.Sprintf("%v", items[i+2])
			data.Add("importantData", pItem3)
		} else {
			pItem3 := "#"
			data.Add("importantData", pItem3)
		}
		client := &http.Client{}
		r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
		if err != nil {
			log.Fatal(err)
		}
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		res, err := client.Do(r)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.Status)
		defer res.Body.Close()
		i = i + 3 //Given that we have used three consecutive string elements, skip forward to avoid stuffing in duplicate strings.
	}

}
