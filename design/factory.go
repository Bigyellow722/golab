package design


type OperationFactory interface {

	GetResult(num1 float64,num2 float64)float64
}

type AddOperation struct {
	OperationFactory
}

func (opf *AddOperation)GetResult(num1 float64,num2 float64) float64{
	var res float64 = 0
	switch  {
	case "+":
		res = num1+num2
	case "-":
		res = num1-num2
	case "*":
		res = num1*num2
	case "/":
		if num2==0{
			panic("")
		}



	}
	return res
}

func main(){
	oper Operation
	oper =
}