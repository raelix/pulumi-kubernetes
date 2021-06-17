// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRole":
		r = &ClusterRole{}
	case "kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBinding":
		r = &ClusterRoleBinding{}
	case "kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleBindingList":
		r = &ClusterRoleBindingList{}
	case "kubernetes:rbac.authorization.k8s.io/v1beta1:ClusterRoleList":
		r = &ClusterRoleList{}
	case "kubernetes:rbac.authorization.k8s.io/v1beta1:Role":
		r = &Role{}
	case "kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBinding":
		r = &RoleBinding{}
	case "kubernetes:rbac.authorization.k8s.io/v1beta1:RoleBindingList":
		r = &RoleBindingList{}
	case "kubernetes:rbac.authorization.k8s.io/v1beta1:RoleList":
		r = &RoleList{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := kubernetes.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"kubernetes",
		"rbac.authorization.k8s.io/v1beta1",
		&module{version},
	)
}
