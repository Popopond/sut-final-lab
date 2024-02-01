package entity_test

import (
	"testing"
	"github.com/Popopond/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomersAge (t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`Age is invalid`, func(t *testing.T){
		cus := entity.Customers{
			Name:"Piwpiw",
			Age :"",
			CustomerID :"CM12345678",
		}
		ok, err := govalidator.ValidateStruct(cus)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Age is not integer"))
	})
}