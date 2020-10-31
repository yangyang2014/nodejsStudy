package main
import (
	"fmt"
)

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError)Error() string {

	strFormat := 'dividee: %d divider: 0'
	return fmt.Sprintf(strFormat,de.dividee)
	
}

func Divide(varDividee int, varDivider int)(result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError {
			divicee: varDividee,
			divider：varDivider，
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee/varDivider,""
	}
}

func main() {
	if result,errorMsg := Divide(100,10);errorMsg = "" {
		fmt.Println("100/10 = ",result)
	}
	if result,errorMsg := Divide(100,0);errorMsg = "" {
		fmt.Println("100/0 = ",errorMsg)
	}
	

}
