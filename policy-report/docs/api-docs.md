# API Reference

## Packages
- [reports.x-k8s.io/v1beta2](#reportsx-k8siov1beta2)


## reports.x-k8s.io/v1beta2

Package v1beta2 contains API Schema definitions for the policy v1beta2 API group

Package v1beta2 contains API Schema definitions for the policy v1beta2 API group

### Resource Types
- [ClusterPolicyReport](#clusterpolicyreport)
- [ClusterPolicyReportList](#clusterpolicyreportlist)
- [PolicyReport](#policyreport)
- [PolicyReportList](#policyreportlist)



#### ClusterPolicyReport



ClusterPolicyReport is the Schema for the clusterpolicyreports API



_Appears in:_
- [ClusterPolicyReportList](#clusterpolicyreportlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `reports.x-k8s.io/v1beta2` | | |
| `kind` _string_ | `ClusterPolicyReport` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `source` _string_ | Source is an identifier for the source e.g. a policy engine that manages this report.<br />Use this field if all the results are produced by a single policy engine.<br />If the results are produced by multiple sources e.g. different engines or scanners,<br />then use the Source field at the PolicyReportResult level. |  |  |
| `scope` _[ObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#objectreference-v1-core)_ | Scope is an optional reference to the report scope (e.g. a Deployment, Namespace, or Node) |  |  |
| `scopeSelector` _[LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#labelselector-v1-meta)_ | ScopeSelector is an optional selector for multiple scopes (e.g. Pods).<br />Either one of, or none of, but not both of, Scope or ScopeSelector should be specified. |  |  |
| `configuration` _[PolicyReportConfiguration](#policyreportconfiguration)_ | Configuration is an optional field which can be used to specify<br />a contract between PolicyReport generators and consumers |  |  |
| `summary` _[PolicyReportSummary](#policyreportsummary)_ | PolicyReportSummary provides a summary of results |  |  |
| `results` _[PolicyReportResult](#policyreportresult) array_ | PolicyReportResult provides result details |  |  |


#### ClusterPolicyReportList



ClusterPolicyReportList contains a list of ClusterPolicyReport





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `reports.x-k8s.io/v1beta2` | | |
| `kind` _string_ | `ClusterPolicyReportList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[ClusterPolicyReport](#clusterpolicyreport) array_ |  |  |  |


#### Limits







_Appears in:_
- [PolicyReportConfiguration](#policyreportconfiguration)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `maxResults` _integer_ | MaxResults is the maximum number of results contained in the report |  |  |
| `statusFilter` _[StatusFilter](#statusfilter) array_ | StatusFilter indicates that the PolicyReport contains only those reports with statuses specified in this list |  | Enum: [pass fail warn error skip] <br /> |


#### PolicyReport



PolicyReport is the Schema for the policyreports API



_Appears in:_
- [PolicyReportList](#policyreportlist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `reports.x-k8s.io/v1beta2` | | |
| `kind` _string_ | `PolicyReport` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `source` _string_ | Source is an identifier for the source e.g. a policy engine that manages this report.<br />Use this field if all the results are produced by a single policy engine.<br />If the results are produced by multiple sources e.g. different engines or scanners,<br />then use the Source field at the PolicyReportResult level. |  |  |
| `scope` _[ObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#objectreference-v1-core)_ | Scope is an optional reference to the report scope (e.g. a Deployment, Namespace, or Node) |  |  |
| `scopeSelector` _[LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#labelselector-v1-meta)_ | ScopeSelector is an optional selector for multiple scopes (e.g. Pods).<br />Either one of, or none of, but not both of, Scope or ScopeSelector should be specified. |  |  |
| `configuration` _[PolicyReportConfiguration](#policyreportconfiguration)_ | Configuration is an optional field which can be used to specify<br />a contract between PolicyReport generators and consumers |  |  |
| `summary` _[PolicyReportSummary](#policyreportsummary)_ | PolicyReportSummary provides a summary of results |  |  |
| `results` _[PolicyReportResult](#policyreportresult) array_ | PolicyReportResult provides result details |  |  |


#### PolicyReportConfiguration







_Appears in:_
- [ClusterPolicyReport](#clusterpolicyreport)
- [PolicyReport](#policyreport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `limits` _[Limits](#limits)_ |  |  |  |


#### PolicyReportList



PolicyReportList contains a list of PolicyReport





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `reports.x-k8s.io/v1beta2` | | |
| `kind` _string_ | `PolicyReportList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[PolicyReport](#policyreport) array_ |  |  |  |


#### PolicyReportResult



PolicyReportResult provides the result for an individual policy



_Appears in:_
- [ClusterPolicyReport](#clusterpolicyreport)
- [PolicyReport](#policyreport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `source` _string_ | Source is an identifier for the policy engine that manages this report<br />If the Source is specified at this level, it will override the Source<br />field set at the PolicyReport level |  |  |
| `policy` _string_ | Policy is the name or identifier of the policy |  |  |
| `rule` _string_ | Rule is the name or identifier of the rule within the policy |  |  |
| `category` _string_ | Category indicates policy category |  |  |
| `severity` _[PolicyResultSeverity](#policyresultseverity)_ | Severity indicates policy check result criticality |  | Enum: [critical high low medium info] <br /> |
| `timestamp` _[Timestamp](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#timestamp-v1-meta)_ | Timestamp indicates the time the result was found |  |  |
| `result` _[PolicyResult](#policyresult)_ | Result indicates the outcome of the policy rule execution |  | Enum: [pass fail warn error skip] <br /> |
| `scored` _boolean_ | Scored indicates if this result is scored |  |  |
| `resources` _[ObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#objectreference-v1-core) array_ | Subjects is an optional reference to the checked Kubernetes resources |  |  |
| `resourceSelector` _[LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.29/#labelselector-v1-meta)_ | ResourceSelector is an optional label selector for checked Kubernetes resources.<br />For example, a policy result may apply to all pods that match a label.<br />Either a Subject or a ResourceSelector can be specified. If neither are provided, the<br />result is assumed to be for the policy report scope. |  |  |
| `message` _string_ | Description is a short user friendly message for the policy rule |  |  |
| `properties` _object (keys:string, values:string)_ | Properties provides additional information for the policy rule |  |  |


#### PolicyReportSummary



PolicyReportSummary provides a status count summary



_Appears in:_
- [ClusterPolicyReport](#clusterpolicyreport)
- [PolicyReport](#policyreport)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `pass` _integer_ | Pass provides the count of policies whose requirements were met |  |  |
| `fail` _integer_ | Fail provides the count of policies whose requirements were not met |  |  |
| `warn` _integer_ | Warn provides the count of non-scored policies whose requirements were not met |  |  |
| `error` _integer_ | Error provides the count of policies that could not be evaluated |  |  |
| `skip` _integer_ | Skip indicates the count of policies that were not selected for evaluation |  |  |


#### PolicyResult

_Underlying type:_ _string_

PolicyResult has one of the following values:
  - pass: the policy requirements are met
  - fail: the policy requirements are not met
  - warn: the policy requirements are not met and the policy is not scored
  - error: the policy could not be evaluated
  - skip: the policy was not selected based on user inputs or applicability

_Validation:_
- Enum: [pass fail warn error skip]

_Appears in:_
- [PolicyReportResult](#policyreportresult)



#### PolicyResultSeverity

_Underlying type:_ _string_

PolicyResultSeverity has one of the following values:
  - critical
  - high
  - low
  - medium
  - info

_Validation:_
- Enum: [critical high low medium info]

_Appears in:_
- [PolicyReportResult](#policyreportresult)



#### StatusFilter

_Underlying type:_ _string_

StatusFilter is used by PolicyReport generators to write only those reports whose status is specified by the filters

_Validation:_
- Enum: [pass fail warn error skip]

_Appears in:_
- [Limits](#limits)



