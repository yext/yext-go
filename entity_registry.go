package yext

import (
	"encoding/json"
	"fmt"
)

func defaultEntityRegistry() Registry {
	registry := make(Registry)
	registry.Register(string(ENTITYTYPE_LOCATION), &Location{})
	registry.Register(string(ENTITYTYPE_EVENT), &Event{})
	return registry
}

func toEntityTypes(entities []interface{}, registry Registry) ([]Entity, error) {
	var types = []Entity{}
	for _, entityInterface := range entities {
		entity, err := toEntityType(entityInterface, registry)
		if err != nil {
			return nil, err
		}
		types = append(types, entity)
	}
	return types, nil
}

func toEntityType(entity interface{}, registry Registry) (Entity, error) {
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

	entityObj, err := registry.Create(entityType.(string))
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
