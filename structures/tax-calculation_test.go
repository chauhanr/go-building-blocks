package structures

import "testing"

var taxCalculationTest = []struct{
	income int
	taxValue int64
}{
	{
	    2100,
	     0,
	},{
		2400,
		28,
	},{
	     2900,
	      100,
	},{
		2500,
		42,
	}, {
	      3300,
	      161,
	},
}

func TestCalculateTax(t *testing.T) {

	for _, testCase := range taxCalculationTest{
		income := testCase.income
		tax, err:= CalculateTax(income)
		if err != nil {
			t.Errorf("There was an error calculating tax %s ", err.Error())
		}else{
			intTax := int64(tax)
			if testCase.taxValue != intTax{
				t.Errorf("Tax value should be %f but was calculated as %f for income %d", testCase.taxValue, tax, income)
			}else{
				t.Logf("Tax successfully calculated as %f for income %d", tax, income)
			}
		}
	}
}
