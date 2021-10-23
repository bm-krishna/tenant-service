package models

import "time"

type Tenant struct {
	Name string
	// dob               time.Time
	LoanAmount float32
	SSN        int32
	LoanCount  int16
	// compnayLeagalName string
	// companyAddress    Address
	// person            Person
}

type Person struct {
	firstName   string
	lastName    string
	dob         time.Time
	ssn         int32
	homeAddress Address
	email       string
}
type Address struct {
	line1   string
	line2   string
	zipcode int32
	state   string
	country string
}
