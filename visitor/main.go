package main

import "fmt"

type Instrument interface {
	Accept(visitor InstrumentVisitor)
}

type InstrumentVisitor interface {
	VisitStock(s *Stock)
	VisitBond(b *Bond)
	VisitOption(o *Option)
}

type Stock struct {
	Symbol string
	Price  float64
	Volume int
}

func (s *Stock) Accept(visitor InstrumentVisitor) {
	visitor.VisitStock(s)
}

type Bond struct {
	ID              string
	FaceValue       float64
	CouponRate      float64
	YearsToMaturity int
}

func (b *Bond) Accept(visitor InstrumentVisitor) {
	visitor.VisitBond(b)

}

type Option struct {
	ContractID     string
	Type           string
	StrikePrice    float64
	ExpirationDate string
}

func (o *Option) Accept(visitor InstrumentVisitor) {
	visitor.VisitOption(o)
}

type ValuationVisitor struct {
	TotalValue float64
}

func (v *ValuationVisitor) VisitStock(s *Stock) {
	value := s.Price * float64(s.Volume)
	fmt.Printf("Calculating Value for Stock %s: %.2f\n", s.Symbol, value)
	v.TotalValue += value
}

func (v *ValuationVisitor) VisitBond(b *Bond) {
	value := b.FaceValue * (1 + b.CouponRate*float64(b.YearsToMaturity))
	fmt.Printf("Calculating Value for Bond %s: %.2f (simplified)\n", b.ID, value)
	v.TotalValue += value
}

func (v *ValuationVisitor) VisitOption(o *Option) {
	fmt.Printf("Calculating Value for Option %s (%s): Requires complex model (e.g., Black-Scholes).\n", o.ContractID, o.Type)

}

type RiskAssessmentVisitor struct {
	OverallRiskScore float64
}

func (r *RiskAssessmentVisitor) VisitStock(s *Stock) {
	risk := float64(s.Volume) * 0.1
	fmt.Printf("Assessing Risk for Stock %s: Risk Score = %.2f\n", s.Symbol, risk)
	r.OverallRiskScore += risk
}

func (r *RiskAssessmentVisitor) VisitBond(b *Bond) {
	risk := float64(b.YearsToMaturity) * 0.5
	fmt.Printf("Assessing Risk for Bond %s: Risk Score = %.2f\n", b.ID, risk)
	r.OverallRiskScore += risk
}

func (r *RiskAssessmentVisitor) VisitOption(o *Option) {
	risk := 5.0
	fmt.Printf("Assessing Risk for Option %s (%s): Risk Score = %.2f (High Leverage)\n", o.ContractID, o.Type, risk)
	r.OverallRiskScore += risk
}

type ReportingVisitor struct{}

func (rp *ReportingVisitor) VisitStock(s *Stock) {
	fmt.Printf("Report Item: Stock - Symbol: %s, Price: %.2f, Volume: %d\n", s.Symbol, s.Price, s.Volume)
}

func (rp *ReportingVisitor) VisitBond(b *Bond) {
	fmt.Printf("Report Item: Bond - ID: %s, Face Value: %.2f, Coupon: %.2f%%, Maturity: %d yrs\n", b.ID, b.FaceValue, b.CouponRate*100, b.YearsToMaturity)
}

func (rp *ReportingVisitor) VisitOption(o *Option) {
	fmt.Printf("Report Item: Option - Contract ID: %s, Type: %s, Strike: %.2f, Expiry: %s\n", o.ContractID, o.Type, o.StrikePrice, o.ExpirationDate)
}

func main() {
	portfolio := []Instrument{
		&Stock{Symbol: "AAPL", Price: 170.0, Volume: 100},
		&Bond{ID: "CORP-B-001", FaceValue: 1000.0, CouponRate: 0.05, YearsToMaturity: 10},
		&Option{ContractID: "AAPL-C-200-JUL25", Type: "Call", StrikePrice: 200.0, ExpirationDate: "2025-07-18"},
		&Stock{Symbol: "GOOG", Price: 1800.0, Volume: 50},
		&Bond{ID: "GOV-B-002", FaceValue: 5000.0, CouponRate: 0.03, YearsToMaturity: 5},
	}

	fmt.Println("--- Running Valuation ---")
	valuationVisitor := &ValuationVisitor{}
	for _, instrument := range portfolio {
		instrument.Accept(valuationVisitor)
	}
	fmt.Printf("Total Portfolio Value (simplified): %.2f\n", valuationVisitor.TotalValue)

	fmt.Println("\n--- Running Risk Assessment ---")
	riskVisitor := &RiskAssessmentVisitor{}
	for _, instrument := range portfolio {
		instrument.Accept(riskVisitor)
	}
	fmt.Printf("Overall Portfolio Risk Score (simplified): %.2f\n", riskVisitor.OverallRiskScore)

	fmt.Println("\n--- Running Reporting ---")
	reportingVisitor := &ReportingVisitor{}
	for _, instrument := range portfolio {
		instrument.Accept(reportingVisitor)
	}
}
