package structures


func CalculateTax(income int) (float64, error){
	tax := 0.0
	prevSlabValue := 0

	for _, slab := range taxSlabs{
		if prevSlabValue <= income && income < slab.SlabValue{
			tax += (float64)(slab.TaxValue/100) * (float64)(income - prevSlabValue)
			return tax, nil
		}else{
			tax += (float64)(slab.TaxValue/100) * (float64)(slab.SlabValue - prevSlabValue)
		}
		prevSlabValue = slab.SlabValue
	}
	return tax, nil
}


var taxSlabs = []struct{
	SlabValue int
	TaxValue float64
}{
	{
		2200,
		0,
	}, {
		2700,
		14,
	},{
		3200,
		15,
	},{
		3700,
		16,
	}, {
		4200,
		17,
	}, {
		102200,
		70,
	},
}