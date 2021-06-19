package yext

import (
	"encoding/json"
	"fmt"
)

type EntityRegistry Registry

func defaultEntityRegistry() *EntityRegistry {
	registry := make(Registry)
	registry.Register(string(ENTITYTYPE_LOCATION), &LocationEntity{})
	registry.Register(string(ENTITYTYPE_EVENT), &EventEntity{})
	registry.Register(string(ENTITYTYPE_RESTAURANT), &RestaurantEntity{})
	registry.Register(string(ENTITYTYPE_HEALTHCAREPROFESSIONAL), &HealthcareProfessionalEntity{})
	registry.Register(string(ENTITYTYPE_HEALTHCAREFACILITY), &HealthcareFacilityEntity{})
	registry.Register(string(ENTITYTYPE_ATM), &ATMEntity{})
	registry.Register(string(ENTITYTYPE_HOTEL), &HotelEntity{})
	entityRegistry := EntityRegistry(registry)
	return &entityRegistry
}

func (r *EntityRegistry) RegisterEntity(t EntityType, entity interface{}) {
	registry := Registry(*r)
	registry.Register(string(t), entity)
}

func (r *EntityRegistry) InitializeEntity(t EntityType) (Entity, error) {
	registry := Registry(*r)
	i, err := registry.Initialize(string(t))
	if err != nil {
		return nil, err
	}
	return i.(Entity), nil
}

func (r *EntityRegistry) ToEntityTypes(entities []interface{}) ([]Entity, error) {
	var types = []Entity{}
	for _, entityInterface := range entities {
		entity, err := r.ToEntityType(entityInterface)
		if err != nil {
			return nil, err
		}
		types = append(types, entity)
	}
	return types, nil
}

func (r *EntityRegistry) ToEntityType(entity interface{}) (Entity, error) {
	// Determine Entity Type
	var entityValsByKey = entity.(map[string]interface{})
	meta, ok := entityValsByKey["meta"]
	if !ok {
		return nil, fmt.Errorf("Unable to find meta attribute in %v\nFor Entity: %v", entityValsByKey, entity)
	}

	var metaByKey = meta.(map[string]interface{})
	entityType, ok := metaByKey["entityType"]
	if !ok {
		return nil, fmt.Errorf("Unable to find entityType attribute in %v\nFor Entity: %v", metaByKey, entity)
	}

	var registry = Registry(*r)
	entityObj, err := registry.Initialize(entityType.(string))
	if err != nil {
		// Unable to create an instace of entityType, use RawEntity instead
		entityObj = &RawEntity{}
	}

	// Convert into struct of Entity Type
	entityJSON, err := json.Marshal(entityValsByKey)
	if err != nil {
		return nil, fmt.Errorf("Marshaling entity to JSON: %s\nFor Entity: %v", err, entity)
	}

	err = json.Unmarshal(entityJSON, &entityObj)
	if err != nil {
		return nil, fmt.Errorf("Unmarshaling entity JSON: %s\nFor Entity: %v", err, entity)
	}
	return entityObj.(Entity), nil
}
