package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_negative_name_test(t *testing.T) {
	g := NewGomegaWithT(t)

	u := Employee{
		Name:       "",
		Email:      "Putjat1454@gmail.com",
		EmployeeID: "M123456",
	}
	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("name not blank"))
}
