package controllers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	bsemver "github.com/blang/semver/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/rand"
	ctrl "sigs.k8s.io/controller-runtime"
	crfinalizer "sigs.k8s.io/controller-runtime/pkg/finalizer"

	"github.com/operator-framework/operator-registry/alpha/declcfg"
	"github.com/operator-framework/operator-registry/alpha/property"

	ocv1alpha1 "github.com/operator-framework/operator-controller/api/v1alpha1"
	"github.com/operator-framework/operator-controller/internal/bundleutil"
	"github.com/operator-framework/operator-controller/internal/controllers"
	"github.com/operator-framework/operator-controller/internal/resolve"
	"github.com/operator-framework/operator-controller/internal/rukpak/source"
)

func TestClusterExtensionRegistryV1DisallowDependencies(t *testing.T) {
	ctx := context.Background()
	cl := newClient(t)

	for _, tt := range []struct {
		name    string
		bundle  declcfg.Bundle
		wantErr string
	}{
		{
			name: "package with no dependencies",
			bundle: declcfg.Bundle{
				Name:    "fake-catalog/no-dependencies-package/alpha/1.0.0",
				Package: "no-dependencies-package",
				Image:   "quay.io/fake-catalog/no-dependencies-package@sha256:3e281e587de3d03011440685fc4fb782672beab044c1ebadc42788ce05a21c35",
				Properties: []property.Property{
					{Type: property.TypePackage, Value: json.RawMessage(`{"packageName":"no-dependencies-package","version":"1.0.0"}`)},
				},
			},
		},
		{
			name: "package with olm.package.required property",
			bundle: declcfg.Bundle{
				Name:    "fake-catalog/package-required-test/alpha/1.0.0",
				Package: "package-required-test",
				Image:   "quay.io/fake-catalog/package-required-test@sha256:3e281e587de3d03011440685fc4fb782672beab044c1ebadc42788ce05a21c35",
				Properties: []property.Property{
					{Type: property.TypePackage, Value: json.RawMessage(`{"packageName":"package-required-test","version":"1.0.0"}`)},
					{Type: property.TypePackageRequired, Value: json.RawMessage("content-is-not-relevant")},
				},
			},
			wantErr: `bundle "fake-catalog/package-required-test/alpha/1.0.0" has a dependency declared via property "olm.package.required" which is currently not supported`,
		},
		{
			name: "package with olm.gvk.required property",
			bundle: declcfg.Bundle{
				Name:    "fake-catalog/gvk-required-test/alpha/1.0.0",
				Package: "gvk-required-test",
				Image:   "quay.io/fake-catalog/gvk-required-test@sha256:3e281e587de3d03011440685fc4fb782672beab044c1ebadc42788ce05a21c35",
				Properties: []property.Property{
					{Type: property.TypePackage, Value: json.RawMessage(`{"packageName":"gvk-required-test","version":"1.0.0"}`)},
					{Type: property.TypeGVKRequired, Value: json.RawMessage(`content-is-not-relevant`)},
				},
			},
			wantErr: `bundle "fake-catalog/gvk-required-test/alpha/1.0.0" has a dependency declared via property "olm.gvk.required" which is currently not supported`,
		},
		{
			name: "package with olm.constraint property",
			bundle: declcfg.Bundle{
				Name:    "fake-catalog/constraint-test/alpha/1.0.0",
				Package: "constraint-test",
				Image:   "quay.io/fake-catalog/constraint-test@sha256:3e281e587de3d03011440685fc4fb782672beab044c1ebadc42788ce05a21c35",
				Properties: []property.Property{
					{Type: property.TypePackage, Value: json.RawMessage(`{"packageName":"constraint-test","version":"1.0.0"}`)},
					{Type: property.TypeConstraint, Value: json.RawMessage(`content-is-not-relevant`)},
				},
			},
			wantErr: `bundle "fake-catalog/constraint-test/alpha/1.0.0" has a dependency declared via property "olm.constraint" which is currently not supported`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				require.NoError(t, cl.DeleteAllOf(ctx, &ocv1alpha1.ClusterExtension{}))
			}()
			resolver := resolve.Func(func(_ context.Context, _ *ocv1alpha1.ClusterExtension, _ *ocv1alpha1.BundleMetadata) (*declcfg.Bundle, *bsemver.Version, *declcfg.Deprecation, error) {
				v, err := bundleutil.GetVersion(tt.bundle)
				if err != nil {
					return nil, nil, nil, err
				}
				return &tt.bundle, v, nil, nil
			})
			mockUnpacker := unpacker.(*MockUnpacker)
			// Set up the Unpack method to return a result with StatePending
			mockUnpacker.On("Unpack", mock.Anything, mock.AnythingOfType("*v1alpha2.BundleDeployment")).Return(&source.Result{
				State: source.StatePending,
			}, nil)

			reconciler := &controllers.ClusterExtensionReconciler{
				Client:                cl,
				Resolver:              resolver,
				ActionClientGetter:    helmClientGetter,
				Unpacker:              unpacker,
				InstalledBundleGetter: &MockInstalledBundleGetter{},
				Finalizers:            crfinalizer.NewFinalizers(),
			}

			installNamespace := fmt.Sprintf("test-ns-%s", rand.String(8))
			serviceAccount := fmt.Sprintf("test-sa-%s", rand.String(8))
			extKey := types.NamespacedName{Name: fmt.Sprintf("cluster-extension-test-%s", rand.String(8))}
			clusterExtension := &ocv1alpha1.ClusterExtension{
				ObjectMeta: metav1.ObjectMeta{Name: extKey.Name},
				Spec: ocv1alpha1.ClusterExtensionSpec{
					PackageName:      tt.bundle.Package,
					InstallNamespace: installNamespace,
					ServiceAccount: ocv1alpha1.ServiceAccountReference{
						Name: serviceAccount,
					},
				},
			}
			require.NoError(t, cl.Create(ctx, clusterExtension))

			res, err := reconciler.Reconcile(ctx, ctrl.Request{NamespacedName: extKey})
			require.Equal(t, ctrl.Result{}, res)
			if tt.wantErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr)

				// In case of an error we want it to be included in the installed condition
				require.NoError(t, cl.Get(ctx, extKey, clusterExtension))
				cond := apimeta.FindStatusCondition(clusterExtension.Status.Conditions, ocv1alpha1.TypeInstalled)
				require.NotNil(t, cond)
				require.Equal(t, metav1.ConditionFalse, cond.Status)
				require.Equal(t, ocv1alpha1.ReasonInstallationFailed, cond.Reason)
				require.Equal(t, tt.wantErr, cond.Message)
			}
		})
	}
}
