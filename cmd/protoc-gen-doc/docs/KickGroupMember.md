# KickGroupMember

### Feature Introduction

* kick a group member


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------- |
| KickGroupMember | [KickGroupMemberReq](#openim.sdk.group.KickGroupMemberReq) | [KickGroupMemberResp](#openim.sdk.group.KickGroupMemberResp) | kick a group member |
 

#### KickGroupMemberReq
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to kick |
| kickedUserIDs | [string](#string) | repeated | user id you want to kick |
| reason | [string](#string) |  | kick reason |
 

#### KickGroupMemberResp
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


