package main

type Set struct {
	collections []interface{}
}

func (s *Set) Collections() []interface{} {
	return s.collections
}

// Adds elements to the set.
// If the element is already present in the set, false is returned, otherwise true.
func (s *Set) Add(value interface{}) bool {
	if _, exist := s.Contains(value); !exist {
		s.collections = append(s.collections, value)
		return true
	}
	return false
}

// Returns index, true if the set contains the specified element.
// Otherwise returns index, false.
func (s *Set) Contains(value interface{}) (int, bool) {
	for k, v := range s.collections {
		if v == value {
			return k, true
		}
	}
	return 0, false
}

// Removes the specified element from the set and returns true.
// If there is no element and it is not deleted, returns false.
func (s *Set) Remove(value interface{}) bool {
	if i, exist := s.Contains(value); exist {
		length := s.Size()

		s.collections[i] = s.collections[length-1]
		s.collections = s.collections[:length-1]

		return true
	}
	return false
}

func (s *Set) Size() int {
	return len(s.collections)
}

// Returns the set obtained by the operation of combining it with the specified one.
func (s *Set) Union(set *Set) *Set {
	union := &Set{}

	for _, v := range s.Collections() {
		union.Add(v)
	}

	for _, v := range set.Collections() {
		union.Add(v)
	}
	return union
}

// Returns the set obtained by the operation of its intersection with the specified one.
func (s *Set) Intersection(set *Set) *Set {
	intersection := &Set{}

	for _, v := range s.Collections() {
		if _, exist := set.Contains(v); exist {
			intersection.Add(v)
		}
	}
	return intersection
}

// Returns a set that is the difference between the current one and the specified one.
func (s *Set) Difference(set *Set) *Set {
	difference := &Set{}

	for _, v := range s.Collections() {
		if _, exist := set.Contains(v); !exist {
			difference.Add(v)
		}
	}
	return difference
}

// Returns true if the second set is a subset of the first, otherwise returns false.
func (s *Set) Subset(set *Set) bool {
	if s.Size() > set.Size() {
		return false
	}

	return s.Difference(set).Size() == 0
}

// Returns a set that is the symmetric difference of the current one with the specified one.
func (s *Set) SymmetricDifference(set *Set) *Set {
	a := s.Difference(set)
	b := set.Difference(s)

	return a.Union(b)
}
