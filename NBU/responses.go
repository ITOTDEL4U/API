package NBU

import "fmt"

type AssetData  struct {
    Code         int      `json:"r030"`
    Name         string   `json:"txt"`
    Rate         float64  `json:"rate"`
    ShortName    string   `json:"cc"`
    Exchangedate string   `json:"exchangedate" `
}

func (ad AssetData) INFO () string {

    return fmt.Sprintf("[Exchangedate] %s | [ShortName] %s | [Name] %s | [Rate] %.6f ", ad.Exchangedate, ad.ShortName, ad.Name, ad.Rate)

}