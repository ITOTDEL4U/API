# API for site https://bank.gov.ua/ Exchange Rates
parse JSON for CURRENT date exaple json: https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?date=20221118&json

we have http client, response and roundtrip for loging

example use:

import ( 
 "fmt"   
 "github.com/ITOTDEL4U/API/NBU"   
 "log"  
 "time"
)

func main(){

   nbuClient, err := NBU.NewClient(time.Second * 10)
   if err != nil{
       log.Fatal(err)
   }

   assets,err := nbuClient.GetAssets()

   if err != nil{
       log.Fatal(err)
   }

  for  _, str := range assets  {
      fmt.Println(str.INFO())
    }

  }
