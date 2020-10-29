// Copyright 2020 OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package datareceivers

import (
	"context"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/testbed/testbed"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/mockdatareceivers/mockawsxrayreceiver"
)

// MockAwsXrayDataReceiver implements AwsXray format receiver.
type MockAwsXrayDataReceiver struct {
	testbed.DataReceiverBase
	receiver component.TracesReceiver
}

// NewMockAwsXrayDataReceiver creates a new  MockDataReceiver
func NewMockAwsXrayDataReceiver(port int) *MockAwsXrayDataReceiver {
	return &MockAwsXrayDataReceiver{DataReceiverBase: testbed.DataReceiverBase{Port: port}}
}

//Start listening on the specified port
func (ar *MockAwsXrayDataReceiver) Start(tc consumer.TracesConsumer, _ consumer.MetricsConsumer, _ consumer.LogsConsumer) error {
	var err error
	os.Setenv("AWS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY")

	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	certs, err := ioutil.ReadFile("../mockdatareceivers/mockawsxrayreceiver/server.crt")

	if err != nil {
		log.Fatalf("Failed to append %q to RootCAs: %v", "../mockdatareceivers/mockawsxrayreceiver/server.crt", err)
	}

	// Append our cert to the system pool
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Println("No certs appended, using system certs only")
	}

	mockDatareceiverCFG := mockawsxrayreceiver.Config{
		Endpoint: fmt.Sprintf("localhost:%d", ar.Port),
		TLSCredentials: &configtls.TLSSetting{
			CertFile: "../mockdatareceivers/mockawsxrayreceiver/server.crt",
			KeyFile:  "../mockdatareceivers/mockawsxrayreceiver/server.key",
		},
	}
	params := component.ReceiverCreateParams{Logger: zap.L()}
	ar.receiver, err = mockawsxrayreceiver.New(tc, params, &mockDatareceiverCFG)

	if err != nil {
		return err
	}

	return ar.receiver.Start(context.Background(), ar)
}

func (ar *MockAwsXrayDataReceiver) Stop() error {
	ar.receiver.Shutdown(context.Background())
	return nil
}

func (ar *MockAwsXrayDataReceiver) GenConfigYAMLStr() string {
	// Note that this generates an exporter config for agent.
	return fmt.Sprintf(`
  awsxray:
    local_mode: true
    endpoint: localhost:%d
    no_verify_ssl: true
    region: us-west-2`, ar.Port)
}

func (ar *MockAwsXrayDataReceiver) ProtocolName() string {
	return "awsxray"
}
