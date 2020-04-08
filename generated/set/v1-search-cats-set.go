// Code generated by genny. DO NOT EDIT.
// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package set

import (
	"sort"

	v1 "github.com/stackrox/rox/generated/api/v1"
)

// If you want to add a set for your custom type, simply add another go generate line along with the
// existing ones. If you're creating a set for a primitive type, you can follow the example of "string"
// and create the generated file in this package.
// For non-primitive sets, please make the generated code files go outside this package.
// Sometimes, you might need to create it in the same package where it is defined to avoid import cycles.
// The permission set is an example of how to do that.
// You can also specify the -imp command to specify additional imports in your generated file, if required.

// v1.SearchCategory represents a generic type that we want to have a set of.

// V1SearchCategorySet will get translated to generic sets.
type V1SearchCategorySet map[v1.SearchCategory]struct{}

// Add adds an element of type v1.SearchCategory.
func (k *V1SearchCategorySet) Add(i v1.SearchCategory) bool {
	if *k == nil {
		*k = make(map[v1.SearchCategory]struct{})
	}

	oldLen := len(*k)
	(*k)[i] = struct{}{}
	return len(*k) > oldLen
}

// AddMatching is a utility function that adds all the elements that match the given function to the set.
func (k *V1SearchCategorySet) AddMatching(matchFunc func(v1.SearchCategory) bool, elems ...v1.SearchCategory) bool {
	oldLen := len(*k)
	for _, elem := range elems {
		if !matchFunc(elem) {
			continue
		}
		if *k == nil {
			*k = make(map[v1.SearchCategory]struct{})
		}
		(*k)[elem] = struct{}{}
	}
	return len(*k) > oldLen
}

// AddAll adds all elements of type v1.SearchCategory. The return value is true if any new element
// was added.
func (k *V1SearchCategorySet) AddAll(is ...v1.SearchCategory) bool {
	if len(is) == 0 {
		return false
	}
	if *k == nil {
		*k = make(map[v1.SearchCategory]struct{})
	}

	oldLen := len(*k)
	for _, i := range is {
		(*k)[i] = struct{}{}
	}
	return len(*k) > oldLen
}

// Remove removes an element of type v1.SearchCategory.
func (k *V1SearchCategorySet) Remove(i v1.SearchCategory) bool {
	if len(*k) == 0 {
		return false
	}

	oldLen := len(*k)
	delete(*k, i)
	return len(*k) < oldLen
}

// RemoveAll removes the given elements.
func (k *V1SearchCategorySet) RemoveAll(is ...v1.SearchCategory) bool {
	if len(*k) == 0 {
		return false
	}

	oldLen := len(*k)
	for _, i := range is {
		delete(*k, i)
	}
	return len(*k) < oldLen
}

// RemoveMatching removes all elements that match a given predicate.
func (k *V1SearchCategorySet) RemoveMatching(pred func(v1.SearchCategory) bool) bool {
	if len(*k) == 0 {
		return false
	}

	oldLen := len(*k)
	for elem := range *k {
		if pred(elem) {
			delete(*k, elem)
		}
	}
	return len(*k) < oldLen
}

// Contains returns whether the set contains an element of type v1.SearchCategory.
func (k V1SearchCategorySet) Contains(i v1.SearchCategory) bool {
	_, ok := k[i]
	return ok
}

// Cardinality returns the number of elements in the set.
func (k V1SearchCategorySet) Cardinality() int {
	return len(k)
}

// IsEmpty returns whether the underlying set is empty (includes uninitialized).
func (k V1SearchCategorySet) IsEmpty() bool {
	return len(k) == 0
}

// Clone returns a copy of this set.
func (k V1SearchCategorySet) Clone() V1SearchCategorySet {
	if k == nil {
		return nil
	}
	cloned := make(map[v1.SearchCategory]struct{}, len(k))
	for elem := range k {
		cloned[elem] = struct{}{}
	}
	return cloned
}

// Difference returns a new set with all elements of k not in other.
func (k V1SearchCategorySet) Difference(other V1SearchCategorySet) V1SearchCategorySet {
	if len(k) == 0 || len(other) == 0 {
		return k.Clone()
	}

	retained := make(map[v1.SearchCategory]struct{}, len(k))
	for elem := range k {
		if !other.Contains(elem) {
			retained[elem] = struct{}{}
		}
	}
	return retained
}

// Intersect returns a new set with the intersection of the members of both sets.
func (k V1SearchCategorySet) Intersect(other V1SearchCategorySet) V1SearchCategorySet {
	maxIntLen := len(k)
	smaller, larger := k, other
	if l := len(other); l < maxIntLen {
		maxIntLen = l
		smaller, larger = larger, smaller
	}
	if maxIntLen == 0 {
		return nil
	}

	retained := make(map[v1.SearchCategory]struct{}, maxIntLen)
	for elem := range smaller {
		if _, ok := larger[elem]; ok {
			retained[elem] = struct{}{}
		}
	}
	return retained
}

// Union returns a new set with the union of the members of both sets.
func (k V1SearchCategorySet) Union(other V1SearchCategorySet) V1SearchCategorySet {
	if len(k) == 0 {
		return other.Clone()
	} else if len(other) == 0 {
		return k.Clone()
	}

	underlying := make(map[v1.SearchCategory]struct{}, len(k)+len(other))
	for elem := range k {
		underlying[elem] = struct{}{}
	}
	for elem := range other {
		underlying[elem] = struct{}{}
	}
	return underlying
}

// Equal returns a bool if the sets are equal
func (k V1SearchCategorySet) Equal(other V1SearchCategorySet) bool {
	thisL, otherL := len(k), len(other)
	if thisL == 0 && otherL == 0 {
		return true
	}
	if thisL != otherL {
		return false
	}
	for elem := range k {
		if _, ok := other[elem]; !ok {
			return false
		}
	}
	return true
}

// AsSlice returns a slice of the elements in the set. The order is unspecified.
func (k V1SearchCategorySet) AsSlice() []v1.SearchCategory {
	if len(k) == 0 {
		return nil
	}
	elems := make([]v1.SearchCategory, 0, len(k))
	for elem := range k {
		elems = append(elems, elem)
	}
	return elems
}

// GetArbitraryElem returns an arbitrary element from the set.
// This can be useful if, for example, you know the set has exactly one
// element, and you want to pull it out.
// If the set is empty, the zero value is returned.
func (k V1SearchCategorySet) GetArbitraryElem() (arbitraryElem v1.SearchCategory) {
	for elem := range k {
		arbitraryElem = elem
		break
	}
	return arbitraryElem
}

// AsSortedSlice returns a slice of the elements in the set, sorted using the passed less function.
func (k V1SearchCategorySet) AsSortedSlice(less func(i, j v1.SearchCategory) bool) []v1.SearchCategory {
	slice := k.AsSlice()
	if len(slice) < 2 {
		return slice
	}
	// Since we're generating the code, we might as well use sort.Sort
	// and avoid paying the reflection penalty of sort.Slice.
	sortable := &sortableV1SearchCategorySlice{slice: slice, less: less}
	sort.Sort(sortable)
	return sortable.slice
}

// Clear empties the set
func (k *V1SearchCategorySet) Clear() {
	*k = nil
}

// Freeze returns a new, frozen version of the set.
func (k V1SearchCategorySet) Freeze() FrozenV1SearchCategorySet {
	return NewFrozenV1SearchCategorySetFromMap(k)
}

// NewV1SearchCategorySet returns a new thread unsafe set with the given key type.
func NewV1SearchCategorySet(initial ...v1.SearchCategory) V1SearchCategorySet {
	underlying := make(map[v1.SearchCategory]struct{}, len(initial))
	for _, elem := range initial {
		underlying[elem] = struct{}{}
	}
	return underlying
}

type sortableV1SearchCategorySlice struct {
	slice []v1.SearchCategory
	less  func(i, j v1.SearchCategory) bool
}

func (s *sortableV1SearchCategorySlice) Len() int {
	return len(s.slice)
}

func (s *sortableV1SearchCategorySlice) Less(i, j int) bool {
	return s.less(s.slice[i], s.slice[j])
}

func (s *sortableV1SearchCategorySlice) Swap(i, j int) {
	s.slice[j], s.slice[i] = s.slice[i], s.slice[j]
}

// A FrozenV1SearchCategorySet is a frozen set of v1.SearchCategory elements, which
// cannot be modified after creation. This allows users to use it as if it were
// a "const" data structure, and also makes it slightly more optimal since
// we don't have to lock accesses to it.
type FrozenV1SearchCategorySet struct {
	underlying map[v1.SearchCategory]struct{}
}

// NewFrozenV1SearchCategorySetFromMap returns a new frozen set from the set-style map.
func NewFrozenV1SearchCategorySetFromMap(m map[v1.SearchCategory]struct{}) FrozenV1SearchCategorySet {
	if len(m) == 0 {
		return FrozenV1SearchCategorySet{}
	}
	underlying := make(map[v1.SearchCategory]struct{}, len(m))
	for elem := range m {
		underlying[elem] = struct{}{}
	}
	return FrozenV1SearchCategorySet{
		underlying: underlying,
	}
}

// NewFrozenV1SearchCategorySet returns a new frozen set with the provided elements.
func NewFrozenV1SearchCategorySet(elements ...v1.SearchCategory) FrozenV1SearchCategorySet {
	underlying := make(map[v1.SearchCategory]struct{}, len(elements))
	for _, elem := range elements {
		underlying[elem] = struct{}{}
	}
	return FrozenV1SearchCategorySet{
		underlying: underlying,
	}
}

// Contains returns whether the set contains the element.
func (k FrozenV1SearchCategorySet) Contains(elem v1.SearchCategory) bool {
	_, ok := k.underlying[elem]
	return ok
}

// Cardinality returns the cardinality of the set.
func (k FrozenV1SearchCategorySet) Cardinality() int {
	return len(k.underlying)
}

// IsEmpty returns whether the underlying set is empty (includes uninitialized).
func (k FrozenV1SearchCategorySet) IsEmpty() bool {
	return len(k.underlying) == 0
}

// AsSlice returns the elements of the set. The order is unspecified.
func (k FrozenV1SearchCategorySet) AsSlice() []v1.SearchCategory {
	if len(k.underlying) == 0 {
		return nil
	}
	slice := make([]v1.SearchCategory, 0, len(k.underlying))
	for elem := range k.underlying {
		slice = append(slice, elem)
	}
	return slice
}

// AsSortedSlice returns the elements of the set as a sorted slice.
func (k FrozenV1SearchCategorySet) AsSortedSlice(less func(i, j v1.SearchCategory) bool) []v1.SearchCategory {
	slice := k.AsSlice()
	if len(slice) < 2 {
		return slice
	}
	// Since we're generating the code, we might as well use sort.Sort
	// and avoid paying the reflection penalty of sort.Slice.
	sortable := &sortableV1SearchCategorySlice{slice: slice, less: less}
	sort.Sort(sortable)
	return sortable.slice
}
