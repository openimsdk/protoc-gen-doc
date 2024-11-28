# KickGroupMember

### 简要描述

- kick a group member

| Method Name | Request Type | Response Type |
| ----------- | ------------ | ------------- |
| KickGroupMember | [KickGroupMemberReq](#openim.sdk.group.KickGroupMemberReq) | [KickGroupMemberResp](#openim.sdk.group.KickGroupMemberResp) |

### KickGroupMemberReq
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to kick |
| kickedUserIDs | [string](#string) | repeated | user id you want to kick |
| reason | [string](#string) |  | kick reason |


### KickGroupMemberResp
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |


