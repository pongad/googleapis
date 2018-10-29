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

// [START language_syntax_text]

func sampleAnalyzeSyntax(content string) error {
	ctx := context.Background()
	c, err := language.NewClient(ctx)
	if err != nil {
		return err
	}

	// content := "Your text to analyze, e.g. Hello, world!"
	req := &languagepb.AnalyzeSyntaxRequest{
		Document: &languagepb.Document{
			Type: languagepb.Document_PLAIN_TEXT,
			Source: &languagepb.Document_Content{
				Content: content,
			},
		},
	}
	resp, err := c.AnalyzeSyntax(ctx, req)
	if err != nil {
		return err
	}

	tokens := resp.GetTokens()
	for _, token := range tokens {
		fmt.Printf("Part of speech: %v\n", token.GetPartOfSpeech().GetTag())
		fmt.Printf("Text:\n", token.GetText())
	}
	return nil
}

// [END language_syntax_text]

func main() {
	content := flag.String("content", "Your text to analyze, e.g. Hello, world!", "")
	flag.Parse()
	if err := sampleAnalyzeSyntax(*content); err != nil {
		log.Fatal(err)
	}
}
