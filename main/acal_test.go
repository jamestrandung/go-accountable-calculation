package main

import (
	"github.com/jamestrandung/go-accountable-calculation/acal"
	"github.com/jamestrandung/go-accountable-calculation/boolean"
	"github.com/jamestrandung/go-accountable-calculation/float"
	"testing"
)

type fare struct {
	numPickUp           float64
	perPickUpPrice      float64
	numDropOff          float64
	perDropOffPrice     float64
	numStopSurcharge    float64
	baseFare            float64
	distanceCoefficient float64
	distance            float64
	timeCoefficient     float64
	duration            float64
	nsnsFare            float64
	platformFee         float64
	weightBasedFee      float64
	finalFare           float64
}

func bareCalculation() *fare {
	f := &fare{}

	f.numPickUp = 1.0
	f.perPickUpPrice = 3.0

	f.numDropOff = 3.0
	f.perDropOffPrice = 2.0

	f.numStopSurcharge = f.numPickUp*f.perPickUpPrice + f.numDropOff*f.perDropOffPrice

	f.baseFare = 1.0
	f.distanceCoefficient = 2.0
	f.distance = 2.0
	f.timeCoefficient = 1.0
	f.duration = 1.0

	f.nsnsFare = f.baseFare + f.distanceCoefficient*f.distance + f.timeCoefficient*f.duration

	f.platformFee = 2.0
	f.weightBasedFee = 2.0

	f.finalFare = f.nsnsFare + f.platformFee + f.weightBasedFee + f.numStopSurcharge
	return f
}

type fareWithFormula struct {
	numPickUp           float.Simple
	perPickUpPrice      float.Simple
	numDropOff          float.Simple
	perDropOffPrice     float.Simple
	numStopSurcharge    float.Simple
	baseFare            float.Simple
	distanceCoefficient float.Simple
	distance            float.Simple
	timeCoefficient     float.Simple
	duration            float.Simple
	nsnsFare            float.Simple
	platformFee         float.Simple
	weightBasedFee      float.Simple
	finalFare           float.Simple
}

func formulaCalculation() fareWithFormula {
	f := fareWithFormula{}

	f.numPickUp = float.MakeSimpleFromFloat("NumPickUp", 1)
	f.perPickUpPrice = float.MakeSimpleFromFloat("PerPickUpPrice", 3)
	f.numDropOff = float.MakeSimpleFromFloat("NumDropOff", 3)
	f.perDropOffPrice = float.MakeSimpleFromFloat("PerDropOffPrice", 2)

	acal.SourceRequest.Apply(f.numPickUp, f.numDropOff)
	acal.SourceHardcode.Apply(f.perPickUpPrice, f.perDropOffPrice)

	f.numStopSurcharge = f.numPickUp.Multiply(f.perPickUpPrice).Plus(f.numDropOff.Multiply(f.perDropOffPrice)).Anchor("numStopSurcharge")

	f.baseFare = float.MakeSimpleFromFloat("BaseFare", 1)
	f.distanceCoefficient = float.MakeSimpleFromFloat("DistanceCoefficient", 2)
	f.distance = float.MakeSimpleFromFloat("Distance", 2)
	f.timeCoefficient = float.MakeSimpleFromFloat("TimeCoefficient", 1)
	f.duration = float.MakeSimpleFromFloat("Duration", 1)

	acal.SourceRequest.Apply(f.distance, f.duration)
	acal.SourceHardcode.Apply(f.baseFare, f.distanceCoefficient, f.timeCoefficient)

	f.nsnsFare = f.baseFare.Plus(f.distanceCoefficient.Multiply(f.distance)).Plus(f.timeCoefficient.Multiply(f.duration)).Anchor("nsnsFare")

	f.platformFee = float.MakeSimpleFromFloat("PlatformFee", 2)
	f.weightBasedFee = float.MakeSimpleFromFloat("WeightBasedFee", 2)
	acal.SourceHardcode.Apply(f.platformFee, f.weightBasedFee)

	f.finalFare = f.nsnsFare.Plus(f.platformFee).Plus(f.weightBasedFee).Plus(f.numStopSurcharge).Anchor("finalFare")
	return f
}

type complexFareWithFormula struct {
	NumPickUp           float.Simple
	PerPickUpPrice      float.Simple
	NumDropOff          float.Simple
	PerDropOffPrice     float.Simple
	NumStopSurcharge    float.Simple
	BaseFare            float.Simple
	DistanceCoefficient float.Simple
	Distance            float.Simple
	TimeCoefficient     float.Simple
	Duration            float.Simple
	NsnsFare            float.Simple
	PlatformFee         float.Simple
	WeightBasedFee      float.Simple
	FinalFare           float.Progressive
}

func complexFareCalculation() complexFareWithFormula {
	f := complexFareWithFormula{}

	f.NumPickUp = float.MakeSimpleFromFloat("NumPickUp", 1)
	f.PerPickUpPrice = float.MakeSimpleFromFloat("PerPickUpPrice", 3)
	f.NumDropOff = float.MakeSimpleFromFloat("NumDropOff", 3)
	f.PerDropOffPrice = float.MakeSimpleFromFloat("PerDropOffPrice", 2)

	acal.SourceRequest.Apply(f.NumPickUp, f.NumDropOff)
	acal.SourceHardcode.Apply(f.PerPickUpPrice, f.PerDropOffPrice)

	f.NumStopSurcharge = f.NumPickUp.Multiply(f.PerPickUpPrice).Plus(f.NumDropOff.Multiply(f.PerDropOffPrice)).Anchor("NumStopSurcharge")

	f.BaseFare = float.MakeSimpleFromFloat("BaseFare", 1)
	f.DistanceCoefficient = float.MakeSimpleFromFloat("DistanceCoefficient", 2)
	f.Distance = float.MakeSimpleFromFloat("Distance", 2)
	f.TimeCoefficient = float.MakeSimpleFromFloat("TimeCoefficient", 1)
	f.Duration = float.MakeSimpleFromFloat("Duration", 1)

	acal.SourceRequest.Apply(f.Distance, f.Duration)
	acal.SourceHardcode.Apply(f.BaseFare, f.DistanceCoefficient, f.TimeCoefficient)

	f.NsnsFare = f.BaseFare.Plus(f.DistanceCoefficient.Multiply(f.Distance)).Plus(f.TimeCoefficient.Multiply(f.Duration)).Anchor("NsnsFare")

	f.FinalFare = float.MakeProgressive("FinalFare")

	f.FinalFare.Update(f.NsnsFare.Plus(f.NumStopSurcharge))

	boolean.If(
		boolean.MakeSimple("IsPlatformFeeEnabled", true), func(criteria boolean.Interface) {
			defer acal.ApplyProgressiveCondition(criteria, f.FinalFare)()

			PlatformFee := float.MakeSimpleFromFloat("PlatformFee", 2)
			PlatformFee.From(acal.SourceHardcode)

			f.FinalFare.Update(f.FinalFare.Plus(PlatformFee))

			acal.ApplyStaticCondition(criteria, PlatformFee)
		},
	)

	f.WeightBasedFee = float.MakeSimpleFromFloat("WeightBasedFee", 2)
	f.WeightBasedFee.From(acal.SourceHardcode)

	f.FinalFare.Update(f.FinalFare.Plus(f.WeightBasedFee))

	return f
}

func simpleCalculation() {
	BaseFare := float.MakeSimpleFromFloat("BaseFare", 1)

	FinalFare := float.MakeProgressive("FinalFare")

	FinalFare.Update(BaseFare)

	boolean.If(
		boolean.MakeSimple("IsPlatformFeeEnabled", true), func(criteria boolean.Interface) {
			defer acal.ApplyProgressiveCondition(criteria, FinalFare)()

			PlatformFee := float.MakeSimpleFromFloat("PlatformFee", 2)
			PlatformFee.From(acal.SourceHardcode)

			FinalFare.Update(FinalFare.Plus(PlatformFee))

			acal.ApplyStaticCondition(criteria, PlatformFee)
		},
	)
}

func testCondition() {
	BaseFare := float.MakeSimpleFromFloat("BaseFare", 1)

	FinalFare := float.MakeProgressive("FinalFare")

	FinalFare.Update(BaseFare)

	boolean.If(
		boolean.MakeSimple("IsPlatformFeeEnabled", true), func(criteria boolean.Interface) {
			defer acal.ApplyProgressiveCondition(criteria, FinalFare)()

			PlatformFee := float.MakeSimpleFromFloat("PlatformFee", 2)
			PlatformFee.From(acal.SourceHardcode)

			FinalFare.Update(FinalFare.Plus(PlatformFee))
		},
	)

	boolean.If(
		boolean.MakeSimple("WeightBaseFeeEnabled", true), func(criteria boolean.Interface) {
			defer acal.ApplyProgressiveCondition(criteria, FinalFare, FinalFare)()

			WeightBaseFee := float.MakeSimpleFromFloat("WeightBaseFee", 2)
			WeightBaseFee.From(acal.SourceHardcode)

			FinalFare.Update(FinalFare.Plus(WeightBaseFee))
		},
	)
}

func BenchmarkCondition(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testCondition()
	}
}

func Test_doFormulaCalculation(t *testing.T) {
	testCondition()
}

func BenchmarkCalculationUsingPrimitive(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		bareCalculation()
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
