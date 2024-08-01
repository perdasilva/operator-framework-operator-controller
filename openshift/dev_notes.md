## Notes for developers

### Information on the downstreaming process can be found here:
[https://github.com/openshift/operator-framework-tooling/blob/main/README.md](https://github.com/openshift/operator-framework-tooling/blob/main/README.md)

### Running on OpenShift Local

#### Prerequisites
 * [oc](https://docs.openshift.com/container-platform/4.16/cli_reference/openshift_cli/getting-started-cli.html)
 * [OpenShift Local/crc](https://developers.redhat.com/products/openshift-local/overview?source=sso)
 * [CI registry access](https://docs.ci.openshift.org/docs/how-tos/use-registries-in-build-farm/#how-do-i-log-in-to-pull-images-that-require-authentication)

#### Considerations

To run your local downstream operator-controller on OpenShift local, we'll need to:
 1. build the downstream o-c image
 2. upload the image to the OpenShift image repo
 3. update the manifests to point to the internally hosted image
 4. scale down CVO to avoid our manifest changes being overridden (only important after GA)
 5. apply the manifests

