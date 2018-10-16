package main

import (
	"log"

	yext "gopkg.in/yext/yext-go.v2"
)

type MorganLocation struct {
	yext.LocationEntity
	CMultiText *string `json:"c_multiText"`
}

type HealthcareProfessional struct {
	yext.Location
}

type PersonEntity interface {
	MyType() string
}

type Teacher struct {
	Person
	Blah string
}

func (t *Teacher) MyType() string {
	return "Teacher"
}

type Person struct {
	Name string
}

func (p *Person) MyType() string {
	return "Person"
}

func whatAmI(it interface{}) {

	if _, ok := it.(*Teacher); ok {
		log.Println("is a teacher")
	}
	if _, ok := it.(*Person); ok {
		log.Println("is a person")
	}
}

func main() {
	// teacher := &Teacher{
	// 	Person: Person{
	// 		Name: "Catherine",
	// 	},
	// 	Blah: "Blah",
	// }
	// whatAmI(teacher)
	client := yext.NewClient(yext.NewDefaultConfig().WithApiKey("e929153c956b051cea51ec289bfd2383"))
	client.EntityService.RegisterEntity(yext.ENTITYTYPE_LOCATION, &MorganLocation{})
	client.EntityService.RegisterEntity("HEALTHCARE_PROFESSIONAL", &HealthcareProfessional{})
	client.EntityService.RegisterEntity("HEALTHCARE_FACILITY", &HealthcareProfessional{})
	client.EntityService.RegisterEntity("ATM", &HealthcareProfessional{})
	// entities, err := client.EntityService.ListAll(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// log.Printf("ListAll: Got %d entities", len(entities))
	// for _, entity := range entities {
	// 	switch entity.(type) {
	// 	case *MorganLocation:
	// 		log.Println("I am a Morgan Location")
	// 		morganEntity := entity.(*MorganLocation)
	// 		log.Printf("PROFILE description: %s", GetString(morganEntity.Description))
	// 		log.Printf("CUSTOM multi: %s", GetString(morganEntity.CMultiText))
	// 	}
	// }

	entity, _, err := client.EntityService.Get("CTG")
	if err != nil {
		log.Fatal(err)
	}
	morganEntity := entity.(*MorganLocation)
	log.Printf("Get: Got %s", morganEntity.GetName())

	// morganEntity.Name = yext.String(morganEntity.GetName() + "- 1")
	// update := &MorganLocation{
	// 	Location: yext.Location{
	// 		// EntityMeta: yext.EntityMeta{
	// 		// 	Id: yext.String("CTG"),
	// 		// },
	// 		Name: yext.String("CTG"),
	// 	},
	// }
	// _, err = client.EntityService.Edit(update)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Edit: Edited %s", morganEntity.GetName())

	// morganEntity = &MorganLocation{
	// 	Location: yext.Location{
	// 		Name: yext.String("Yext Consulting"),
	// 		EntityMeta: &yext.EntityMeta{
	// 			//Id:         yext.String("CTG2"),
	// 			EntityType:  yext.ENTITYTYPE_LOCATION,
	// 			CategoryIds: &[]string{"2374", "708"},
	// 		},
	// 		MainPhone: yext.String("8888888888"),
	// Address: &yext.Address{
	// 	Line1:      yext.String("7900 Westpark"),
	// 	City:       yext.String("McLean"),
	// 	Region:     yext.String("VA"),
	// 	PostalCode: yext.String("22102"),
	// },

	// 		FeaturedMessage: &yext.FeaturedMessage{
	// 			Url:         yext.String("www.yext.com"),
	// 			Description: yext.String("Yext Consulting"),
	// 		},
	// 	},
	// }
	//
	// _, err = client.EntityService.Create(morganEntity)
	// if err != nil {
	// 	log.Fatalf("Create error: %s", err)
	// }
	// log.Printf("Create: Created %s", morganEntity.GetEntityId())

}

func GetString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
