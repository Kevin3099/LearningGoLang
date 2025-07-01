package main

import (
	"fmt"
)

func main() {
	var items []string

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Add item")
		fmt.Println("2 - View items")
		fmt.Println("3 - Update item")
		fmt.Println("4 - Delete item")
		fmt.Println("0 - Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var input string
			fmt.Print("Enter item to add: ")
			fmt.Scan(&input)
			items = append(items, input)
			fmt.Println("Added.")

		case 2:
			fmt.Println("Current items:")
			for i, item := range items {
				fmt.Printf("%d: %s\n", i, item)
			}

		case 3:
			var index int
			var newValue string
			fmt.Print("Enter index to update: ")
			fmt.Scan(&index)
			fmt.Print("Enter new value: ")
			fmt.Scan(&newValue)
			items[index] = newValue
			fmt.Println("Updated.")

		case 4:
			var index int
			fmt.Print("Enter index to delete: ")
			fmt.Scan(&index)
			items = append(items[:index], items[index+1:]...)
			fmt.Println("Deleted.")

		case 0:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
