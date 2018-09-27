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
func (m *MorganLocation) EntityId() string {
	return m.Location.EntityId()
}

func (m *MorganLocation) Type() yext.EntityType {
	return m.Location.Type()
}

func (m *MorganLocation) PathName() string {
	return m.Location.PathName()
}

func (m *MorganLocation) Copy() yext.Entity {
	return deepcopy.Copy(m).(*MorganLocation)
}

func main() {
	client := yext.NewClient(yext.NewDefaultConfig().WithApiKey("e929153c956b051cea51ec289bfd2383"))
	client.EntityService.RegisterEntity(yext.ENTITYTYPE_LOCATION, &MorganLocation{})
	entities, err := client.EntityService.ListAll(nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ListAll: Got %d entities", len(entities))
	for _, entity := range entities {

		morganEntity := entity.(*MorganLocation)
		log.Printf("PROFILE description: %s", GetString(morganEntity.Description))
		log.Printf("CUSTOM multi: %s", GetString(morganEntity.CMultiText))
	}

}

func GetString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
