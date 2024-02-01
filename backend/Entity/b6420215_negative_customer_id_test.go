package entity_test

import (
	"testing"
	"github.com/Popopond/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomersID (t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`ID is invalid`, func(t *testing.T){
		cus := entity.Customers{
			Name:"Piwpiw",
			Age :21,
			CustomerID :"0012345678",
		}
		ok, err := govalidator.ValidateStruct(cus)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("ID is invalid"))
	})
}