# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [group.proto](#group-proto)
    - [ChangeGroupMemberMuteReq](#openim-sdk-group-ChangeGroupMemberMuteReq)
    - [ChangeGroupMemberMuteResp](#openim-sdk-group-ChangeGroupMemberMuteResp)
    - [ChangeGroupMuteReq](#openim-sdk-group-ChangeGroupMuteReq)
    - [ChangeGroupMuteResp](#openim-sdk-group-ChangeGroupMuteResp)
    - [DismissGroupReq](#openim-sdk-group-DismissGroupReq)
    - [DismissGroupResp](#openim-sdk-group-DismissGroupResp)
    - [GetUsersInGroupReq](#openim-sdk-group-GetUsersInGroupReq)
    - [GetUsersInGroupResp](#openim-sdk-group-GetUsersInGroupResp)
    - [InviteUserToGroupReq](#openim-sdk-group-InviteUserToGroupReq)
    - [InviteUserToGroupResp](#openim-sdk-group-InviteUserToGroupResp)
    - [IsJoinGroupReq](#openim-sdk-group-IsJoinGroupReq)
    - [IsJoinGroupResp](#openim-sdk-group-IsJoinGroupResp)
    - [KickGroupMemberReq](#openim-sdk-group-KickGroupMemberReq)
    - [KickGroupMemberResp](#openim-sdk-group-KickGroupMemberResp)
    - [QuitGroupReq](#openim-sdk-group-QuitGroupReq)
    - [QuitGroupResp](#openim-sdk-group-QuitGroupResp)
    - [SetGroupInfoReq](#openim-sdk-group-SetGroupInfoReq)
    - [SetGroupInfoResp](#openim-sdk-group-SetGroupInfoResp)
    - [SetGroupMemberInfoReq](#openim-sdk-group-SetGroupMemberInfoReq)
    - [SetGroupMemberInfoResp](#openim-sdk-group-SetGroupMemberInfoResp)
    - [TransferGroupOwnerReq](#openim-sdk-group-TransferGroupOwnerReq)
    - [TransferGroupOwnerResp](#openim-sdk-group-TransferGroupOwnerResp)
  
    - [Group](#openim-sdk-group-Group)
  
- [Scalar Value Types](#scalar-value-types)



<a name="group-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## group.proto
<h1> Group1 </h1>


<a name="openim-sdk-group-ChangeGroupMemberMuteReq"></a>

### ChangeGroupMemberMuteReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to mute |
| userID | [string](#string) |  | user id you want to mute |
| mutedSeconds | [uint32](#uint32) |  | mute time (unit: seconds), 0 cancels the mute |






<a name="openim-sdk-group-ChangeGroupMemberMuteResp"></a>

### ChangeGroupMemberMuteResp







<a name="openim-sdk-group-ChangeGroupMuteReq"></a>

### ChangeGroupMuteReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to mute |
| mute | [bool](#bool) |  | mute or cancel mute |






<a name="openim-sdk-group-ChangeGroupMuteResp"></a>

### ChangeGroupMuteResp







<a name="openim-sdk-group-DismissGroupReq"></a>

### DismissGroupReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you dismissed |






<a name="openim-sdk-group-DismissGroupResp"></a>

### DismissGroupResp







<a name="openim-sdk-group-GetUsersInGroupReq"></a>

### GetUsersInGroupReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to get |
| userIDs | [string](#string) | repeated | user id you want to get |






<a name="openim-sdk-group-GetUsersInGroupResp"></a>

### GetUsersInGroupResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userIDs | [string](#string) | repeated | user id in the group |






<a name="openim-sdk-group-InviteUserToGroupReq"></a>

### InviteUserToGroupReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to invite |
| userIDs | [string](#string) | repeated | user id you want to invite |
| reason | [string](#string) |  | invite reason |






<a name="openim-sdk-group-InviteUserToGroupResp"></a>

### InviteUserToGroupResp







<a name="openim-sdk-group-IsJoinGroupReq"></a>

### IsJoinGroupReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to check |
| userID | [string](#string) |  | user id you want to check |






<a name="openim-sdk-group-IsJoinGroupResp"></a>

### IsJoinGroupResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| joined | [bool](#bool) |  | whether to join the group |






<a name="openim-sdk-group-KickGroupMemberReq"></a>

### KickGroupMemberReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to kick |
| kickedUserIDs | [string](#string) | repeated | user id you want to kick |
| reason | [string](#string) |  | kick reason |






<a name="openim-sdk-group-KickGroupMemberResp"></a>

### KickGroupMemberResp







<a name="openim-sdk-group-QuitGroupReq"></a>

### QuitGroupReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you quit |






<a name="openim-sdk-group-QuitGroupResp"></a>

### QuitGroupResp







<a name="openim-sdk-group-SetGroupInfoReq"></a>

### SetGroupInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to set |
| groupName | [string](#string) | optional | group name |
| notification | [string](#string) | optional | group notification |
| introduction | [string](#string) | optional | group introduction |
| faceURL | [string](#string) | optional | group face url |
| ex | [string](#string) | optional | extension field |
| needVerification | [int32](#int32) | optional | whether to verify the group |
| lookMemberInfo | [int32](#int32) | optional | whether to view the group member information |
| applyMemberFriend | [int32](#int32) | optional | whether to allow group members to add friends |






<a name="openim-sdk-group-SetGroupInfoResp"></a>

### SetGroupInfoResp







<a name="openim-sdk-group-SetGroupMemberInfoReq"></a>

### SetGroupMemberInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to set |
| userID | [string](#string) |  | user id you want to set |
| nickname | [string](#string) | optional | user nickname |
| faceURL | [string](#string) | optional | user face url |
| roleLevel | [int32](#int32) | optional | user role level |
| ex | [string](#string) | optional | extension field |






<a name="openim-sdk-group-SetGroupMemberInfoResp"></a>

### SetGroupMemberInfoResp







<a name="openim-sdk-group-TransferGroupOwnerReq"></a>

### TransferGroupOwnerReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupID | [string](#string) |  | group id you want to transfer |
| ownerUserID | [string](#string) |  | new owner user id |






<a name="openim-sdk-group-TransferGroupOwnerResp"></a>

### TransferGroupOwnerResp






 

 

 


<a name="openim-sdk-group-Group"></a>

### Group


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| QuitGroup | [QuitGroupReq](#openim-sdk-group-QuitGroupReq) | [QuitGroupResp](#openim-sdk-group-QuitGroupResp) | This is a leading comment for a message<br>This is a leading comment for a message<br>This is a leading comment for a message<br>This is a leading comment for a message |
| DismissGroup | [DismissGroupReq](#openim-sdk-group-DismissGroupReq) | [DismissGroupResp](#openim-sdk-group-DismissGroupResp) | dismiss a group |
| ChangeGroupMute | [ChangeGroupMuteReq](#openim-sdk-group-ChangeGroupMuteReq) | [ChangeGroupMuteResp](#openim-sdk-group-ChangeGroupMuteResp) | mute or cancel mute a group |
| ChangeGroupMemberMute | [ChangeGroupMemberMuteReq](#openim-sdk-group-ChangeGroupMemberMuteReq) | [ChangeGroupMemberMuteResp](#openim-sdk-group-ChangeGroupMemberMuteResp) | mute or cancel mute a group member |
| TransferGroupOwner | [TransferGroupOwnerReq](#openim-sdk-group-TransferGroupOwnerReq) | [TransferGroupOwnerResp](#openim-sdk-group-TransferGroupOwnerResp) | transfer group owner |
| KickGroupMember | [KickGroupMemberReq](#openim-sdk-group-KickGroupMemberReq) | [KickGroupMemberResp](#openim-sdk-group-KickGroupMemberResp) | kick a group member |
| SetGroupInfo | [SetGroupInfoReq](#openim-sdk-group-SetGroupInfoReq) | [SetGroupInfoResp](#openim-sdk-group-SetGroupInfoResp) | set group information |
| SetGroupMemberInfo | [SetGroupMemberInfoReq](#openim-sdk-group-SetGroupMemberInfoReq) | [SetGroupMemberInfoResp](#openim-sdk-group-SetGroupMemberInfoResp) | set group member information |
| IsJoinGroup | [IsJoinGroupReq](#openim-sdk-group-IsJoinGroupReq) | [IsJoinGroupResp](#openim-sdk-group-IsJoinGroupResp) | is join group |
| GetUsersInGroup | [GetUsersInGroupReq](#openim-sdk-group-GetUsersInGroupReq) | [GetUsersInGroupResp](#openim-sdk-group-GetUsersInGroupResp) | get users in group |
| InviteUserToGroup | [InviteUserToGroupReq](#openim-sdk-group-InviteUserToGroupReq) | [InviteUserToGroupResp](#openim-sdk-group-InviteUserToGroupResp) | invite user to group |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

