<?php
/*
 * Copyright 2018 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * DO NOT EDIT! This is a generated sample ("Request",  "dialogflow_create_entity_type")
 */

// [START dialogflow_create_entity_type]

namespace Google\Cloud\Dialogflow\V2;

function sampleCreateEntityType()
{
    // [START dialogflow_create_entity_type_core]

    $entityTypesClient = new EntityTypesClient();

    $formattedParent = $entityTypesClient->projectAgentName('[PROJECT]');
    $kind = EntityType_Kind::KIND_LIST;
    $entityType = new EntityType();
    $entityType->setKind($kind);

    try {
        $response = $entityTypesClient->createEntityType($formattedParent, $entityType);
        printf('Entity type created with display name: %s'.PHP_EOL, $response->getDisplayName());
    } finally {
        $entityTypesClient->close();
    }

    // [END dialogflow_create_entity_type_core]
}
// [END dialogflow_create_entity_type]

sampleCreateEntityType();
