# IsJoinGroup

### Feature Introduction

- is join group

| Method Name | Request Type | Response Type |
| ----------- | ------------ | ------------- |
| IsJoinGroup | [IsJoinGroupReq](#openim.sdk.group.IsJoinGroupReq) | [IsJoinGroupResp](#openim.sdk.group.IsJoinGroupResp) |

### IsJoinGroupReq
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to check |
| userID | [string](#string) |  | user id you want to check |


### IsJoinGroupResp
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| joined | [bool](#bool) |  | whether to join the group |


