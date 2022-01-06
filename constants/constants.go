package constants

const (
	TaxRateForRAW                                    = 12.50
	ImportDuty                                       = 10.00
	ImportDutyLimit1                                 = 100
	ImportDutyLimit2                                 = 200
	SurchargeAmountForFinalCostUptoImportDutyLimit1  = 5
	SurchargeAmountForFinalCostUptoImportDutyLimit2  = 10
	SurchargeRateForFinalCostExceedeImportDutyLimit2 = 5.00
	TaxRateForManufacturedItemOnItemCost             = 12.50
	//Combined =ItemCost +12.5% Item Cost
	TaxRateForManufactureItemOnCombined = 2.00
	Accept                              = "y"
	Deny                                = "n"
)
