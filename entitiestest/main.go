package main

import (
	"log"

	"github.com/mohae/deepcopy"
	yext "gopkg.in/yext/yext-go.v2"
)

type MorganLocation struct {
	yext.Location
	CMultiText *string `json:"c_multiText"`
}

//TODO: Revisit this...what if m.Location is nil?
func (m *MorganLocation) GetEntityId() string {
	return m.Location.GetEntityId()
}

func (m *MorganLocation) GetEntityType() yext.EntityType {
	return m.Location.GetEntityType()
}

func (m *MorganLocation) Copy() yext.Entity {
	return deepcopy.Copy(m).(*MorganLocation)
}

func main() {
	client := yext.NewClient(yext.NewDefaultConfig().WithApiKey("e929153c956b051cea51ec289bfd2383"))
	client.EntityService.RegisterEntity(yext.ENTITYTYPE_LOCATION, &MorganLocation{})
	// entities, err := client.EntityService.ListAll(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// log.Printf("ListAll: Got %d entities", len(entities))
	//for _, entity := range entities {
	//morganEntity := entity.(*MorganLocation)
	//log.Printf("PROFILE description: %s", GetString(morganEntity.Description))
	//log.Printf("CUSTOM multi: %s", GetString(morganEntity.CMultiText))
	//}

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

	morganEntity = &MorganLocation{
		Location: yext.Location{
			Name: yext.String("Yext Consulting"),
			EntityMeta: &yext.EntityMeta{
				//Id:         yext.String("CTG2"),
				EntityType:  yext.ENTITYTYPE_LOCATION,
				CategoryIds: &[]string{"2374", "708"},
			},
			MainPhone: yext.String("8888888888"),
			// Address: &yext.Address{
			// 	Line1:      yext.String("7900 Westpark"),
			// 	City:       yext.String("McLean"),
			// 	Region:     yext.String("VA"),
			// 	PostalCode: yext.String("22102"),
			// },

			FeaturedMessage: &yext.FeaturedMessage{
				Url:         yext.String("www.yext.com"),
				Description: yext.String("Yext Consulting"),
			},
		},
	}

	_, err = client.EntityService.Create(morganEntity)
	if err != nil {
		log.Fatalf("Create error: %s", err)
	}
	log.Printf("Create: Created %s", morganEntity.GetEntityId())

}

func GetString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
