package behavior_chain

import "log"

type Filter interface {
	filter(cm string) error
}

type NumFilter struct {
	next Filter
}

func (num *NumFilter) filter(cm string) error {
	var err error

	log.Println("num filter")
	if err != nil {
		return err
	}

	// 岂不是每次都需要判断再next，优雅一点？
	if num.next != nil {
		return num.next.filter(cm)
	}
	return nil
}

type EnFilter struct {
	next Filter
}

func (en *EnFilter) filter(cm string) error {
	var err error

	log.Println("en filter")
	if err != nil {
		return err
	}

	if en.next != nil {
		return en.next.filter(cm)
	}

	return nil
}
