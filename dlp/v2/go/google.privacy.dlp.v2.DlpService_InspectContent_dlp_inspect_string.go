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

	dlp "cloud.google.com/go/dlp/apiv2"
	dlppb "google.golang.org/genproto/googleapis/privacy/dlp/v2"
)

// [START dlp_inspect_string]

func sampleInspectContent(value string) error {
	ctx := context.Background()
	c, err := dlp.NewClient(ctx)
	if err != nil {
		return err
	}

	// value := "My name is Gary and my email is gary@example.com"
	req := &dlppb.InspectContentRequest{
		Item: &dlppb.ContentItem{
			DataItem: &dlppb.ContentItem_Value{
				Value: value,
			},
		},
		InspectConfig: &dlppb.InspectConfig{
			InfoTypes: []*dlppb.InfoType{
				&dlppb.InfoType{
					Name: "PHONE_NUMBER",
				},
				&dlppb.InfoType{
					Name: "EMAIL_ADDRESS",
				},
				&dlppb.InfoType{
					Name: "CREDIT_CARD_NUMBER",
				},
			},
			MinLikelihood: dlppb.Likelihood_LIKELIHOOD_UNSPECIFIED,
			IncludeQuote:  true,
			Limits: &dlppb.InspectConfig_FindingLimits{
				MaxFindingsPerRequest: 0,
			},
		},
	}
	resp, err := c.InspectContent(ctx, req)
	if err != nil {
		return err
	}

	findings := resp.GetResult().GetFindings()
	fmt.Printf("Findings:\n")
	for _, finding := range findings {
		fmt.Printf("Quote: %v\n", finding.GetQuote())
		fmt.Printf("Info type: %v\n", finding.GetInfoType().GetName())
		fmt.Printf("Likelihood: %v\n", finding.GetLikelihood())
	}
	return nil
}

// [END dlp_inspect_string]

func main() {
	value := flag.String("value", "My name is Gary and my email is gary@example.com", "")
	flag.Parse()
	if err := sampleInspectContent(*value); err != nil {
		log.Fatal(err)
	}
}
