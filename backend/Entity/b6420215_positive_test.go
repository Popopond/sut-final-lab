package entity_test

import (
	"testing"
	"github.com/Popopond/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomers (t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`Correct`, func(t *testing.T){
		cus := entity.Customers{
			Name:"Piwpiw",
			Age :21,
			CustomerID :"CM12345678",
		}
		ok, err := govalidator.ValidateStruct(cus)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}