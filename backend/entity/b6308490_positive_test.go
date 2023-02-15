package entity
import (
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
)

func Test_positive(t *testing.T){
	g := NewGomegaWithT(t)

	u := Employee{
		Name : "Pattanasak" ,    
		Email : "Putjat1454@gmail.com",   
		EmployeeID :"M123456",
	}
	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}