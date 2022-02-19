package behavior_chain

import "testing"

func TestFilter(t *testing.T) {

	var filter = &NumFilter{
		next: &EnFilter{},
	}

	_ = filter.filter("123")
}
