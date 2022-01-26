package resource

type ResourceList interface {
	Append(res Resource) ResourceList
	AppendList(list ResourceList) ResourceList
	Slice(first, second int) ResourceList
	At(idx int) Resource
	Set(idx int, res Resource)
	Length() int
}

type ResourceChangeList interface {
	Append(change ResourceChange) ResourceChangeList
	AppendList(list ResourceChangeList) ResourceChangeList
	Slice(first, second int) ResourceChangeList
	At(idx int) ResourceChange
	Set(idx int, change ResourceChange)
	Length() int
}

type NameList interface {
	Append(res Name) NameList
	AppendList(list NameList) NameList
	Slice(first, second int) NameList
	At(idx int) Name
	Set(idx int, name Name)
	Length() int
}

type ReferenceList interface {
	Append(ref Reference) ReferenceList
	AppendList(list ReferenceList) ReferenceList
	Slice(first, second int) ReferenceList
	At(idx int) Reference
	Set(idx int, ref Reference)
	Length() int
}

type ParentNameList interface {
	Append(res Name) ParentNameList
	AppendList(list ParentNameList) ParentNameList
	Slice(first, second int) ParentNameList
	At(idx int) Name
	Set(idx int, name Name)
	Length() int
}

type ParentReferenceList interface {
	Append(ref Reference) ParentReferenceList
	AppendList(list ParentReferenceList) ParentReferenceList
	Slice(first, second int) ParentReferenceList
	At(idx int) Reference
	Set(idx int, ref Reference)
	Length() int
}

type ResourceMap interface {
	Get(name Name) Resource
	Set(change Resource)
	Delete(name Name)
	Length() int
	ForEach(cb func(Name, Resource) bool)
}

type ResourceChangeMap interface {
	Get(name Name) ResourceChange
	Set(change ResourceChange)
	Delete(name Name)
	Length() int
	ForEach(cb func(Name, ResourceChange) bool)
}
