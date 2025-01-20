package cloudian

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"k8s.io/utils/ptr"
)

// QoS is the Cloudian API's term for limits on quotas, counts and rates enforced on a `User`
type QoS struct {
	// Max storage quota
	StorageQuota *int64
	// Warning limit storage quota
	StorageQuotaWarning *int64
	// Max storage quota in number of objects
	StorageQuotaCount *int64
	// Warning limit storage quota in number of objects
	StorageQuotaCountWarning *int64
	// Max nr of HTTP requests per minute
	RequestRatePrMin *int64
	// Warning limit nr of HTTP requests per minute
	RequestRatePrMinWarning *int64
	// Max inbound datarate in ByteSize per minute
	DataRatePrMinInbound *int64
	// Warning limit inbound datarate in ByteSize per minute
	DataRatePrMinInboundWarning *int64
	// Max outbound datarate in ByteSize per minute
	DataRatePrMinOutbound *int64
	// Warning limit outbound datarate in ByteSize per minute
	DataRatePrMinOutboundWarning *int64
}

// CreateQuota sets the QoS limits for a `User`. To change QoS limits, a delete and recreate is necessary.
func (client Client) CreateQuota(ctx context.Context, user User, qos QoS) error {
	intStr := func(i *int64) string { return strconv.FormatInt(ptr.Deref(i, -1), 10) }

	resp, err := client.newRequest(ctx).
		SetQueryParam("userId", user.UserID).
		SetQueryParam("groupId", user.GroupID).
		SetQueryParam("hlStorageQuotaKBytes", intStr(qos.StorageQuota)).
		SetQueryParam("wlStorageQuotaKBytes", intStr(qos.StorageQuotaWarning)).
		SetQueryParam("hlStorageQuotaCount", intStr(qos.StorageQuotaCount)).
		SetQueryParam("wlStorageQuotaCount", intStr(qos.StorageQuotaCountWarning)).
		SetQueryParam("hlRequestRate", intStr(qos.RequestRatePrMin)).
		SetQueryParam("wlRequestRate", intStr(qos.RequestRatePrMinWarning)).
		SetQueryParam("hlDataKBytesIn", intStr(qos.DataRatePrMinInbound)).
		SetQueryParam("wlDataKBytesIn", intStr(qos.DataRatePrMinInboundWarning)).
		SetQueryParam("hlDataKBytesOut", intStr(qos.DataRatePrMinOutbound)).
		SetQueryParam("wlDataKBytesOut", intStr(qos.DataRatePrMinOutboundWarning)).
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

func (client Client) GetQuota(ctx context.Context, user User) (*QoS, error) {
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

		extract := func(field string) *int64 {
			for _, item := range data.QOSLimitList {
				if item.Type != field {
					continue
				}

				if item.Value == -1 {
					return nil
				}
				return &item.Value
			}
			return nil // Should be unreachable
		}

		return &QoS{
			StorageQuota:                 extract("STORAGE_QUOTA_KBYTES_LH"),
			StorageQuotaWarning:          extract("STORAGE_QUOTA_KBYTES_LW"),
			StorageQuotaCount:            extract("STORAGE_QUOTA_COUNT_LH"),
			StorageQuotaCountWarning:     extract("STORAGE_QUOTA_COUNT_LW"),
			RequestRatePrMin:             extract("REQUEST_RATE_LH"),
			RequestRatePrMinWarning:      extract("REQUEST_RATE_LW"),
			DataRatePrMinInbound:         extract("DATAKBYTES_IN_LH"),
			DataRatePrMinInboundWarning:  extract("DATAKBYTES_IN_LW"),
			DataRatePrMinOutbound:        extract("DATAKBYTES_OUT_LH"),
			DataRatePrMinOutboundWarning: extract("DATAKBYTES_OUT_LW"),
		}, nil
	default:
		return nil, fmt.Errorf("SET quota unexpected status: %d", resp.StatusCode())
	}
}


// err := c.CreateQuota(context.TODO(), cloudian.User{
// 	GroupID: "tenant1",
// 	UserID:  "*",
// }, cloudian.QoS{
// 	StorageQuota:                 nil,
// 	StorageQuotaWarning:          ptr.To(int64(1)),
// 	StorageQuotaCount:            ptr.To(int64(2)),
// 	StorageQuotaCountWarning:     ptr.To(int64(3)),
// 	RequestRatePrMin:             ptr.To(int64(4)),
// 	RequestRatePrMinWarning:      ptr.To(int64(5)),
// 	DataRatePrMinInbound:         ptr.To(int64(6)),
// 	DataRatePrMinInboundWarning:  ptr.To(int64(7)),
// 	DataRatePrMinOutbound:        ptr.To(int64(8)),
// 	DataRatePrMinOutboundWarning: ptr.To(int64(math.MaxInt64)),
// })
// fmt.Println(err)

// q, err := c.GetQuota(context.TODO(), cloudian.User{
// 	GroupID: "tenant1",
// 	UserID:  "*",
// })
// b, e2 := json.Marshal(q)
// fmt.Println(string(b), err, e2)