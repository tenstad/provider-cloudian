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
	// StorageQuotaKBytes is the hard limit for total stored data in KiB.
	StorageQuotaKBytes *int64
	// StorageQuotaKBytesWarning is the warning limit for total stored data in KiB.
	StorageQuotaKBytesWarning *int64
	// StorageQuotaCount is the hard limit for total number of objects.
	StorageQuotaCount *int64
	// StorageQuotaCountWarning is the warning limit for total number of objects.
	StorageQuotaCountWarning *int64
	// RequestsPerMin is the hard limit for number of HTTP requests per minute.
	RequestsPerMin *int64
	// RequestsPerMinWarning is the warning limit for number of HTTP requests per minute.
	RequestsPerMinWarning *int64
	// InboundKBytesPerMin is the hard limit for inbound data per minute in KiB.
	InboundKBytesPerMin *int64
	// InboundKBytesPerMin is the warning limit for inbound data per minute in KiB.
	InboundKBytesPerMinWarning *int64
	// OutboundKBytesPerMin is the hard limit for outbound data per minute in KiB.
	OutboundKBytesPerMin *int64
	// OutboundKBytesPerMinWarning is the warning limit for outbound data per minute in KiB.
	OutboundKBytesPerMinWarning *int64
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
		SetQueryParam("hlStorageQuotaKBytes", intStr(qos.StorageQuotaKBytes)).
		SetQueryParam("wlStorageQuotaKBytes", intStr(qos.StorageQuotaKBytesWarning)).
		SetQueryParam("hlStorageQuotaCount", intStr(qos.StorageQuotaCount)).
		SetQueryParam("wlStorageQuotaCount", intStr(qos.StorageQuotaCountWarning)).
		SetQueryParam("hlRequestRate", intStr(qos.RequestsPerMin)).
		SetQueryParam("wlRequestRate", intStr(qos.RequestsPerMinWarning)).
		SetQueryParam("hlDataKBytesIn", intStr(qos.InboundKBytesPerMin)).
		SetQueryParam("wlDataKBytesIn", intStr(qos.InboundKBytesPerMinWarning)).
		SetQueryParam("hlDataKBytesOut", intStr(qos.OutboundKBytesPerMin)).
		SetQueryParam("wlDataKBytesOut", intStr(qos.OutboundKBytesPerMinWarning)).
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
				qos.StorageQuotaKBytes = v
			case "STORAGE_QUOTA_KBYTES_LW":
				qos.StorageQuotaKBytesWarning = v
			case "STORAGE_QUOTA_COUNT_LH":
				qos.StorageQuotaCount = v
			case "STORAGE_QUOTA_COUNT_LW":
				qos.StorageQuotaCountWarning = v
			case "REQUEST_RATE_LH":
				qos.RequestsPerMin = v
			case "REQUEST_RATE_LW":
				qos.RequestsPerMinWarning = v
			case "DATAKBYTES_IN_LH":
				qos.InboundKBytesPerMin = v
			case "DATAKBYTES_IN_LW":
				qos.InboundKBytesPerMinWarning = v
			case "DATAKBYTES_OUT_LH":
				qos.OutboundKBytesPerMin = v
			case "DATAKBYTES_OUT_LW":
				qos.OutboundKBytesPerMinWarning = v
			}
		}
		return &qos, nil
	default:
		return nil, fmt.Errorf("SET quota unexpected status: %d", resp.StatusCode())
	}
}
