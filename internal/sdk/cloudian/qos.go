package cloudian

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"k8s.io/utils/ptr"
)



// QualityOfService configures soft (warning) and hard limits for a Group or User.
type QualityOfService struct {
	Soft QualityOfServiceLimits
	Hard QualityOfServiceLimits
}

// QualityOfService configures limits.
type QualityOfServiceLimits struct {
	// StorageQuotaKBytes is the limit for total stored data in KiB.
	StorageQuotaKBytes *int64
	// StorageQuotaCount is the limit for total number of objects.
	StorageQuotaCount *int64
	// RequestsPerMin is the limit for number of HTTP requests per minute.
	RequestsPerMin *int64
	// InboundKBytesPerMin is the limit for inbound data per minute in KiB.
	InboundKBytesPerMin *int64
	// OutboundKBytesPerMin is the limit for outbound data per minute in KiB.
	OutboundKBytesPerMin *int64
}

// CreateQuota sets the QoS limits for a `User`. To change QoS limits, a delete and recreate is necessary.
func (client Client) CreateQuota(ctx context.Context, user User, qos QualityOfService) error {
	intStr := func(i *int64) string {
		v := ptr.Deref(i, -1)
		if v < -1 {
			v = -1
		}
		return strconv.FormatInt(v, 10)
	}

	resp, err := client.newRequest(ctx).
		SetQueryParam("userId", user.UserID).
		SetQueryParam("groupId", user.GroupID).
		SetQueryParam("hlStorageQuotaKBytes", intStr(qos.Hard.StorageQuotaKBytes)).
		SetQueryParam("wlStorageQuotaKBytes", intStr(qos.Soft.StorageQuotaKBytes)).
		SetQueryParam("hlStorageQuotaCount", intStr(qos.Hard.StorageQuotaCount)).
		SetQueryParam("wlStorageQuotaCount", intStr(qos.Soft.StorageQuotaCount)).
		SetQueryParam("hlRequestRate", intStr(qos.Hard.RequestsPerMin)).
		SetQueryParam("wlRequestRate", intStr(qos.Soft.RequestsPerMin)).
		SetQueryParam("hlDataKBytesIn", intStr(qos.Hard.InboundKBytesPerMin)).
		SetQueryParam("wlDataKBytesIn", intStr(qos.Soft.InboundKBytesPerMin)).
		SetQueryParam("hlDataKBytesOut", intStr(qos.Hard.OutboundKBytesPerMin)).
		SetQueryParam("wlDataKBytesOut", intStr(qos.Soft.OutboundKBytesPerMin)).
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
		var data struct {
			QOSLimitList []struct {
				Type  string
				Value int64
			} `json:"qosLimitList"`
		}

		if err := json.Unmarshal(resp.Body(), &data); err != nil {
			return nil, err
		}

		qos := QualityOfService{}
		for _, item := range data.QOSLimitList {
			if item.Value < 0 {
				continue
			}

			v := &item.Value
			switch item.Type {
			case "STORAGE_QUOTA_KBYTES_LH":
				qos.Hard.StorageQuotaKBytes = v
			case "STORAGE_QUOTA_KBYTES_LW":
				qos.Soft.StorageQuotaKBytes = v
			case "STORAGE_QUOTA_COUNT_LH":
				qos.Hard.StorageQuotaCount = v
			case "STORAGE_QUOTA_COUNT_LW":
				qos.Soft.StorageQuotaCount = v
			case "REQUEST_RATE_LH":
				qos.Hard.RequestsPerMin = v
			case "REQUEST_RATE_LW":
				qos.Soft.RequestsPerMin = v
			case "DATAKBYTES_IN_LH":
				qos.Hard.InboundKBytesPerMin = v
			case "DATAKBYTES_IN_LW":
				qos.Soft.InboundKBytesPerMin = v
			case "DATAKBYTES_OUT_LH":
				qos.Hard.OutboundKBytesPerMin = v
			case "DATAKBYTES_OUT_LW":
				qos.Soft.OutboundKBytesPerMin = v
			}
		}
		return &qos, nil
	default:
		return nil, fmt.Errorf("SET quota unexpected status: %d", resp.StatusCode())
	}
}
