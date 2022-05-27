package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Pallinder/go-randomdata"
	"github.com/valyala/fasthttp"
)

var (
	client     fasthttp.Client
	threads    int = 250
	postalCode int = 1337
)

func main() {
	wg := sync.WaitGroup{}
	goroutines := make(chan struct{}, threads)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			goroutines <- struct{}{}
			for {
				sign()
			}
			//<-goroutines
		}()
	}

	wg.Wait()
}

func sign() {
	firstName := randomdata.FirstName(randomdata.Male)
	lastName := randomdata.LastName()
	email := fmt.Sprintf("%v@gmail.com", randomdata.SillyName())

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.Header.SetMethod("POST")
	req.SetRequestURI("https://www.alp.org.au/Umbraco/ML/FormComponent/DoAction")
	req.SetBody([]byte(fmt.Sprintf("Bootstrap[Theme]=2&Bootstrap[StagesFields][0][0][Name]=&Bootstrap[StagesFields][0][0][Value]=&Bootstrap[StagesFields][0][0][Placeholder]=&Bootstrap[StagesFields][0][0][Required]=false&Bootstrap[StagesFields][0][0][Hidden]=false&Bootstrap[StagesFields][0][0][Rows]=&Bootstrap[StagesFields][0][0][Type]=0&Bootstrap[PostSubmitFunc]=&Bootstrap[UmbracoContentId]=1544&Bootstrap[UmbracoContentName]=Petition-TAFEcuts&Bootstrap[UmbracoContentUrl]=/components-forms/petition-tafe-cuts/&Bootstrap[UmbracoContentPageId]=1543&Bootstrap[UmbracoContentPageName]=NoTAFECuts&Bootstrap[UmbracoContentPageUrl]=/petitions/no-tafe-cuts/&Bootstrap[ComponentClassName]=&Bootstrap[ComponentId]=FormComponent2f9e5333-2ea9-4fa7-a1db-4fda53fd4c09&SubmitButton[Link][LinkUrl]=&SubmitButton[Link][LinkTarget]=&SubmitButton[Link][Caption]=&SubmitButton[BackgroundColour]=&SubmitButton[TextColour]=&SubmitButton[ClassName]=&SubmitButton[ParentClassName]=&SubmitButton[Alignment]=2&SubmitButton[Display]=1&SubmitButton[Theme]=0&State=0&Action=0&Theme=2&Stage=0&Fields[0][Name]=firstname&Fields[0][Value]=&Fields[0][Placeholder]=Firstname&Fields[0][Required]=true&Fields[0][Hidden]=false&Fields[0][Rows]=1&Fields[0][Type]=0&Fields[1][Name]=lastname&Fields[1][Value]=&Fields[1][Placeholder]=Lastname&Fields[1][Required]=true&Fields[1][Hidden]=false&Fields[1][Rows]=1&Fields[1][Type]=0&Fields[2][Name]=email&Fields[2][Value]=&Fields[2][Placeholder]=Email&Fields[2][Required]=true&Fields[2][Hidden]=false&Fields[2][Rows]=1&Fields[2][Type]=0&Fields[3][Name]=postcode&Fields[3][Value]=&Fields[3][Placeholder]=Postcode&Fields[3][Required]=true&Fields[3][Hidden]=false&Fields[3][Rows]=1&Fields[3][Type]=0&Fields[4][Name]=mobile&Fields[4][Value]=&Fields[4][Placeholder]=Mobile&Fields[4][Required]=false&Fields[4][Hidden]=false&Fields[4][Rows]=1&Fields[4][Type]=0&StagesFields[0][0][Name]=&StagesFields[0][0][Value]=&StagesFields[0][0][Placeholder]=&StagesFields[0][0][Required]=false&StagesFields[0][0][Hidden]=false&StagesFields[0][0][Rows]=&StagesFields[0][0][Type]=0&FieldValues={\"firstname\":\"%v\",\"lastname\":\"%v\",\"email\":\"%v\",\"postcode\":\"%v\"}&StagesFieldValues=&Count=950&CountLabel={0}+people+have+signed,+will+you?&Errors=&Metadata=&tempJD8I9f0asdf0isd0fsop[0][firstname]=%v&tempJD8I9f0asdf0isd0fsop[1][lastname]=%v&tempJD8I9f0asdf0isd0fsop[2][email]=%v&tempJD8I9f0asdf0isd0fsop[3][postcode]=%v", firstName, lastName, email, postalCode, firstName, lastName, email, postalCode)))

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	err := client.Do(req, res)
	if err != nil {
		return
	}

	if strings.Contains(string(res.Body()), "people have signed") {
		fmt.Println("Successfully signed..")
	} else {
		fmt.Println(string(res.Body()))
	}
}
