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
 * DO NOT EDIT! This is a generated sample ("Request",  "language_entities_text")
 */

// [START language_entities_text]

namespace Google\Cloud\Language\V1;

function sampleAnalyzeEntities($content)
{
    // [START language_entities_text_core]

    $languageServiceClient = new LanguageServiceClient();

    // $content = 'Your text to analyze, e.g. Hello, world!';
    $type = Document_Type::PLAIN_TEXT;
    $document = new Document();
    $document->setType($type);
    $document->setContent($content);

    try {
        $response = $languageServiceClient->analyzeEntities($document);
        $entities = $response->getEntities();
        foreach ($entities as $entity) {
            printf('Entity name: %s'.PHP_EOL, $entity->getName());
            printf('Entity type: %s'.PHP_EOL, $entity->getType());
        }
    } finally {
        $languageServiceClient->close();
    }

    // [END language_entities_text_core]
}
// [END language_entities_text]

$opts = [
    'content::',
];

$defaultOptions = [
    'content' => 'Your text to analyze, e.g. Hello, world!',
];

$options = getopt('', $opts);
$options += $defaultOptions;

$content = $options['content'];

sampleAnalyzeEntities($content);
