package main

import (
    "fmt"
    "github.com/jamestrandung/go-accountable-calculation/acal"
    "github.com/jamestrandung/go-accountable-calculation/float"
)

func main() {
    price := float.MakeSimpleFromFloat("ProductPrice", 2.5)
    quantity := float.MakeSimpleFromInt("Quantity", 3)

    // Give the variable an identity by anchoring it with a fixed name
    orderValue := price.Mul(quantity).Anchor("OrderValue")

    fmt.Println(acal.ToString(orderValue))

    // {
    //  "OrderValue": {
    //    "Value": "7.5",
    //    "Source": "static_calculation",
    //    "Formula": {
    //      "Category": "TwoValMiddleOp",
    //      "Operation": "*",
    //      "Operands": [
    //        {
    //          "Name": "ProductPrice"
    //        },
    //        {
    //          "Name": "Quantity"
    //        }
    //      ]
    //    }
    //  },
    //  "ProductPrice": {
    //    "Value": "2.5",
    //    "Source": "unknown"
    //  },
    //  "Quantity": {
    //    "Value": "3",
    //    "Source": "unknown"
    //  }
    //}
}
