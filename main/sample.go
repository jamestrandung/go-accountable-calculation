package main

import (
	"fmt"

	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
	"github.com/jamestrandung/go-accountable-calculation/float"
)

func main() {
	totalFare := float.MakeProgressive("TotalFare")

	// Step 1: begin with starting fare as a minimum
	startingFare := float.MakeSimple("StartingFare", 1)
	startingFare.From(acal.SourceHardcode)
	startingFare.Tag(acal.NewTag("RandomTag", 3))

	totalFare.Update(startingFare)

	// Step 2: add surcharge for each additional stop
	additionalStopSurcharge := calculateAdditionalStopSurcharge()

	totalFare.Update(totalFare.Add(additionalStopSurcharge))

	// Step 3: add surcharge for inconvenience
	inconvenienceSurcharge := calculateInconvenienceSurcharge()

	totalFare.Update(totalFare.Add(inconvenienceSurcharge))

	// Step 4: add platform fee if enabled
	boolean.If(
		boolean.MakeSimple("IsPlatformFeeEnabled", true), func(criteria boolean.Value) {
			defer acal.ApplyProgressiveCondition(criteria, totalFare)()

			platformFee := float.MakeSimple("PlatformFee", 2)
			platformFee.From(acal.SourceHardcode)

			totalFare.Update(totalFare.Add(platformFee))

			acal.ApplyStaticCondition(criteria, platformFee)
		},
	)

	// Step 5: add VAT
	vatPercentage := float.MakeSimple("VATPercentage", 10)
	vatPercentage.From(acal.SourceHardcode)

	// TotalFare = TotalFare + TotalFare * VATPercentage / 100
	totalFare.Update(totalFare.Mul(vatPercentage).Div(float.Hundred).Add(totalFare))

	totalFare.Tag(acal.NewTag("AnotherRandomTag", true))

	fmt.Println(acal.ToString(totalFare))
}

func calculateAdditionalStopSurcharge() float.Simple {
	numPickUp := float.MakeSimple("NumPickUp", 1)
	perPickUpPrice := float.MakeSimple("PerPickUpPrice", 3)
	numDropOff := float.MakeSimple("NumDropOff", 3)
	perDropOffPrice := float.MakeSimple("PerDropOffPrice", 2)

	acal.SourceRequest.Apply(numPickUp, numDropOff)
	acal.SourceHardcode.Apply(perPickUpPrice, perDropOffPrice)

	pickUpSurcharge := numPickUp.Sub(float.One).Mul(perPickUpPrice)
	dropOffSurcharge := numDropOff.Sub(float.One).Mul(perDropOffPrice)

	// AdditionalStopSurcharge = (NumPickUp - 1) * PerPickUpPrice + (NumDropOff - 1) * perDropOffPrice
	return pickUpSurcharge.Add(dropOffSurcharge).Anchor("AdditionalStopSurcharge")
}

func calculateInconvenienceSurcharge() float.Simple {
	loudMouthSurcharge := float.MakeSimple("LoudMouthSurcharge", 2)
	loudMouthCount := float.MakeSimple("LoudMouthCount", 2)
	smellyShoeSurcharge := float.MakeSimple("SmellyShoeSurcharge", 1)
	smellyShoeCount := float.MakeSimple("SmellyShoeCount", 1)

	acal.SourceRequest.Apply(loudMouthCount, smellyShoeCount)
	acal.SourceHardcode.Apply(loudMouthSurcharge, smellyShoeSurcharge)

	// InconvenienceSurcharge = LoudMouthSurcharge * LoudMouthCount + SmellyShoeSurcharge * SmellyShoeCount
	return loudMouthSurcharge.Mul(loudMouthCount).Add(smellyShoeSurcharge.Mul(smellyShoeCount)).Anchor("InconvenienceSurcharge")
}
