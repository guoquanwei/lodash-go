package lodash

import "reflect"

// TODO
// min  max  sum avg random [minBy maxBy sumBy]

func Min(output interface{}, input interface{}) (err error) {
	useInput, isChain := chainArgConvert(input)
	inputRv := reflect.ValueOf(useInput)

	err = CheckKindErr(`Min`, isChain, reflect.ValueOf(output).Kind().String(), inputRv.Kind().String())
	if err != nil {
		return err
	}
	return nil
	//chainOutputNoSlice(output, input, isChain, inputRv.Index().Interface())
}
