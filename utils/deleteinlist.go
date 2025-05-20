package utils

func DeleteInList[T any](list []T, index int) []T {
	newList := list[:index]
	newList = append(newList, list[index+1:]...)
	return newList;
}