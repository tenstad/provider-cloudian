package qualityofservicelimits

import (
	"github.com/statnett/provider-cloudian/apis/user/v1alpha1"
	"github.com/statnett/provider-cloudian/internal/sdk/cloudian"
	"k8s.io/utils/ptr"
)

func ToCloudianQOS(warning *v1alpha1.QualityOfServiceLimits, hard *v1alpha1.QualityOfServiceLimits) (cloudian.QualityOfService, error) {
	var err error
	qos := cloudian.QualityOfService{}

	if qos.Warning, err = ToCloudianLimits(warning); err != nil {
		return cloudian.QualityOfService{}, err
	}
	if qos.Hard, err = ToCloudianLimits(hard); err != nil {
		return cloudian.QualityOfService{}, err
	}
	return qos, nil
}

func ToCloudianLimits(limits *v1alpha1.QualityOfServiceLimits) (cloudian.QualityOfServiceLimits, error) {
	if limits == nil {
		return cloudian.QualityOfServiceLimits{}, nil
	}

	var err error
	qosl := cloudian.QualityOfServiceLimits{}

	if qosl.StorageQuotaKiBs, err = limits.StorageQuotaBytes.ToKiB(); err != nil {
		return cloudian.QualityOfServiceLimits{}, err
	}
	if qosl.InboundKiBsPerMin, err = limits.InboundBytesPerMin.ToKiB(); err != nil {
		return cloudian.QualityOfServiceLimits{}, err
	}
	if qosl.OutboundKiBsPerMin, err = limits.OutboundBytesPerMin.ToKiB(); err != nil {
		return cloudian.QualityOfServiceLimits{}, err
	}

	if limits.StorageQuotaCount != nil {
		qosl.StorageQuotaCount = ptr.To(int64(*limits.StorageQuotaCount))
	}
	if limits.RequestsPerMin != nil {
		qosl.RequestsPerMin = ptr.To(int64(*limits.RequestsPerMin))
	}

	return qosl, nil
}

func LimitsEqual(a cloudian.QualityOfServiceLimits, b cloudian.QualityOfServiceLimits) bool {
	return ptr.Equal(a.InboundKiBsPerMin, b.InboundKiBsPerMin) &&
		ptr.Equal(a.OutboundKiBsPerMin, b.OutboundKiBsPerMin) &&
		ptr.Equal(a.RequestsPerMin, b.RequestsPerMin) &&
		ptr.Equal(a.StorageQuotaCount, b.StorageQuotaCount) &&
		ptr.Equal(a.StorageQuotaKiBs, b.StorageQuotaKiBs)
}
