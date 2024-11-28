# IsJoinGroup

### 简要描述

- is join group

| Method Name | Request Type | Response Type |
| ----------- | ------------ | ------------- |
| IsJoinGroup | [IsJoinGroupReq](#openim.sdk.group.IsJoinGroupReq) | [IsJoinGroupResp](#openim.sdk.group.IsJoinGroupResp) |

### 请求参数
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to check |
| userID | [string](#string) |  | user id you want to check |


### 响应参数
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| joined | [bool](#bool) |  | whether to join the group |


