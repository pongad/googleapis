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
 * DO NOT EDIT! This is a generated sample ("Request",  "speech_transcribe_sync_gcs")
 */

// [START speech_transcribe_sync_gcs]

namespace Google\Cloud\Speech\V1;

function sampleRecognize($uri)
{
    // [START speech_transcribe_sync_gcs_core]

    $speechClient = new SpeechClient();

    // $uri = 'Path to audio file in GCS, e.g. gs://my-bucket/audio.wav';
    $encoding = RecognitionConfig_AudioEncoding::LINEAR16;
    $sampleRateHertz = 16000;
    $languageCode = 'en-US';
    $config = new RecognitionConfig();
    $config->setEncoding($encoding);
    $config->setSampleRateHertz($sampleRateHertz);
    $config->setLanguageCode($languageCode);
    $audio = new RecognitionAudio();
    $audio->setUri($uri);

    try {
        $response = $speechClient->recognize($config, $audio);
        foreach ($response->getResults() as $result) {
            foreach ($result->getAlternatives() as $alternative) {
                printf('Transcript: %s'.PHP_EOL, $alternative->getTranscript());
            }
        }
    } finally {
        $speechClient->close();
    }

    // [END speech_transcribe_sync_gcs_core]
}
// [END speech_transcribe_sync_gcs]

$opts = [
    'uri::',
];

$defaultOptions = [
    'uri' => 'Path to audio file in GCS, e.g. gs://my-bucket/audio.wav',
];

$options = getopt('', $opts);
$options += $defaultOptions;

$uri = $options['uri'];

sampleRecognize($uri);
