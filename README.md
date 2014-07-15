# FluidSurveys API wrapper
***Golang wrapper for the FluidSurveys V3 API***

## Introduction

This library provides a pure go interface for the [FluidSurveys API](http://docs.fluidsurveys.com/#api-v3-documentation).

## Documentation

View the lastest release API documentation [FluidSurveys API](http://docs.fluidsurveys.com/#api-v3-documentation).

## Example
``` go
package main

import (
	"encoding/json"
	"github.com/collinglass/fluidsurveys-go/fluidsurveys"
	"log"
)

func main() {

	log.Println("Running..")

	fluidsurveys.Fluidsurveys("your_email", "your_pasword")
	data := map[string]interface{}{
		"name": "New Survey123",
	}

	// Create survey
	result, err := fluidsurveys.Create("surveys", data)
	if err != nil {
		log.Println(err)
	}
	// convert json to go and use it
```

## API
Entity types are puralized to match collections. You can interact with `templates`, `surveys`, `collectors`, `contacts`, `embeds`, `contact-lists`, & `webhooks`.

### Create( entityType, data )
Creates a new entity, such as `surveys`.

### CreateChild ( parentType, parentId, childType, data )
Creates a child entity under a parent, such as `collectors` under `surveys`.

### Delete( entityType, id )
Deletes an entity.

### Get( entityType, id )
Gets an entity.

### GetChild( parentType, parentId, childType, childId )
Gets a child entity under a parent, such as `collectors` under `surveys`.

### List( entityType, args )
Gets a paginated list of entities.

### ListChildren( parentType, parentId, childType, args )
Gets a paginated list of children under a parent.

### Update( entityType, id, data )
Updates an entity.

## Caveats
- Getting a 'survey' will retrieve the 'entity', and it's 'structure'
- Updating a 'survey' supports modifying the 'entity', and/or the 'structure'

## Issues

The V3 API, and this wrapper are 'beta'. Please report issues you find!

Using create() for a template will result in an error (API bug)
License

Copyright (c) 2014 Fluidware Licensed under the MIT license.