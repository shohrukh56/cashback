package main

func main()  {

}
func cashback(amount int ) int {
	//если сумма больще 3000 ттогда кешбэк 5% от покупки
	const bound  = 3000  // граница с которой начисляется кешбэе
	const percent  = 5  //
	result :=0
	if amount >= bound{
		result =amount*percent/100
	}
	return result

}
