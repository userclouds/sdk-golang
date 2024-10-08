# Changelog

## 1.6.1 - 26-09-2024

- add policy secrets SDK methods
- minor changes to prepare for upgrade to go 1.23.1

## 1.6.0 - 12-09-2024

- add ObjectStore CRUD methods
- add Truncated flag to ExecuteAccessorResponse
- add Version to Transformer
- add UpdateTransformerRequest
- add GetTransformerByVersion method
- add UpdateTransformer method
- add SearchIndexed and AccessPolicy to Column
- add TokenAccessPolicy to ColumnOutputConfig
- add AreColumnAccessPoliciesOverridden, IsAutoGenerated, UseSearchIndex to Accessor
- add Schemas to ColumnInputConfig
- add ShimObjectStore
- client cache improvements
- logging improvements
- add onprem region and universe

## 1.5.0 - 30-07-2024

- Update userstore sample to exercise multi-key execute accessor pagination

## 1.4.0 - 18-07-2024

- Deprecate v1.0.0 and v1.1.0
- require DataType and remove requirement for Type for Column
- require InputDataType and OutputDataType, and remove requirement for InputType, InputConstraints, OutputType, and OutputConstraints for Transformer
- enforce length limit of 128 characters for ColumnDataType Description
- expand allowable length to 256 characters for ResourceID Name

## 1.3.0 - 28-06-2024

- Retry getting access token on EOF type network errors which occur when connection is lost.
- Breaking change: RetryNetworkErrors option for jsonclient now takes a boolean, and the option is on by default

## 1.2.0 - 09-04-2024

- Update userstore sample to exercise partial update columns
- Add methods for creating, retrieving, updating, and deleting ColumnDataTypes
- Add DataType field to Column that refers to a ColumnDataType
- Add InputDataType and OutputDataType fields to Transformer that refer to ColumnDataTypes
- Update userstore sample to interact with ColumnDataTypes
- Breaking change: Add additional boolean parameter to ListAccessors an ListMutators for requesting all versions

## 1.1.0 - 20-03-2024

- Breaking change: idp/userstore/ColumnField parameter "Optional" has been changed to "Required", with fields not required by default
- Add InputConstraints and OutputConstraints parameters of type idp/userstore/ColumnConstraints to idp/policy/Transformer
- Add pagination support for chained logical filter queries (query,logical_operator,query,logical_operator,query...)

## 1.0.0 - 31-01-2024

- Breaking change: Return ErrXXXNotFound error when getting a HTTP 404 from authz endpoints
- Breaking change: Move prefix argument for NewRedisClientCacheProvider to be optional KeyPrefixRedis(prefix) instead of required
- Add validation of non empty pointer strings
- Add column constraints implementation
- Adding "ID" as an optional field for user creation
