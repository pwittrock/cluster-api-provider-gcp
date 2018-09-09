// Package monitoring provides access to the Stackdriver Monitoring API.
//
// This package is DEPRECATED. Use package cloud.google.com/go/monitoring/apiv3 instead.
//
// See https://cloud.google.com/monitoring/api/
//
// Usage example:
//
//   import "google.golang.org/api/monitoring/v3"
//   ...
//   monitoringService, err := monitoring.New(oauthHttpClient)
package monitoring // import "google.golang.org/api/monitoring/v3"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "monitoring:v3"
const apiName = "monitoring"
const apiVersion = "v3"
const basePath = "https://monitoring.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and write monitoring data for all of your Google and third-party
	// Cloud and API projects
	MonitoringScope = "https://www.googleapis.com/auth/monitoring"

	// View monitoring data for all of your Google Cloud and third-party
	// projects
	MonitoringReadScope = "https://www.googleapis.com/auth/monitoring.read"

	// Publish metric data to your Google Cloud projects
	MonitoringWriteScope = "https://www.googleapis.com/auth/monitoring.write"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	s.UptimeCheckIps = NewUptimeCheckIpsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService

	UptimeCheckIps *UptimeCheckIpsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.AlertPolicies = NewProjectsAlertPoliciesService(s)
	rs.CollectdTimeSeries = NewProjectsCollectdTimeSeriesService(s)
	rs.Groups = NewProjectsGroupsService(s)
	rs.MetricDescriptors = NewProjectsMetricDescriptorsService(s)
	rs.MonitoredResourceDescriptors = NewProjectsMonitoredResourceDescriptorsService(s)
	rs.NotificationChannelDescriptors = NewProjectsNotificationChannelDescriptorsService(s)
	rs.NotificationChannels = NewProjectsNotificationChannelsService(s)
	rs.TimeSeries = NewProjectsTimeSeriesService(s)
	rs.UptimeCheckConfigs = NewProjectsUptimeCheckConfigsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	AlertPolicies *ProjectsAlertPoliciesService

	CollectdTimeSeries *ProjectsCollectdTimeSeriesService

	Groups *ProjectsGroupsService

	MetricDescriptors *ProjectsMetricDescriptorsService

	MonitoredResourceDescriptors *ProjectsMonitoredResourceDescriptorsService

	NotificationChannelDescriptors *ProjectsNotificationChannelDescriptorsService

	NotificationChannels *ProjectsNotificationChannelsService

	TimeSeries *ProjectsTimeSeriesService

	UptimeCheckConfigs *ProjectsUptimeCheckConfigsService
}

func NewProjectsAlertPoliciesService(s *Service) *ProjectsAlertPoliciesService {
	rs := &ProjectsAlertPoliciesService{s: s}
	return rs
}

type ProjectsAlertPoliciesService struct {
	s *Service
}

func NewProjectsCollectdTimeSeriesService(s *Service) *ProjectsCollectdTimeSeriesService {
	rs := &ProjectsCollectdTimeSeriesService{s: s}
	return rs
}

type ProjectsCollectdTimeSeriesService struct {
	s *Service
}

func NewProjectsGroupsService(s *Service) *ProjectsGroupsService {
	rs := &ProjectsGroupsService{s: s}
	rs.Members = NewProjectsGroupsMembersService(s)
	return rs
}

type ProjectsGroupsService struct {
	s *Service

	Members *ProjectsGroupsMembersService
}

func NewProjectsGroupsMembersService(s *Service) *ProjectsGroupsMembersService {
	rs := &ProjectsGroupsMembersService{s: s}
	return rs
}

type ProjectsGroupsMembersService struct {
	s *Service
}

func NewProjectsMetricDescriptorsService(s *Service) *ProjectsMetricDescriptorsService {
	rs := &ProjectsMetricDescriptorsService{s: s}
	return rs
}

type ProjectsMetricDescriptorsService struct {
	s *Service
}

func NewProjectsMonitoredResourceDescriptorsService(s *Service) *ProjectsMonitoredResourceDescriptorsService {
	rs := &ProjectsMonitoredResourceDescriptorsService{s: s}
	return rs
}

type ProjectsMonitoredResourceDescriptorsService struct {
	s *Service
}

func NewProjectsNotificationChannelDescriptorsService(s *Service) *ProjectsNotificationChannelDescriptorsService {
	rs := &ProjectsNotificationChannelDescriptorsService{s: s}
	return rs
}

type ProjectsNotificationChannelDescriptorsService struct {
	s *Service
}

func NewProjectsNotificationChannelsService(s *Service) *ProjectsNotificationChannelsService {
	rs := &ProjectsNotificationChannelsService{s: s}
	return rs
}

type ProjectsNotificationChannelsService struct {
	s *Service
}

func NewProjectsTimeSeriesService(s *Service) *ProjectsTimeSeriesService {
	rs := &ProjectsTimeSeriesService{s: s}
	return rs
}

type ProjectsTimeSeriesService struct {
	s *Service
}

func NewProjectsUptimeCheckConfigsService(s *Service) *ProjectsUptimeCheckConfigsService {
	rs := &ProjectsUptimeCheckConfigsService{s: s}
	return rs
}

type ProjectsUptimeCheckConfigsService struct {
	s *Service
}

func NewUptimeCheckIpsService(s *Service) *UptimeCheckIpsService {
	rs := &UptimeCheckIpsService{s: s}
	return rs
}

type UptimeCheckIpsService struct {
	s *Service
}

// Aggregation: Describes how to combine multiple time series to provide
// different views of the data. Aggregation consists of an alignment
// step on individual time series (alignment_period and
// per_series_aligner) followed by an optional reduction step of the
// data across the aligned time series (cross_series_reducer and
// group_by_fields). For more details, see Aggregation.
type Aggregation struct {
	// AlignmentPeriod: The alignment period for per-time series alignment.
	// If present, alignmentPeriod must be at least 60 seconds. After
	// per-time series alignment, each time series will contain data points
	// only on the period boundaries. If perSeriesAligner is not specified
	// or equals ALIGN_NONE, then this field is ignored. If perSeriesAligner
	// is specified and does not equal ALIGN_NONE, then this field must be
	// defined; otherwise an error is returned.
	AlignmentPeriod string `json:"alignmentPeriod,omitempty"`

	// CrossSeriesReducer: The approach to be used to combine time series.
	// Not all reducer functions may be applied to all time series,
	// depending on the metric type and the value type of the original time
	// series. Reduction may change the metric type of value type of the
	// time series.Time series data must be aligned in order to perform
	// cross-time series reduction. If crossSeriesReducer is specified, then
	// perSeriesAligner must be specified and not equal ALIGN_NONE and
	// alignmentPeriod must be specified; otherwise, an error is returned.
	//
	// Possible values:
	//   "REDUCE_NONE" - No cross-time series reduction. The output of the
	// aligner is returned.
	//   "REDUCE_MEAN" - Reduce by computing the mean across time series for
	// each alignment period. This reducer is valid for delta and gauge
	// metrics with numeric or distribution values. The value type of the
	// output is DOUBLE.
	//   "REDUCE_MIN" - Reduce by computing the minimum across time series
	// for each alignment period. This reducer is valid for delta and gauge
	// metrics with numeric values. The value type of the output is the same
	// as the value type of the input.
	//   "REDUCE_MAX" - Reduce by computing the maximum across time series
	// for each alignment period. This reducer is valid for delta and gauge
	// metrics with numeric values. The value type of the output is the same
	// as the value type of the input.
	//   "REDUCE_SUM" - Reduce by computing the sum across time series for
	// each alignment period. This reducer is valid for delta and gauge
	// metrics with numeric and distribution values. The value type of the
	// output is the same as the value type of the input.
	//   "REDUCE_STDDEV" - Reduce by computing the standard deviation across
	// time series for each alignment period. This reducer is valid for
	// delta and gauge metrics with numeric or distribution values. The
	// value type of the output is DOUBLE.
	//   "REDUCE_COUNT" - Reduce by computing the count of data points
	// across time series for each alignment period. This reducer is valid
	// for delta and gauge metrics of numeric, Boolean, distribution, and
	// string value type. The value type of the output is INT64.
	//   "REDUCE_COUNT_TRUE" - Reduce by computing the count of True-valued
	// data points across time series for each alignment period. This
	// reducer is valid for delta and gauge metrics of Boolean value type.
	// The value type of the output is INT64.
	//   "REDUCE_COUNT_FALSE" - Reduce by computing the count of
	// False-valued data points across time series for each alignment
	// period. This reducer is valid for delta and gauge metrics of Boolean
	// value type. The value type of the output is INT64.
	//   "REDUCE_FRACTION_TRUE" - Reduce by computing the fraction of
	// True-valued data points across time series for each alignment period.
	// This reducer is valid for delta and gauge metrics of Boolean value
	// type. The output value is in the range 0, 1 and has value type
	// DOUBLE.
	//   "REDUCE_PERCENTILE_99" - Reduce by computing 99th percentile of
	// data points across time series for each alignment period. This
	// reducer is valid for gauge and delta metrics of numeric and
	// distribution type. The value of the output is DOUBLE
	//   "REDUCE_PERCENTILE_95" - Reduce by computing 95th percentile of
	// data points across time series for each alignment period. This
	// reducer is valid for gauge and delta metrics of numeric and
	// distribution type. The value of the output is DOUBLE
	//   "REDUCE_PERCENTILE_50" - Reduce by computing 50th percentile of
	// data points across time series for each alignment period. This
	// reducer is valid for gauge and delta metrics of numeric and
	// distribution type. The value of the output is DOUBLE
	//   "REDUCE_PERCENTILE_05" - Reduce by computing 5th percentile of data
	// points across time series for each alignment period. This reducer is
	// valid for gauge and delta metrics of numeric and distribution type.
	// The value of the output is DOUBLE
	CrossSeriesReducer string `json:"crossSeriesReducer,omitempty"`

	// GroupByFields: The set of fields to preserve when crossSeriesReducer
	// is specified. The groupByFields determine how the time series are
	// partitioned into subsets prior to applying the aggregation function.
	// Each subset contains time series that have the same value for each of
	// the grouping fields. Each individual time series is a member of
	// exactly one subset. The crossSeriesReducer is applied to each subset
	// of time series. It is not possible to reduce across different
	// resource types, so this field implicitly contains resource.type.
	// Fields not specified in groupByFields are aggregated away. If
	// groupByFields is not specified and all the time series have the same
	// resource type, then the time series are aggregated into a single
	// output time series. If crossSeriesReducer is not defined, this field
	// is ignored.
	GroupByFields []string `json:"groupByFields,omitempty"`

	// PerSeriesAligner: The approach to be used to align individual time
	// series. Not all alignment functions may be applied to all time
	// series, depending on the metric type and value type of the original
	// time series. Alignment may change the metric type or the value type
	// of the time series.Time series data must be aligned in order to
	// perform cross-time series reduction. If crossSeriesReducer is
	// specified, then perSeriesAligner must be specified and not equal
	// ALIGN_NONE and alignmentPeriod must be specified; otherwise, an error
	// is returned.
	//
	// Possible values:
	//   "ALIGN_NONE" - No alignment. Raw data is returned. Not valid if
	// cross-time series reduction is requested. The value type of the
	// result is the same as the value type of the input.
	//   "ALIGN_DELTA" - Align and convert to delta metric type. This
	// alignment is valid for cumulative metrics and delta metrics. Aligning
	// an existing delta metric to a delta metric requires that the
	// alignment period be increased. The value type of the result is the
	// same as the value type of the input.One can think of this aligner as
	// a rate but without time units; that is, the output is conceptually
	// (second_point - first_point).
	//   "ALIGN_RATE" - Align and convert to a rate. This alignment is valid
	// for cumulative metrics and delta metrics with numeric values. The
	// output is a gauge metric with value type DOUBLE.One can think of this
	// aligner as conceptually providing the slope of the line that passes
	// through the value at the start and end of the window. In other words,
	// this is conceptually ((y1 - y0)/(t1 - t0)), and the output unit is
	// one that has a "/time" dimension.If, by rate, you are looking for
	// percentage change, see the ALIGN_PERCENT_CHANGE aligner option.
	//   "ALIGN_INTERPOLATE" - Align by interpolating between adjacent
	// points around the period boundary. This alignment is valid for gauge
	// metrics with numeric values. The value type of the result is the same
	// as the value type of the input.
	//   "ALIGN_NEXT_OLDER" - Align by shifting the oldest data point before
	// the period boundary to the boundary. This alignment is valid for
	// gauge metrics. The value type of the result is the same as the value
	// type of the input.
	//   "ALIGN_MIN" - Align time series via aggregation. The resulting data
	// point in the alignment period is the minimum of all data points in
	// the period. This alignment is valid for gauge and delta metrics with
	// numeric values. The value type of the result is the same as the value
	// type of the input.
	//   "ALIGN_MAX" - Align time series via aggregation. The resulting data
	// point in the alignment period is the maximum of all data points in
	// the period. This alignment is valid for gauge and delta metrics with
	// numeric values. The value type of the result is the same as the value
	// type of the input.
	//   "ALIGN_MEAN" - Align time series via aggregation. The resulting
	// data point in the alignment period is the average or arithmetic mean
	// of all data points in the period. This alignment is valid for gauge
	// and delta metrics with numeric values. The value type of the output
	// is DOUBLE.
	//   "ALIGN_COUNT" - Align time series via aggregation. The resulting
	// data point in the alignment period is the count of all data points in
	// the period. This alignment is valid for gauge and delta metrics with
	// numeric or Boolean values. The value type of the output is INT64.
	//   "ALIGN_SUM" - Align time series via aggregation. The resulting data
	// point in the alignment period is the sum of all data points in the
	// period. This alignment is valid for gauge and delta metrics with
	// numeric and distribution values. The value type of the output is the
	// same as the value type of the input.
	//   "ALIGN_STDDEV" - Align time series via aggregation. The resulting
	// data point in the alignment period is the standard deviation of all
	// data points in the period. This alignment is valid for gauge and
	// delta metrics with numeric values. The value type of the output is
	// DOUBLE.
	//   "ALIGN_COUNT_TRUE" - Align time series via aggregation. The
	// resulting data point in the alignment period is the count of
	// True-valued data points in the period. This alignment is valid for
	// gauge metrics with Boolean values. The value type of the output is
	// INT64.
	//   "ALIGN_COUNT_FALSE" - Align time series via aggregation. The
	// resulting data point in the alignment period is the count of
	// False-valued data points in the period. This alignment is valid for
	// gauge metrics with Boolean values. The value type of the output is
	// INT64.
	//   "ALIGN_FRACTION_TRUE" - Align time series via aggregation. The
	// resulting data point in the alignment period is the fraction of
	// True-valued data points in the period. This alignment is valid for
	// gauge metrics with Boolean values. The output value is in the range
	// 0, 1 and has value type DOUBLE.
	//   "ALIGN_PERCENTILE_99" - Align time series via aggregation. The
	// resulting data point in the alignment period is the 99th percentile
	// of all data points in the period. This alignment is valid for gauge
	// and delta metrics with distribution values. The output is a gauge
	// metric with value type DOUBLE.
	//   "ALIGN_PERCENTILE_95" - Align time series via aggregation. The
	// resulting data point in the alignment period is the 95th percentile
	// of all data points in the period. This alignment is valid for gauge
	// and delta metrics with distribution values. The output is a gauge
	// metric with value type DOUBLE.
	//   "ALIGN_PERCENTILE_50" - Align time series via aggregation. The
	// resulting data point in the alignment period is the 50th percentile
	// of all data points in the period. This alignment is valid for gauge
	// and delta metrics with distribution values. The output is a gauge
	// metric with value type DOUBLE.
	//   "ALIGN_PERCENTILE_05" - Align time series via aggregation. The
	// resulting data point in the alignment period is the 5th percentile of
	// all data points in the period. This alignment is valid for gauge and
	// delta metrics with distribution values. The output is a gauge metric
	// with value type DOUBLE.
	//   "ALIGN_PERCENT_CHANGE" - Align and convert to a percentage change.
	// This alignment is valid for gauge and delta metrics with numeric
	// values. This alignment conceptually computes the equivalent of
	// "((current - previous)/previous)*100" where previous value is
	// determined based on the alignmentPeriod. In the event that previous
	// is 0 the calculated value is infinity with the exception that if both
	// (current - previous) and previous are 0 the calculated value is 0. A
	// 10 minute moving mean is computed at each point of the time window
	// prior to the above calculation to smooth the metric and prevent false
	// positives from very short lived spikes. Only applicable for data that
	// is >= 0. Any values < 0 are treated as no data. While delta metrics
	// are accepted by this alignment special care should be taken that the
	// values for the metric will always be positive. The output is a gauge
	// metric with value type DOUBLE.
	PerSeriesAligner string `json:"perSeriesAligner,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AlignmentPeriod") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AlignmentPeriod") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Aggregation) MarshalJSON() ([]byte, error) {
	type NoMethod Aggregation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AlertPolicy: A description of the conditions under which some aspect
// of your system is considered to be "unhealthy" and the ways to notify
// people or services about this state. For an overview of alert
// policies, see Introduction to Alerting.
type AlertPolicy struct {
	// Combiner: How to combine the results of multiple conditions to
	// determine if an incident should be opened.
	//
	// Possible values:
	//   "COMBINE_UNSPECIFIED" - An unspecified combiner.
	//   "AND" - Combine conditions using the logical AND operator. An
	// incident is created only if all conditions are met simultaneously.
	// This combiner is satisfied if all conditions are met, even if they
	// are met on completely different resources.
	//   "OR" - Combine conditions using the logical OR operator. An
	// incident is created if any of the listed conditions is met.
	//   "AND_WITH_MATCHING_RESOURCE" - Combine conditions using logical AND
	// operator, but unlike the regular AND option, an incident is created
	// only if all conditions are met simultaneously on at least one
	// resource.
	Combiner string `json:"combiner,omitempty"`

	// Conditions: A list of conditions for the policy. The conditions are
	// combined by AND or OR according to the combiner field. If the
	// combined conditions evaluate to true, then an incident is created. A
	// policy can have from one to six conditions.
	Conditions []*Condition `json:"conditions,omitempty"`

	// CreationRecord: A read-only record of the creation of the alerting
	// policy. If provided in a call to create or update, this field will be
	// ignored.
	CreationRecord *MutationRecord `json:"creationRecord,omitempty"`

	// DisplayName: A short name or phrase used to identify the policy in
	// dashboards, notifications, and incidents. To avoid confusion, don't
	// use the same display name for multiple policies in the same project.
	// The name is limited to 512 Unicode characters.
	DisplayName string `json:"displayName,omitempty"`

	// Documentation: Documentation that is included with notifications and
	// incidents related to this policy. Best practice is for the
	// documentation to include information to help responders understand,
	// mitigate, escalate, and correct the underlying problems detected by
	// the alerting policy. Notification channels that have limited capacity
	// might not show this documentation.
	Documentation *Documentation `json:"documentation,omitempty"`

	// Enabled: Whether or not the policy is enabled. On write, the default
	// interpretation if unset is that the policy is enabled. On read,
	// clients should not make any assumption about the state if it has not
	// been populated. The field should always be populated on List and Get
	// operations, unless a field projection has been specified that strips
	// it out.
	Enabled bool `json:"enabled,omitempty"`

	// MutationRecord: A read-only record of the most recent change to the
	// alerting policy. If provided in a call to create or update, this
	// field will be ignored.
	MutationRecord *MutationRecord `json:"mutationRecord,omitempty"`

	// Name: Required if the policy exists. The resource name for this
	// policy. The syntax
	// is:
	// projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]
	// [ALERT_POLIC
	// Y_ID] is assigned by Stackdriver Monitoring when the policy is
	// created. When calling the alertPolicies.create method, do not include
	// the name field in the alerting policy passed as part of the request.
	Name string `json:"name,omitempty"`

	// NotificationChannels: Identifies the notification channels to which
	// notifications should be sent when incidents are opened or closed or
	// when new violations occur on an already opened incident. Each element
	// of this array corresponds to the name field in each of the
	// NotificationChannel objects that are returned from the
	// ListNotificationChannels method. The syntax of the entries in this
	// field is:
	// projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]
	//
	NotificationChannels []string `json:"notificationChannels,omitempty"`

	// UserLabels: User-supplied key/value data to be used for organizing
	// and identifying the AlertPolicy objects.The field can contain up to
	// 64 entries. Each key and value is limited to 63 Unicode characters or
	// 128 bytes, whichever is smaller. Labels and values can contain only
	// lowercase letters, numerals, underscores, and dashes. Keys must begin
	// with a letter.
	UserLabels map[string]string `json:"userLabels,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Combiner") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Combiner") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AlertPolicy) MarshalJSON() ([]byte, error) {
	type NoMethod AlertPolicy
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BasicAuthentication: A type of authentication to perform against the
// specified resource or URL that uses username and password. Currently,
// only Basic authentication is supported in Uptime Monitoring.
type BasicAuthentication struct {
	// Password: The password to authenticate.
	Password string `json:"password,omitempty"`

	// Username: The username to authenticate.
	Username string `json:"username,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Password") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Password") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *BasicAuthentication) MarshalJSON() ([]byte, error) {
	type NoMethod BasicAuthentication
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// BucketOptions: BucketOptions describes the bucket boundaries used to
// create a histogram for the distribution. The buckets can be in a
// linear sequence, an exponential sequence, or each bucket can be
// specified explicitly. BucketOptions does not include the number of
// values in each bucket.A bucket has an inclusive lower bound and
// exclusive upper bound for the values that are counted for that
// bucket. The upper bound of a bucket must be strictly greater than the
// lower bound. The sequence of N buckets for a distribution consists of
// an underflow bucket (number 0), zero or more finite buckets (number 1
// through N - 2) and an overflow bucket (number N - 1). The buckets are
// contiguous: the lower bound of bucket i (i > 0) is the same as the
// upper bound of bucket i - 1. The buckets span the whole range of
// finite values: lower bound of the underflow bucket is -infinity and
// the upper bound of the overflow bucket is +infinity. The finite
// buckets are so-called because both bounds are finite.
type BucketOptions struct {
	// ExplicitBuckets: The explicit buckets.
	ExplicitBuckets *Explicit `json:"explicitBuckets,omitempty"`

	// ExponentialBuckets: The exponential buckets.
	ExponentialBuckets *Exponential `json:"exponentialBuckets,omitempty"`

	// LinearBuckets: The linear bucket.
	LinearBuckets *Linear `json:"linearBuckets,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExplicitBuckets") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExplicitBuckets") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *BucketOptions) MarshalJSON() ([]byte, error) {
	type NoMethod BucketOptions
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CollectdPayload: A collection of data points sent from a
// collectd-based plugin. See the collectd documentation for more
// information.
type CollectdPayload struct {
	// EndTime: The end time of the interval.
	EndTime string `json:"endTime,omitempty"`

	// Metadata: The measurement metadata. Example: "process_id" -> 12345
	Metadata map[string]TypedValue `json:"metadata,omitempty"`

	// Plugin: The name of the plugin. Example: "disk".
	Plugin string `json:"plugin,omitempty"`

	// PluginInstance: The instance name of the plugin Example: "hdcl".
	PluginInstance string `json:"pluginInstance,omitempty"`

	// StartTime: The start time of the interval.
	StartTime string `json:"startTime,omitempty"`

	// Type: The measurement type. Example: "memory".
	Type string `json:"type,omitempty"`

	// TypeInstance: The measurement type instance. Example: "used".
	TypeInstance string `json:"typeInstance,omitempty"`

	// Values: The measured values during this time interval. Each value
	// must have a different dataSourceName.
	Values []*CollectdValue `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CollectdPayload) MarshalJSON() ([]byte, error) {
	type NoMethod CollectdPayload
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CollectdPayloadError: Describes the error status for payloads that
// were not written.
type CollectdPayloadError struct {
	// Error: Records the error status for the payload. If this field is
	// present, the partial errors for nested values won't be populated.
	Error *Status `json:"error,omitempty"`

	// Index: The zero-based index in
	// CreateCollectdTimeSeriesRequest.collectd_payloads.
	Index int64 `json:"index,omitempty"`

	// ValueErrors: Records the error status for values that were not
	// written due to an error.Failed payloads for which nothing is written
	// will not include partial value errors.
	ValueErrors []*CollectdValueError `json:"valueErrors,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Error") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CollectdPayloadError) MarshalJSON() ([]byte, error) {
	type NoMethod CollectdPayloadError
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CollectdValue: A single data point from a collectd-based plugin.
type CollectdValue struct {
	// DataSourceName: The data source for the collectd value. For example
	// there are two data sources for network measurements: "rx" and "tx".
	DataSourceName string `json:"dataSourceName,omitempty"`

	// DataSourceType: The type of measurement.
	//
	// Possible values:
	//   "UNSPECIFIED_DATA_SOURCE_TYPE" - An unspecified data source type.
	// This corresponds to
	// google.api.MetricDescriptor.MetricKind.METRIC_KIND_UNSPECIFIED.
	//   "GAUGE" - An instantaneous measurement of a varying quantity. This
	// corresponds to google.api.MetricDescriptor.MetricKind.GAUGE.
	//   "COUNTER" - A cumulative value over time. This corresponds to
	// google.api.MetricDescriptor.MetricKind.CUMULATIVE.
	//   "DERIVE" - A rate of change of the measurement.
	//   "ABSOLUTE" - An amount of change since the last measurement
	// interval. This corresponds to
	// google.api.MetricDescriptor.MetricKind.DELTA.
	DataSourceType string `json:"dataSourceType,omitempty"`

	// Value: The measurement value.
	Value *TypedValue `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DataSourceName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DataSourceName") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CollectdValue) MarshalJSON() ([]byte, error) {
	type NoMethod CollectdValue
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CollectdValueError: Describes the error status for values that were
// not written.
type CollectdValueError struct {
	// Error: Records the error status for the value.
	Error *Status `json:"error,omitempty"`

	// Index: The zero-based index in CollectdPayload.values within the
	// parent CreateCollectdTimeSeriesRequest.collectd_payloads.
	Index int64 `json:"index,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Error") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CollectdValueError) MarshalJSON() ([]byte, error) {
	type NoMethod CollectdValueError
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Condition: A condition is a true/false test that determines when an
// alerting policy should open an incident. If a condition evaluates to
// true, it signifies that something is wrong.
type Condition struct {
	// ConditionAbsent: A condition that checks that a time series continues
	// to receive new data points.
	ConditionAbsent *MetricAbsence `json:"conditionAbsent,omitempty"`

	// ConditionThreshold: A condition that compares a time series against a
	// threshold.
	ConditionThreshold *MetricThreshold `json:"conditionThreshold,omitempty"`

	// DisplayName: A short name or phrase used to identify the condition in
	// dashboards, notifications, and incidents. To avoid confusion, don't
	// use the same display name for multiple conditions in the same policy.
	DisplayName string `json:"displayName,omitempty"`

	// Name: Required if the condition exists. The unique resource name for
	// this condition. Its syntax
	// is:
	// projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDIT
	// ION_ID]
	// [CONDITION_ID] is assigned by Stackdriver Monitoring when the
	// condition is created as part of a new or updated alerting policy.When
	// calling the alertPolicies.create method, do not include the name
	// field in the conditions of the requested alerting policy. Stackdriver
	// Monitoring creates the condition identifiers and includes them in the
	// new policy.When calling the alertPolicies.update method to update a
	// policy, including a condition name causes the existing condition to
	// be updated. Conditions without names are added to the updated policy.
	// Existing conditions are deleted if they are not updated.Best practice
	// is to preserve [CONDITION_ID] if you make only small changes, such as
	// those to condition thresholds, durations, or trigger values.
	// Otherwise, treat the change as a new condition and let the existing
	// condition be deleted.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ConditionAbsent") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ConditionAbsent") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Condition) MarshalJSON() ([]byte, error) {
	type NoMethod Condition
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ContentMatcher: Used to perform string matching. Currently, this
// matches on the exact content. In the future, it can be expanded to
// allow for regular expressions and more complex matching.
type ContentMatcher struct {
	// Content: String content to match (max 1024 bytes)
	Content string `json:"content,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Content") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ContentMatcher) MarshalJSON() ([]byte, error) {
	type NoMethod ContentMatcher
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateCollectdTimeSeriesRequest: The CreateCollectdTimeSeries
// request.
type CreateCollectdTimeSeriesRequest struct {
	// CollectdPayloads: The collectd payloads representing the time series
	// data. You must not include more than a single point for each time
	// series, so no two payloads can have the same values for all of the
	// fields plugin, plugin_instance, type, and type_instance.
	CollectdPayloads []*CollectdPayload `json:"collectdPayloads,omitempty"`

	// CollectdVersion: The version of collectd that collected the data.
	// Example: "5.3.0-192.el6".
	CollectdVersion string `json:"collectdVersion,omitempty"`

	// Resource: The monitored resource associated with the time series.
	Resource *MonitoredResource `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CollectdPayloads") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CollectdPayloads") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CreateCollectdTimeSeriesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreateCollectdTimeSeriesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateCollectdTimeSeriesResponse: The CreateCollectdTimeSeries
// response.
type CreateCollectdTimeSeriesResponse struct {
	// PayloadErrors: Records the error status for points that were not
	// written due to an error.Failed requests for which nothing is written
	// will return an error response instead.
	PayloadErrors []*CollectdPayloadError `json:"payloadErrors,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "PayloadErrors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PayloadErrors") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateCollectdTimeSeriesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod CreateCollectdTimeSeriesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateTimeSeriesRequest: The CreateTimeSeries request.
type CreateTimeSeriesRequest struct {
	// TimeSeries: The new data to be added to a list of time series. Adds
	// at most one data point to each of several time series. The new data
	// point must be more recent than any other point in its time series.
	// Each TimeSeries value must fully specify a unique time series by
	// supplying all label values for the metric and the monitored resource.
	TimeSeries []*TimeSeries `json:"timeSeries,omitempty"`

	// ForceSendFields is a list of field names (e.g. "TimeSeries") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "TimeSeries") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateTimeSeriesRequest) MarshalJSON() ([]byte, error) {
	type NoMethod CreateTimeSeriesRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Distribution: Distribution contains summary statistics for a
// population of values. It optionally contains a histogram representing
// the distribution of those values across a set of buckets.The summary
// statistics are the count, mean, sum of the squared deviation from the
// mean, the minimum, and the maximum of the set of population of
// values. The histogram is based on a sequence of buckets and gives a
// count of values that fall into each bucket. The boundaries of the
// buckets are given either explicitly or by formulas for buckets of
// fixed or exponentially increasing widths.Although it is not
// forbidden, it is generally a bad idea to include non-finite values
// (infinities or NaNs) in the population of values, as this will render
// the mean and sum_of_squared_deviation fields meaningless.
type Distribution struct {
	// BucketCounts: Required in the Stackdriver Monitoring API v3. The
	// values for each bucket specified in bucket_options. The sum of the
	// values in bucketCounts must equal the value in the count field of the
	// Distribution object. The order of the bucket counts follows the
	// numbering schemes described for the three bucket types. The underflow
	// bucket has number 0; the finite buckets, if any, have numbers 1
	// through N-2; and the overflow bucket has number N-1. The size of
	// bucket_counts must not be greater than N. If the size is less than N,
	// then the remaining buckets are assigned values of zero.
	BucketCounts googleapi.Int64s `json:"bucketCounts,omitempty"`

	// BucketOptions: Required in the Stackdriver Monitoring API v3. Defines
	// the histogram bucket boundaries.
	BucketOptions *BucketOptions `json:"bucketOptions,omitempty"`

	// Count: The number of values in the population. Must be non-negative.
	// This value must equal the sum of the values in bucket_counts if a
	// histogram is provided.
	Count int64 `json:"count,omitempty,string"`

	// Exemplars: Must be in increasing order of value field.
	Exemplars []*Exemplar `json:"exemplars,omitempty"`

	// Mean: The arithmetic mean of the values in the population. If count
	// is zero then this field must be zero.
	Mean float64 `json:"mean,omitempty"`

	// Range: If specified, contains the range of the population values. The
	// field must not be present if the count is zero. This field is
	// presently ignored by the Stackdriver Monitoring API v3.
	Range *Range `json:"range,omitempty"`

	// SumOfSquaredDeviation: The sum of squared deviations from the mean of
	// the values in the population. For values x_i this
	// is:
	// Sum[i=1..n]((x_i - mean)^2)
	// Knuth, "The Art of Computer Programming", Vol. 2, page 323, 3rd
	// edition describes Welford's method for accumulating this sum in one
	// pass.If count is zero then this field must be zero.
	SumOfSquaredDeviation float64 `json:"sumOfSquaredDeviation,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BucketCounts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BucketCounts") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Distribution) MarshalJSON() ([]byte, error) {
	type NoMethod Distribution
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Distribution) UnmarshalJSON(data []byte) error {
	type NoMethod Distribution
	var s1 struct {
		Mean                  gensupport.JSONFloat64 `json:"mean"`
		SumOfSquaredDeviation gensupport.JSONFloat64 `json:"sumOfSquaredDeviation"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Mean = float64(s1.Mean)
	s.SumOfSquaredDeviation = float64(s1.SumOfSquaredDeviation)
	return nil
}

// Documentation: A content string and a MIME type that describes the
// content string's format.
type Documentation struct {
	// Content: The text of the documentation, interpreted according to
	// mime_type. The content may not exceed 8,192 Unicode characters and
	// may not exceed more than 10,240 bytes when encoded in UTF-8 format,
	// whichever is smaller.
	Content string `json:"content,omitempty"`

	// MimeType: The format of the content field. Presently, only the value
	// "text/markdown" is supported. See Markdown
	// (https://en.wikipedia.org/wiki/Markdown) for more information.
	MimeType string `json:"mimeType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Content") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Documentation) MarshalJSON() ([]byte, error) {
	type NoMethod Documentation
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DroppedLabels: A set of (label, value) pairs which were dropped
// during aggregation, attached to google.api.Distribution.Exemplars in
// google.api.Distribution values during aggregation.These values are
// used in combination with the label values that remain on the
// aggregated Distribution timeseries to construct the full label set
// for the exemplar values. The resulting full label set may be used to
// identify the specific task/job/instance (for example) which may be
// contributing to a long-tail, while allowing the storage savings of
// only storing aggregated distribution values for a large group.Note
// that there are no guarantees on ordering of the labels from
// exemplar-to-exemplar and from distribution-to-distribution in the
// same stream, and there may be duplicates. It is up to clients to
// resolve any ambiguities.
type DroppedLabels struct {
	// Label: Map from label to its value, for all labels dropped in any
	// aggregation.
	Label map[string]string `json:"label,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Label") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Label") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DroppedLabels) MarshalJSON() ([]byte, error) {
	type NoMethod DroppedLabels
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated empty messages in your APIs. A typical example is to use
// it as the request or the response type of an API method. For
// instance:
// service Foo {
//   rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
// }
// The JSON representation for Empty is empty JSON object {}.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// Exemplar: Exemplars are example points that may be used to annotate
// aggregated distribution values. They are metadata that gives
// information about a particular value added to a Distribution bucket,
// such as a trace ID that was active when a value was added. They may
// contain further information, such as a example values and timestamps,
// origin, etc.
type Exemplar struct {
	// Attachments: Contextual information about the example value. Examples
	// are:Trace ID:
	// type.googleapis.com/google.devtools.cloudtrace.v1.TraceLiteral
	// string: type.googleapis.com/google.protobuf.StringValueLabels dropped
	// during aggregation:
	// type.googleapis.com/google.monitoring.v3.DroppedLabelsThere may be
	// only a single attachment of any given message type in a single
	// exemplar, and this is enforced by the system.
	Attachments []googleapi.RawMessage `json:"attachments,omitempty"`

	// Timestamp: The observation (sampling) time of the above value.
	Timestamp string `json:"timestamp,omitempty"`

	// Value: Value of the exemplar point. This value determines to which
	// bucket the exemplar belongs.
	Value float64 `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Attachments") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Attachments") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Exemplar) MarshalJSON() ([]byte, error) {
	type NoMethod Exemplar
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Exemplar) UnmarshalJSON(data []byte) error {
	type NoMethod Exemplar
	var s1 struct {
		Value gensupport.JSONFloat64 `json:"value"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Value = float64(s1.Value)
	return nil
}

// Explicit: Specifies a set of buckets with arbitrary widths.There are
// size(bounds) + 1 (= N) buckets. Bucket i has the following
// boundaries:Upper bound (0 <= i < N-1): boundsi  Lower bound (1 <= i <
// N); boundsi - 1The bounds field must contain at least one element. If
// bounds has only one element, then there are no finite buckets, and
// that single element is the common boundary of the overflow and
// underflow buckets.
type Explicit struct {
	// Bounds: The values must be monotonically increasing.
	Bounds []float64 `json:"bounds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bounds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Bounds") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Explicit) MarshalJSON() ([]byte, error) {
	type NoMethod Explicit
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Exponential: Specifies an exponential sequence of buckets that have a
// width that is proportional to the value of the lower bound. Each
// bucket represents a constant relative uncertainty on a specific value
// in the bucket.There are num_finite_buckets + 2 (= N) buckets. Bucket
// i has the following boundaries:Upper bound (0 <= i < N-1): scale *
// (growth_factor ^ i).  Lower bound (1 <= i < N): scale *
// (growth_factor ^ (i - 1)).
type Exponential struct {
	// GrowthFactor: Must be greater than 1.
	GrowthFactor float64 `json:"growthFactor,omitempty"`

	// NumFiniteBuckets: Must be greater than 0.
	NumFiniteBuckets int64 `json:"numFiniteBuckets,omitempty"`

	// Scale: Must be greater than 0.
	Scale float64 `json:"scale,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GrowthFactor") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GrowthFactor") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Exponential) MarshalJSON() ([]byte, error) {
	type NoMethod Exponential
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Exponential) UnmarshalJSON(data []byte) error {
	type NoMethod Exponential
	var s1 struct {
		GrowthFactor gensupport.JSONFloat64 `json:"growthFactor"`
		Scale        gensupport.JSONFloat64 `json:"scale"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.GrowthFactor = float64(s1.GrowthFactor)
	s.Scale = float64(s1.Scale)
	return nil
}

// Field: A single field of a message type.
type Field struct {
	// Cardinality: The field cardinality.
	//
	// Possible values:
	//   "CARDINALITY_UNKNOWN" - For fields with unknown cardinality.
	//   "CARDINALITY_OPTIONAL" - For optional fields.
	//   "CARDINALITY_REQUIRED" - For required fields. Proto2 syntax only.
	//   "CARDINALITY_REPEATED" - For repeated fields.
	Cardinality string `json:"cardinality,omitempty"`

	// DefaultValue: The string value of the default value of this field.
	// Proto2 syntax only.
	DefaultValue string `json:"defaultValue,omitempty"`

	// JsonName: The field JSON name.
	JsonName string `json:"jsonName,omitempty"`

	// Kind: The field type.
	//
	// Possible values:
	//   "TYPE_UNKNOWN" - Field type unknown.
	//   "TYPE_DOUBLE" - Field type double.
	//   "TYPE_FLOAT" - Field type float.
	//   "TYPE_INT64" - Field type int64.
	//   "TYPE_UINT64" - Field type uint64.
	//   "TYPE_INT32" - Field type int32.
	//   "TYPE_FIXED64" - Field type fixed64.
	//   "TYPE_FIXED32" - Field type fixed32.
	//   "TYPE_BOOL" - Field type bool.
	//   "TYPE_STRING" - Field type string.
	//   "TYPE_GROUP" - Field type group. Proto2 syntax only, and
	// deprecated.
	//   "TYPE_MESSAGE" - Field type message.
	//   "TYPE_BYTES" - Field type bytes.
	//   "TYPE_UINT32" - Field type uint32.
	//   "TYPE_ENUM" - Field type enum.
	//   "TYPE_SFIXED32" - Field type sfixed32.
	//   "TYPE_SFIXED64" - Field type sfixed64.
	//   "TYPE_SINT32" - Field type sint32.
	//   "TYPE_SINT64" - Field type sint64.
	Kind string `json:"kind,omitempty"`

	// Name: The field name.
	Name string `json:"name,omitempty"`

	// Number: The field number.
	Number int64 `json:"number,omitempty"`

	// OneofIndex: The index of the field type in Type.oneofs, for message
	// or enumeration types. The first type has index 1; zero means the type
	// is not in the list.
	OneofIndex int64 `json:"oneofIndex,omitempty"`

	// Options: The protocol buffer options.
	Options []*Option `json:"options,omitempty"`

	// Packed: Whether to use alternative packed wire representation.
	Packed bool `json:"packed,omitempty"`

	// TypeUrl: The field type URL, without the scheme, for message or
	// enumeration types. Example:
	// "type.googleapis.com/google.protobuf.Timestamp".
	TypeUrl string `json:"typeUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cardinality") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Cardinality") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Field) MarshalJSON() ([]byte, error) {
	type NoMethod Field
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GetNotificationChannelVerificationCodeRequest: The
// GetNotificationChannelVerificationCode request.
type GetNotificationChannelVerificationCodeRequest struct {
	// ExpireTime: The desired expiration time. If specified, the API will
	// guarantee that the returned code will not be valid after the
	// specified timestamp; however, the API cannot guarantee that the
	// returned code will be valid for at least as long as the requested
	// time (the API puts an upper bound on the amount of time for which a
	// code may be valid). If omitted, a default expiration will be used,
	// which may be less than the max permissible expiration (so specifying
	// an expiration may extend the code's lifetime over omitting an
	// expiration, even though the API does impose an upper limit on the
	// maximum expiration that is permitted).
	ExpireTime string `json:"expireTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExpireTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExpireTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GetNotificationChannelVerificationCodeRequest) MarshalJSON() ([]byte, error) {
	type NoMethod GetNotificationChannelVerificationCodeRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GetNotificationChannelVerificationCodeResponse: The
// GetNotificationChannelVerificationCode request.
type GetNotificationChannelVerificationCodeResponse struct {
	// Code: The verification code, which may be used to verify other
	// channels that have an equivalent identity (i.e. other channels of the
	// same type with the same fingerprint such as other email channels with
	// the same email address or other sms channels with the same number).
	Code string `json:"code,omitempty"`

	// ExpireTime: The expiration time associated with the code that was
	// returned. If an expiration was provided in the request, this is the
	// minimum of the requested expiration in the request and the max
	// permitted expiration.
	ExpireTime string `json:"expireTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GetNotificationChannelVerificationCodeResponse) MarshalJSON() ([]byte, error) {
	type NoMethod GetNotificationChannelVerificationCodeResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Group: The description of a dynamic collection of monitored
// resources. Each group has a filter that is matched against monitored
// resources and their associated metadata. If a group's filter matches
// an available monitored resource, then that resource is a member of
// that group. Groups can contain any number of monitored resources, and
// each monitored resource can be a member of any number of
// groups.Groups can be nested in parent-child hierarchies. The
// parentName field identifies an optional parent for each group. If a
// group has a parent, then the only monitored resources available to be
// matched by the group's filter are the resources contained in the
// parent group. In other words, a group contains the monitored
// resources that match its filter and the filters of all the group's
// ancestors. A group without a parent can contain any monitored
// resource.For example, consider an infrastructure running a set of
// instances with two user-defined tags: "environment" and "role". A
// parent group has a filter, environment="production". A child of that
// parent group has a filter, role="transcoder". The parent group
// contains all instances in the production environment, regardless of
// their roles. The child group contains instances that have the
// transcoder role and are in the production environment.The monitored
// resources contained in a group can change at any moment, depending on
// what resources exist and what filters are associated with the group
// and its ancestors.
type Group struct {
	// DisplayName: A user-assigned name for this group, used only for
	// display purposes.
	DisplayName string `json:"displayName,omitempty"`

	// Filter: The filter used to determine which monitored resources belong
	// to this group.
	Filter string `json:"filter,omitempty"`

	// IsCluster: If true, the members of this group are considered to be a
	// cluster. The system can perform additional analysis on groups that
	// are clusters.
	IsCluster bool `json:"isCluster,omitempty"`

	// Name: Output only. The name of this group. The format is
	// "projects/{project_id_or_number}/groups/{group_id}". When creating a
	// group, this field is ignored and a new name is created consisting of
	// the project specified in the call to CreateGroup and a unique
	// {group_id} that is generated automatically.
	Name string `json:"name,omitempty"`

	// ParentName: The name of the group's parent, if it has one. The format
	// is "projects/{project_id_or_number}/groups/{group_id}". For groups
	// with no parent, parentName is the empty string, "".
	ParentName string `json:"parentName,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Group) MarshalJSON() ([]byte, error) {
	type NoMethod Group
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HttpCheck: Information involved in an HTTP/HTTPS uptime check
// request.
type HttpCheck struct {
	// AuthInfo: The authentication information. Optional when creating an
	// HTTP check; defaults to empty.
	AuthInfo *BasicAuthentication `json:"authInfo,omitempty"`

	// Headers: The list of headers to send as part of the uptime check
	// request. If two headers have the same key and different values, they
	// should be entered as a single header, with the value being a
	// comma-separated list of all the desired values as described at
	// https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering
	// two separate headers with the same key in a Create call will cause
	// the first to be overwritten by the second. The maximum number of
	// headers allowed is 100.
	Headers map[string]string `json:"headers,omitempty"`

	// MaskHeaders: Boolean specifiying whether to encrypt the header
	// information. Encryption should be specified for any headers related
	// to authentication that you do not wish to be seen when retrieving the
	// configuration. The server will be responsible for encrypting the
	// headers. On Get/List calls, if mask_headers is set to True then the
	// headers will be obscured with ******.
	MaskHeaders bool `json:"maskHeaders,omitempty"`

	// Path: The path to the page to run the check against. Will be combined
	// with the host (specified within the MonitoredResource) and port to
	// construct the full URL. Optional (defaults to "/").
	Path string `json:"path,omitempty"`

	// Port: The port to the page to run the check against. Will be combined
	// with host (specified within the MonitoredResource) and path to
	// construct the full URL. Optional (defaults to 80 without SSL, or 443
	// with SSL).
	Port int64 `json:"port,omitempty"`

	// UseSsl: If true, use HTTPS instead of HTTP to run the check.
	UseSsl bool `json:"useSsl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AuthInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuthInfo") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HttpCheck) MarshalJSON() ([]byte, error) {
	type NoMethod HttpCheck
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// InternalChecker: Nimbus InternalCheckers.
type InternalChecker struct {
	// CheckerId: The checker ID.
	CheckerId string `json:"checkerId,omitempty"`

	// DisplayName: The checker's human-readable name.
	DisplayName string `json:"displayName,omitempty"`

	// GcpZone: The GCP zone the uptime check should egress from. Only
	// respected for internal uptime checks, where internal_network is
	// specified.
	GcpZone string `json:"gcpZone,omitempty"`

	// Network: The internal network to perform this uptime check on.
	Network string `json:"network,omitempty"`

	// ProjectId: The GCP project ID. Not necessarily the same as the
	// project_id for the config.
	ProjectId string `json:"projectId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CheckerId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CheckerId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *InternalChecker) MarshalJSON() ([]byte, error) {
	type NoMethod InternalChecker
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LabelDescriptor: A description of a label.
type LabelDescriptor struct {
	// Description: A human-readable description for the label.
	Description string `json:"description,omitempty"`

	// Key: The label key.
	Key string `json:"key,omitempty"`

	// ValueType: The type of data that can be assigned to the label.
	//
	// Possible values:
	//   "STRING" - A variable-length string. This is the default.
	//   "BOOL" - Boolean; true or false.
	//   "INT64" - A 64-bit signed integer.
	ValueType string `json:"valueType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LabelDescriptor) MarshalJSON() ([]byte, error) {
	type NoMethod LabelDescriptor
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Linear: Specifies a linear sequence of buckets that all have the same
// width (except overflow and underflow). Each bucket represents a
// constant absolute uncertainty on the specific value in the
// bucket.There are num_finite_buckets + 2 (= N) buckets. Bucket i has
// the following boundaries:Upper bound (0 <= i < N-1): offset + (width
// * i).  Lower bound (1 <= i < N): offset + (width * (i - 1)).
type Linear struct {
	// NumFiniteBuckets: Must be greater than 0.
	NumFiniteBuckets int64 `json:"numFiniteBuckets,omitempty"`

	// Offset: Lower bound of the first bucket.
	Offset float64 `json:"offset,omitempty"`

	// Width: Must be greater than 0.
	Width float64 `json:"width,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NumFiniteBuckets") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NumFiniteBuckets") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Linear) MarshalJSON() ([]byte, error) {
	type NoMethod Linear
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Linear) UnmarshalJSON(data []byte) error {
	type NoMethod Linear
	var s1 struct {
		Offset gensupport.JSONFloat64 `json:"offset"`
		Width  gensupport.JSONFloat64 `json:"width"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Offset = float64(s1.Offset)
	s.Width = float64(s1.Width)
	return nil
}

// ListAlertPoliciesResponse: The protocol for the ListAlertPolicies
// response.
type ListAlertPoliciesResponse struct {
	// AlertPolicies: The returned alert policies.
	AlertPolicies []*AlertPolicy `json:"alertPolicies,omitempty"`

	// NextPageToken: If there might be more results than were returned,
	// then this field is set to a non-empty value. To see the additional
	// results, use that value as pageToken in the next call to this method.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AlertPolicies") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AlertPolicies") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListAlertPoliciesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListAlertPoliciesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListGroupMembersResponse: The ListGroupMembers response.
type ListGroupMembersResponse struct {
	// Members: A set of monitored resources in the group.
	Members []*MonitoredResource `json:"members,omitempty"`

	// NextPageToken: If there are more results than have been returned,
	// then this field is set to a non-empty value. To see the additional
	// results, use that value as pageToken in the next call to this method.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalSize: The total number of elements matching this request.
	TotalSize int64 `json:"totalSize,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Members") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Members") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListGroupMembersResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListGroupMembersResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListGroupsResponse: The ListGroups response.
type ListGroupsResponse struct {
	// Group: The groups that match the specified filters.
	Group []*Group `json:"group,omitempty"`

	// NextPageToken: If there are more results than have been returned,
	// then this field is set to a non-empty value. To see the additional
	// results, use that value as pageToken in the next call to this method.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Group") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Group") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListGroupsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListGroupsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListMetricDescriptorsResponse: The ListMetricDescriptors response.
type ListMetricDescriptorsResponse struct {
	// MetricDescriptors: The metric descriptors that are available to the
	// project and that match the value of filter, if present.
	MetricDescriptors []*MetricDescriptor `json:"metricDescriptors,omitempty"`

	// NextPageToken: If there are more results than have been returned,
	// then this field is set to a non-empty value. To see the additional
	// results, use that value as pageToken in the next call to this method.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "MetricDescriptors")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MetricDescriptors") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListMetricDescriptorsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListMetricDescriptorsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListMonitoredResourceDescriptorsResponse: The
// ListMonitoredResourceDescriptors response.
type ListMonitoredResourceDescriptorsResponse struct {
	// NextPageToken: If there are more results than have been returned,
	// then this field is set to a non-empty value. To see the additional
	// results, use that value as pageToken in the next call to this method.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ResourceDescriptors: The monitored resource descriptors that are
	// available to this project and that match filter, if present.
	ResourceDescriptors []*MonitoredResourceDescriptor `json:"resourceDescriptors,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListMonitoredResourceDescriptorsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListMonitoredResourceDescriptorsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListNotificationChannelDescriptorsResponse: The
// ListNotificationChannelDescriptors response.
type ListNotificationChannelDescriptorsResponse struct {
	// ChannelDescriptors: The monitored resource descriptors supported for
	// the specified project, optionally filtered.
	ChannelDescriptors []*NotificationChannelDescriptor `json:"channelDescriptors,omitempty"`

	// NextPageToken: If not empty, indicates that there may be more results
	// that match the request. Use the value in the page_token field in a
	// subsequent request to fetch the next set of results. If empty, all
	// results have been returned.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ChannelDescriptors")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ChannelDescriptors") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListNotificationChannelDescriptorsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListNotificationChannelDescriptorsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListNotificationChannelsResponse: The ListNotificationChannels
// response.
type ListNotificationChannelsResponse struct {
	// NextPageToken: If not empty, indicates that there may be more results
	// that match the request. Use the value in the page_token field in a
	// subsequent request to fetch the next set of results. If empty, all
	// results have been returned.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// NotificationChannels: The notification channels defined for the
	// specified project.
	NotificationChannels []*NotificationChannel `json:"notificationChannels,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListNotificationChannelsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListNotificationChannelsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListTimeSeriesResponse: The ListTimeSeries response.
type ListTimeSeriesResponse struct {
	// ExecutionErrors: Query execution errors that may have caused the time
	// series data returned to be incomplete.
	ExecutionErrors []*Status `json:"executionErrors,omitempty"`

	// NextPageToken: If there are more results than have been returned,
	// then this field is set to a non-empty value. To see the additional
	// results, use that value as pageToken in the next call to this method.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TimeSeries: One or more time series that match the filter included in
	// the request.
	TimeSeries []*TimeSeries `json:"timeSeries,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ExecutionErrors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExecutionErrors") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListTimeSeriesResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListTimeSeriesResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListUptimeCheckConfigsResponse: The protocol for the
// ListUptimeCheckConfigs response.
type ListUptimeCheckConfigsResponse struct {
	// NextPageToken: This field represents the pagination token to retrieve
	// the next page of results. If the value is empty, it means no further
	// results for the request. To retrieve the next page of results, the
	// value of the next_page_token is passed to the subsequent List method
	// call (in the request message's page_token field).
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalSize: The total number of uptime check configurations for the
	// project, irrespective of any pagination.
	TotalSize int64 `json:"totalSize,omitempty"`

	// UptimeCheckConfigs: The returned uptime check configurations.
	UptimeCheckConfigs []*UptimeCheckConfig `json:"uptimeCheckConfigs,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListUptimeCheckConfigsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListUptimeCheckConfigsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListUptimeCheckIpsResponse: The protocol for the ListUptimeCheckIps
// response.
type ListUptimeCheckIpsResponse struct {
	// NextPageToken: This field represents the pagination token to retrieve
	// the next page of results. If the value is empty, it means no further
	// results for the request. To retrieve the next page of results, the
	// value of the next_page_token is passed to the subsequent List method
	// call (in the request message's page_token field). NOTE: this field is
	// not yet implemented
	NextPageToken string `json:"nextPageToken,omitempty"`

	// UptimeCheckIps: The returned list of IP addresses (including region
	// and location) that the checkers run from.
	UptimeCheckIps []*UptimeCheckIp `json:"uptimeCheckIps,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListUptimeCheckIpsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListUptimeCheckIpsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Metric: A specific metric, identified by specifying values for all of
// the labels of a MetricDescriptor.
type Metric struct {
	// Labels: The set of label values that uniquely identify this metric.
	// All labels listed in the MetricDescriptor must be assigned values.
	Labels map[string]string `json:"labels,omitempty"`

	// Type: An existing metric type, see google.api.MetricDescriptor. For
	// example, custom.googleapis.com/invoice/paid/amount.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Labels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Labels") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Metric) MarshalJSON() ([]byte, error) {
	type NoMethod Metric
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MetricAbsence: A condition type that checks that monitored resources
// are reporting data. The configuration defines a metric and a set of
// monitored resources. The predicate is considered in violation when a
// time series for the specified metric of a monitored resource does not
// include any data in the specified duration.
type MetricAbsence struct {
	// Aggregations: Specifies the alignment of data points in individual
	// time series as well as how to combine the retrieved time series
	// together (such as when aggregating multiple streams on each resource
	// to a single stream for each resource or when aggregating streams
	// across all members of a group of resrouces). Multiple aggregations
	// are applied in the order specified.This field is similar to the one
	// in the MetricService.ListTimeSeries request. It is advisable to use
	// the ListTimeSeries method when debugging this field.
	Aggregations []*Aggregation `json:"aggregations,omitempty"`

	// Duration: The amount of time that a time series must fail to report
	// new data to be considered failing. Currently, only values that are a
	// multiple of a minute--e.g. 60, 120, or 300 seconds--are supported. If
	// an invalid value is given, an error will be returned. The
	// Duration.nanos field is ignored.
	Duration string `json:"duration,omitempty"`

	// Filter: A filter that identifies which time series should be compared
	// with the threshold.The filter is similar to the one that is specified
	// in the MetricService.ListTimeSeries request (that call is useful to
	// verify the time series that will be retrieved / processed) and must
	// specify the metric type and optionally may contain restrictions on
	// resource type, resource labels, and metric labels. This field may not
	// exceed 2048 Unicode characters in length.
	Filter string `json:"filter,omitempty"`

	// Trigger: The number/percent of time series for which the comparison
	// must hold in order for the condition to trigger. If unspecified, then
	// the condition will trigger if the comparison is true for any of the
	// time series that have been identified by filter and aggregations.
	Trigger *Trigger `json:"trigger,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Aggregations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Aggregations") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MetricAbsence) MarshalJSON() ([]byte, error) {
	type NoMethod MetricAbsence
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MetricDescriptor: Defines a metric type and its schema. Once a metric
// descriptor is created, deleting or altering it stops data collection
// and makes the metric type's existing data unusable.
type MetricDescriptor struct {
	// Description: A detailed description of the metric, which can be used
	// in documentation.
	Description string `json:"description,omitempty"`

	// DisplayName: A concise name for the metric, which can be displayed in
	// user interfaces. Use sentence case without an ending period, for
	// example "Request count". This field is optional but it is recommended
	// to be set for any metrics associated with user-visible concepts, such
	// as Quota.
	DisplayName string `json:"displayName,omitempty"`

	// Labels: The set of labels that can be used to describe a specific
	// instance of this metric type. For example, the
	// appengine.googleapis.com/http/server/response_latencies metric type
	// has a label for the HTTP response code, response_code, so you can
	// look at latencies for successful responses or just for responses that
	// failed.
	Labels []*LabelDescriptor `json:"labels,omitempty"`

	// Metadata: Optional. Metadata which can be used to guide usage of the
	// metric.
	Metadata *MetricDescriptorMetadata `json:"metadata,omitempty"`

	// MetricKind: Whether the metric records instantaneous values, changes
	// to a value, etc. Some combinations of metric_kind and value_type
	// might not be supported.
	//
	// Possible values:
	//   "METRIC_KIND_UNSPECIFIED" - Do not use this default value.
	//   "GAUGE" - An instantaneous measurement of a value.
	//   "DELTA" - The change in a value during a time interval.
	//   "CUMULATIVE" - A value accumulated over a time interval. Cumulative
	// measurements in a time series should have the same start time and
	// increasing end times, until an event resets the cumulative value to
	// zero and sets a new start time for the following points.
	MetricKind string `json:"metricKind,omitempty"`

	// Name: The resource name of the metric descriptor.
	Name string `json:"name,omitempty"`

	// Type: The metric type, including its DNS name prefix. The type is not
	// URL-encoded. All user-defined metric types have the DNS name
	// custom.googleapis.com or external.googleapis.com. Metric types should
	// use a natural hierarchical grouping. For
	// example:
	// "custom.googleapis.com/invoice/paid/amount"
	// "external.googlea
	// pis.com/prometheus/up"
	// "appengine.googleapis.com/http/server/response_
	// latencies"
	//
	Type string `json:"type,omitempty"`

	// Unit: The unit in which the metric value is reported. It is only
	// applicable if the value_type is INT64, DOUBLE, or DISTRIBUTION. The
	// supported units are a subset of The Unified Code for Units of Measure
	// (http://unitsofmeasure.org/ucum.html) standard:Basic units (UNIT)
	// bit bit
	// By byte
	// s second
	// min minute
	// h hour
	// d dayPrefixes (PREFIX)
	// k kilo (10**3)
	// M mega (10**6)
	// G giga (10**9)
	// T tera (10**12)
	// P peta (10**15)
	// E exa (10**18)
	// Z zetta (10**21)
	// Y yotta (10**24)
	// m milli (10**-3)
	// u micro (10**-6)
	// n nano (10**-9)
	// p pico (10**-12)
	// f femto (10**-15)
	// a atto (10**-18)
	// z zepto (10**-21)
	// y yocto (10**-24)
	// Ki kibi (2**10)
	// Mi mebi (2**20)
	// Gi gibi (2**30)
	// Ti tebi (2**40)GrammarThe grammar also includes these connectors:
	// / division (as an infix operator, e.g. 1/s).
	// . multiplication (as an infix operator, e.g. GBy.d)The grammar for a
	// unit is as follows:
	// Expression = Component { "." Component } { "/" Component }
	// ;
	//
	// Component = ( [ PREFIX ] UNIT | "%" ) [ Annotation ]
	//           | Annotation
	//           | "1"
	//           ;
	//
	// Annotation = "{" NAME "}" ;
	// Notes:
	// Annotation is just a comment if it follows a UNIT and is  equivalent
	// to 1 if it is used alone. For examples,  {requests}/s == 1/s,
	// By{transmitted}/s == By/s.
	// NAME is a sequence of non-blank printable ASCII characters not
	// containing '{' or '}'.
	// 1 represents dimensionless value 1, such as in 1/s.
	// % represents dimensionless value 1/100, and annotates values giving
	// a percentage.
	Unit string `json:"unit,omitempty"`

	// ValueType: Whether the measurement is an integer, a floating-point
	// number, etc. Some combinations of metric_kind and value_type might
	// not be supported.
	//
	// Possible values:
	//   "VALUE_TYPE_UNSPECIFIED" - Do not use this default value.
	//   "BOOL" - The value is a boolean. This value type can be used only
	// if the metric kind is GAUGE.
	//   "INT64" - The value is a signed 64-bit integer.
	//   "DOUBLE" - The value is a double precision floating point number.
	//   "STRING" - The value is a text string. This value type can be used
	// only if the metric kind is GAUGE.
	//   "DISTRIBUTION" - The value is a Distribution.
	//   "MONEY" - The value is money.
	ValueType string `json:"valueType,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MetricDescriptor) MarshalJSON() ([]byte, error) {
	type NoMethod MetricDescriptor
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MetricDescriptorMetadata: Additional annotations that can be used to
// guide the usage of a metric.
type MetricDescriptorMetadata struct {
	// IngestDelay: The delay of data points caused by ingestion. Data
	// points older than this age are guaranteed to be ingested and
	// available to be read, excluding data loss due to errors.
	IngestDelay string `json:"ingestDelay,omitempty"`

	// LaunchStage: The launch stage of the metric definition.
	//
	// Possible values:
	//   "LAUNCH_STAGE_UNSPECIFIED" - Do not use this default value.
	//   "EARLY_ACCESS" - Early Access features are limited to a closed
	// group of testers. To use these features, you must sign up in advance
	// and sign a Trusted Tester agreement (which includes confidentiality
	// provisions). These features may be unstable, changed in
	// backward-incompatible ways, and are not guaranteed to be released.
	//   "ALPHA" - Alpha is a limited availability test for releases before
	// they are cleared for widespread use. By Alpha, all significant design
	// issues are resolved and we are in the process of verifying
	// functionality. Alpha customers need to apply for access, agree to
	// applicable terms, and have their projects whitelisted. Alpha releases
	// don’t have to be feature complete, no SLAs are provided, and there
	// are no technical support obligations, but they will be far enough
	// along that customers can actually use them in test environments or
	// for limited-use tests -- just like they would in normal production
	// cases.
	//   "BETA" - Beta is the point at which we are ready to open a release
	// for any customer to use. There are no SLA or technical support
	// obligations in a Beta release. Products will be complete from a
	// feature perspective, but may have some open outstanding issues. Beta
	// releases are suitable for limited production use cases.
	//   "GA" - GA features are open to all developers and are considered
	// stable and fully qualified for production use.
	//   "DEPRECATED" - Deprecated features are scheduled to be shut down
	// and removed. For more information, see the “Deprecation Policy”
	// section of our Terms of Service (https://cloud.google.com/terms/) and
	// the Google Cloud Platform Subject to the Deprecation Policy
	// (https://cloud.google.com/terms/deprecation) documentation.
	LaunchStage string `json:"launchStage,omitempty"`

	// SamplePeriod: The sampling period of metric data points. For metrics
	// which are written periodically, consecutive data points are stored at
	// this time interval, excluding data loss due to errors. Metrics with a
	// higher granularity have a smaller sampling period.
	SamplePeriod string `json:"samplePeriod,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IngestDelay") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IngestDelay") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MetricDescriptorMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod MetricDescriptorMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MetricThreshold: A condition type that compares a collection of time
// series against a threshold.
type MetricThreshold struct {
	// Aggregations: Specifies the alignment of data points in individual
	// time series as well as how to combine the retrieved time series
	// together (such as when aggregating multiple streams on each resource
	// to a single stream for each resource or when aggregating streams
	// across all members of a group of resrouces). Multiple aggregations
	// are applied in the order specified.This field is similar to the one
	// in the MetricService.ListTimeSeries request. It is advisable to use
	// the ListTimeSeries method when debugging this field.
	Aggregations []*Aggregation `json:"aggregations,omitempty"`

	// Comparison: The comparison to apply between the time series
	// (indicated by filter and aggregation) and the threshold (indicated by
	// threshold_value). The comparison is applied on each time series, with
	// the time series on the left-hand side and the threshold on the
	// right-hand side.Only COMPARISON_LT and COMPARISON_GT are supported
	// currently.
	//
	// Possible values:
	//   "COMPARISON_UNSPECIFIED" - No ordering relationship is specified.
	//   "COMPARISON_GT" - The left argument is greater than the right
	// argument.
	//   "COMPARISON_GE" - The left argument is greater than or equal to the
	// right argument.
	//   "COMPARISON_LT" - The left argument is less than the right
	// argument.
	//   "COMPARISON_LE" - The left argument is less than or equal to the
	// right argument.
	//   "COMPARISON_EQ" - The left argument is equal to the right argument.
	//   "COMPARISON_NE" - The left argument is not equal to the right
	// argument.
	Comparison string `json:"comparison,omitempty"`

	// DenominatorAggregations: Specifies the alignment of data points in
	// individual time series selected by denominatorFilter as well as how
	// to combine the retrieved time series together (such as when
	// aggregating multiple streams on each resource to a single stream for
	// each resource or when aggregating streams across all members of a
	// group of resources).When computing ratios, the aggregations and
	// denominator_aggregations fields must use the same alignment period
	// and produce time series that have the same periodicity and
	// labels.This field is similar to the one in the
	// MetricService.ListTimeSeries request. It is advisable to use the
	// ListTimeSeries method when debugging this field.
	DenominatorAggregations []*Aggregation `json:"denominatorAggregations,omitempty"`

	// DenominatorFilter: A filter that identifies a time series that should
	// be used as the denominator of a ratio that will be compared with the
	// threshold. If a denominator_filter is specified, the time series
	// specified by the filter field will be used as the numerator.The
	// filter is similar to the one that is specified in the
	// MetricService.ListTimeSeries request (that call is useful to verify
	// the time series that will be retrieved / processed) and must specify
	// the metric type and optionally may contain restrictions on resource
	// type, resource labels, and metric labels. This field may not exceed
	// 2048 Unicode characters in length.
	DenominatorFilter string `json:"denominatorFilter,omitempty"`

	// Duration: The amount of time that a time series must violate the
	// threshold to be considered failing. Currently, only values that are a
	// multiple of a minute--e.g., 0, 60, 120, or 300 seconds--are
	// supported. If an invalid value is given, an error will be returned.
	// When choosing a duration, it is useful to keep in mind the frequency
	// of the underlying time series data (which may also be affected by any
	// alignments specified in the aggregations field); a good duration is
	// long enough so that a single outlier does not generate spurious
	// alerts, but short enough that unhealthy states are detected and
	// alerted on quickly.
	Duration string `json:"duration,omitempty"`

	// Filter: A filter that identifies which time series should be compared
	// with the threshold.The filter is similar to the one that is specified
	// in the MetricService.ListTimeSeries request (that call is useful to
	// verify the time series that will be retrieved / processed) and must
	// specify the metric type and optionally may contain restrictions on
	// resource type, resource labels, and metric labels. This field may not
	// exceed 2048 Unicode characters in length.
	Filter string `json:"filter,omitempty"`

	// ThresholdValue: A value against which to compare the time series.
	ThresholdValue float64 `json:"thresholdValue,omitempty"`

	// Trigger: The number/percent of time series for which the comparison
	// must hold in order for the condition to trigger. If unspecified, then
	// the condition will trigger if the comparison is true for any of the
	// time series that have been identified by filter and aggregations, or
	// by the ratio, if denominator_filter and denominator_aggregations are
	// specified.
	Trigger *Trigger `json:"trigger,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Aggregations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Aggregations") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MetricThreshold) MarshalJSON() ([]byte, error) {
	type NoMethod MetricThreshold
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *MetricThreshold) UnmarshalJSON(data []byte) error {
	type NoMethod MetricThreshold
	var s1 struct {
		ThresholdValue gensupport.JSONFloat64 `json:"thresholdValue"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.ThresholdValue = float64(s1.ThresholdValue)
	return nil
}

// MonitoredResource: An object representing a resource that can be used
// for monitoring, logging, billing, or other purposes. Examples include
// virtual machine instances, databases, and storage devices such as
// disks. The type field identifies a MonitoredResourceDescriptor object
// that describes the resource's schema. Information in the labels field
// identifies the actual resource and its attributes according to the
// schema. For example, a particular Compute Engine VM instance could be
// represented by the following object, because the
// MonitoredResourceDescriptor for "gce_instance" has labels
// "instance_id" and "zone":
// { "type": "gce_instance",
//   "labels": { "instance_id": "12345678901234",
//               "zone": "us-central1-a" }}
//
type MonitoredResource struct {
	// Labels: Required. Values for all of the labels listed in the
	// associated monitored resource descriptor. For example, Compute Engine
	// VM instances use the labels "project_id", "instance_id", and "zone".
	Labels map[string]string `json:"labels,omitempty"`

	// Type: Required. The monitored resource type. This field must match
	// the type field of a MonitoredResourceDescriptor object. For example,
	// the type of a Compute Engine VM instance is gce_instance.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Labels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Labels") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MonitoredResource) MarshalJSON() ([]byte, error) {
	type NoMethod MonitoredResource
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MonitoredResourceDescriptor: An object that describes the schema of a
// MonitoredResource object using a type name and a set of labels. For
// example, the monitored resource descriptor for Google Compute Engine
// VM instances has a type of "gce_instance" and specifies the use of
// the labels "instance_id" and "zone" to identify particular VM
// instances.Different APIs can support different monitored resource
// types. APIs generally provide a list method that returns the
// monitored resource descriptors used by the API.
type MonitoredResourceDescriptor struct {
	// Description: Optional. A detailed description of the monitored
	// resource type that might be used in documentation.
	Description string `json:"description,omitempty"`

	// DisplayName: Optional. A concise name for the monitored resource type
	// that might be displayed in user interfaces. It should be a Title
	// Cased Noun Phrase, without any article or other determiners. For
	// example, "Google Cloud SQL Database".
	DisplayName string `json:"displayName,omitempty"`

	// Labels: Required. A set of labels used to describe instances of this
	// monitored resource type. For example, an individual Google Cloud SQL
	// database is identified by values for the labels "database_id" and
	// "zone".
	Labels []*LabelDescriptor `json:"labels,omitempty"`

	// Name: Optional. The resource name of the monitored resource
	// descriptor:
	// "projects/{project_id}/monitoredResourceDescriptors/{type}" where
	// {type} is the value of the type field in this object and {project_id}
	// is a project ID that provides API-specific context for accessing the
	// type. APIs that do not use project information can use the resource
	// name format "monitoredResourceDescriptors/{type}".
	Name string `json:"name,omitempty"`

	// Type: Required. The monitored resource type. For example, the type
	// "cloudsql_database" represents databases in Google Cloud SQL. The
	// maximum length of this value is 256 characters.
	Type string `json:"type,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MonitoredResourceDescriptor) MarshalJSON() ([]byte, error) {
	type NoMethod MonitoredResourceDescriptor
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MonitoredResourceMetadata: Auxiliary metadata for a MonitoredResource
// object. MonitoredResource objects contain the minimum set of
// information to uniquely identify a monitored resource instance. There
// is some other useful auxiliary metadata. Monitoring and Logging use
// an ingestion pipeline to extract metadata for cloud resources of all
// types, and store the metadata in this message.
type MonitoredResourceMetadata struct {
	// SystemLabels: Output only. Values for predefined system metadata
	// labels. System labels are a kind of metadata extracted by Google,
	// including "machine_image", "vpc", "subnet_id", "security_group",
	// "name", etc. System label values can be only strings, Boolean values,
	// or a list of strings. For example:
	// { "name": "my-test-instance",
	//   "security_group": ["a", "b", "c"],
	//   "spot_instance": false }
	//
	SystemLabels googleapi.RawMessage `json:"systemLabels,omitempty"`

	// UserLabels: Output only. A map of user-defined metadata labels.
	UserLabels map[string]string `json:"userLabels,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SystemLabels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SystemLabels") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MonitoredResourceMetadata) MarshalJSON() ([]byte, error) {
	type NoMethod MonitoredResourceMetadata
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// MutationRecord: Describes a change made to a configuration.
type MutationRecord struct {
	// MutateTime: When the change occurred.
	MutateTime string `json:"mutateTime,omitempty"`

	// MutatedBy: The email address of the user making the change.
	MutatedBy string `json:"mutatedBy,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MutateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MutateTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *MutationRecord) MarshalJSON() ([]byte, error) {
	type NoMethod MutationRecord
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NotificationChannel: A NotificationChannel is a medium through which
// an alert is delivered when a policy violation is detected. Examples
// of channels include email, SMS, and third-party messaging
// applications. Fields containing sensitive information like
// authentication tokens or contact info are only partially populated on
// retrieval.
type NotificationChannel struct {
	// Description: An optional human-readable description of this
	// notification channel. This description may provide additional
	// details, beyond the display name, for the channel. This may not
	// exceeed 1024 Unicode characters.
	Description string `json:"description,omitempty"`

	// DisplayName: An optional human-readable name for this notification
	// channel. It is recommended that you specify a non-empty and unique
	// name in order to make it easier to identify the channels in your
	// project, though this is not enforced. The display name is limited to
	// 512 Unicode characters.
	DisplayName string `json:"displayName,omitempty"`

	// Enabled: Whether notifications are forwarded to the described
	// channel. This makes it possible to disable delivery of notifications
	// to a particular channel without removing the channel from all
	// alerting policies that reference the channel. This is a more
	// convenient approach when the change is temporary and you want to
	// receive notifications from the same set of alerting policies on the
	// channel at some point in the future.
	Enabled bool `json:"enabled,omitempty"`

	// Labels: Configuration fields that define the channel and its
	// behavior. The permissible and required labels are specified in the
	// NotificationChannelDescriptor.labels of the
	// NotificationChannelDescriptor corresponding to the type field.
	Labels map[string]string `json:"labels,omitempty"`

	// Name: The full REST resource name for this channel. The syntax
	// is:
	// projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]
	// The [CHANNEL_ID] is automatically assigned by the server on creation.
	Name string `json:"name,omitempty"`

	// Type: The type of the notification channel. This field matches the
	// value of the NotificationChannelDescriptor.type field.
	Type string `json:"type,omitempty"`

	// UserLabels: User-supplied key/value data that does not need to
	// conform to the corresponding NotificationChannelDescriptor's schema,
	// unlike the labels field. This field is intended to be used for
	// organizing and identifying the NotificationChannel objects.The field
	// can contain up to 64 entries. Each key and value is limited to 63
	// Unicode characters or 128 bytes, whichever is smaller. Labels and
	// values can contain only lowercase letters, numerals, underscores, and
	// dashes. Keys must begin with a letter.
	UserLabels map[string]string `json:"userLabels,omitempty"`

	// VerificationStatus: Indicates whether this channel has been verified
	// or not. On a ListNotificationChannels or GetNotificationChannel
	// operation, this field is expected to be populated.If the value is
	// UNVERIFIED, then it indicates that the channel is non-functioning (it
	// both requires verification and lacks verification); otherwise, it is
	// assumed that the channel works.If the channel is neither VERIFIED nor
	// UNVERIFIED, it implies that the channel is of a type that does not
	// require verification or that this specific channel has been exempted
	// from verification because it was created prior to verification being
	// required for channels of this type.This field cannot be modified
	// using a standard UpdateNotificationChannel operation. To change the
	// value of this field, you must call VerifyNotificationChannel.
	//
	// Possible values:
	//   "VERIFICATION_STATUS_UNSPECIFIED" - Sentinel value used to indicate
	// that the state is unknown, omitted, or is not applicable (as in the
	// case of channels that neither support nor require verification in
	// order to function).
	//   "UNVERIFIED" - The channel has yet to be verified and requires
	// verification to function. Note that this state also applies to the
	// case where the verification process has been initiated by sending a
	// verification code but where the verification code has not been
	// submitted to complete the process.
	//   "VERIFIED" - It has been proven that notifications can be received
	// on this notification channel and that someone on the project has
	// access to messages that are delivered to that channel.
	VerificationStatus string `json:"verificationStatus,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NotificationChannel) MarshalJSON() ([]byte, error) {
	type NoMethod NotificationChannel
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// NotificationChannelDescriptor: A description of a notification
// channel. The descriptor includes the properties of the channel and
// the set of labels or fields that must be specified to configure
// channels of a given type.
type NotificationChannelDescriptor struct {
	// Description: A human-readable description of the notification channel
	// type. The description may include a description of the properties of
	// the channel and pointers to external documentation.
	Description string `json:"description,omitempty"`

	// DisplayName: A human-readable name for the notification channel type.
	// This form of the name is suitable for a user interface.
	DisplayName string `json:"displayName,omitempty"`

	// Labels: The set of labels that must be defined to identify a
	// particular channel of the corresponding type. Each label includes a
	// description for how that field should be populated.
	Labels []*LabelDescriptor `json:"labels,omitempty"`

	// Name: The full REST resource name for this descriptor. The syntax
	// is:
	// projects/[PROJECT_ID]/notificationChannelDescriptors/[TYPE]
	// In the above, [TYPE] is the value of the type field.
	Name string `json:"name,omitempty"`

	// SupportedTiers: The tiers that support this notification channel; the
	// project service tier must be one of the supported_tiers.
	//
	// Possible values:
	//   "SERVICE_TIER_UNSPECIFIED" - An invalid sentinel value, used to
	// indicate that a tier has not been provided explicitly.
	//   "SERVICE_TIER_BASIC" - The Stackdriver Basic tier, a free tier of
	// service that provides basic features, a moderate allotment of logs,
	// and access to built-in metrics. A number of features are not
	// available in this tier. For more details, see the service tiers
	// documentation (https://cloud.google.com/monitoring/accounts/tiers).
	//   "SERVICE_TIER_PREMIUM" - The Stackdriver Premium tier, a higher,
	// more expensive tier of service that provides access to all
	// Stackdriver features, lets you use Stackdriver with AWS accounts, and
	// has a larger allotments for logs and metrics. For more details, see
	// the service tiers documentation
	// (https://cloud.google.com/monitoring/accounts/tiers).
	SupportedTiers []string `json:"supportedTiers,omitempty"`

	// Type: The type of notification channel, such as "email", "sms", etc.
	// Notification channel types are globally unique.
	Type string `json:"type,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Description") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NotificationChannelDescriptor) MarshalJSON() ([]byte, error) {
	type NoMethod NotificationChannelDescriptor
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Option: A protocol buffer option, which can be attached to a message,
// field, enumeration, etc.
type Option struct {
	// Name: The option's name. For protobuf built-in options (options
	// defined in descriptor.proto), this is the short name. For example,
	// "map_entry". For custom options, it should be the fully-qualified
	// name. For example, "google.api.http".
	Name string `json:"name,omitempty"`

	// Value: The option's value packed in an Any message. If the value is a
	// primitive, the corresponding wrapper type defined in
	// google/protobuf/wrappers.proto should be used. If the value is an
	// enum, it should be stored as an int32 value using the
	// google.protobuf.Int32Value type.
	Value googleapi.RawMessage `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Option) MarshalJSON() ([]byte, error) {
	type NoMethod Option
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Point: A single data point in a time series.
type Point struct {
	// Interval: The time interval to which the data point applies. For
	// GAUGE metrics, only the end time of the interval is used. For DELTA
	// metrics, the start and end time should specify a non-zero interval,
	// with subsequent points specifying contiguous and non-overlapping
	// intervals. For CUMULATIVE metrics, the start and end time should
	// specify a non-zero interval, with subsequent points specifying the
	// same start time and increasing end times, until an event resets the
	// cumulative value to zero and sets a new start time for the following
	// points.
	Interval *TimeInterval `json:"interval,omitempty"`

	// Value: The value of the data point.
	Value *TypedValue `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Interval") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Interval") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Point) MarshalJSON() ([]byte, error) {
	type NoMethod Point
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Range: The range of the population values.
type Range struct {
	// Max: The maximum of the population values.
	Max float64 `json:"max,omitempty"`

	// Min: The minimum of the population values.
	Min float64 `json:"min,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Max") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Max") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Range) MarshalJSON() ([]byte, error) {
	type NoMethod Range
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Range) UnmarshalJSON(data []byte) error {
	type NoMethod Range
	var s1 struct {
		Max gensupport.JSONFloat64 `json:"max"`
		Min gensupport.JSONFloat64 `json:"min"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Max = float64(s1.Max)
	s.Min = float64(s1.Min)
	return nil
}

// ResourceGroup: The resource submessage for group checks. It can be
// used instead of a monitored resource, when multiple resources are
// being monitored.
type ResourceGroup struct {
	// GroupId: The group of resources being monitored. Should be only the
	// group_id, not projects/<project_id>/groups/<group_id>.
	GroupId string `json:"groupId,omitempty"`

	// ResourceType: The resource type of the group members.
	//
	// Possible values:
	//   "RESOURCE_TYPE_UNSPECIFIED" - Default value (not valid).
	//   "INSTANCE" - A group of instances from Google Cloud Platform (GCP)
	// or Amazon Web Services (AWS).
	//   "AWS_ELB_LOAD_BALANCER" - A group of Amazon ELB load balancers.
	ResourceType string `json:"resourceType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GroupId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GroupId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResourceGroup) MarshalJSON() ([]byte, error) {
	type NoMethod ResourceGroup
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SendNotificationChannelVerificationCodeRequest: The
// SendNotificationChannelVerificationCode request.
type SendNotificationChannelVerificationCodeRequest struct {
}

// SourceContext: SourceContext represents information about the source
// of a protobuf element, like the file in which it is defined.
type SourceContext struct {
	// FileName: The path-qualified name of the .proto file that contained
	// the associated protobuf element. For example:
	// "google/protobuf/source_context.proto".
	FileName string `json:"fileName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FileName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FileName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SourceContext) MarshalJSON() ([]byte, error) {
	type NoMethod SourceContext
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SpanContext: The context of a span, attached to
// google.api.Distribution.Exemplars in google.api.Distribution values
// during aggregation.It contains the name of a span with format:
// projects/PROJECT_ID/traces/TRACE_ID/spans/SPAN_ID
type SpanContext struct {
	// SpanName: The resource name of the span in the following
	// format:
	// projects/[PROJECT_ID]/traces/[TRACE_ID]/spans/[SPAN_ID]
	// TRACE_
	// ID is a unique identifier for a trace within a project; it is a
	// 32-character hexadecimal encoding of a 16-byte array.SPAN_ID is a
	// unique identifier for a span within a trace; it is a 16-character
	// hexadecimal encoding of an 8-byte array.
	SpanName string `json:"spanName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SpanName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SpanName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SpanContext) MarshalJSON() ([]byte, error) {
	type NoMethod SpanContext
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Status: The Status type defines a logical error model that is
// suitable for different programming environments, including REST APIs
// and RPC APIs. It is used by gRPC (https://github.com/grpc). The error
// model is designed to be:
// Simple to use and understand for most users
// Flexible enough to meet unexpected needsOverviewThe Status message
// contains three pieces of data: error code, error message, and error
// details. The error code should be an enum value of google.rpc.Code,
// but it may accept additional error codes if needed. The error message
// should be a developer-facing English message that helps developers
// understand and resolve the error. If a localized user-facing error
// message is needed, put the localized message in the error details or
// localize it in the client. The optional error details may contain
// arbitrary information about the error. There is a predefined set of
// error detail types in the package google.rpc that can be used for
// common error conditions.Language mappingThe Status message is the
// logical representation of the error model, but it is not necessarily
// the actual wire format. When the Status message is exposed in
// different client libraries and different wire protocols, it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions in Java, but more likely mapped to some error codes in
// C.Other usesThe error model and the Status message can be used in a
// variety of environments, either with or without APIs, to provide a
// consistent developer experience across different environments.Example
// uses of this error model include:
// Partial errors. If a service needs to return partial errors to the
// client, it may embed the Status in the normal response to indicate
// the partial errors.
// Workflow errors. A typical workflow has multiple steps. Each step may
// have a Status message for error reporting.
// Batch operations. If a client uses batch request and batch response,
// the Status message should be used directly inside batch response, one
// for each error sub-response.
// Asynchronous operations. If an API call embeds asynchronous operation
// results in its response, the status of those operations should be
// represented directly using the Status message.
// Logging. If some API errors are stored in logs, the message Status
// could be used directly after any stripping needed for
// security/privacy reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details. There is a
	// common set of message types for APIs to use.
	Details []googleapi.RawMessage `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any user-facing error message should be localized and sent
	// in the google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type NoMethod Status
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TcpCheck: Information required for a TCP uptime check request.
type TcpCheck struct {
	// Port: The port to the page to run the check against. Will be combined
	// with host (specified within the MonitoredResource) to construct the
	// full URL. Required.
	Port int64 `json:"port,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Port") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Port") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TcpCheck) MarshalJSON() ([]byte, error) {
	type NoMethod TcpCheck
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimeInterval: A time interval extending just after a start time
// through an end time. If the start time is the same as the end time,
// then the interval represents a single point in time.
type TimeInterval struct {
	// EndTime: Required. The end of the time interval.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: Optional. The beginning of the time interval. The default
	// value for the start time is the end time. The start time must not be
	// later than the end time.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TimeInterval) MarshalJSON() ([]byte, error) {
	type NoMethod TimeInterval
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TimeSeries: A collection of data points that describes the
// time-varying values of a metric. A time series is identified by a
// combination of a fully-specified monitored resource and a
// fully-specified metric. This type is used for both listing and
// creating time series.
type TimeSeries struct {
	// Metadata: Output only. The associated monitored resource metadata.
	// When reading a a timeseries, this field will include metadata labels
	// that are explicitly named in the reduction. When creating a
	// timeseries, this field is ignored.
	Metadata *MonitoredResourceMetadata `json:"metadata,omitempty"`

	// Metric: The associated metric. A fully-specified metric used to
	// identify the time series.
	Metric *Metric `json:"metric,omitempty"`

	// MetricKind: The metric kind of the time series. When listing time
	// series, this metric kind might be different from the metric kind of
	// the associated metric if this time series is an alignment or
	// reduction of other time series.When creating a time series, this
	// field is optional. If present, it must be the same as the metric kind
	// of the associated metric. If the associated metric's descriptor must
	// be auto-created, then this field specifies the metric kind of the new
	// descriptor and must be either GAUGE (the default) or CUMULATIVE.
	//
	// Possible values:
	//   "METRIC_KIND_UNSPECIFIED" - Do not use this default value.
	//   "GAUGE" - An instantaneous measurement of a value.
	//   "DELTA" - The change in a value during a time interval.
	//   "CUMULATIVE" - A value accumulated over a time interval. Cumulative
	// measurements in a time series should have the same start time and
	// increasing end times, until an event resets the cumulative value to
	// zero and sets a new start time for the following points.
	MetricKind string `json:"metricKind,omitempty"`

	// Points: The data points of this time series. When listing time
	// series, points are returned in reverse time order.When creating a
	// time series, this field must contain exactly one point and the
	// point's type must be the same as the value type of the associated
	// metric. If the associated metric's descriptor must be auto-created,
	// then the value type of the descriptor is determined by the point's
	// type, which must be BOOL, INT64, DOUBLE, or DISTRIBUTION.
	Points []*Point `json:"points,omitempty"`

	// Resource: The associated monitored resource. Custom metrics can use
	// only certain monitored resource types in their time series data.
	Resource *MonitoredResource `json:"resource,omitempty"`

	// ValueType: The value type of the time series. When listing time
	// series, this value type might be different from the value type of the
	// associated metric if this time series is an alignment or reduction of
	// other time series.When creating a time series, this field is
	// optional. If present, it must be the same as the type of the data in
	// the points field.
	//
	// Possible values:
	//   "VALUE_TYPE_UNSPECIFIED" - Do not use this default value.
	//   "BOOL" - The value is a boolean. This value type can be used only
	// if the metric kind is GAUGE.
	//   "INT64" - The value is a signed 64-bit integer.
	//   "DOUBLE" - The value is a double precision floating point number.
	//   "STRING" - The value is a text string. This value type can be used
	// only if the metric kind is GAUGE.
	//   "DISTRIBUTION" - The value is a Distribution.
	//   "MONEY" - The value is money.
	ValueType string `json:"valueType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Metadata") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Metadata") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TimeSeries) MarshalJSON() ([]byte, error) {
	type NoMethod TimeSeries
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Trigger: Specifies how many time series must fail a predicate to
// trigger a condition. If not specified, then a {count: 1} trigger is
// used.
type Trigger struct {
	// Count: The absolute number of time series that must fail the
	// predicate for the condition to be triggered.
	Count int64 `json:"count,omitempty"`

	// Percent: The percentage of time series that must fail the predicate
	// for the condition to be triggered.
	Percent float64 `json:"percent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Count") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Count") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Trigger) MarshalJSON() ([]byte, error) {
	type NoMethod Trigger
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Trigger) UnmarshalJSON(data []byte) error {
	type NoMethod Trigger
	var s1 struct {
		Percent gensupport.JSONFloat64 `json:"percent"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Percent = float64(s1.Percent)
	return nil
}

// Type: A protocol buffer message type.
type Type struct {
	// Fields: The list of fields.
	Fields []*Field `json:"fields,omitempty"`

	// Name: The fully qualified message name.
	Name string `json:"name,omitempty"`

	// Oneofs: The list of types appearing in oneof definitions in this
	// type.
	Oneofs []string `json:"oneofs,omitempty"`

	// Options: The protocol buffer options.
	Options []*Option `json:"options,omitempty"`

	// SourceContext: The source context.
	SourceContext *SourceContext `json:"sourceContext,omitempty"`

	// Syntax: The source syntax.
	//
	// Possible values:
	//   "SYNTAX_PROTO2" - Syntax proto2.
	//   "SYNTAX_PROTO3" - Syntax proto3.
	Syntax string `json:"syntax,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Fields") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Fields") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Type) MarshalJSON() ([]byte, error) {
	type NoMethod Type
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TypedValue: A single strongly-typed value.
type TypedValue struct {
	// BoolValue: A Boolean value: true or false.
	BoolValue *bool `json:"boolValue,omitempty"`

	// DistributionValue: A distribution value.
	DistributionValue *Distribution `json:"distributionValue,omitempty"`

	// DoubleValue: A 64-bit double-precision floating-point number. Its
	// magnitude is approximately &plusmn;10<sup>&plusmn;300</sup> and it
	// has 16 significant digits of precision.
	DoubleValue *float64 `json:"doubleValue,omitempty"`

	// Int64Value: A 64-bit integer. Its range is approximately
	// &plusmn;9.2x10<sup>18</sup>.
	Int64Value *int64 `json:"int64Value,omitempty,string"`

	// StringValue: A variable-length string value.
	StringValue *string `json:"stringValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BoolValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BoolValue") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *TypedValue) MarshalJSON() ([]byte, error) {
	type NoMethod TypedValue
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *TypedValue) UnmarshalJSON(data []byte) error {
	type NoMethod TypedValue
	var s1 struct {
		DoubleValue *gensupport.JSONFloat64 `json:"doubleValue"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	if s1.DoubleValue != nil {
		s.DoubleValue = (*float64)(s1.DoubleValue)
	}
	return nil
}

// UptimeCheckConfig: This message configures which resources and
// services to monitor for availability.
type UptimeCheckConfig struct {
	// ContentMatchers: The expected content on the page the check is run
	// against. Currently, only the first entry in the list is supported,
	// and other entries will be ignored. The server will look for an exact
	// match of the string in the page response's content. This field is
	// optional and should only be specified if a content match is required.
	ContentMatchers []*ContentMatcher `json:"contentMatchers,omitempty"`

	// DisplayName: A human-friendly name for the uptime check
	// configuration. The display name should be unique within a Stackdriver
	// Account in order to make it easier to identify; however, uniqueness
	// is not enforced. Required.
	DisplayName string `json:"displayName,omitempty"`

	// HttpCheck: Contains information needed to make an HTTP or HTTPS
	// check.
	HttpCheck *HttpCheck `json:"httpCheck,omitempty"`

	// InternalCheckers: The internal checkers that this check will egress
	// from. If is_internal is true and this list is empty, the check will
	// egress from all InternalCheckers configured for the project that owns
	// this CheckConfig.
	InternalCheckers []*InternalChecker `json:"internalCheckers,omitempty"`

	// IsInternal: Denotes whether this is a check that egresses from
	// InternalCheckers.
	IsInternal bool `json:"isInternal,omitempty"`

	// MonitoredResource: The monitored resource
	// (https://cloud.google.com/monitoring/api/resources) associated with
	// the configuration. The following monitored resource types are
	// supported for uptime checks:  uptime_url  gce_instance  gae_app
	// aws_ec2_instance  aws_elb_load_balancer
	MonitoredResource *MonitoredResource `json:"monitoredResource,omitempty"`

	// Name: A unique resource name for this UptimeCheckConfig. The format
	// is:projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].This
	// field should be omitted when creating the uptime check configuration;
	// on create, the resource name is assigned by the server and included
	// in the response.
	Name string `json:"name,omitempty"`

	// Period: How often, in seconds, the uptime check is performed.
	// Currently, the only supported values are 60s (1 minute), 300s (5
	// minutes), 600s (10 minutes), and 900s (15 minutes). Optional,
	// defaults to 300s.
	Period string `json:"period,omitempty"`

	// ResourceGroup: The group resource associated with the configuration.
	ResourceGroup *ResourceGroup `json:"resourceGroup,omitempty"`

	// SelectedRegions: The list of regions from which the check will be
	// run. If this field is specified, enough regions to include a minimum
	// of 3 locations must be provided, or an error message is returned. Not
	// specifying this field will result in uptime checks running from all
	// regions.
	//
	// Possible values:
	//   "REGION_UNSPECIFIED" - Default value if no region is specified.
	// Will result in uptime checks running from all regions.
	//   "USA" - Allows checks to run from locations within the United
	// States of America.
	//   "EUROPE" - Allows checks to run from locations within the continent
	// of Europe.
	//   "SOUTH_AMERICA" - Allows checks to run from locations within the
	// continent of South America.
	//   "ASIA_PACIFIC" - Allows checks to run from locations within the
	// Asia Pacific area (ex: Singapore).
	SelectedRegions []string `json:"selectedRegions,omitempty"`

	// TcpCheck: Contains information needed to make a TCP check.
	TcpCheck *TcpCheck `json:"tcpCheck,omitempty"`

	// Timeout: The maximum amount of time to wait for the request to
	// complete (must be between 1 and 60 seconds). Required.
	Timeout string `json:"timeout,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ContentMatchers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContentMatchers") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *UptimeCheckConfig) MarshalJSON() ([]byte, error) {
	type NoMethod UptimeCheckConfig
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UptimeCheckIp: Contains the region, location, and list of IP
// addresses where checkers in the location run from.
type UptimeCheckIp struct {
	// IpAddress: The IP address from which the uptime check originates.
	// This is a full IP address (not an IP address range). Most IP
	// addresses, as of this publication, are in IPv4 format; however, one
	// should not rely on the IP addresses being in IPv4 format indefinitely
	// and should support interpreting this field in either IPv4 or IPv6
	// format.
	IpAddress string `json:"ipAddress,omitempty"`

	// Location: A more specific location within the region that typically
	// encodes a particular city/town/metro (and its containing
	// state/province or country) within the broader umbrella region
	// category.
	Location string `json:"location,omitempty"`

	// Region: A broad region category in which the IP address is located.
	//
	// Possible values:
	//   "REGION_UNSPECIFIED" - Default value if no region is specified.
	// Will result in uptime checks running from all regions.
	//   "USA" - Allows checks to run from locations within the United
	// States of America.
	//   "EUROPE" - Allows checks to run from locations within the continent
	// of Europe.
	//   "SOUTH_AMERICA" - Allows checks to run from locations within the
	// continent of South America.
	//   "ASIA_PACIFIC" - Allows checks to run from locations within the
	// Asia Pacific area (ex: Singapore).
	Region string `json:"region,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IpAddress") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IpAddress") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UptimeCheckIp) MarshalJSON() ([]byte, error) {
	type NoMethod UptimeCheckIp
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// VerifyNotificationChannelRequest: The VerifyNotificationChannel
// request.
type VerifyNotificationChannelRequest struct {
	// Code: The verification code that was delivered to the channel as a
	// result of invoking the SendNotificationChannelVerificationCode API
	// method or that was retrieved from a verified channel via
	// GetNotificationChannelVerificationCode. For example, one might have
	// "G-123456" or "TKNZGhhd2EyN3I1MnRnMjRv" (in general, one is only
	// guaranteed that the code is valid UTF-8; one should not make any
	// assumptions regarding the structure or format of the code).
	Code string `json:"code,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *VerifyNotificationChannelRequest) MarshalJSON() ([]byte, error) {
	type NoMethod VerifyNotificationChannelRequest
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "monitoring.projects.alertPolicies.create":

type ProjectsAlertPoliciesCreateCall struct {
	s           *Service
	name        string
	alertpolicy *AlertPolicy
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Create: Creates a new alerting policy.
func (r *ProjectsAlertPoliciesService) Create(name string, alertpolicy *AlertPolicy) *ProjectsAlertPoliciesCreateCall {
	c := &ProjectsAlertPoliciesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.alertpolicy = alertpolicy
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAlertPoliciesCreateCall) Fields(s ...googleapi.Field) *ProjectsAlertPoliciesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAlertPoliciesCreateCall) Context(ctx context.Context) *ProjectsAlertPoliciesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAlertPoliciesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAlertPoliciesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.alertpolicy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/alertPolicies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.alertPolicies.create" call.
// Exactly one of *AlertPolicy or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AlertPolicy.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAlertPoliciesCreateCall) Do(opts ...googleapi.CallOption) (*AlertPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AlertPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new alerting policy.",
	//   "flatPath": "v3/projects/{projectsId}/alertPolicies",
	//   "httpMethod": "POST",
	//   "id": "monitoring.projects.alertPolicies.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The project in which to create the alerting policy. The format is projects/[PROJECT_ID].Note that this field names the parent container in which the alerting policy will be written, not the name of the created policy. The alerting policy that is returned will have a name that contains a normalized representation of this name as a prefix but adds a suffix of the form /alertPolicies/[POLICY_ID], identifying the policy in the container.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/alertPolicies",
	//   "request": {
	//     "$ref": "AlertPolicy"
	//   },
	//   "response": {
	//     "$ref": "AlertPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.alertPolicies.delete":

type ProjectsAlertPoliciesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes an alerting policy.
func (r *ProjectsAlertPoliciesService) Delete(name string) *ProjectsAlertPoliciesDeleteCall {
	c := &ProjectsAlertPoliciesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAlertPoliciesDeleteCall) Fields(s ...googleapi.Field) *ProjectsAlertPoliciesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAlertPoliciesDeleteCall) Context(ctx context.Context) *ProjectsAlertPoliciesDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAlertPoliciesDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAlertPoliciesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.alertPolicies.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsAlertPoliciesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes an alerting policy.",
	//   "flatPath": "v3/projects/{projectsId}/alertPolicies/{alertPoliciesId}",
	//   "httpMethod": "DELETE",
	//   "id": "monitoring.projects.alertPolicies.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The alerting policy to delete. The format is:\nprojects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]\nFor more information, see AlertPolicy.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/alertPolicies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.alertPolicies.get":

type ProjectsAlertPoliciesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a single alerting policy.
func (r *ProjectsAlertPoliciesService) Get(name string) *ProjectsAlertPoliciesGetCall {
	c := &ProjectsAlertPoliciesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAlertPoliciesGetCall) Fields(s ...googleapi.Field) *ProjectsAlertPoliciesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAlertPoliciesGetCall) IfNoneMatch(entityTag string) *ProjectsAlertPoliciesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAlertPoliciesGetCall) Context(ctx context.Context) *ProjectsAlertPoliciesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAlertPoliciesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAlertPoliciesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.alertPolicies.get" call.
// Exactly one of *AlertPolicy or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AlertPolicy.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAlertPoliciesGetCall) Do(opts ...googleapi.CallOption) (*AlertPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AlertPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a single alerting policy.",
	//   "flatPath": "v3/projects/{projectsId}/alertPolicies/{alertPoliciesId}",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.alertPolicies.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The alerting policy to retrieve. The format is\nprojects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]\n",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/alertPolicies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "AlertPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// method id "monitoring.projects.alertPolicies.list":

type ProjectsAlertPoliciesListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the existing alerting policies for the project.
func (r *ProjectsAlertPoliciesService) List(name string) *ProjectsAlertPoliciesListCall {
	c := &ProjectsAlertPoliciesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": If provided, this field
// specifies the criteria that must be met by alert policies to be
// included in the response.For more details, see sorting and filtering.
func (c *ProjectsAlertPoliciesListCall) Filter(filter string) *ProjectsAlertPoliciesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// OrderBy sets the optional parameter "orderBy": A comma-separated list
// of fields by which to sort the result. Supports the same set of field
// references as the filter field. Entries can be prefixed with a minus
// sign to sort by the field in descending order.For more details, see
// sorting and filtering.
func (c *ProjectsAlertPoliciesListCall) OrderBy(orderBy string) *ProjectsAlertPoliciesListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return in a single response.
func (c *ProjectsAlertPoliciesListCall) PageSize(pageSize int64) *ProjectsAlertPoliciesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If this field is
// not empty then it must contain the nextPageToken value returned by a
// previous call to this method. Using this field causes the method to
// return more results from the previous method call.
func (c *ProjectsAlertPoliciesListCall) PageToken(pageToken string) *ProjectsAlertPoliciesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAlertPoliciesListCall) Fields(s ...googleapi.Field) *ProjectsAlertPoliciesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsAlertPoliciesListCall) IfNoneMatch(entityTag string) *ProjectsAlertPoliciesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAlertPoliciesListCall) Context(ctx context.Context) *ProjectsAlertPoliciesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAlertPoliciesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAlertPoliciesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/alertPolicies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.alertPolicies.list" call.
// Exactly one of *ListAlertPoliciesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListAlertPoliciesResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAlertPoliciesListCall) Do(opts ...googleapi.CallOption) (*ListAlertPoliciesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListAlertPoliciesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the existing alerting policies for the project.",
	//   "flatPath": "v3/projects/{projectsId}/alertPolicies",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.alertPolicies.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "If provided, this field specifies the criteria that must be met by alert policies to be included in the response.For more details, see sorting and filtering.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The project whose alert policies are to be listed. The format is\nprojects/[PROJECT_ID]\nNote that this field names the parent container in which the alerting policies to be listed are stored. To retrieve a single alerting policy by name, use the GetAlertPolicy operation, instead.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "A comma-separated list of fields by which to sort the result. Supports the same set of field references as the filter field. Entries can be prefixed with a minus sign to sort by the field in descending order.For more details, see sorting and filtering.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of results to return in a single response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If this field is not empty then it must contain the nextPageToken value returned by a previous call to this method. Using this field causes the method to return more results from the previous method call.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/alertPolicies",
	//   "response": {
	//     "$ref": "ListAlertPoliciesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsAlertPoliciesListCall) Pages(ctx context.Context, f func(*ListAlertPoliciesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "monitoring.projects.alertPolicies.patch":

type ProjectsAlertPoliciesPatchCall struct {
	s           *Service
	name        string
	alertpolicy *AlertPolicy
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// Patch: Updates an alerting policy. You can either replace the entire
// policy with a new one or replace only certain fields in the current
// alerting policy by specifying the fields to be updated via
// updateMask. Returns the updated alerting policy.
func (r *ProjectsAlertPoliciesService) Patch(name string, alertpolicy *AlertPolicy) *ProjectsAlertPoliciesPatchCall {
	c := &ProjectsAlertPoliciesPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.alertpolicy = alertpolicy
	return c
}

// UpdateMask sets the optional parameter "updateMask": A list of
// alerting policy field names. If this field is not empty, each listed
// field in the existing alerting policy is set to the value of the
// corresponding field in the supplied policy (alert_policy), or to the
// field's default value if the field is not in the supplied alerting
// policy. Fields not listed retain their previous value.Examples of
// valid field masks include display_name, documentation,
// documentation.content, documentation.mime_type, user_labels,
// user_label.nameofkey, enabled, conditions, combiner, etc.If this
// field is empty, then the supplied alerting policy replaces the
// existing policy. It is the same as deleting the existing policy and
// adding the supplied policy, except for the following:
// The new policy will have the same [ALERT_POLICY_ID] as the former
// policy. This gives you continuity with the former policy in your
// notifications and incidents.
// Conditions in the new policy will keep their former [CONDITION_ID] if
// the supplied condition includes the name field with that
// [CONDITION_ID]. If the supplied condition omits the name field, then
// a new [CONDITION_ID] is created.
func (c *ProjectsAlertPoliciesPatchCall) UpdateMask(updateMask string) *ProjectsAlertPoliciesPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAlertPoliciesPatchCall) Fields(s ...googleapi.Field) *ProjectsAlertPoliciesPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAlertPoliciesPatchCall) Context(ctx context.Context) *ProjectsAlertPoliciesPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsAlertPoliciesPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsAlertPoliciesPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.alertpolicy)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.alertPolicies.patch" call.
// Exactly one of *AlertPolicy or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *AlertPolicy.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsAlertPoliciesPatchCall) Do(opts ...googleapi.CallOption) (*AlertPolicy, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AlertPolicy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an alerting policy. You can either replace the entire policy with a new one or replace only certain fields in the current alerting policy by specifying the fields to be updated via updateMask. Returns the updated alerting policy.",
	//   "flatPath": "v3/projects/{projectsId}/alertPolicies/{alertPoliciesId}",
	//   "httpMethod": "PATCH",
	//   "id": "monitoring.projects.alertPolicies.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required if the policy exists. The resource name for this policy. The syntax is:\nprojects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]\n[ALERT_POLICY_ID] is assigned by Stackdriver Monitoring when the policy is created. When calling the alertPolicies.create method, do not include the name field in the alerting policy passed as part of the request.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/alertPolicies/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. A list of alerting policy field names. If this field is not empty, each listed field in the existing alerting policy is set to the value of the corresponding field in the supplied policy (alert_policy), or to the field's default value if the field is not in the supplied alerting policy. Fields not listed retain their previous value.Examples of valid field masks include display_name, documentation, documentation.content, documentation.mime_type, user_labels, user_label.nameofkey, enabled, conditions, combiner, etc.If this field is empty, then the supplied alerting policy replaces the existing policy. It is the same as deleting the existing policy and adding the supplied policy, except for the following:\nThe new policy will have the same [ALERT_POLICY_ID] as the former policy. This gives you continuity with the former policy in your notifications and incidents.\nConditions in the new policy will keep their former [CONDITION_ID] if the supplied condition includes the name field with that [CONDITION_ID]. If the supplied condition omits the name field, then a new [CONDITION_ID] is created.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "request": {
	//     "$ref": "AlertPolicy"
	//   },
	//   "response": {
	//     "$ref": "AlertPolicy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.collectdTimeSeries.create":

type ProjectsCollectdTimeSeriesCreateCall struct {
	s                               *Service
	name                            string
	createcollectdtimeseriesrequest *CreateCollectdTimeSeriesRequest
	urlParams_                      gensupport.URLParams
	ctx_                            context.Context
	header_                         http.Header
}

// Create: Stackdriver Monitoring Agent only: Creates a new time
// series.<aside class="caution">This method is only for use by the
// Stackdriver Monitoring Agent. Use projects.timeSeries.create
// instead.</aside>
func (r *ProjectsCollectdTimeSeriesService) Create(name string, createcollectdtimeseriesrequest *CreateCollectdTimeSeriesRequest) *ProjectsCollectdTimeSeriesCreateCall {
	c := &ProjectsCollectdTimeSeriesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.createcollectdtimeseriesrequest = createcollectdtimeseriesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsCollectdTimeSeriesCreateCall) Fields(s ...googleapi.Field) *ProjectsCollectdTimeSeriesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsCollectdTimeSeriesCreateCall) Context(ctx context.Context) *ProjectsCollectdTimeSeriesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsCollectdTimeSeriesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsCollectdTimeSeriesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createcollectdtimeseriesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/collectdTimeSeries")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.collectdTimeSeries.create" call.
// Exactly one of *CreateCollectdTimeSeriesResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *CreateCollectdTimeSeriesResponse.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsCollectdTimeSeriesCreateCall) Do(opts ...googleapi.CallOption) (*CreateCollectdTimeSeriesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CreateCollectdTimeSeriesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Stackdriver Monitoring Agent only: Creates a new time series.\u003caside class=\"caution\"\u003eThis method is only for use by the Stackdriver Monitoring Agent. Use projects.timeSeries.create instead.\u003c/aside\u003e",
	//   "flatPath": "v3/projects/{projectsId}/collectdTimeSeries",
	//   "httpMethod": "POST",
	//   "id": "monitoring.projects.collectdTimeSeries.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The project in which to create the time series. The format is \"projects/PROJECT_ID_OR_NUMBER\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/collectdTimeSeries",
	//   "request": {
	//     "$ref": "CreateCollectdTimeSeriesRequest"
	//   },
	//   "response": {
	//     "$ref": "CreateCollectdTimeSeriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.write"
	//   ]
	// }

}

// method id "monitoring.projects.groups.create":

type ProjectsGroupsCreateCall struct {
	s          *Service
	name       string
	group      *Group
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Create: Creates a new group.
func (r *ProjectsGroupsService) Create(name string, group *Group) *ProjectsGroupsCreateCall {
	c := &ProjectsGroupsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.group = group
	return c
}

// ValidateOnly sets the optional parameter "validateOnly": If true,
// validate this request but do not create the group.
func (c *ProjectsGroupsCreateCall) ValidateOnly(validateOnly bool) *ProjectsGroupsCreateCall {
	c.urlParams_.Set("validateOnly", fmt.Sprint(validateOnly))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGroupsCreateCall) Fields(s ...googleapi.Field) *ProjectsGroupsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGroupsCreateCall) Context(ctx context.Context) *ProjectsGroupsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGroupsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGroupsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/groups")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.groups.create" call.
// Exactly one of *Group or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Group.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsGroupsCreateCall) Do(opts ...googleapi.CallOption) (*Group, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Group{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new group.",
	//   "flatPath": "v3/projects/{projectsId}/groups",
	//   "httpMethod": "POST",
	//   "id": "monitoring.projects.groups.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The project in which to create the group. The format is \"projects/{project_id_or_number}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "validateOnly": {
	//       "description": "If true, validate this request but do not create the group.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "v3/{+name}/groups",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.groups.delete":

type ProjectsGroupsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes an existing group.
func (r *ProjectsGroupsService) Delete(name string) *ProjectsGroupsDeleteCall {
	c := &ProjectsGroupsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGroupsDeleteCall) Fields(s ...googleapi.Field) *ProjectsGroupsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGroupsDeleteCall) Context(ctx context.Context) *ProjectsGroupsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGroupsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGroupsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.groups.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsGroupsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes an existing group.",
	//   "flatPath": "v3/projects/{projectsId}/groups/{groupsId}",
	//   "httpMethod": "DELETE",
	//   "id": "monitoring.projects.groups.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The group to delete. The format is \"projects/{project_id_or_number}/groups/{group_id}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/groups/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.groups.get":

type ProjectsGroupsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a single group.
func (r *ProjectsGroupsService) Get(name string) *ProjectsGroupsGetCall {
	c := &ProjectsGroupsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGroupsGetCall) Fields(s ...googleapi.Field) *ProjectsGroupsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGroupsGetCall) IfNoneMatch(entityTag string) *ProjectsGroupsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGroupsGetCall) Context(ctx context.Context) *ProjectsGroupsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGroupsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGroupsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.groups.get" call.
// Exactly one of *Group or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Group.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsGroupsGetCall) Do(opts ...googleapi.CallOption) (*Group, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Group{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a single group.",
	//   "flatPath": "v3/projects/{projectsId}/groups/{groupsId}",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.groups.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The group to retrieve. The format is \"projects/{project_id_or_number}/groups/{group_id}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/groups/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// method id "monitoring.projects.groups.list":

type ProjectsGroupsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the existing groups.
func (r *ProjectsGroupsService) List(name string) *ProjectsGroupsListCall {
	c := &ProjectsGroupsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// AncestorsOfGroup sets the optional parameter "ancestorsOfGroup": A
// group name: "projects/{project_id_or_number}/groups/{group_id}".
// Returns groups that are ancestors of the specified group. The groups
// are returned in order, starting with the immediate parent and ending
// with the most distant ancestor. If the specified group has no
// immediate parent, the results are empty.
func (c *ProjectsGroupsListCall) AncestorsOfGroup(ancestorsOfGroup string) *ProjectsGroupsListCall {
	c.urlParams_.Set("ancestorsOfGroup", ancestorsOfGroup)
	return c
}

// ChildrenOfGroup sets the optional parameter "childrenOfGroup": A
// group name: "projects/{project_id_or_number}/groups/{group_id}".
// Returns groups whose parentName field contains the group name. If no
// groups have this parent, the results are empty.
func (c *ProjectsGroupsListCall) ChildrenOfGroup(childrenOfGroup string) *ProjectsGroupsListCall {
	c.urlParams_.Set("childrenOfGroup", childrenOfGroup)
	return c
}

// DescendantsOfGroup sets the optional parameter "descendantsOfGroup":
// A group name: "projects/{project_id_or_number}/groups/{group_id}".
// Returns the descendants of the specified group. This is a superset of
// the results returned by the childrenOfGroup filter, and includes
// children-of-children, and so forth.
func (c *ProjectsGroupsListCall) DescendantsOfGroup(descendantsOfGroup string) *ProjectsGroupsListCall {
	c.urlParams_.Set("descendantsOfGroup", descendantsOfGroup)
	return c
}

// PageSize sets the optional parameter "pageSize": A positive number
// that is the maximum number of results to return.
func (c *ProjectsGroupsListCall) PageSize(pageSize int64) *ProjectsGroupsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If this field is
// not empty then it must contain the nextPageToken value returned by a
// previous call to this method. Using this field causes the method to
// return additional results from the previous method call.
func (c *ProjectsGroupsListCall) PageToken(pageToken string) *ProjectsGroupsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGroupsListCall) Fields(s ...googleapi.Field) *ProjectsGroupsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGroupsListCall) IfNoneMatch(entityTag string) *ProjectsGroupsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGroupsListCall) Context(ctx context.Context) *ProjectsGroupsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGroupsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGroupsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/groups")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.groups.list" call.
// Exactly one of *ListGroupsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListGroupsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsGroupsListCall) Do(opts ...googleapi.CallOption) (*ListGroupsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListGroupsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the existing groups.",
	//   "flatPath": "v3/projects/{projectsId}/groups",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.groups.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "ancestorsOfGroup": {
	//       "description": "A group name: \"projects/{project_id_or_number}/groups/{group_id}\". Returns groups that are ancestors of the specified group. The groups are returned in order, starting with the immediate parent and ending with the most distant ancestor. If the specified group has no immediate parent, the results are empty.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "childrenOfGroup": {
	//       "description": "A group name: \"projects/{project_id_or_number}/groups/{group_id}\". Returns groups whose parentName field contains the group name. If no groups have this parent, the results are empty.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "descendantsOfGroup": {
	//       "description": "A group name: \"projects/{project_id_or_number}/groups/{group_id}\". Returns the descendants of the specified group. This is a superset of the results returned by the childrenOfGroup filter, and includes children-of-children, and so forth.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The project whose groups are to be listed. The format is \"projects/{project_id_or_number}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "A positive number that is the maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If this field is not empty then it must contain the nextPageToken value returned by a previous call to this method. Using this field causes the method to return additional results from the previous method call.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/groups",
	//   "response": {
	//     "$ref": "ListGroupsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsGroupsListCall) Pages(ctx context.Context, f func(*ListGroupsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "monitoring.projects.groups.update":

type ProjectsGroupsUpdateCall struct {
	s          *Service
	name       string
	group      *Group
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Updates an existing group. You can change any group
// attributes except name.
func (r *ProjectsGroupsService) Update(name string, group *Group) *ProjectsGroupsUpdateCall {
	c := &ProjectsGroupsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.group = group
	return c
}

// ValidateOnly sets the optional parameter "validateOnly": If true,
// validate this request but do not update the existing group.
func (c *ProjectsGroupsUpdateCall) ValidateOnly(validateOnly bool) *ProjectsGroupsUpdateCall {
	c.urlParams_.Set("validateOnly", fmt.Sprint(validateOnly))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGroupsUpdateCall) Fields(s ...googleapi.Field) *ProjectsGroupsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGroupsUpdateCall) Context(ctx context.Context) *ProjectsGroupsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGroupsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGroupsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.groups.update" call.
// Exactly one of *Group or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Group.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsGroupsUpdateCall) Do(opts ...googleapi.CallOption) (*Group, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Group{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing group. You can change any group attributes except name.",
	//   "flatPath": "v3/projects/{projectsId}/groups/{groupsId}",
	//   "httpMethod": "PUT",
	//   "id": "monitoring.projects.groups.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Output only. The name of this group. The format is \"projects/{project_id_or_number}/groups/{group_id}\". When creating a group, this field is ignored and a new name is created consisting of the project specified in the call to CreateGroup and a unique {group_id} that is generated automatically.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/groups/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "validateOnly": {
	//       "description": "If true, validate this request but do not update the existing group.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.groups.members.list":

type ProjectsGroupsMembersListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the monitored resources that are members of a group.
func (r *ProjectsGroupsMembersService) List(name string) *ProjectsGroupsMembersListCall {
	c := &ProjectsGroupsMembersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": An optional list filter
// describing the members to be returned. The filter may reference the
// type, labels, and metadata of monitored resources that comprise the
// group. For example, to return only resources representing Compute
// Engine VM instances, use this filter:
// resource.type = "gce_instance"
func (c *ProjectsGroupsMembersListCall) Filter(filter string) *ProjectsGroupsMembersListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// IntervalEndTime sets the optional parameter "interval.endTime":
// Required. The end of the time interval.
func (c *ProjectsGroupsMembersListCall) IntervalEndTime(intervalEndTime string) *ProjectsGroupsMembersListCall {
	c.urlParams_.Set("interval.endTime", intervalEndTime)
	return c
}

// IntervalStartTime sets the optional parameter "interval.startTime":
// The beginning of the time interval. The default value for the start
// time is the end time. The start time must not be later than the end
// time.
func (c *ProjectsGroupsMembersListCall) IntervalStartTime(intervalStartTime string) *ProjectsGroupsMembersListCall {
	c.urlParams_.Set("interval.startTime", intervalStartTime)
	return c
}

// PageSize sets the optional parameter "pageSize": A positive number
// that is the maximum number of results to return.
func (c *ProjectsGroupsMembersListCall) PageSize(pageSize int64) *ProjectsGroupsMembersListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If this field is
// not empty then it must contain the nextPageToken value returned by a
// previous call to this method. Using this field causes the method to
// return additional results from the previous method call.
func (c *ProjectsGroupsMembersListCall) PageToken(pageToken string) *ProjectsGroupsMembersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGroupsMembersListCall) Fields(s ...googleapi.Field) *ProjectsGroupsMembersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGroupsMembersListCall) IfNoneMatch(entityTag string) *ProjectsGroupsMembersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGroupsMembersListCall) Context(ctx context.Context) *ProjectsGroupsMembersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsGroupsMembersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsGroupsMembersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/members")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.groups.members.list" call.
// Exactly one of *ListGroupMembersResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListGroupMembersResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsGroupsMembersListCall) Do(opts ...googleapi.CallOption) (*ListGroupMembersResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListGroupMembersResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the monitored resources that are members of a group.",
	//   "flatPath": "v3/projects/{projectsId}/groups/{groupsId}/members",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.groups.members.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "An optional list filter describing the members to be returned. The filter may reference the type, labels, and metadata of monitored resources that comprise the group. For example, to return only resources representing Compute Engine VM instances, use this filter:\nresource.type = \"gce_instance\"\n",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "interval.endTime": {
	//       "description": "Required. The end of the time interval.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "interval.startTime": {
	//       "description": "Optional. The beginning of the time interval. The default value for the start time is the end time. The start time must not be later than the end time.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The group whose members are listed. The format is \"projects/{project_id_or_number}/groups/{group_id}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/groups/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "A positive number that is the maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If this field is not empty then it must contain the nextPageToken value returned by a previous call to this method. Using this field causes the method to return additional results from the previous method call.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/members",
	//   "response": {
	//     "$ref": "ListGroupMembersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsGroupsMembersListCall) Pages(ctx context.Context, f func(*ListGroupMembersResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "monitoring.projects.metricDescriptors.create":

type ProjectsMetricDescriptorsCreateCall struct {
	s                *Service
	name             string
	metricdescriptor *MetricDescriptor
	urlParams_       gensupport.URLParams
	ctx_             context.Context
	header_          http.Header
}

// Create: Creates a new metric descriptor. User-created metric
// descriptors define custom metrics.
func (r *ProjectsMetricDescriptorsService) Create(name string, metricdescriptor *MetricDescriptor) *ProjectsMetricDescriptorsCreateCall {
	c := &ProjectsMetricDescriptorsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.metricdescriptor = metricdescriptor
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricDescriptorsCreateCall) Fields(s ...googleapi.Field) *ProjectsMetricDescriptorsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMetricDescriptorsCreateCall) Context(ctx context.Context) *ProjectsMetricDescriptorsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsMetricDescriptorsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsMetricDescriptorsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.metricdescriptor)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/metricDescriptors")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.metricDescriptors.create" call.
// Exactly one of *MetricDescriptor or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *MetricDescriptor.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsMetricDescriptorsCreateCall) Do(opts ...googleapi.CallOption) (*MetricDescriptor, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &MetricDescriptor{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new metric descriptor. User-created metric descriptors define custom metrics.",
	//   "flatPath": "v3/projects/{projectsId}/metricDescriptors",
	//   "httpMethod": "POST",
	//   "id": "monitoring.projects.metricDescriptors.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The project on which to execute the request. The format is \"projects/{project_id_or_number}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/metricDescriptors",
	//   "request": {
	//     "$ref": "MetricDescriptor"
	//   },
	//   "response": {
	//     "$ref": "MetricDescriptor"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.write"
	//   ]
	// }

}

// method id "monitoring.projects.metricDescriptors.delete":

type ProjectsMetricDescriptorsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a metric descriptor. Only user-created custom metrics
// can be deleted.
func (r *ProjectsMetricDescriptorsService) Delete(name string) *ProjectsMetricDescriptorsDeleteCall {
	c := &ProjectsMetricDescriptorsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricDescriptorsDeleteCall) Fields(s ...googleapi.Field) *ProjectsMetricDescriptorsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMetricDescriptorsDeleteCall) Context(ctx context.Context) *ProjectsMetricDescriptorsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsMetricDescriptorsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsMetricDescriptorsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.metricDescriptors.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsMetricDescriptorsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a metric descriptor. Only user-created custom metrics can be deleted.",
	//   "flatPath": "v3/projects/{projectsId}/metricDescriptors/{metricDescriptorsId}",
	//   "httpMethod": "DELETE",
	//   "id": "monitoring.projects.metricDescriptors.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The metric descriptor on which to execute the request. The format is \"projects/{project_id_or_number}/metricDescriptors/{metric_id}\". An example of {metric_id} is: \"custom.googleapis.com/my_test_metric\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/metricDescriptors/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.metricDescriptors.get":

type ProjectsMetricDescriptorsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a single metric descriptor. This method does not require a
// Stackdriver account.
func (r *ProjectsMetricDescriptorsService) Get(name string) *ProjectsMetricDescriptorsGetCall {
	c := &ProjectsMetricDescriptorsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricDescriptorsGetCall) Fields(s ...googleapi.Field) *ProjectsMetricDescriptorsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsMetricDescriptorsGetCall) IfNoneMatch(entityTag string) *ProjectsMetricDescriptorsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMetricDescriptorsGetCall) Context(ctx context.Context) *ProjectsMetricDescriptorsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsMetricDescriptorsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsMetricDescriptorsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.metricDescriptors.get" call.
// Exactly one of *MetricDescriptor or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *MetricDescriptor.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsMetricDescriptorsGetCall) Do(opts ...googleapi.CallOption) (*MetricDescriptor, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &MetricDescriptor{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a single metric descriptor. This method does not require a Stackdriver account.",
	//   "flatPath": "v3/projects/{projectsId}/metricDescriptors/{metricDescriptorsId}",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.metricDescriptors.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The metric descriptor on which to execute the request. The format is \"projects/{project_id_or_number}/metricDescriptors/{metric_id}\". An example value of {metric_id} is \"compute.googleapis.com/instance/disk/read_bytes_count\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/metricDescriptors/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "MetricDescriptor"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read",
	//     "https://www.googleapis.com/auth/monitoring.write"
	//   ]
	// }

}

// method id "monitoring.projects.metricDescriptors.list":

type ProjectsMetricDescriptorsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists metric descriptors that match a filter. This method does
// not require a Stackdriver account.
func (r *ProjectsMetricDescriptorsService) List(name string) *ProjectsMetricDescriptorsListCall {
	c := &ProjectsMetricDescriptorsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": If this field is empty,
// all custom and system-defined metric descriptors are returned.
// Otherwise, the filter specifies which metric descriptors are to be
// returned. For example, the following filter matches all custom
// metrics:
// metric.type = starts_with("custom.googleapis.com/")
func (c *ProjectsMetricDescriptorsListCall) Filter(filter string) *ProjectsMetricDescriptorsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": A positive number
// that is the maximum number of results to return.
func (c *ProjectsMetricDescriptorsListCall) PageSize(pageSize int64) *ProjectsMetricDescriptorsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If this field is
// not empty then it must contain the nextPageToken value returned by a
// previous call to this method. Using this field causes the method to
// return additional results from the previous method call.
func (c *ProjectsMetricDescriptorsListCall) PageToken(pageToken string) *ProjectsMetricDescriptorsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricDescriptorsListCall) Fields(s ...googleapi.Field) *ProjectsMetricDescriptorsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsMetricDescriptorsListCall) IfNoneMatch(entityTag string) *ProjectsMetricDescriptorsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMetricDescriptorsListCall) Context(ctx context.Context) *ProjectsMetricDescriptorsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsMetricDescriptorsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsMetricDescriptorsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/metricDescriptors")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.metricDescriptors.list" call.
// Exactly one of *ListMetricDescriptorsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListMetricDescriptorsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsMetricDescriptorsListCall) Do(opts ...googleapi.CallOption) (*ListMetricDescriptorsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListMetricDescriptorsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists metric descriptors that match a filter. This method does not require a Stackdriver account.",
	//   "flatPath": "v3/projects/{projectsId}/metricDescriptors",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.metricDescriptors.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "If this field is empty, all custom and system-defined metric descriptors are returned. Otherwise, the filter specifies which metric descriptors are to be returned. For example, the following filter matches all custom metrics:\nmetric.type = starts_with(\"custom.googleapis.com/\")\n",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The project on which to execute the request. The format is \"projects/{project_id_or_number}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "A positive number that is the maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If this field is not empty then it must contain the nextPageToken value returned by a previous call to this method. Using this field causes the method to return additional results from the previous method call.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/metricDescriptors",
	//   "response": {
	//     "$ref": "ListMetricDescriptorsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read",
	//     "https://www.googleapis.com/auth/monitoring.write"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsMetricDescriptorsListCall) Pages(ctx context.Context, f func(*ListMetricDescriptorsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "monitoring.projects.monitoredResourceDescriptors.get":

type ProjectsMonitoredResourceDescriptorsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a single monitored resource descriptor. This method does
// not require a Stackdriver account.
func (r *ProjectsMonitoredResourceDescriptorsService) Get(name string) *ProjectsMonitoredResourceDescriptorsGetCall {
	c := &ProjectsMonitoredResourceDescriptorsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMonitoredResourceDescriptorsGetCall) Fields(s ...googleapi.Field) *ProjectsMonitoredResourceDescriptorsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsMonitoredResourceDescriptorsGetCall) IfNoneMatch(entityTag string) *ProjectsMonitoredResourceDescriptorsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMonitoredResourceDescriptorsGetCall) Context(ctx context.Context) *ProjectsMonitoredResourceDescriptorsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsMonitoredResourceDescriptorsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsMonitoredResourceDescriptorsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.monitoredResourceDescriptors.get" call.
// Exactly one of *MonitoredResourceDescriptor or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *MonitoredResourceDescriptor.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsMonitoredResourceDescriptorsGetCall) Do(opts ...googleapi.CallOption) (*MonitoredResourceDescriptor, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &MonitoredResourceDescriptor{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a single monitored resource descriptor. This method does not require a Stackdriver account.",
	//   "flatPath": "v3/projects/{projectsId}/monitoredResourceDescriptors/{monitoredResourceDescriptorsId}",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.monitoredResourceDescriptors.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The monitored resource descriptor to get. The format is \"projects/{project_id_or_number}/monitoredResourceDescriptors/{resource_type}\". The {resource_type} is a predefined type, such as cloudsql_database.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/monitoredResourceDescriptors/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "MonitoredResourceDescriptor"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read",
	//     "https://www.googleapis.com/auth/monitoring.write"
	//   ]
	// }

}

// method id "monitoring.projects.monitoredResourceDescriptors.list":

type ProjectsMonitoredResourceDescriptorsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists monitored resource descriptors that match a filter. This
// method does not require a Stackdriver account.
func (r *ProjectsMonitoredResourceDescriptorsService) List(name string) *ProjectsMonitoredResourceDescriptorsListCall {
	c := &ProjectsMonitoredResourceDescriptorsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": An optional filter
// describing the descriptors to be returned. The filter can reference
// the descriptor's type and labels. For example, the following filter
// returns only Google Compute Engine descriptors that have an id
// label:
// resource.type = starts_with("gce_") AND resource.label:id
func (c *ProjectsMonitoredResourceDescriptorsListCall) Filter(filter string) *ProjectsMonitoredResourceDescriptorsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": A positive number
// that is the maximum number of results to return.
func (c *ProjectsMonitoredResourceDescriptorsListCall) PageSize(pageSize int64) *ProjectsMonitoredResourceDescriptorsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If this field is
// not empty then it must contain the nextPageToken value returned by a
// previous call to this method. Using this field causes the method to
// return additional results from the previous method call.
func (c *ProjectsMonitoredResourceDescriptorsListCall) PageToken(pageToken string) *ProjectsMonitoredResourceDescriptorsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMonitoredResourceDescriptorsListCall) Fields(s ...googleapi.Field) *ProjectsMonitoredResourceDescriptorsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsMonitoredResourceDescriptorsListCall) IfNoneMatch(entityTag string) *ProjectsMonitoredResourceDescriptorsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMonitoredResourceDescriptorsListCall) Context(ctx context.Context) *ProjectsMonitoredResourceDescriptorsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsMonitoredResourceDescriptorsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsMonitoredResourceDescriptorsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/monitoredResourceDescriptors")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.monitoredResourceDescriptors.list" call.
// Exactly one of *ListMonitoredResourceDescriptorsResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *ListMonitoredResourceDescriptorsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsMonitoredResourceDescriptorsListCall) Do(opts ...googleapi.CallOption) (*ListMonitoredResourceDescriptorsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListMonitoredResourceDescriptorsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists monitored resource descriptors that match a filter. This method does not require a Stackdriver account.",
	//   "flatPath": "v3/projects/{projectsId}/monitoredResourceDescriptors",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.monitoredResourceDescriptors.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "An optional filter describing the descriptors to be returned. The filter can reference the descriptor's type and labels. For example, the following filter returns only Google Compute Engine descriptors that have an id label:\nresource.type = starts_with(\"gce_\") AND resource.label:id\n",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The project on which to execute the request. The format is \"projects/{project_id_or_number}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "A positive number that is the maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If this field is not empty then it must contain the nextPageToken value returned by a previous call to this method. Using this field causes the method to return additional results from the previous method call.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/monitoredResourceDescriptors",
	//   "response": {
	//     "$ref": "ListMonitoredResourceDescriptorsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read",
	//     "https://www.googleapis.com/auth/monitoring.write"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsMonitoredResourceDescriptorsListCall) Pages(ctx context.Context, f func(*ListMonitoredResourceDescriptorsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "monitoring.projects.notificationChannelDescriptors.get":

type ProjectsNotificationChannelDescriptorsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a single channel descriptor. The descriptor indicates which
// fields are expected / permitted for a notification channel of the
// given type.
func (r *ProjectsNotificationChannelDescriptorsService) Get(name string) *ProjectsNotificationChannelDescriptorsGetCall {
	c := &ProjectsNotificationChannelDescriptorsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotificationChannelDescriptorsGetCall) Fields(s ...googleapi.Field) *ProjectsNotificationChannelDescriptorsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsNotificationChannelDescriptorsGetCall) IfNoneMatch(entityTag string) *ProjectsNotificationChannelDescriptorsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotificationChannelDescriptorsGetCall) Context(ctx context.Context) *ProjectsNotificationChannelDescriptorsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsNotificationChannelDescriptorsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsNotificationChannelDescriptorsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.notificationChannelDescriptors.get" call.
// Exactly one of *NotificationChannelDescriptor or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *NotificationChannelDescriptor.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsNotificationChannelDescriptorsGetCall) Do(opts ...googleapi.CallOption) (*NotificationChannelDescriptor, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &NotificationChannelDescriptor{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a single channel descriptor. The descriptor indicates which fields are expected / permitted for a notification channel of the given type.",
	//   "flatPath": "v3/projects/{projectsId}/notificationChannelDescriptors/{notificationChannelDescriptorsId}",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.notificationChannelDescriptors.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The channel type for which to execute the request. The format is projects/[PROJECT_ID]/notificationChannelDescriptors/{channel_type}.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notificationChannelDescriptors/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "NotificationChannelDescriptor"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// method id "monitoring.projects.notificationChannelDescriptors.list":

type ProjectsNotificationChannelDescriptorsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the descriptors for supported channel types. The use of
// descriptors makes it possible for new channel types to be dynamically
// added.
func (r *ProjectsNotificationChannelDescriptorsService) List(name string) *ProjectsNotificationChannelDescriptorsListCall {
	c := &ProjectsNotificationChannelDescriptorsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return in a single response. If not set to a positive
// number, a reasonable value will be chosen by the service.
func (c *ProjectsNotificationChannelDescriptorsListCall) PageSize(pageSize int64) *ProjectsNotificationChannelDescriptorsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If non-empty,
// page_token must contain a value returned as the next_page_token in a
// previous response to request the next set of results.
func (c *ProjectsNotificationChannelDescriptorsListCall) PageToken(pageToken string) *ProjectsNotificationChannelDescriptorsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotificationChannelDescriptorsListCall) Fields(s ...googleapi.Field) *ProjectsNotificationChannelDescriptorsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsNotificationChannelDescriptorsListCall) IfNoneMatch(entityTag string) *ProjectsNotificationChannelDescriptorsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotificationChannelDescriptorsListCall) Context(ctx context.Context) *ProjectsNotificationChannelDescriptorsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsNotificationChannelDescriptorsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsNotificationChannelDescriptorsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/notificationChannelDescriptors")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.notificationChannelDescriptors.list" call.
// Exactly one of *ListNotificationChannelDescriptorsResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *ListNotificationChannelDescriptorsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsNotificationChannelDescriptorsListCall) Do(opts ...googleapi.CallOption) (*ListNotificationChannelDescriptorsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListNotificationChannelDescriptorsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the descriptors for supported channel types. The use of descriptors makes it possible for new channel types to be dynamically added.",
	//   "flatPath": "v3/projects/{projectsId}/notificationChannelDescriptors",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.notificationChannelDescriptors.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The REST resource name of the parent from which to retrieve the notification channel descriptors. The expected syntax is:\nprojects/[PROJECT_ID]\nNote that this names the parent container in which to look for the descriptors; to retrieve a single descriptor by name, use the GetNotificationChannelDescriptor operation, instead.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of results to return in a single response. If not set to a positive number, a reasonable value will be chosen by the service.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If non-empty, page_token must contain a value returned as the next_page_token in a previous response to request the next set of results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/notificationChannelDescriptors",
	//   "response": {
	//     "$ref": "ListNotificationChannelDescriptorsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsNotificationChannelDescriptorsListCall) Pages(ctx context.Context, f func(*ListNotificationChannelDescriptorsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "monitoring.projects.notificationChannels.create":

type ProjectsNotificationChannelsCreateCall struct {
	s                   *Service
	name                string
	notificationchannel *NotificationChannel
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Create: Creates a new notification channel, representing a single
// notification endpoint such as an email address, SMS number, or
// pagerduty service.
func (r *ProjectsNotificationChannelsService) Create(name string, notificationchannel *NotificationChannel) *ProjectsNotificationChannelsCreateCall {
	c := &ProjectsNotificationChannelsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.notificationchannel = notificationchannel
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotificationChannelsCreateCall) Fields(s ...googleapi.Field) *ProjectsNotificationChannelsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotificationChannelsCreateCall) Context(ctx context.Context) *ProjectsNotificationChannelsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsNotificationChannelsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsNotificationChannelsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.notificationchannel)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/notificationChannels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.notificationChannels.create" call.
// Exactly one of *NotificationChannel or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *NotificationChannel.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsNotificationChannelsCreateCall) Do(opts ...googleapi.CallOption) (*NotificationChannel, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &NotificationChannel{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new notification channel, representing a single notification endpoint such as an email address, SMS number, or pagerduty service.",
	//   "flatPath": "v3/projects/{projectsId}/notificationChannels",
	//   "httpMethod": "POST",
	//   "id": "monitoring.projects.notificationChannels.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The project on which to execute the request. The format is:\nprojects/[PROJECT_ID]\nNote that this names the container into which the channel will be written. This does not name the newly created channel. The resulting channel's name will have a normalized version of this field as a prefix, but will add /notificationChannels/[CHANNEL_ID] to identify the channel.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/notificationChannels",
	//   "request": {
	//     "$ref": "NotificationChannel"
	//   },
	//   "response": {
	//     "$ref": "NotificationChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.notificationChannels.delete":

type ProjectsNotificationChannelsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a notification channel.
func (r *ProjectsNotificationChannelsService) Delete(name string) *ProjectsNotificationChannelsDeleteCall {
	c := &ProjectsNotificationChannelsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Force sets the optional parameter "force": If true, the notification
// channel will be deleted regardless of its use in alert policies (the
// policies will be updated to remove the channel). If false, channels
// that are still referenced by an existing alerting policy will fail to
// be deleted in a delete operation.
func (c *ProjectsNotificationChannelsDeleteCall) Force(force bool) *ProjectsNotificationChannelsDeleteCall {
	c.urlParams_.Set("force", fmt.Sprint(force))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotificationChannelsDeleteCall) Fields(s ...googleapi.Field) *ProjectsNotificationChannelsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotificationChannelsDeleteCall) Context(ctx context.Context) *ProjectsNotificationChannelsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsNotificationChannelsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsNotificationChannelsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.notificationChannels.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsNotificationChannelsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a notification channel.",
	//   "flatPath": "v3/projects/{projectsId}/notificationChannels/{notificationChannelsId}",
	//   "httpMethod": "DELETE",
	//   "id": "monitoring.projects.notificationChannels.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "force": {
	//       "description": "If true, the notification channel will be deleted regardless of its use in alert policies (the policies will be updated to remove the channel). If false, channels that are still referenced by an existing alerting policy will fail to be deleted in a delete operation.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "name": {
	//       "description": "The channel for which to execute the request. The format is projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID].",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notificationChannels/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.notificationChannels.get":

type ProjectsNotificationChannelsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a single notification channel. The channel includes the
// relevant configuration details with which the channel was created.
// However, the response may truncate or omit passwords, API keys, or
// other private key matter and thus the response may not be 100%
// identical to the information that was supplied in the call to the
// create method.
func (r *ProjectsNotificationChannelsService) Get(name string) *ProjectsNotificationChannelsGetCall {
	c := &ProjectsNotificationChannelsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotificationChannelsGetCall) Fields(s ...googleapi.Field) *ProjectsNotificationChannelsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsNotificationChannelsGetCall) IfNoneMatch(entityTag string) *ProjectsNotificationChannelsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotificationChannelsGetCall) Context(ctx context.Context) *ProjectsNotificationChannelsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsNotificationChannelsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsNotificationChannelsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.notificationChannels.get" call.
// Exactly one of *NotificationChannel or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *NotificationChannel.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsNotificationChannelsGetCall) Do(opts ...googleapi.CallOption) (*NotificationChannel, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &NotificationChannel{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a single notification channel. The channel includes the relevant configuration details with which the channel was created. However, the response may truncate or omit passwords, API keys, or other private key matter and thus the response may not be 100% identical to the information that was supplied in the call to the create method.",
	//   "flatPath": "v3/projects/{projectsId}/notificationChannels/{notificationChannelsId}",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.notificationChannels.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The channel for which to execute the request. The format is projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID].",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notificationChannels/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "NotificationChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// method id "monitoring.projects.notificationChannels.getVerificationCode":

type ProjectsNotificationChannelsGetVerificationCodeCall struct {
	s                                             *Service
	name                                          string
	getnotificationchannelverificationcoderequest *GetNotificationChannelVerificationCodeRequest
	urlParams_                                    gensupport.URLParams
	ctx_                                          context.Context
	header_                                       http.Header
}

// GetVerificationCode: Requests a verification code for an already
// verified channel that can then be used in a call to
// VerifyNotificationChannel() on a different channel with an equivalent
// identity in the same or in a different project. This makes it
// possible to copy a channel between projects without requiring manual
// reverification of the channel. If the channel is not in the verified
// state, this method will fail (in other words, this may only be used
// if the SendNotificationChannelVerificationCode and
// VerifyNotificationChannel paths have already been used to put the
// given channel into the verified state).There is no guarantee that the
// verification codes returned by this method will be of a similar
// structure or form as the ones that are delivered to the channel via
// SendNotificationChannelVerificationCode; while
// VerifyNotificationChannel() will recognize both the codes delivered
// via SendNotificationChannelVerificationCode() and returned from
// GetNotificationChannelVerificationCode(), it is typically the case
// that the verification codes delivered via
// SendNotificationChannelVerificationCode() will be shorter and also
// have a shorter expiration (e.g. codes such as "G-123456") whereas
// GetVerificationCode() will typically return a much longer, websafe
// base 64 encoded string that has a longer expiration time.
func (r *ProjectsNotificationChannelsService) GetVerificationCode(name string, getnotificationchannelverificationcoderequest *GetNotificationChannelVerificationCodeRequest) *ProjectsNotificationChannelsGetVerificationCodeCall {
	c := &ProjectsNotificationChannelsGetVerificationCodeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.getnotificationchannelverificationcoderequest = getnotificationchannelverificationcoderequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotificationChannelsGetVerificationCodeCall) Fields(s ...googleapi.Field) *ProjectsNotificationChannelsGetVerificationCodeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotificationChannelsGetVerificationCodeCall) Context(ctx context.Context) *ProjectsNotificationChannelsGetVerificationCodeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsNotificationChannelsGetVerificationCodeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsNotificationChannelsGetVerificationCodeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getnotificationchannelverificationcoderequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}:getVerificationCode")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.notificationChannels.getVerificationCode" call.
// Exactly one of *GetNotificationChannelVerificationCodeResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GetNotificationChannelVerificationCodeResponse.ServerResponse.Header
// or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsNotificationChannelsGetVerificationCodeCall) Do(opts ...googleapi.CallOption) (*GetNotificationChannelVerificationCodeResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GetNotificationChannelVerificationCodeResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Requests a verification code for an already verified channel that can then be used in a call to VerifyNotificationChannel() on a different channel with an equivalent identity in the same or in a different project. This makes it possible to copy a channel between projects without requiring manual reverification of the channel. If the channel is not in the verified state, this method will fail (in other words, this may only be used if the SendNotificationChannelVerificationCode and VerifyNotificationChannel paths have already been used to put the given channel into the verified state).There is no guarantee that the verification codes returned by this method will be of a similar structure or form as the ones that are delivered to the channel via SendNotificationChannelVerificationCode; while VerifyNotificationChannel() will recognize both the codes delivered via SendNotificationChannelVerificationCode() and returned from GetNotificationChannelVerificationCode(), it is typically the case that the verification codes delivered via SendNotificationChannelVerificationCode() will be shorter and also have a shorter expiration (e.g. codes such as \"G-123456\") whereas GetVerificationCode() will typically return a much longer, websafe base 64 encoded string that has a longer expiration time.",
	//   "flatPath": "v3/projects/{projectsId}/notificationChannels/{notificationChannelsId}:getVerificationCode",
	//   "httpMethod": "POST",
	//   "id": "monitoring.projects.notificationChannels.getVerificationCode",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The notification channel for which a verification code is to be generated and retrieved. This must name a channel that is already verified; if the specified channel is not verified, the request will fail.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notificationChannels/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}:getVerificationCode",
	//   "request": {
	//     "$ref": "GetNotificationChannelVerificationCodeRequest"
	//   },
	//   "response": {
	//     "$ref": "GetNotificationChannelVerificationCodeResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.notificationChannels.list":

type ProjectsNotificationChannelsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the notification channels that have been created for the
// project.
func (r *ProjectsNotificationChannelsService) List(name string) *ProjectsNotificationChannelsListCall {
	c := &ProjectsNotificationChannelsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": If provided, this field
// specifies the criteria that must be met by notification channels to
// be included in the response.For more details, see sorting and
// filtering.
func (c *ProjectsNotificationChannelsListCall) Filter(filter string) *ProjectsNotificationChannelsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// OrderBy sets the optional parameter "orderBy": A comma-separated list
// of fields by which to sort the result. Supports the same set of
// fields as in filter. Entries can be prefixed with a minus sign to
// sort in descending rather than ascending order.For more details, see
// sorting and filtering.
func (c *ProjectsNotificationChannelsListCall) OrderBy(orderBy string) *ProjectsNotificationChannelsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return in a single response. If not set to a positive
// number, a reasonable value will be chosen by the service.
func (c *ProjectsNotificationChannelsListCall) PageSize(pageSize int64) *ProjectsNotificationChannelsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If non-empty,
// page_token must contain a value returned as the next_page_token in a
// previous response to request the next set of results.
func (c *ProjectsNotificationChannelsListCall) PageToken(pageToken string) *ProjectsNotificationChannelsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotificationChannelsListCall) Fields(s ...googleapi.Field) *ProjectsNotificationChannelsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsNotificationChannelsListCall) IfNoneMatch(entityTag string) *ProjectsNotificationChannelsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotificationChannelsListCall) Context(ctx context.Context) *ProjectsNotificationChannelsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsNotificationChannelsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsNotificationChannelsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/notificationChannels")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.notificationChannels.list" call.
// Exactly one of *ListNotificationChannelsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListNotificationChannelsResponse.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsNotificationChannelsListCall) Do(opts ...googleapi.CallOption) (*ListNotificationChannelsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListNotificationChannelsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the notification channels that have been created for the project.",
	//   "flatPath": "v3/projects/{projectsId}/notificationChannels",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.notificationChannels.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "If provided, this field specifies the criteria that must be met by notification channels to be included in the response.For more details, see sorting and filtering.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The project on which to execute the request. The format is projects/[PROJECT_ID]. That is, this names the container in which to look for the notification channels; it does not name a specific channel. To query a specific channel by REST resource name, use the GetNotificationChannel operation.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "A comma-separated list of fields by which to sort the result. Supports the same set of fields as in filter. Entries can be prefixed with a minus sign to sort in descending rather than ascending order.For more details, see sorting and filtering.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of results to return in a single response. If not set to a positive number, a reasonable value will be chosen by the service.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If non-empty, page_token must contain a value returned as the next_page_token in a previous response to request the next set of results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/notificationChannels",
	//   "response": {
	//     "$ref": "ListNotificationChannelsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsNotificationChannelsListCall) Pages(ctx context.Context, f func(*ListNotificationChannelsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "monitoring.projects.notificationChannels.patch":

type ProjectsNotificationChannelsPatchCall struct {
	s                   *Service
	name                string
	notificationchannel *NotificationChannel
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Patch: Updates a notification channel. Fields not specified in the
// field mask remain unchanged.
func (r *ProjectsNotificationChannelsService) Patch(name string, notificationchannel *NotificationChannel) *ProjectsNotificationChannelsPatchCall {
	c := &ProjectsNotificationChannelsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.notificationchannel = notificationchannel
	return c
}

// UpdateMask sets the optional parameter "updateMask": The fields to
// update.
func (c *ProjectsNotificationChannelsPatchCall) UpdateMask(updateMask string) *ProjectsNotificationChannelsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotificationChannelsPatchCall) Fields(s ...googleapi.Field) *ProjectsNotificationChannelsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotificationChannelsPatchCall) Context(ctx context.Context) *ProjectsNotificationChannelsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsNotificationChannelsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsNotificationChannelsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.notificationchannel)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.notificationChannels.patch" call.
// Exactly one of *NotificationChannel or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *NotificationChannel.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsNotificationChannelsPatchCall) Do(opts ...googleapi.CallOption) (*NotificationChannel, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &NotificationChannel{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a notification channel. Fields not specified in the field mask remain unchanged.",
	//   "flatPath": "v3/projects/{projectsId}/notificationChannels/{notificationChannelsId}",
	//   "httpMethod": "PATCH",
	//   "id": "monitoring.projects.notificationChannels.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The full REST resource name for this channel. The syntax is:\nprojects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]\nThe [CHANNEL_ID] is automatically assigned by the server on creation.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notificationChannels/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "The fields to update.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "request": {
	//     "$ref": "NotificationChannel"
	//   },
	//   "response": {
	//     "$ref": "NotificationChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.notificationChannels.sendVerificationCode":

type ProjectsNotificationChannelsSendVerificationCodeCall struct {
	s                                              *Service
	name                                           string
	sendnotificationchannelverificationcoderequest *SendNotificationChannelVerificationCodeRequest
	urlParams_                                     gensupport.URLParams
	ctx_                                           context.Context
	header_                                        http.Header
}

// SendVerificationCode: Causes a verification code to be delivered to
// the channel. The code can then be supplied in
// VerifyNotificationChannel to verify the channel.
func (r *ProjectsNotificationChannelsService) SendVerificationCode(name string, sendnotificationchannelverificationcoderequest *SendNotificationChannelVerificationCodeRequest) *ProjectsNotificationChannelsSendVerificationCodeCall {
	c := &ProjectsNotificationChannelsSendVerificationCodeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.sendnotificationchannelverificationcoderequest = sendnotificationchannelverificationcoderequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotificationChannelsSendVerificationCodeCall) Fields(s ...googleapi.Field) *ProjectsNotificationChannelsSendVerificationCodeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotificationChannelsSendVerificationCodeCall) Context(ctx context.Context) *ProjectsNotificationChannelsSendVerificationCodeCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsNotificationChannelsSendVerificationCodeCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsNotificationChannelsSendVerificationCodeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.sendnotificationchannelverificationcoderequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}:sendVerificationCode")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.notificationChannels.sendVerificationCode" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsNotificationChannelsSendVerificationCodeCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Causes a verification code to be delivered to the channel. The code can then be supplied in VerifyNotificationChannel to verify the channel.",
	//   "flatPath": "v3/projects/{projectsId}/notificationChannels/{notificationChannelsId}:sendVerificationCode",
	//   "httpMethod": "POST",
	//   "id": "monitoring.projects.notificationChannels.sendVerificationCode",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The notification channel to which to send a verification code.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notificationChannels/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}:sendVerificationCode",
	//   "request": {
	//     "$ref": "SendNotificationChannelVerificationCodeRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.notificationChannels.verify":

type ProjectsNotificationChannelsVerifyCall struct {
	s                                *Service
	name                             string
	verifynotificationchannelrequest *VerifyNotificationChannelRequest
	urlParams_                       gensupport.URLParams
	ctx_                             context.Context
	header_                          http.Header
}

// Verify: Verifies a NotificationChannel by proving receipt of the code
// delivered to the channel as a result of calling
// SendNotificationChannelVerificationCode.
func (r *ProjectsNotificationChannelsService) Verify(name string, verifynotificationchannelrequest *VerifyNotificationChannelRequest) *ProjectsNotificationChannelsVerifyCall {
	c := &ProjectsNotificationChannelsVerifyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.verifynotificationchannelrequest = verifynotificationchannelrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotificationChannelsVerifyCall) Fields(s ...googleapi.Field) *ProjectsNotificationChannelsVerifyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotificationChannelsVerifyCall) Context(ctx context.Context) *ProjectsNotificationChannelsVerifyCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsNotificationChannelsVerifyCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsNotificationChannelsVerifyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.verifynotificationchannelrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}:verify")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.notificationChannels.verify" call.
// Exactly one of *NotificationChannel or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *NotificationChannel.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsNotificationChannelsVerifyCall) Do(opts ...googleapi.CallOption) (*NotificationChannel, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &NotificationChannel{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Verifies a NotificationChannel by proving receipt of the code delivered to the channel as a result of calling SendNotificationChannelVerificationCode.",
	//   "flatPath": "v3/projects/{projectsId}/notificationChannels/{notificationChannelsId}:verify",
	//   "httpMethod": "POST",
	//   "id": "monitoring.projects.notificationChannels.verify",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The notification channel to verify.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notificationChannels/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}:verify",
	//   "request": {
	//     "$ref": "VerifyNotificationChannelRequest"
	//   },
	//   "response": {
	//     "$ref": "NotificationChannel"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.timeSeries.create":

type ProjectsTimeSeriesCreateCall struct {
	s                       *Service
	name                    string
	createtimeseriesrequest *CreateTimeSeriesRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
	header_                 http.Header
}

// Create: Creates or adds data to one or more time series. The response
// is empty if all time series in the request were written. If any time
// series could not be written, a corresponding failure message is
// included in the error response.
func (r *ProjectsTimeSeriesService) Create(name string, createtimeseriesrequest *CreateTimeSeriesRequest) *ProjectsTimeSeriesCreateCall {
	c := &ProjectsTimeSeriesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.createtimeseriesrequest = createtimeseriesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTimeSeriesCreateCall) Fields(s ...googleapi.Field) *ProjectsTimeSeriesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTimeSeriesCreateCall) Context(ctx context.Context) *ProjectsTimeSeriesCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTimeSeriesCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTimeSeriesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createtimeseriesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/timeSeries")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.timeSeries.create" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsTimeSeriesCreateCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates or adds data to one or more time series. The response is empty if all time series in the request were written. If any time series could not be written, a corresponding failure message is included in the error response.",
	//   "flatPath": "v3/projects/{projectsId}/timeSeries",
	//   "httpMethod": "POST",
	//   "id": "monitoring.projects.timeSeries.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The project on which to execute the request. The format is \"projects/{project_id_or_number}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/timeSeries",
	//   "request": {
	//     "$ref": "CreateTimeSeriesRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.write"
	//   ]
	// }

}

// method id "monitoring.projects.timeSeries.list":

type ProjectsTimeSeriesListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists time series that match a filter. This method does not
// require a Stackdriver account.
func (r *ProjectsTimeSeriesService) List(name string) *ProjectsTimeSeriesListCall {
	c := &ProjectsTimeSeriesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// AggregationAlignmentPeriod sets the optional parameter
// "aggregation.alignmentPeriod": The alignment period for per-time
// series alignment. If present, alignmentPeriod must be at least 60
// seconds. After per-time series alignment, each time series will
// contain data points only on the period boundaries. If
// perSeriesAligner is not specified or equals ALIGN_NONE, then this
// field is ignored. If perSeriesAligner is specified and does not equal
// ALIGN_NONE, then this field must be defined; otherwise an error is
// returned.
func (c *ProjectsTimeSeriesListCall) AggregationAlignmentPeriod(aggregationAlignmentPeriod string) *ProjectsTimeSeriesListCall {
	c.urlParams_.Set("aggregation.alignmentPeriod", aggregationAlignmentPeriod)
	return c
}

// AggregationCrossSeriesReducer sets the optional parameter
// "aggregation.crossSeriesReducer": The approach to be used to combine
// time series. Not all reducer functions may be applied to all time
// series, depending on the metric type and the value type of the
// original time series. Reduction may change the metric type of value
// type of the time series.Time series data must be aligned in order to
// perform cross-time series reduction. If crossSeriesReducer is
// specified, then perSeriesAligner must be specified and not equal
// ALIGN_NONE and alignmentPeriod must be specified; otherwise, an error
// is returned.
//
// Possible values:
//   "REDUCE_NONE"
//   "REDUCE_MEAN"
//   "REDUCE_MIN"
//   "REDUCE_MAX"
//   "REDUCE_SUM"
//   "REDUCE_STDDEV"
//   "REDUCE_COUNT"
//   "REDUCE_COUNT_TRUE"
//   "REDUCE_COUNT_FALSE"
//   "REDUCE_FRACTION_TRUE"
//   "REDUCE_PERCENTILE_99"
//   "REDUCE_PERCENTILE_95"
//   "REDUCE_PERCENTILE_50"
//   "REDUCE_PERCENTILE_05"
func (c *ProjectsTimeSeriesListCall) AggregationCrossSeriesReducer(aggregationCrossSeriesReducer string) *ProjectsTimeSeriesListCall {
	c.urlParams_.Set("aggregation.crossSeriesReducer", aggregationCrossSeriesReducer)
	return c
}

// AggregationGroupByFields sets the optional parameter
// "aggregation.groupByFields": The set of fields to preserve when
// crossSeriesReducer is specified. The groupByFields determine how the
// time series are partitioned into subsets prior to applying the
// aggregation function. Each subset contains time series that have the
// same value for each of the grouping fields. Each individual time
// series is a member of exactly one subset. The crossSeriesReducer is
// applied to each subset of time series. It is not possible to reduce
// across different resource types, so this field implicitly contains
// resource.type. Fields not specified in groupByFields are aggregated
// away. If groupByFields is not specified and all the time series have
// the same resource type, then the time series are aggregated into a
// single output time series. If crossSeriesReducer is not defined, this
// field is ignored.
func (c *ProjectsTimeSeriesListCall) AggregationGroupByFields(aggregationGroupByFields ...string) *ProjectsTimeSeriesListCall {
	c.urlParams_.SetMulti("aggregation.groupByFields", append([]string{}, aggregationGroupByFields...))
	return c
}

// AggregationPerSeriesAligner sets the optional parameter
// "aggregation.perSeriesAligner": The approach to be used to align
// individual time series. Not all alignment functions may be applied to
// all time series, depending on the metric type and value type of the
// original time series. Alignment may change the metric type or the
// value type of the time series.Time series data must be aligned in
// order to perform cross-time series reduction. If crossSeriesReducer
// is specified, then perSeriesAligner must be specified and not equal
// ALIGN_NONE and alignmentPeriod must be specified; otherwise, an error
// is returned.
//
// Possible values:
//   "ALIGN_NONE"
//   "ALIGN_DELTA"
//   "ALIGN_RATE"
//   "ALIGN_INTERPOLATE"
//   "ALIGN_NEXT_OLDER"
//   "ALIGN_MIN"
//   "ALIGN_MAX"
//   "ALIGN_MEAN"
//   "ALIGN_COUNT"
//   "ALIGN_SUM"
//   "ALIGN_STDDEV"
//   "ALIGN_COUNT_TRUE"
//   "ALIGN_COUNT_FALSE"
//   "ALIGN_FRACTION_TRUE"
//   "ALIGN_PERCENTILE_99"
//   "ALIGN_PERCENTILE_95"
//   "ALIGN_PERCENTILE_50"
//   "ALIGN_PERCENTILE_05"
//   "ALIGN_PERCENT_CHANGE"
func (c *ProjectsTimeSeriesListCall) AggregationPerSeriesAligner(aggregationPerSeriesAligner string) *ProjectsTimeSeriesListCall {
	c.urlParams_.Set("aggregation.perSeriesAligner", aggregationPerSeriesAligner)
	return c
}

// Filter sets the optional parameter "filter": A monitoring filter that
// specifies which time series should be returned. The filter must
// specify a single metric type, and can additionally specify metric
// labels and other information. For example:
// metric.type = "compute.googleapis.com/instance/cpu/usage_time" AND
//     metric.label.instance_name = "my-instance-name"
func (c *ProjectsTimeSeriesListCall) Filter(filter string) *ProjectsTimeSeriesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// IntervalEndTime sets the optional parameter "interval.endTime":
// Required. The end of the time interval.
func (c *ProjectsTimeSeriesListCall) IntervalEndTime(intervalEndTime string) *ProjectsTimeSeriesListCall {
	c.urlParams_.Set("interval.endTime", intervalEndTime)
	return c
}

// IntervalStartTime sets the optional parameter "interval.startTime":
// The beginning of the time interval. The default value for the start
// time is the end time. The start time must not be later than the end
// time.
func (c *ProjectsTimeSeriesListCall) IntervalStartTime(intervalStartTime string) *ProjectsTimeSeriesListCall {
	c.urlParams_.Set("interval.startTime", intervalStartTime)
	return c
}

// OrderBy sets the optional parameter "orderBy": Unsupported: must be
// left blank. The points in each time series are returned in reverse
// time order.
func (c *ProjectsTimeSeriesListCall) OrderBy(orderBy string) *ProjectsTimeSeriesListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": A positive number
// that is the maximum number of results to return. If page_size is
// empty or more than 100,000 results, the effective page_size is
// 100,000 results. If view is set to FULL, this is the maximum number
// of Points returned. If view is set to HEADERS, this is the maximum
// number of TimeSeries returned.
func (c *ProjectsTimeSeriesListCall) PageSize(pageSize int64) *ProjectsTimeSeriesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If this field is
// not empty then it must contain the nextPageToken value returned by a
// previous call to this method. Using this field causes the method to
// return additional results from the previous method call.
func (c *ProjectsTimeSeriesListCall) PageToken(pageToken string) *ProjectsTimeSeriesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// View sets the optional parameter "view": Specifies which information
// is returned about the time series.
//
// Possible values:
//   "FULL"
//   "HEADERS"
func (c *ProjectsTimeSeriesListCall) View(view string) *ProjectsTimeSeriesListCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTimeSeriesListCall) Fields(s ...googleapi.Field) *ProjectsTimeSeriesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsTimeSeriesListCall) IfNoneMatch(entityTag string) *ProjectsTimeSeriesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsTimeSeriesListCall) Context(ctx context.Context) *ProjectsTimeSeriesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsTimeSeriesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsTimeSeriesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}/timeSeries")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.timeSeries.list" call.
// Exactly one of *ListTimeSeriesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListTimeSeriesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsTimeSeriesListCall) Do(opts ...googleapi.CallOption) (*ListTimeSeriesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListTimeSeriesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists time series that match a filter. This method does not require a Stackdriver account.",
	//   "flatPath": "v3/projects/{projectsId}/timeSeries",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.timeSeries.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "aggregation.alignmentPeriod": {
	//       "description": "The alignment period for per-time series alignment. If present, alignmentPeriod must be at least 60 seconds. After per-time series alignment, each time series will contain data points only on the period boundaries. If perSeriesAligner is not specified or equals ALIGN_NONE, then this field is ignored. If perSeriesAligner is specified and does not equal ALIGN_NONE, then this field must be defined; otherwise an error is returned.",
	//       "format": "google-duration",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "aggregation.crossSeriesReducer": {
	//       "description": "The approach to be used to combine time series. Not all reducer functions may be applied to all time series, depending on the metric type and the value type of the original time series. Reduction may change the metric type of value type of the time series.Time series data must be aligned in order to perform cross-time series reduction. If crossSeriesReducer is specified, then perSeriesAligner must be specified and not equal ALIGN_NONE and alignmentPeriod must be specified; otherwise, an error is returned.",
	//       "enum": [
	//         "REDUCE_NONE",
	//         "REDUCE_MEAN",
	//         "REDUCE_MIN",
	//         "REDUCE_MAX",
	//         "REDUCE_SUM",
	//         "REDUCE_STDDEV",
	//         "REDUCE_COUNT",
	//         "REDUCE_COUNT_TRUE",
	//         "REDUCE_COUNT_FALSE",
	//         "REDUCE_FRACTION_TRUE",
	//         "REDUCE_PERCENTILE_99",
	//         "REDUCE_PERCENTILE_95",
	//         "REDUCE_PERCENTILE_50",
	//         "REDUCE_PERCENTILE_05"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "aggregation.groupByFields": {
	//       "description": "The set of fields to preserve when crossSeriesReducer is specified. The groupByFields determine how the time series are partitioned into subsets prior to applying the aggregation function. Each subset contains time series that have the same value for each of the grouping fields. Each individual time series is a member of exactly one subset. The crossSeriesReducer is applied to each subset of time series. It is not possible to reduce across different resource types, so this field implicitly contains resource.type. Fields not specified in groupByFields are aggregated away. If groupByFields is not specified and all the time series have the same resource type, then the time series are aggregated into a single output time series. If crossSeriesReducer is not defined, this field is ignored.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "aggregation.perSeriesAligner": {
	//       "description": "The approach to be used to align individual time series. Not all alignment functions may be applied to all time series, depending on the metric type and value type of the original time series. Alignment may change the metric type or the value type of the time series.Time series data must be aligned in order to perform cross-time series reduction. If crossSeriesReducer is specified, then perSeriesAligner must be specified and not equal ALIGN_NONE and alignmentPeriod must be specified; otherwise, an error is returned.",
	//       "enum": [
	//         "ALIGN_NONE",
	//         "ALIGN_DELTA",
	//         "ALIGN_RATE",
	//         "ALIGN_INTERPOLATE",
	//         "ALIGN_NEXT_OLDER",
	//         "ALIGN_MIN",
	//         "ALIGN_MAX",
	//         "ALIGN_MEAN",
	//         "ALIGN_COUNT",
	//         "ALIGN_SUM",
	//         "ALIGN_STDDEV",
	//         "ALIGN_COUNT_TRUE",
	//         "ALIGN_COUNT_FALSE",
	//         "ALIGN_FRACTION_TRUE",
	//         "ALIGN_PERCENTILE_99",
	//         "ALIGN_PERCENTILE_95",
	//         "ALIGN_PERCENTILE_50",
	//         "ALIGN_PERCENTILE_05",
	//         "ALIGN_PERCENT_CHANGE"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "A monitoring filter that specifies which time series should be returned. The filter must specify a single metric type, and can additionally specify metric labels and other information. For example:\nmetric.type = \"compute.googleapis.com/instance/cpu/usage_time\" AND\n    metric.label.instance_name = \"my-instance-name\"\n",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "interval.endTime": {
	//       "description": "Required. The end of the time interval.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "interval.startTime": {
	//       "description": "Optional. The beginning of the time interval. The default value for the start time is the end time. The start time must not be later than the end time.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The project on which to execute the request. The format is \"projects/{project_id_or_number}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "Unsupported: must be left blank. The points in each time series are returned in reverse time order.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "A positive number that is the maximum number of results to return. If page_size is empty or more than 100,000 results, the effective page_size is 100,000 results. If view is set to FULL, this is the maximum number of Points returned. If view is set to HEADERS, this is the maximum number of TimeSeries returned.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If this field is not empty then it must contain the nextPageToken value returned by a previous call to this method. Using this field causes the method to return additional results from the previous method call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Specifies which information is returned about the time series.",
	//       "enum": [
	//         "FULL",
	//         "HEADERS"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}/timeSeries",
	//   "response": {
	//     "$ref": "ListTimeSeriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsTimeSeriesListCall) Pages(ctx context.Context, f func(*ListTimeSeriesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "monitoring.projects.uptimeCheckConfigs.create":

type ProjectsUptimeCheckConfigsCreateCall struct {
	s                 *Service
	parent            string
	uptimecheckconfig *UptimeCheckConfig
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Create: Creates a new uptime check configuration.
func (r *ProjectsUptimeCheckConfigsService) Create(parent string, uptimecheckconfig *UptimeCheckConfig) *ProjectsUptimeCheckConfigsCreateCall {
	c := &ProjectsUptimeCheckConfigsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.uptimecheckconfig = uptimecheckconfig
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUptimeCheckConfigsCreateCall) Fields(s ...googleapi.Field) *ProjectsUptimeCheckConfigsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsUptimeCheckConfigsCreateCall) Context(ctx context.Context) *ProjectsUptimeCheckConfigsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsUptimeCheckConfigsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsUptimeCheckConfigsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.uptimecheckconfig)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+parent}/uptimeCheckConfigs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.uptimeCheckConfigs.create" call.
// Exactly one of *UptimeCheckConfig or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *UptimeCheckConfig.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsUptimeCheckConfigsCreateCall) Do(opts ...googleapi.CallOption) (*UptimeCheckConfig, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &UptimeCheckConfig{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new uptime check configuration.",
	//   "flatPath": "v3/projects/{projectsId}/uptimeCheckConfigs",
	//   "httpMethod": "POST",
	//   "id": "monitoring.projects.uptimeCheckConfigs.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The project in which to create the uptime check. The format  is projects/[PROJECT_ID].",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+parent}/uptimeCheckConfigs",
	//   "request": {
	//     "$ref": "UptimeCheckConfig"
	//   },
	//   "response": {
	//     "$ref": "UptimeCheckConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.uptimeCheckConfigs.delete":

type ProjectsUptimeCheckConfigsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes an uptime check configuration. Note that this method
// will fail if the uptime check configuration is referenced by an alert
// policy or other dependent configs that would be rendered invalid by
// the deletion.
func (r *ProjectsUptimeCheckConfigsService) Delete(name string) *ProjectsUptimeCheckConfigsDeleteCall {
	c := &ProjectsUptimeCheckConfigsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUptimeCheckConfigsDeleteCall) Fields(s ...googleapi.Field) *ProjectsUptimeCheckConfigsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsUptimeCheckConfigsDeleteCall) Context(ctx context.Context) *ProjectsUptimeCheckConfigsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsUptimeCheckConfigsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsUptimeCheckConfigsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.uptimeCheckConfigs.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsUptimeCheckConfigsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes an uptime check configuration. Note that this method will fail if the uptime check configuration is referenced by an alert policy or other dependent configs that would be rendered invalid by the deletion.",
	//   "flatPath": "v3/projects/{projectsId}/uptimeCheckConfigs/{uptimeCheckConfigsId}",
	//   "httpMethod": "DELETE",
	//   "id": "monitoring.projects.uptimeCheckConfigs.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The uptime check configuration to delete. The format  is projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/uptimeCheckConfigs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.projects.uptimeCheckConfigs.get":

type ProjectsUptimeCheckConfigsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a single uptime check configuration.
func (r *ProjectsUptimeCheckConfigsService) Get(name string) *ProjectsUptimeCheckConfigsGetCall {
	c := &ProjectsUptimeCheckConfigsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUptimeCheckConfigsGetCall) Fields(s ...googleapi.Field) *ProjectsUptimeCheckConfigsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsUptimeCheckConfigsGetCall) IfNoneMatch(entityTag string) *ProjectsUptimeCheckConfigsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsUptimeCheckConfigsGetCall) Context(ctx context.Context) *ProjectsUptimeCheckConfigsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsUptimeCheckConfigsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsUptimeCheckConfigsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.uptimeCheckConfigs.get" call.
// Exactly one of *UptimeCheckConfig or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *UptimeCheckConfig.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsUptimeCheckConfigsGetCall) Do(opts ...googleapi.CallOption) (*UptimeCheckConfig, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &UptimeCheckConfig{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a single uptime check configuration.",
	//   "flatPath": "v3/projects/{projectsId}/uptimeCheckConfigs/{uptimeCheckConfigsId}",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.uptimeCheckConfigs.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The uptime check configuration to retrieve. The format  is projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/uptimeCheckConfigs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "response": {
	//     "$ref": "UptimeCheckConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// method id "monitoring.projects.uptimeCheckConfigs.list":

type ProjectsUptimeCheckConfigsListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the existing valid uptime check configurations for the
// project, leaving out any invalid configurations.
func (r *ProjectsUptimeCheckConfigsService) List(parent string) *ProjectsUptimeCheckConfigsListCall {
	c := &ProjectsUptimeCheckConfigsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return in a single response. The server may further
// constrain the maximum number of results returned in a single page. If
// the page_size is <=0, the server will decide the number of results to
// be returned.
func (c *ProjectsUptimeCheckConfigsListCall) PageSize(pageSize int64) *ProjectsUptimeCheckConfigsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If this field is
// not empty then it must contain the nextPageToken value returned by a
// previous call to this method. Using this field causes the method to
// return more results from the previous method call.
func (c *ProjectsUptimeCheckConfigsListCall) PageToken(pageToken string) *ProjectsUptimeCheckConfigsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUptimeCheckConfigsListCall) Fields(s ...googleapi.Field) *ProjectsUptimeCheckConfigsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsUptimeCheckConfigsListCall) IfNoneMatch(entityTag string) *ProjectsUptimeCheckConfigsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsUptimeCheckConfigsListCall) Context(ctx context.Context) *ProjectsUptimeCheckConfigsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsUptimeCheckConfigsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsUptimeCheckConfigsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+parent}/uptimeCheckConfigs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.uptimeCheckConfigs.list" call.
// Exactly one of *ListUptimeCheckConfigsResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListUptimeCheckConfigsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsUptimeCheckConfigsListCall) Do(opts ...googleapi.CallOption) (*ListUptimeCheckConfigsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListUptimeCheckConfigsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the existing valid uptime check configurations for the project, leaving out any invalid configurations.",
	//   "flatPath": "v3/projects/{projectsId}/uptimeCheckConfigs",
	//   "httpMethod": "GET",
	//   "id": "monitoring.projects.uptimeCheckConfigs.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results to return in a single response. The server may further constrain the maximum number of results returned in a single page. If the page_size is \u003c=0, the server will decide the number of results to be returned.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If this field is not empty then it must contain the nextPageToken value returned by a previous call to this method. Using this field causes the method to return more results from the previous method call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The project whose uptime check configurations are listed. The format  is projects/[PROJECT_ID].",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+parent}/uptimeCheckConfigs",
	//   "response": {
	//     "$ref": "ListUptimeCheckConfigsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsUptimeCheckConfigsListCall) Pages(ctx context.Context, f func(*ListUptimeCheckConfigsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "monitoring.projects.uptimeCheckConfigs.patch":

type ProjectsUptimeCheckConfigsPatchCall struct {
	s                 *Service
	name              string
	uptimecheckconfig *UptimeCheckConfig
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Patch: Updates an uptime check configuration. You can either replace
// the entire configuration with a new one or replace only certain
// fields in the current configuration by specifying the fields to be
// updated via "updateMask". Returns the updated configuration.
func (r *ProjectsUptimeCheckConfigsService) Patch(name string, uptimecheckconfig *UptimeCheckConfig) *ProjectsUptimeCheckConfigsPatchCall {
	c := &ProjectsUptimeCheckConfigsPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.uptimecheckconfig = uptimecheckconfig
	return c
}

// UpdateMask sets the optional parameter "updateMask": If present, only
// the listed fields in the current uptime check configuration are
// updated with values from the new configuration. If this field is
// empty, then the current configuration is completely replaced with the
// new configuration.
func (c *ProjectsUptimeCheckConfigsPatchCall) UpdateMask(updateMask string) *ProjectsUptimeCheckConfigsPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUptimeCheckConfigsPatchCall) Fields(s ...googleapi.Field) *ProjectsUptimeCheckConfigsPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsUptimeCheckConfigsPatchCall) Context(ctx context.Context) *ProjectsUptimeCheckConfigsPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ProjectsUptimeCheckConfigsPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ProjectsUptimeCheckConfigsPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.uptimecheckconfig)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.projects.uptimeCheckConfigs.patch" call.
// Exactly one of *UptimeCheckConfig or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *UptimeCheckConfig.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsUptimeCheckConfigsPatchCall) Do(opts ...googleapi.CallOption) (*UptimeCheckConfig, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &UptimeCheckConfig{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an uptime check configuration. You can either replace the entire configuration with a new one or replace only certain fields in the current configuration by specifying the fields to be updated via \"updateMask\". Returns the updated configuration.",
	//   "flatPath": "v3/projects/{projectsId}/uptimeCheckConfigs/{uptimeCheckConfigsId}",
	//   "httpMethod": "PATCH",
	//   "id": "monitoring.projects.uptimeCheckConfigs.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "A unique resource name for this UptimeCheckConfig. The format is:projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].This field should be omitted when creating the uptime check configuration; on create, the resource name is assigned by the server and included in the response.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/uptimeCheckConfigs/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Optional. If present, only the listed fields in the current uptime check configuration are updated with values from the new configuration. If this field is empty, then the current configuration is completely replaced with the new configuration.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/{+name}",
	//   "request": {
	//     "$ref": "UptimeCheckConfig"
	//   },
	//   "response": {
	//     "$ref": "UptimeCheckConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}

// method id "monitoring.uptimeCheckIps.list":

type UptimeCheckIpsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns the list of IPs that checkers run from
func (r *UptimeCheckIpsService) List() *UptimeCheckIpsListCall {
	c := &UptimeCheckIpsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return in a single response. The server may further
// constrain the maximum number of results returned in a single page. If
// the page_size is <=0, the server will decide the number of results to
// be returned. NOTE: this field is not yet implemented
func (c *UptimeCheckIpsListCall) PageSize(pageSize int64) *UptimeCheckIpsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If this field is
// not empty then it must contain the nextPageToken value returned by a
// previous call to this method. Using this field causes the method to
// return more results from the previous method call. NOTE: this field
// is not yet implemented
func (c *UptimeCheckIpsListCall) PageToken(pageToken string) *UptimeCheckIpsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UptimeCheckIpsListCall) Fields(s ...googleapi.Field) *UptimeCheckIpsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *UptimeCheckIpsListCall) IfNoneMatch(entityTag string) *UptimeCheckIpsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UptimeCheckIpsListCall) Context(ctx context.Context) *UptimeCheckIpsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UptimeCheckIpsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UptimeCheckIpsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v3/uptimeCheckIps")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "monitoring.uptimeCheckIps.list" call.
// Exactly one of *ListUptimeCheckIpsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListUptimeCheckIpsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *UptimeCheckIpsListCall) Do(opts ...googleapi.CallOption) (*ListUptimeCheckIpsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListUptimeCheckIpsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the list of IPs that checkers run from",
	//   "flatPath": "v3/uptimeCheckIps",
	//   "httpMethod": "GET",
	//   "id": "monitoring.uptimeCheckIps.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results to return in a single response. The server may further constrain the maximum number of results returned in a single page. If the page_size is \u003c=0, the server will decide the number of results to be returned. NOTE: this field is not yet implemented",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If this field is not empty then it must contain the nextPageToken value returned by a previous call to this method. Using this field causes the method to return more results from the previous method call. NOTE: this field is not yet implemented",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v3/uptimeCheckIps",
	//   "response": {
	//     "$ref": "ListUptimeCheckIpsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/monitoring",
	//     "https://www.googleapis.com/auth/monitoring.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *UptimeCheckIpsListCall) Pages(ctx context.Context, f func(*ListUptimeCheckIpsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
