# GetUsersInGroup

### 简要描述

- get users in group

| Method Name | Request Type | Response Type |
| ----------- | ------------ | ------------- |
| GetUsersInGroup | [GetUsersInGroupReq](#openim.sdk.group.GetUsersInGroupReq) | [GetUsersInGroupResp](#openim.sdk.group.GetUsersInGroupResp) |

### 请求参数
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to get |
| userIDs | [string](#string) | repeated | user id you want to get |


### 响应参数
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userIDs | [string](#string) | repeated | user id in the group |


