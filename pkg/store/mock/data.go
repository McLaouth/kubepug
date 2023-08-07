package mock

const (
	MockValidData = `
		[
		{
			"group": "extensions",
			"version": "v1beta1",
			"kind": "DaemonSet",
			"description": "DEPRECATED - This group version of DaemonSet is deprecated by apps/v1beta2/DaemonSet. See the release notes for\nmore information.\nDaemonSet represents the configuration of a daemon set.",
			"introduced_version": {
				"version_major": 1,
				"version_minor": 1
			},
			"deprecated_version": {
				"version_major": 1,
				"version_minor": 8
			},
			"removed_version": {
				"version_major": 1,
				"version_minor": 16
			},
			"replacement": {
				"group": "apps",
				"version": "v1",
				"kind": "DaemonSet"
			}
		},
		{
			"group": "",
			"version": "v1",
			"kind": "BlahPod",
			"description": "BlahPod is a deprecated Pod.",
			"introduced_version": {
				"version_major": 1,
				"version_minor": 1
			},
			"deprecated_version": {
				"version_major": 1,
				"version_minor": 8
			},
			"removed_version": {
				"version_major": 1,
				"version_minor": 16
			},
			"replacement": {
				"group": "",
				"version": "v1",
				"kind": "Pod"
			}
		},
		{
			"group": "discovery.k8s.io",
			"version": "v1beta1",
			"kind": "EndpointSliceList",
			"description": "EndpointSliceList represents a list of endpoint slices",
			"introduced_version": {
				"version_major": 1,
				"version_minor": 16
			},
			"deprecated_version": {
				"version_major": 1,
				"version_minor": 21
			},
			"removed_version": {
				"version_major": 1,
				"version_minor": 25
			},
			"replacement": {
				"group": "discovery.k8s.io",
				"version": "v1",
				"kind": "EndpointSlice"
			}
		},
		{
			"group": "admission.k8s.io",
			"version": "v1beta1",
			"kind": "AdmissionReview",
			"description": "AdmissionReview describes an admission review request/response.",
			"introduced_version": {
				"version_major": 1,
				"version_minor": 9
			},
			"deprecated_version": {
				"version_major": 1,
				"version_minor": 19
			},
			"removed_version": {
				"version_major": 1,
				"version_minor": 22
			},
			"replacement": {
				"group": "admission.k8s.io",
				"version": "v1",
				"kind": "AdmissionReview"
			}
		}
	]`

	MockInvalidData = `
	[
		{
			"group": "extensions",
			"version": "v1beta1",
			"kind": "DaemonSet",
			"description": "DEPRECATED - This group version of DaemonSet is deprecated by apps/v1beta2/DaemonSet. See the release notes for\nmore information.\nDaemonSet represents the configuration of a daemon set.",
			"introduced_version": {
				"version_major": 1,
				"version_minor": 1
			},
			"replacement": {
				"group": "apps",
				"version": "v1",
				"kind": "DaemonSet"
			`
	MockNoVersions = `
	[
		{
			"group": "extensions",
			"version": "v1beta1",
			"kind": "DaemonSet",
			"description": "DEPRECATED - This group version of DaemonSet is deprecated by apps/v1beta2/DaemonSet. See the release notes for\nmore information.\nDaemonSet represents the configuration of a daemon set.",
			"deprecated_version": {
				"version_major": 0,
				"version_minor": 5
			},
			"removed_version": {
				"version_major": 7,
				"version_minor": 0
			}
		}
	]`
)
