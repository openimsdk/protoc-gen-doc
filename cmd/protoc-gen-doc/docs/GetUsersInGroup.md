# GetUsersInGroup

### 简要描述

- get users in group

| Method Name | Request Type | Response Type |
| ----------- | ------------ | ------------- |
| GetUsersInGroup | [GetUsersInGroupReq](#openim.sdk.group.GetUsersInGroupReq) | [GetUsersInGroupResp](#openim.sdk.group.GetUsersInGroupResp) |

### GetUsersInGroupReq
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to get |
| userIDs | [string](#string) | repeated | user id you want to get |


### GetUsersInGroupResp
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userIDs | [string](#string) | repeated | user id in the group |


