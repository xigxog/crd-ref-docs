# API Reference

## Packages
- [webapp.test.k8s.elastic.co/v1](#webapptestk8selasticcov1)


## webapp.test.k8s.elastic.co/v1

Package v1 contains API Schema definitions for the webapp v1 API group

### Resource Types
- [Embedded](#embedded)
- [Guestbook](#guestbook)
- [GuestbookList](#guestbooklist)
- [Underlying](#underlying)



#### Embedded







| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1`
| `kind` _string_ | `Embedded`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `a` _string_ |  |
| `b` _string_ |  |
| `c` _string_ |  |
| `x` _string_ |  |
| `d` _string_ |  |
| `e` _string_ |  |


#### EmbeddedX





_Appears in:_
- [Embedded](#embedded)
- [Embedded1](#embedded1)
- [Embedded2](#embedded2)
- [Embedded3](#embedded3)
- [Embedded4](#embedded4)

| Field | Description |
| --- | --- |
| `x` _string_ |  |


#### Guestbook



Guestbook is the Schema for the guestbooks API.

_Appears in:_
- [GuestbookList](#guestbooklist)

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1`
| `kind` _string_ | `Guestbook`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[GuestbookSpec](#guestbookspec)_ |  |


#### GuestbookEntry



GuestbookEntry defines an entry in a guest book.

_Appears in:_
- [GuestbookSpec](#guestbookspec)

| Field | Description |
| --- | --- |
| `name` _string_ | Name of the guest (pipe \| should be escaped) |
| `time` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#time-v1-meta)_ | Time of entry |
| `comment` _string_ | Comment by guest. This can be a multi-line comment. <br /><br /> Just like this one. |
| `rating` _[Rating](#rating)_ | Rating provided by the guest |


#### GuestbookHeader

_Underlying type:_ _string_

GuestbookHeaders are strings to include at the top of a page.

_Appears in:_
- [GuestbookSpec](#guestbookspec)



#### GuestbookList



GuestbookList contains a list of Guestbook.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1`
| `kind` _string_ | `GuestbookList`
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `items` _[Guestbook](#guestbook) array_ |  |


#### GuestbookSpec



GuestbookSpec defines the desired state of Guestbook.

_Appears in:_
- [Guestbook](#guestbook)

| Field | Description |
| --- | --- |
| `page` _integer_ | Page indicates the page number |
| `entries` _[GuestbookEntry](#guestbookentry) array_ | Entries contain guest book entries for the page |
| `selector` _[LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#labelselector-v1-meta)_ | Selector selects something |
| `headers` _[GuestbookHeader](#guestbookheader) array_ | Headers contains a list of header items to include in the page |
| `certificateRef` _[SecretObjectReference](https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io/v1beta1.SecretObjectReference)_ | CertificateRef is a reference to a secret containing a certificate |




#### Rating

_Underlying type:_ _integer_

Rating is the rating provided by a guest.

_Appears in:_
- [GuestbookEntry](#guestbookentry)



#### Underlying



Underlying tests that Underlying1's underlying type is Underlying2 instead of string.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1`
| `kind` _string_ | `Underlying`
| `a` _[Underlying1](#underlying1)_ |  |


#### Underlying1

_Underlying type:_ _[Underlying2](#underlying2)_

Underlying1 has an underlying type with an underlying type

_Appears in:_
- [Underlying](#underlying)



#### Underlying2

_Underlying type:_ _string_

Underlying2 is a string alias

_Appears in:_
- [Underlying1](#underlying1)



