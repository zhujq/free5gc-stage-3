/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository_test

import (
	"context"
	"github.com/antihax/optional"
	"github.com/davecgh/go-spew/spew"
	"free5gc/lib/Nudr_DataRepository"
	"free5gc/lib/openapi/models"
	"free5gc/src/udr/udr_context"
	"net/http"
	"testing"
)

// SdmSubscriptions - Create, Query Update and Remove
func TestCreateQueryUpdateRemoveSdmSubscriptions(t *testing.T) {
	runTestServer(t)

	// Set client and set url
	client := setTestClient(t)

	var sdmSubscription = models.SdmSubscription{
		NfInstanceId:      udr_context.UDR_Self().NfId,
		CallbackReference: "http://127.0.0.1/callback",
		AmfServiceName:    models.ServiceName_NUDR_DR,
		SingleNssai: &models.Snssai{
			Sst: 1,
			Sd:  "010203",
		},
		Dnn: "internet",
		PlmnId: &models.PlmnId{
			Mnc: "208",
			Mcc: "93",
		},
	}

	ueId := "imsi-208930123456789"
	var subsUri string
	var subsId string
	_ = subsId

	// Create sdmSubscription
	{
		sdmSubscription, res, err := client.SDMSubscriptionsCollectionApi.CreateSdmSubscriptions(context.TODO(), ueId, sdmSubscription)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if status := res.StatusCode; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}

		subsUri = res.Header.Get("Location")
		spew.Printf("[subsUri_Header_Location] %s\n", subsUri)
		subsId = sdmSubscription.SubscriptionId
		spew.Printf("[subsId] %s\n", sdmSubscription.SubscriptionId)
	}

	var sdmSubscriptionUpdate = models.SdmSubscription{
		NfInstanceId:      udr_context.UDR_Self().NfId,
		CallbackReference: "http://127.0.0.1/callback_2",
		AmfServiceName:    models.ServiceName_NUDR_DR,
		SingleNssai: &models.Snssai{
			Sst: 1,
			Sd:  "010203",
		},
		Dnn: "internet",
		PlmnId: &models.PlmnId{
			Mnc: "208",
			Mcc: "93",
		},
	}

	// Query sdmSubscriptions
	{
		sdmSubscriptionArray, res, err := client.SDMSubscriptionsCollectionApi.Querysdmsubscriptions(context.TODO(), ueId, &Nudr_DataRepository.QuerysdmsubscriptionsParamOpts{})
		if err != nil {
			t.Fatalf(err.Error())
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		spew.Dump(sdmSubscriptionArray)
	}

	// Update sdmSubscriptions
	{
		res, err := client.SDMSubscriptionDocumentApi.Updatesdmsubscriptions(context.TODO(), ueId, subsId, &Nudr_DataRepository.UpdatesdmsubscriptionsParamOpts{
			SdmSubscription: optional.NewInterface(sdmSubscriptionUpdate),
		})
		if err != nil {
			t.Fatalf(err.Error())
		}

		if status := res.StatusCode; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}

	// Query sdmSubscriptions
	{
		sdmSubscriptionArray, res, err := client.SDMSubscriptionsCollectionApi.Querysdmsubscriptions(context.TODO(), ueId, &Nudr_DataRepository.QuerysdmsubscriptionsParamOpts{})
		if err != nil {
			t.Fatalf(err.Error())
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		spew.Dump(sdmSubscriptionArray)
	}

	// Delete sdmSubscriptions
	{
		res, err := client.SDMSubscriptionDocumentApi.RemovesdmSubscriptions(context.TODO(), ueId, subsId)
		if err != nil {
			t.Fatalf(err.Error())
		}

		if status := res.StatusCode; status != http.StatusNoContent {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNoContent)
		}
	}

	// Query sdmSubscriptions
	{
		sdmSubscriptionArray, res, err := client.SDMSubscriptionsCollectionApi.Querysdmsubscriptions(context.TODO(), ueId, &Nudr_DataRepository.QuerysdmsubscriptionsParamOpts{})
		if err != nil {
			t.Fatalf(err.Error())
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		spew.Dump(sdmSubscriptionArray)
	}

}