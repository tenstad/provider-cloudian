package cloudian

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

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
func (qos *QualityOfService) unmarshalJSON(raw []byte) error {
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
		if item.Value < 0 {
			continue
		}

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

// CreateQuota sets the QoS limits for a `User`. To change QoS limits, a delete and recreate is necessary.
func (client Client) CreateQuota(ctx context.Context, user User, qos QualityOfService) error {
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

	params := make(map[string]string, len(rawParams))
	for key, raw := range rawParams {
		val := int64(-1)
		if raw != nil {
			val = *raw
		}
		if val < -1 {
			val = -1
		}
		params[key] = strconv.FormatInt(val, 10)
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
		return fmt.Errorf("SET quota unexpected status: %d", resp.StatusCode())
	}
}

func (client Client) GetQuota(ctx context.Context, user User) (*QualityOfService, error) {
	resp, err := client.newRequest(ctx).
		SetQueryParam("userId", user.UserID).
		SetQueryParam("groupId", user.GroupID).
		Get("/qos/limits")
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case 200:
		qos := &QualityOfService{}
		return qos, qos.unmarshalJSON(resp.Body())
	default:
		return nil, fmt.Errorf("SET quota unexpected status: %d", resp.StatusCode())
	}
}
