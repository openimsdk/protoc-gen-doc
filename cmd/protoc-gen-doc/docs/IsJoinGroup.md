# IsJoinGroup

### Feature Introduction



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------- |
| IsJoinGroup | [IsJoinGroupReq](#openim.sdk.group.IsJoinGroupReq) | [IsJoinGroupResp](#openim.sdk.group.IsJoinGroupResp) | is join group |
 

#### IsJoinGroupReq
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to check |
| userID | [string](#string) |  | user id you want to check |
 

#### IsJoinGroupResp
| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| joined | [bool](#bool) |  | whether to join the group |


