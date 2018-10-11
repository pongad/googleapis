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
 * DO NOT EDIT! This is a generated sample ("Request",  "automl_vision_create_dataset")
 */

// [START automl_vision_create_dataset]

namespace Google\Cloud\Automl\V1beta1;

function sampleCreateDataset($displayName)
{
    // [START automl_vision_create_dataset_core]

    $autoMlClient = new AutoMlClient();

    // $displayName = 'Your Dataset display name';
    $formattedParent = $autoMlClient->locationName('[PROJECT]', '[LOCATION]');
    $dataset = new Dataset();
    $dataset->setDisplayName($displayName);

    try {
        $response = $autoMlClient->createDataset($formattedParent, $dataset);
        $dataset = $response;
        printf('Dataset name: %s'.PHP_EOL, $dataset->getName());
        printf('Dataset display name: %s'.PHP_EOL, $dataset->getDisplayName());
        printf('Image classification dataset metadata:'.PHP_EOL);
        printf('	%s'.PHP_EOL, $dataset->getImageClassificationDatasetMetadata());
        printf('Dataset example count: %s'.PHP_EOL, $dataset->getExampleCount());
        printf('Dataset create time:'.PHP_EOL);
        printf('	seconds: %s'.PHP_EOL, $dataset->getCreateTime()->getSeconds());
        printf('	nanos: %s'.PHP_EOL, $dataset->getCreateTime()->getNanos());
    } finally {
        $autoMlClient->close();
    }

    // [END automl_vision_create_dataset_core]
}
// [END automl_vision_create_dataset]

$opts = [
    'displayName::',
];

$defaultOptions = [
    'displayName' => 'Your Dataset display name',
];

$options = getopt('', $opts);
$options += $defaultOptions;

$displayName = $options['displayName'];

sampleCreateDataset($displayName);
