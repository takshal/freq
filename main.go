package main

import (
	"sync"
	"bufio"
	"net/http"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main(){
	// fmt.Println("\n")
	fmt.Println("frequester tool By tojojo !!")
	fmt.Println("\\__(-_-)__/")
	// fmt.Println("\n")

	colorReset := "\033[0m"
	colorRed := "\033[31m"
    colorGreen := "\033[32m"


	sc := bufio.NewScanner(os.Stdin)

	jobs := make(chan string)
	var wg sync.WaitGroup

	for i:= 0; i < 20; i++{

		wg.Add(1)
		go func(){
			defer wg.Done()
			for domain := range jobs {

				resp, err := http.Get(domain)
				if err != nil{
					continue
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
	      			fmt.Println(err)
	   			}
	   			sb := string(body)
	   			check_result := strings.Contains(sb , "<img src=x onerror=alert(1)>")
	   			// fmt.Println(check_result)
	   			if check_result != false {
	   				fmt.Println(string(colorRed),"Vulnerable To XSS:", domain,string(colorReset))
	   			}else{
	   				fmt.Println(string(colorGreen),"Not Vulnerable To XSS:", domain, string(colorReset))
	   			}

			}
			
   		}()

	}



	for sc.Scan(){
		domain := sc.Text()
		jobs <- domain		
		

	}
	close(jobs)
	wg.Wait()
}
