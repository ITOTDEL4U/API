package main

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
