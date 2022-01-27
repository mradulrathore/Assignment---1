package service

import "log"

const (
	ChunkSize = 100
)

func getItemsFromDB(repo *Repository) *ListItems {
	list, err := repo.GetItems()
	if err != nil {
		log.Println(err)
	}
	return list
}

func Produce(list *ListItems, itemDB chan Item) {
	// split into chunks of 1000
	chunks := GetChunk(list.Items)

	for i := range chunks {
		itemDB <- chunks[i]
		items = append(items, chunks[i])

	}
}

func GetChunk(input []Item) []Item {
	var result []Item

	boundary := len(input)
	index := 0
	for index = 0; boundary >= ChunkSize; index += ChunkSize {
		boundary -= ChunkSize
		lastIndex := index + ChunkSize
		result = append(result, input[index:lastIndex]...)
	}
	boundary = len(input) % ChunkSize
	if boundary > 0 {
		lastIndex := index + boundary
		result = append(result, input[index:lastIndex]...)
	}

	return result
}
