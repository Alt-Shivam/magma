// Package TS29518NamfEventExposure provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package TS29518NamfEventExposure

import (
	externalRef0 "magma/feg/gateway/sbi/specs/TS29503NudmEE"
	externalRef1 "magma/feg/gateway/sbi/specs/TS29571CommonData"
)

const (
	OAuth2ClientCredentialsScopes = "oAuth2ClientCredentials.Scopes"
)

// Defines values for AmfUpdateEventOptionItemOp.
const (
	AmfUpdateEventOptionItemOpReplace AmfUpdateEventOptionItemOp = "replace"
)

// Defines values for AmfUpdateEventSubscriptionItemOp.
const (
	AmfUpdateEventSubscriptionItemOpAdd AmfUpdateEventSubscriptionItemOp = "add"

	AmfUpdateEventSubscriptionItemOpRemove AmfUpdateEventSubscriptionItemOp = "remove"

	AmfUpdateEventSubscriptionItemOpReplace AmfUpdateEventSubscriptionItemOp = "replace"
)

// N5gGuti defines model for 5gGuti.
type N5gGuti string

// AmfCreateEventSubscription defines model for AmfCreateEventSubscription.
type AmfCreateEventSubscription struct {
	OldGuami          *externalRef1.Guami             `json:"oldGuami,omitempty"`
	Subscription      AmfEventSubscription            `json:"subscription"`
	SupportedFeatures *externalRef1.SupportedFeatures `json:"supportedFeatures,omitempty"`
}

// AmfCreatedEventSubscription defines model for AmfCreatedEventSubscription.
type AmfCreatedEventSubscription struct {
	ReportList        *[]AmfEventReport               `json:"reportList,omitempty"`
	Subscription      AmfEventSubscription            `json:"subscription"`
	SubscriptionId    externalRef1.Uri                `json:"subscriptionId"`
	SupportedFeatures *externalRef1.SupportedFeatures `json:"supportedFeatures,omitempty"`
}

// AmfEvent defines model for AmfEvent.
type AmfEvent struct {
	AreaList           *[]AmfEventArea           `json:"areaList,omitempty"`
	ImmediateFlag      *bool                     `json:"immediateFlag,omitempty"`
	LocationFilterList *[]LocationFilter         `json:"locationFilterList,omitempty"`
	ReachabilityFilter *ReachabilityFilter       `json:"reachabilityFilter,omitempty"`
	RefId              *externalRef0.ReferenceId `json:"refId,omitempty"`
	Type               AmfEventType              `json:"type"`
}

// AmfEventArea defines model for AmfEventArea.
type AmfEventArea struct {
	LadnInfo     *LadnInfo                  `json:"ladnInfo,omitempty"`
	PresenceInfo *externalRef1.PresenceInfo `json:"presenceInfo,omitempty"`
}

// AmfEventMode defines model for AmfEventMode.
type AmfEventMode struct {
	Expiry     *externalRef1.DateTime `json:"expiry,omitempty"`
	MaxReports *int                   `json:"maxReports,omitempty"`
	Trigger    AmfEventTrigger        `json:"trigger"`
}

// AmfEventNotification defines model for AmfEventNotification.
type AmfEventNotification struct {
	NotifyCorrelationId           *string           `json:"notifyCorrelationId,omitempty"`
	ReportList                    *[]AmfEventReport `json:"reportList,omitempty"`
	SubsChangeNotifyCorrelationId *string           `json:"subsChangeNotifyCorrelationId,omitempty"`
}

// AmfEventReport defines model for AmfEventReport.
type AmfEventReport struct {
	AccessTypeList *[]externalRef1.AccessType `json:"accessTypeList,omitempty"`
	AnyUe          *bool                      `json:"anyUe,omitempty"`
	AreaList       *[]AmfEventArea            `json:"areaList,omitempty"`
	CmInfoList     *[]CmInfo                  `json:"cmInfoList,omitempty"`
	CommFailure    *CommunicationFailure      `json:"commFailure,omitempty"`
	Gpsi           *externalRef1.Gpsi         `json:"gpsi,omitempty"`
	Location       *externalRef1.UserLocation `json:"location,omitempty"`
	NumberOfUes    *int                       `json:"numberOfUes,omitempty"`
	Pei            *externalRef1.Pei          `json:"pei,omitempty"`
	Reachability   *UeReachability            `json:"reachability,omitempty"`
	RefId          *externalRef0.ReferenceId  `json:"refId,omitempty"`
	RmInfoList     *[]RmInfo                  `json:"rmInfoList,omitempty"`
	State          AmfEventState              `json:"state"`
	SubscriptionId *externalRef1.Uri          `json:"subscriptionId,omitempty"`
	Supi           *externalRef1.Supi         `json:"supi,omitempty"`
	TimeStamp      externalRef1.DateTime      `json:"timeStamp"`
	Timezone       *externalRef1.TimeZone     `json:"timezone,omitempty"`
	Type           AmfEventType               `json:"type"`
}

// AmfEventState defines model for AmfEventState.
type AmfEventState struct {
	Active         bool                      `json:"active"`
	RemainDuration *externalRef1.DurationSec `json:"remainDuration,omitempty"`
	RemainReports  *int                      `json:"remainReports,omitempty"`
}

// AmfEventSubscription defines model for AmfEventSubscription.
type AmfEventSubscription struct {
	AnyUE                         *bool                     `json:"anyUE,omitempty"`
	EventList                     []AmfEvent                `json:"eventList"`
	EventNotifyUri                externalRef1.Uri          `json:"eventNotifyUri"`
	Gpsi                          *externalRef1.Gpsi        `json:"gpsi,omitempty"`
	GroupId                       *externalRef1.GroupId     `json:"groupId,omitempty"`
	NfId                          externalRef1.NfInstanceId `json:"nfId"`
	NotifyCorrelationId           string                    `json:"notifyCorrelationId"`
	Options                       *AmfEventMode             `json:"options,omitempty"`
	Pei                           *externalRef1.Pei         `json:"pei,omitempty"`
	SubsChangeNotifyCorrelationId *string                   `json:"subsChangeNotifyCorrelationId,omitempty"`
	SubsChangeNotifyUri           *externalRef1.Uri         `json:"subsChangeNotifyUri,omitempty"`
	Supi                          *externalRef1.Supi        `json:"supi,omitempty"`
}

// AmfEventTrigger defines model for AmfEventTrigger.
type AmfEventTrigger interface{}

// AmfEventType defines model for AmfEventType.
type AmfEventType interface{}

// AmfUpdateEventOptionItem defines model for AmfUpdateEventOptionItem.
type AmfUpdateEventOptionItem struct {
	Op    AmfUpdateEventOptionItemOp `json:"op"`
	Path  string                     `json:"path"`
	Value externalRef1.DateTime      `json:"value"`
}

// AmfUpdateEventOptionItemOp defines model for AmfUpdateEventOptionItem.Op.
type AmfUpdateEventOptionItemOp string

// AmfUpdateEventSubscriptionItem defines model for AmfUpdateEventSubscriptionItem.
type AmfUpdateEventSubscriptionItem []struct {
	Op    AmfUpdateEventSubscriptionItemOp `json:"op"`
	Path  string                           `json:"path"`
	Value *AmfEvent                        `json:"value,omitempty"`
}

// AmfUpdateEventSubscriptionItemOp defines model for AmfUpdateEventSubscriptionItemOp.
type AmfUpdateEventSubscriptionItemOp string

// AmfUpdatedEventSubscription defines model for AmfUpdatedEventSubscription.
type AmfUpdatedEventSubscription struct {
	Subscription AmfEventSubscription `json:"subscription"`
}

// CmInfo defines model for CmInfo.
type CmInfo struct {
	AccessType externalRef1.AccessType `json:"accessType"`
	CmState    CmState                 `json:"cmState"`
}

// CmState defines model for CmState.
type CmState interface{}

// CommunicationFailure defines model for CommunicationFailure.
type CommunicationFailure struct {
	NasReleaseCode *string                 `json:"nasReleaseCode,omitempty"`
	RanReleaseCode *externalRef1.NgApCause `json:"ranReleaseCode,omitempty"`
}

// LadnInfo defines model for LadnInfo.
type LadnInfo struct {
	Ladn     string                      `json:"ladn"`
	Presence *externalRef1.PresenceState `json:"presence,omitempty"`
}

// LocationFilter defines model for LocationFilter.
type LocationFilter interface{}

// ReachabilityFilter defines model for ReachabilityFilter.
type ReachabilityFilter interface{}

// RmInfo defines model for RmInfo.
type RmInfo struct {
	AccessType externalRef1.AccessType `json:"accessType"`
	RmState    RmState                 `json:"rmState"`
}

// RmState defines model for RmState.
type RmState interface{}

// UeReachability defines model for UeReachability.
type UeReachability interface{}

// CreateSubscriptionJSONBody defines parameters for CreateSubscription.
type CreateSubscriptionJSONBody AmfCreateEventSubscription

// CreateSubscriptionJSONRequestBody defines body for CreateSubscription for application/json ContentType.
type CreateSubscriptionJSONRequestBody CreateSubscriptionJSONBody