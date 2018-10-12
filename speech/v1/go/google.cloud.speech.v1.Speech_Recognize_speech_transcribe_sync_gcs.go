// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

// [START speech_transcribe_sync_gcs]

func sampleRecognize(arg0 string) error {
	ctx := context.Background()
	c, err := speech.NewClient(ctx)
	if err != nil {
		return err
	}

	// arg0 := "Path to audio file in GCS, e.g. gs://my-bucket/audio.wav"
	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
			LanguageCode:    "en-US",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Uri{
				Uri: arg0,
			},
		},
	}
	resp, err := c.Recognize(ctx, req)
	if err != nil {
		return err
	}

	for _, result := range resp.GetResults() {
		for _, alternative := range result.GetAlternatives() {
			fmt.Printf("Transcript: %v\n", alternative.GetTranscript())
		}
	}
	return nil
}

// [END speech_transcribe_sync_gcs]

func main() {
	arg0 := flag.String("arg0", "Path to audio file in GCS, e.g. gs://my-bucket/audio.wav", "")
	flag.Parse()
	if err := sampleRecognize(*arg0); err != nil {
		log.Fatal(err)
	}
}
