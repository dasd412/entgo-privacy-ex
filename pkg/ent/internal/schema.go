// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"privacy-ex/pkg/ent/schema\",\"Package\":\"privacy-ex/pkg/ent\",\"Schemas\":[{\"name\":\"Post\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"author\",\"type\":\"User\",\"field\":\"author_id\",\"ref_name\":\"posts\",\"unique\":true,\"inverse\":true,\"required\":true,\"comment\":\"게시물 작성자와의 관계\"}],\"fields\":[{\"name\":\"title\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"게시물 제목\"},{\"name\":\"content\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":2147483647,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"게시물 내용\"},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"게시물 작성 시간\"},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"게시물 수정 시간\"},{\"name\":\"author_id\",\"type\":{\"Type\":12,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"게시물 작성자 ID\"}],\"annotations\":{\"EntGQL\":{\"MultiOrder\":true,\"MutationInputs\":[{\"IsCreate\":true},{}],\"QueryField\":{},\"RelayConnection\":true}}},{\"name\":\"User\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"posts\",\"type\":\"Post\",\"unique\":true,\"annotations\":{\"EntSQL\":{\"on_delete\":\"CASCADE\"}}}],\"fields\":[{\"name\":\"email\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"이메일\"},{\"name\":\"password\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"해시화된 비밀 번호\"},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"이름\"},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"생성 날짜\"},{\"name\":\"role\",\"type\":{\"Type\":6,\"Ident\":\"user.Role\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"enums\":[{\"N\":\"admin\",\"V\":\"admin\"},{\"N\":\"author\",\"V\":\"author\"},{\"N\":\"guest\",\"V\":\"guest\"}],\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0},\"comment\":\"인가 권한\"}],\"annotations\":{\"EntGQL\":{\"MultiOrder\":true,\"MutationInputs\":[{\"IsCreate\":true},{}],\"QueryField\":{},\"RelayConnection\":true}}}],\"Features\":[\"namedges\",\"privacy\",\"schema/snapshot\"]}"
