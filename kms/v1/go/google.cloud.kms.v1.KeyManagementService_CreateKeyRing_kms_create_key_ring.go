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

	kms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

// [START kms_create_key_ring]

func sampleCreateKeyRing(arg0 string) error {
	ctx := context.Background()
	c, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return err
	}

	// arg0 := "the id"
	req := &kmspb.CreateKeyRingRequest{
		KeyRingId: arg0,
	}
	resp, err := c.CreateKeyRing(ctx, req)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %v\n", resp.GetName())
	return nil
}

// [END kms_create_key_ring]

func main() {
	arg0 := flag.String("arg0", "the id", "")
	flag.Parse()
	if err := sampleCreateKeyRing(*arg0); err != nil {
		log.Fatal(err)
	}
}
