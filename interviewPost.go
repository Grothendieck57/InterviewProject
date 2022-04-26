package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
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
		pItem1 := fmt.Sprintf("%v,", items[i]) //Payload item variable for holding the string that goes into the URL variable.
		data.Set("importantData", pItem1)      //Let's start giving `data` some meat
		for j := 0; j < (len(items)/2); j++ {
			value := i + j
			//Given that we are not at the final element, then the payload item needs to include a comma. Let's also make sure we don't cross the bounds of the `items` array.
			if j < ((len(items)/2)-1) {
				if value < len(items) {
					pItem1 = fmt.Sprintf("%v,", items[i+j])
					data.Add("importantData", pItem1)
				} else {
					pItem1 = "#,"
					data.Add("importantData", pItem1)
				}
			}
			//Now that we are at the last element in the string, let's make sure there isn't a comma left floating there. Still making sure we don't traverse out of `items` bounds.
			if value < len(items) {
				pItem1 = fmt.Sprintf("%v", items[i+j])
				data.Add("importantData", pItem1)
			} else {
				pItem1 = "#"
				data.Add("importantData", pItem1)
			}
		}
		time.Sleep(500 * time.Millisecond) //Let's make a POST request every 500 milliseconds. From testing, the server seems to be happy with this rate.
		client := &http.Client{}
		r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
		if err != nil {
			log.Fatal(err)
		}
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		//r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		res, err := client.Do(r)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.Status)
		defer res.Body.Close()
		i = i + (len(items)/2) //Given that we have used three consecutive string elements, skip forward to avoid stuffing in duplicate strings.
	}

}
