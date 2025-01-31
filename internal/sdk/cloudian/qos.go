package cloudian

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

const DefaultRegion = ""

// QualityOfService configures data limits for a Group or User.
type QualityOfService struct {
	// Warning is the soft limit that triggers a warning.
	Warning QualityOfServiceLimits
	// Hard is the hard limit.
	Hard QualityOfServiceLimits
}

// QualityOfService configures data limits.
type QualityOfServiceLimits struct {
	// StorageQuotaKiBs is the limit for total stored data in KiB.
	StorageQuotaKiBs *int64
	// StorageQuotaCount is the limit for total number of objects.
	StorageQuotaCount *int64
	// RequestsPerMin is the limit for number of HTTP requests per minute.
	RequestsPerMin *int64
	// InboundKiBsPerMin is the limit for inbound data per minute in KiB.
	InboundKiBsPerMin *int64
	// OutboundKiBsPerMin is the limit for outbound data per minute in KiB.
	OutboundKiBsPerMin *int64
}

// nolint: gocyclo
func (qos *QualityOfService) unmarshalQOSList(raw []byte) error {
	var data struct {
		QOSLimitList []struct {
			Type  string `json:"type"`
			Value int64  `json:"value"`
		} `json:"qosLimitList"`
	}

	if err := json.Unmarshal(raw, &data); err != nil {
		return err
	}

	for _, item := range data.QOSLimitList {
		v := &item.Value
		switch item.Type {
		case "STORAGE_QUOTA_KBYTES_LH":
			qos.Hard.StorageQuotaKiBs = v
		case "STORAGE_QUOTA_KBYTES_LW":
			qos.Warning.StorageQuotaKiBs = v
		case "STORAGE_QUOTA_COUNT_LH":
			qos.Hard.StorageQuotaCount = v
		case "STORAGE_QUOTA_COUNT_LW":
			qos.Warning.StorageQuotaCount = v
		case "REQUEST_RATE_LH":
			qos.Hard.RequestsPerMin = v
		case "REQUEST_RATE_LW":
			qos.Warning.RequestsPerMin = v
		case "DATAKBYTES_IN_LH":
			qos.Hard.InboundKiBsPerMin = v
		case "DATAKBYTES_IN_LW":
			qos.Warning.InboundKiBsPerMin = v
		case "DATAKBYTES_OUT_LH":
			qos.Hard.OutboundKiBsPerMin = v
		case "DATAKBYTES_OUT_LW":
			qos.Warning.OutboundKiBsPerMin = v
		}
	}
	return nil
}

func (qos *QualityOfService) queryParams(params map[string]string) error {
	rawParams := map[string]*int64{
		"hlStorageQuotaKBytes": qos.Hard.StorageQuotaKiBs,
		"wlStorageQuotaKBytes": qos.Warning.StorageQuotaKiBs,
		"hlStorageQuotaCount":  qos.Hard.StorageQuotaCount,
		"wlStorageQuotaCount":  qos.Warning.StorageQuotaCount,
		"hlRequestRate":        qos.Hard.RequestsPerMin,
		"wlRequestRate":        qos.Warning.RequestsPerMin,
		"hlDataKBytesIn":       qos.Hard.InboundKiBsPerMin,
		"wlDataKBytesIn":       qos.Warning.InboundKiBsPerMin,
		"hlDataKBytesOut":      qos.Hard.OutboundKiBsPerMin,
		"wlDataKBytesOut":      qos.Warning.OutboundKiBsPerMin,
	}

	for key, raw := range rawParams {
		val := int64(-1)
		if raw != nil {
			val = *raw
		}
		if val < -1 {
			return fmt.Errorf("invalid QoS limit value: %d", val)
		}
		params[key] = strconv.FormatInt(val, 10)
	}
	return nil
}

// SetQOS sets QualityOfService limits for a Group or User, depending on the value of GroupID and UserID.
//
// User-level QoS for a specific user (GroupID="<groupId>", UserID="<userId>")
// Default user-level QoS for a specific group (GroupID="<groupId>", UserID="ALL")
// Default user-level QoS for the whole region (GroupID="*", UserID="ALL")
// Group-level QoS for a specific group (GroupID="<groupId>", UserID="*")
// Default group-level QoS for the whole region (GroupID="ALL", UserID="*")
func (client Client) SetQOS(ctx context.Context, user User, region string, qos QualityOfService) error {
	params := make(map[string]string)
	if err := qos.queryParams(params); err != nil {
		return err
	}

	if region != DefaultRegion {
		params["region"] = region
	}

	resp, err := client.newRequest(ctx).
		SetQueryParam("userId", user.UserID).
		SetQueryParam("groupId", user.GroupID).
		SetQueryParams(params).
		Post("/qos/limits")
	if err != nil {
		return err
	}

	switch resp.StatusCode() {
	case 200:
		return nil
	default:
		return fmt.Errorf("POST quota unexpected status: %d", resp.StatusCode())
	}
}

// SetQOS gets QualityOfService limits for a Group or User, depending on the value of GroupID and UserID.
// See SetQOS for details.
func (client Client) GetQOS(ctx context.Context, user User, region string) (*QualityOfService, error) {
	params := make(map[string]string)
	if region != DefaultRegion {
		params["region"] = region
	}

	resp, err := client.newRequest(ctx).
		SetQueryParam("userId", user.UserID).
		SetQueryParam("groupId", user.GroupID).
		SetQueryParams(params).
		Get("/qos/limits")
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case 200:
		qos := &QualityOfService{}
		return qos, qos.unmarshalQOSList(resp.Body())
	default:
		return nil, fmt.Errorf("GET quota unexpected status: %d", resp.StatusCode())
	}
}
