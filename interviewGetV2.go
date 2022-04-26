package main

//Let's import packages (libs) that will be needed for our program.
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Define a struct for Customer (Used in part 1 of the project)
type Customer struct {
	ID    string
	Name  string
	Date  string
	Array []string
}

//Begin the main.
func main() {

	var match bool  //This boolean will be used as the control element within the endpoint loop.
	var c1 Customer //Palceholder for each customer we ecounter as we step through the endpoint JSON objects.
	var customers []Customer
	var uuids []string
	i := 9                                 //The default "final" index for customers array is nine after the initial GTTP GET request. This number will grow as we encounter new JSONs at the endpoint.
	outFile, _ := os.Create("./items.txt") //This opens up an output file in the current directory of the program. The file will hold a table containing each JSON object with its corresponding values.
	w := bufio.NewWriter(outFile)          //Writer that can interface with our newly opened output file.

	//Let's define variables for the response and error returned by the Get request.
	resp, err := http.Get("http://localhost:3456/items/")
	//Is there an error? If so, redirect to the log for inspection.
	if err != nil {
		log.Fatalln(err)
	}

	//Now let's look at the body contents of the response--trimming off any of the header data and such.
	body, err := ioutil.ReadAll(resp.Body)
	//Again, given an error--let's document it
	if err != nil {
		log.Fatalln(err)
	}

	//Let's Unmarshal the JSON elements into our customer placeholder.
	json.Unmarshal([]byte(body), &customers)

	//Iterate through the endpoint until we finally match the array UUID of the most recent customer encountered with the UUID of the first customer. By this logic, we will have "come full circle" and exhausted all of the elements on the endpoint.
	for match == false {
		urlUuid := fmt.Sprintf("http://localhost:3456/items/%v", customers[i].Array[0]) // Format the string so that it navigates to each new Array UUID allowing us to step through.
		resp, err := http.Get(urlUuid)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		json.Unmarshal([]byte(body), &c1)
		if c1.ID == customers[0].ID {
			match = true
		}
		customers = append(customers, c1)
		i++

	}
	//Let's collect the remaining UUIDs in the arrays that we did not visit.
	for j := 0; j < len(customers); j++ {
		if len(customers[j].Array) > 1 {
			for k := 0; k < len(customers[j].Array); k++ {
				uuids = append(uuids, customers[j].Array[k])
			}
		}
	}

	for l := 0; l < len(customers); l++ {
		fmt.Println(l)
		fmt.Println(len(customers))
		for n := 0; n < len(uuids); n++ {
			if customers[l].ID != uuids[n] {
				urlUuid := fmt.Sprintf("http://localhost:3456/items/%v", uuids[n]) // Format the string so that it navigates to each new Array UUID allowing us to step through.
				resp, err := http.Get(urlUuid)
				if err != nil {
					log.Fatalln(err)
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}
				json.Unmarshal([]byte(body), &c1)
				customers = append(customers, c1)
			}
		}
	}
	fmt.Fprintf(w, "There were %d customers found on the endppoint, each containing 4 items. Making the total number of items %d.", len(customers), (4*len(customers)))
	fmt.Fprintf(w, "%-50v %-50v %-50v %-50v\n", "Name", "ID", "Date", "ArrayUUID")
	for m := 0; m < len(customers); m++ {
		fmt.Fprintf(w, "%-50v %-50v %-50v %-50v\n\n", customers[m].Name, customers[m].ID, customers[m].Date, customers[m].Array)
	}
}
