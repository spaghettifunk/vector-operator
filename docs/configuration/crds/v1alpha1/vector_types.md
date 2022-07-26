## VectorSpec

VectorSpec defines the desired state of Vector

### foo (string, optional) {#vectorspec-foo}

Default: -


## VectorStatus

VectorStatus defines the observed state of Vector


## Vector

Vector is the Schema for the vectors API

###  (metav1.TypeMeta, required) {#vector-}

Default: -

### metadata (metav1.ObjectMeta, optional) {#vector-metadata}

Default: -

### spec (VectorSpec, optional) {#vector-spec}

Default: -

### status (VectorStatus, optional) {#vector-status}

Default: -


## VectorList

VectorList contains a list of Vector

###  (metav1.TypeMeta, required) {#vectorlist-}

Default: -

### metadata (metav1.ListMeta, optional) {#vectorlist-metadata}

Default: -

### items ([]Vector, required) {#vectorlist-items}

Default: -


