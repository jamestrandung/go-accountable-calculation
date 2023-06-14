package main

import (
	"fmt"
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
	"github.com/jamestrandung/go-accountable-calculation/float"
	"testing"
)

type fare struct {
	numPickUp              float64
	perPickUpPrice         float64
	numDropOff             float64
	perDropOffPrice        float64
	numStopSurcharge       float64
	startingFare           float64
	loudMouthSurcharge     float64
	loudMouthCount         float64
	smellyShoeSurcharge    float64
	smellyShoeCount        float64
	inconvenienceSurcharge float64
	platformFee            float64
	weightBasedFee         float64
	totalFare              float64
}

func bareCalculation() *fare {
	f := &fare{}

	f.numPickUp = 1.0
	f.perPickUpPrice = 3.0

	f.numDropOff = 3.0
	f.perDropOffPrice = 2.0

	f.numStopSurcharge = f.numPickUp*f.perPickUpPrice + f.numDropOff*f.perDropOffPrice

	f.startingFare = 1.0
	f.loudMouthSurcharge = 2.0
	f.loudMouthCount = 2.0
	f.smellyShoeSurcharge = 1.0
	f.smellyShoeCount = 1.0

	f.inconvenienceSurcharge = f.startingFare + f.loudMouthSurcharge*f.loudMouthCount + f.smellyShoeSurcharge*f.smellyShoeCount

	f.platformFee = 2.0
	f.weightBasedFee = 2.0

	f.totalFare = f.inconvenienceSurcharge + f.platformFee + f.weightBasedFee + f.numStopSurcharge
	return f
}

type fareWithFormula struct {
	numPickUp              float.Simple
	perPickUpPrice         float.Simple
	numDropOff             float.Simple
	perDropOffPrice        float.Simple
	numStopSurcharge       float.Simple
	startingFare           float.Simple
	loudMouthSurcharge     float.Simple
	loudMouthCount         float.Simple
	smellyShoeSurcharge    float.Simple
	smellyShoeCount        float.Simple
	inconvenienceSurcharge float.Simple
	platformFee            float.Simple
	weightBasedFee         float.Simple
	totalFare              float.Simple
}

func formulaCalculation() fareWithFormula {
	f := fareWithFormula{}

	f.numPickUp = float.MakeSimpleFromFloat("NumPickUp", 1)
	f.perPickUpPrice = float.MakeSimpleFromFloat("PerPickUpPrice", 3)
	f.numDropOff = float.MakeSimpleFromFloat("NumDropOff", 3)
	f.perDropOffPrice = float.MakeSimpleFromFloat("PerDropOffPrice", 2)

	acal.SourceRequest.Apply(f.numPickUp, f.numDropOff)
	acal.SourceHardcode.Apply(f.perPickUpPrice, f.perDropOffPrice)

	f.numStopSurcharge = f.numPickUp.Mul(f.perPickUpPrice).Add(f.numDropOff.Mul(f.perDropOffPrice)).Anchor("NumStopSurcharge")

	f.startingFare = float.MakeSimpleFromFloat("StartingFare", 1)
	f.loudMouthSurcharge = float.MakeSimpleFromFloat("LoudMouthSurcharge", 2)
	f.loudMouthCount = float.MakeSimpleFromFloat("LoudMouthCount", 2)
	f.smellyShoeSurcharge = float.MakeSimpleFromFloat("SmellyShoeSurcharge", 1)
	f.smellyShoeCount = float.MakeSimpleFromFloat("SmellyShoeCount", 1)

	acal.SourceRequest.Apply(f.loudMouthCount, f.smellyShoeCount)
	acal.SourceHardcode.Apply(f.startingFare, f.loudMouthSurcharge, f.smellyShoeSurcharge)

	f.inconvenienceSurcharge = f.startingFare.Add(f.loudMouthSurcharge.Mul(f.loudMouthCount)).Add(f.smellyShoeSurcharge.Mul(f.smellyShoeCount)).Anchor("InconvenienceSurcharge")

	f.platformFee = float.MakeSimpleFromFloat("PlatformFee", 2)
	f.weightBasedFee = float.MakeSimpleFromFloat("WeightBasedFee", 2)
	acal.SourceHardcode.Apply(f.platformFee, f.weightBasedFee)

	f.totalFare = f.inconvenienceSurcharge.Add(f.platformFee).Add(f.weightBasedFee).Add(f.numStopSurcharge).Anchor("TotalFare")
	return f
}

type complexFareWithFormula struct {
	numPickUp              float.Simple
	perPickUpPrice         float.Simple
	numDropOff             float.Simple
	perDropOffPrice        float.Simple
	numStopSurcharge       float.Simple
	startingFare           float.Simple
	loudMouthSurcharge     float.Simple
	loudMouthCount         float.Simple
	smellyShoeSurcharge    float.Simple
	smellyShoeCount        float.Simple
	inconvenienceSurcharge float.Simple
	platformFee            float.Simple
	weightBasedFee         float.Simple
	totalFare              float.Progressive
}

func Test_Something(t *testing.T) {
	c := complexFareCalculation()
	fmt.Println(acal.ToString(c.totalFare))
}

func complexFareCalculation() complexFareWithFormula {
	f := complexFareWithFormula{}

	f.numPickUp = float.MakeSimpleFromFloat("NumPickUp", 1)
	f.perPickUpPrice = float.MakeSimpleFromFloat("PerPickUpPrice", 3)
	f.numDropOff = float.MakeSimpleFromFloat("NumDropOff", 3)
	f.perDropOffPrice = float.MakeSimpleFromFloat("PerDropOffPrice", 2)

	acal.SourceRequest.Apply(f.numPickUp, f.numDropOff)
	acal.SourceHardcode.Apply(f.perPickUpPrice, f.perDropOffPrice)

	f.numStopSurcharge = f.numPickUp.Mul(f.perPickUpPrice).Add(f.numDropOff.Mul(f.perDropOffPrice)).Anchor("NumStopSurcharge")

	f.startingFare = float.MakeSimpleFromFloat("StartingFare", 1)
	f.loudMouthSurcharge = float.MakeSimpleFromFloat("LoudMouthSurcharge", 2)
	f.loudMouthCount = float.MakeSimpleFromFloat("LoudMouthCount", 2)
	f.smellyShoeSurcharge = float.MakeSimpleFromFloat("SmellyShoeSurcharge", 1)
	f.smellyShoeCount = float.MakeSimpleFromFloat("SmellyShoeCount", 1)

	acal.SourceRequest.Apply(f.loudMouthCount, f.smellyShoeCount)
	acal.SourceHardcode.Apply(f.startingFare, f.loudMouthSurcharge, f.smellyShoeSurcharge)

	f.inconvenienceSurcharge = f.startingFare.Add(f.loudMouthSurcharge.Mul(f.loudMouthCount)).Add(f.smellyShoeSurcharge.Mul(f.smellyShoeCount)).Anchor("InconvenienceSurcharge")

	f.totalFare = float.MakeProgressive("TotalFare")

	f.totalFare.Update(f.inconvenienceSurcharge.Add(f.numStopSurcharge))

	boolean.If(
		boolean.MakeSimple("IsPlatformFeeEnabled", true), func(criteria boolean.Interface) {
			defer acal.ApplyProgressiveCondition(criteria, f.totalFare)()

			platformFee := float.MakeSimpleFromFloat("PlatformFee", 2)
			platformFee.From(acal.SourceHardcode)

			f.totalFare.Update(f.totalFare.Add(platformFee))

			acal.ApplyStaticCondition(criteria, platformFee)
		},
	)

	f.weightBasedFee = float.MakeSimpleFromFloat("WeightBasedFee", 2)
	f.weightBasedFee.From(acal.SourceHardcode)

	f.totalFare.Update(f.totalFare.Add(f.weightBasedFee))

	return f
}

func simpleCalculation() {
	startingFare := float.MakeSimpleFromFloat("StartingFare", 1)

	totalFare := float.MakeProgressive("TotalFare")

	totalFare.Update(startingFare)

	boolean.If(
		boolean.MakeSimple("IsPlatformFeeEnabled", true), func(criteria boolean.Interface) {
			defer acal.ApplyProgressiveCondition(criteria, totalFare)()

			platformFee := float.MakeSimpleFromFloat("PlatformFee", 2)
			platformFee.From(acal.SourceHardcode)

			totalFare.Update(totalFare.Add(platformFee))

			acal.ApplyStaticCondition(criteria, platformFee)
		},
	)
}

func testCondition() {
	startingFare := float.MakeSimpleFromFloat("StartingFare", 1)

	totalFare := float.MakeProgressive("TotalFare")

	totalFare.Update(startingFare)

	boolean.If(
		boolean.MakeSimple("IsPlatformFeeEnabled", true), func(criteria boolean.Interface) {
			defer acal.ApplyProgressiveCondition(criteria, totalFare)()

			platformFee := float.MakeSimpleFromFloat("PlatformFee", 2)
			platformFee.From(acal.SourceHardcode)

			totalFare.Update(totalFare.Add(platformFee))
		},
	)

	boolean.If(
		boolean.MakeSimple("WeightBaseFeeEnabled", true), func(criteria boolean.Interface) {
			defer acal.ApplyProgressiveCondition(criteria, totalFare)()

			weightBaseFee := float.MakeSimpleFromFloat("WeightBaseFee", 2)
			weightBaseFee.From(acal.SourceHardcode)

			totalFare.Update(totalFare.Add(weightBaseFee))
		},
	)
}

func BenchmarkCondition(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testCondition()
	}
}

func BenchmarkCalculationUsingPrimitive(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		bareCalculation()
	}
}

func BenchmarkOneCalculation(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		numPickUp := float.MakeSimpleFromFloat("NumPickUp", 1)
		numPickUp.GetName()
		//perPickUpPrice := float.MakeSimpleFromFloat("PerPickUpPrice", 3)
		//
		//numStopSurcharge := numPickUp.Mul(perPickUpPrice)
		//numStopSurcharge.Anchor("NumStopSurcharge")
	}
}

func BenchmarkCalculationUsingFramework(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		formulaCalculation()
	}
}

func BenchmarkComplexFormulaCalculation(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		complexFareCalculation()
	}
}
