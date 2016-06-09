package main

import "time"
import "math/big"

type User struct {
	Id      string     `json:"id"`
	Start       time.Time  `json:"start"`
    End         time.Time  `json:"end"`
    HourlyRate  big.Rat    `json:"hourlyRate"`
    Salary      big.Rat    `json:"salary"`
    IntervalAmt big.Rat    `json:"intervalAmt"`
}
