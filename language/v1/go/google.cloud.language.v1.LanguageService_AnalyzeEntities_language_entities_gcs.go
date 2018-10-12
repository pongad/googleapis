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

	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

// [START language_entities_gcs]

func sampleAnalyzeEntities(arg0 string) error {
	ctx := context.Background()
	c, err := language.NewClient(ctx)
	if err != nil {
		return err
	}

	// arg0 := "Path to text file, e.g. gs://my-bucket/textfile"
	req := &languagepb.AnalyzeEntitiesRequest{
		Document: &languagepb.Document{
			Type: languagepb.Document_PLAIN_TEXT,
			Source: &languagepb.Document_GcsContentUri{
				GcsContentUri: arg0,
			},
		},
	}
	resp, err := c.AnalyzeEntities(ctx, req)
	if err != nil {
		return err
	}

	entities := resp.GetEntities()
	for _, entity := range entities {
		fmt.Printf("Entity name: %v\n", entity.GetName())
		fmt.Printf("Entity type: %v\n", entity.GetType())
	}
	return nil
}

// [END language_entities_gcs]

func main() {
	arg0 := flag.String("arg0", "Path to text file, e.g. gs://my-bucket/textfile", "")
	flag.Parse()
	if err := sampleAnalyzeEntities(*arg0); err != nil {
		log.Fatal(err)
	}
}
