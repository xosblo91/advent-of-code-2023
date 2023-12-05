package sliceutils

func MakeBatches[T any](input []T, batchSize int) [][]T {
	if len(input) == 0 || batchSize == 0 {
		return nil
	}

	batches := make([][]T, 0, (len(input)+batchSize-1)/batchSize)

	for batchSize < len(input) {
		input, batches = input[batchSize:], append(batches, input[0:batchSize:batchSize])
	}
	batches = append(batches, input)

	return batches
}
