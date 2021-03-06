// Copyright 2021 Google LLC
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

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package functions_test

import (
	"context"

	functions "cloud.google.com/go/functions/apiv1"
	"google.golang.org/api/iterator"
	functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func ExampleNewCloudFunctionsClient() {
	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleCloudFunctionsClient_ListFunctions() {
	// import functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &functionspb.ListFunctionsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListFunctions(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleCloudFunctionsClient_GetFunction() {
	// import functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &functionspb.GetFunctionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetFunction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudFunctionsClient_CreateFunction() {
	// import functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &functionspb.CreateFunctionRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateFunction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudFunctionsClient_UpdateFunction() {
	// import functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &functionspb.UpdateFunctionRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateFunction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudFunctionsClient_DeleteFunction() {
	// import functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &functionspb.DeleteFunctionRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteFunction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCloudFunctionsClient_CallFunction() {
	// import functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &functionspb.CallFunctionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CallFunction(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudFunctionsClient_GenerateUploadUrl() {
	// import functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &functionspb.GenerateUploadUrlRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GenerateUploadUrl(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudFunctionsClient_GenerateDownloadUrl() {
	// import functionspb "google.golang.org/genproto/googleapis/cloud/functions/v1"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &functionspb.GenerateDownloadUrlRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GenerateDownloadUrl(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudFunctionsClient_SetIamPolicy() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudFunctionsClient_GetIamPolicy() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCloudFunctionsClient_TestIamPermissions() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := functions.NewCloudFunctionsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
