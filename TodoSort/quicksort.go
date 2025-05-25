package TodoSort

import "TodoApp/models"

func Partition(todos []models.Todo, low, high int) ([]models.Todo, int) {
	pivot := todos[high].Priority
	i := low
	for j := low; j < high; j++ {
		if todos[j].Priority < pivot {
			todos[i], todos[j] = todos[j], todos[i]
			i++
		}
	}
	todos[i], todos[high] = todos[high], todos[i]
	return todos, i
}

// quickSort function for Todo slices
func QuickSort(todos []models.Todo, low, high int) []models.Todo {
	if low < high {
		var p int
		todos, p = Partition(todos, low, high)
		todos = QuickSort(todos, low, p-1)
		todos = QuickSort(todos, p+1, high)
	}
	return todos
}

// quickSortStart is a helper function to start the quicksort process
func QuickSortStart(todos []models.Todo) []models.Todo {
	return QuickSort(todos, 0, len(todos)-1)
}

//参考　https://blog.boot.dev/golang/quick-sort-golang/