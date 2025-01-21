package cloudian

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"k8s.io/utils/ptr"
)

type Limit struct {
	Soft *int64
	Hard *int64
}
type QoSLimit int

const (
	StorageQuotaKBytes QoSLimit = iota
	StorageQuotaCount
	RequestRate
	DataKBytesIn
	DataKBytesOut
)

var qoSLimitName = map[QoSLimit]string{
	StorageQuotaKBytes: "StorageQuotaKBytes",
	StorageQuotaCount:  "StorageQuotaCount",
	RequestRate:        "RequestRate",
	DataKBytesIn:       "DataKBytesIn",
	DataKBytesOut:      "DataKBytesOut",
}
var qoSLimitJSONName = map[string]QoSLimit{
	"STORAGE_QUOTA_KBYTES": StorageQuotaKBytes,
	"STORAGE_QUOTA_COUNT":  StorageQuotaCount,
	"REQUEST_RATE":         RequestRate,
	"DATAKBYTES_IN":        DataKBytesIn,
	"DATAKBYTES_OUT":       DataKBytesOut,
}

func (ql QoSLimit) String() string {
	return qoSLimitName[ql]
}

type QoS map[QoSLimit]Limit

func (q QoS) QueryParams() map[string]string {
	params := make(map[string]string)
	for ql, l := range q {
		params["wl"+ql.String()] = strconv.FormatInt(ptr.Deref(l.Soft, -1), 10)
		params["hl"+ql.String()] = strconv.FormatInt(ptr.Deref(l.Hard, -1), 10)
	}
	return params
}
func (q QoS) UnmarshalJSON(b []byte) error {
	var data struct {
		QOSLimitList []struct {
			Type  string `json:"type"`
			Value *int64 `json:"value"`
		} `json:"qosLimitList"`
	}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	for _, ql := range data.QOSLimitList {
		if before, found := strings.CutSuffix(ql.Type, "_LW"); found {
			l := q[qoSLimitJSONName[before]]
			l.Soft = ql.Value
			q[qoSLimitJSONName[before]] = l
		}
		if before, found := strings.CutSuffix(ql.Type, "_LH"); found {
			l := q[qoSLimitJSONName[before]]
			l.Hard = ql.Value
			q[qoSLimitJSONName[before]] = l
		}
	}
	return nil
}

// CreateQuota sets the QoS limits for a `User`. To change QoS limits, a delete and recreate is necessary.
func (client Client) CreateQuota(ctx context.Context, user User, qos QoS) error {
	resp, err := client.newRequest(ctx).
		SetQueryParam("userId", user.UserID).
		SetQueryParam("groupId", user.GroupID).
		SetQueryParams(qos.QueryParams()).
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
		qos := QoS{}
		return &qos, qos.UnmarshalJSON(resp.Body())
	default:
		return nil, fmt.Errorf("SET quota unexpected status: %d", resp.StatusCode())
	}
}

// err := c.CreateQuota(context.TODO(), cloudian.User{
// 	GroupID: "tenant1",
// 	UserID:  "*",
// }, cloudian.QoS{
// 	cloudian.StorageQuotaKBytes: cloudian.Limit{Soft: ptr.To(int64(1)), Hard: ptr.To(int64(6))},
// 	cloudian.StorageQuotaCount:  cloudian.Limit{Soft: ptr.To(int64(2)), Hard: ptr.To(int64(7))},
// 	cloudian.RequestRate:        cloudian.Limit{Soft: ptr.To(int64(3)), Hard: ptr.To(int64(8))},
// 	cloudian.DataKBytesIn:       cloudian.Limit{Soft: ptr.To(int64(4)), Hard: ptr.To(int64(9))},
// 	cloudian.DataKBytesOut:      cloudian.Limit{Soft: ptr.To(int64(5)), Hard: ptr.To(int64(10))},
// })
// fmt.Println(err)

// q, err := c.GetQuota(context.TODO(), cloudian.User{
// 	GroupID: "tenant1",
// 	UserID:  "*",
// })
// b, e2 := json.Marshal(q)
// fmt.Println(string(b), err, e2)
