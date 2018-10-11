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
 * DO NOT EDIT! This is a generated sample ("Request",  "kms_create_key_ring")
 */

// [START kms_create_key_ring]

namespace Google\Cloud\Kms\V1;

function sampleCreateKeyRing($keyRingId)
{
    // [START kms_create_key_ring_core]

    $keyManagementServiceClient = new KeyManagementServiceClient();

    // $keyRingId = '';
    $formattedParent = $keyManagementServiceClient->locationName('[PROJECT]', '[LOCATION]');
    $keyRing = new KeyRing();

    try {
        $response = $keyManagementServiceClient->createKeyRing($formattedParent, $keyRingId, $keyRing);
        printf('Name: %s'.PHP_EOL, $response->getName());
    } finally {
        $keyManagementServiceClient->close();
    }

    // [END kms_create_key_ring_core]
}
// [END kms_create_key_ring]

$opts = [
    'keyRingId::',
];

$defaultOptions = [
    'keyRingId' => '',
];

$options = getopt('', $opts);
$options += $defaultOptions;

$keyRingId = $options['keyRingId'];

sampleCreateKeyRing($keyRingId);
