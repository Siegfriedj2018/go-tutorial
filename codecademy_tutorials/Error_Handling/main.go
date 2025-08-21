package main

import (
	"errors"
	"fmt"
)

type Item struct {
	name     string
	quantity int
}

type Inventory map[string]Item

const MaxInventorySize = 10

func addItem(inv Inventory, name string, quantity int) error {
	fmt.Printf("Adding %s to inventory...\n", name)
	if _, exists := inv[name]; exists {
		return fmt.Errorf("This %s already exists, Please try again.", name)
	}
	if len(inv) >= MaxInventorySize {
		fullInv := InventoryFullError{
			inventorySize: len(inv),
		}
		return fullInv
	}
	inv[name] = Item{name, quantity}
	return nil
}

func wrapError(err error, message string) error {
	return fmt.Errorf("%s, %w", message, err)
}

func printErrorChain(err error) {
	if err == nil {
		fmt.Println("No Errors occurred.")
		return
	}
	fmt.Printf("Error: %v\n", err)
	if unwrapped := errors.Unwrap(err); unwrapped != nil {
		printErrorChain(errors.Unwrap(err))
	}
}

func removeItem(inv Inventory, name string) error {
	if _, exists := inv[name]; !exists {
		return wrapError(errors.New("Does not exist"), fmt.Sprintf("%s does not exists in inventory, Please try again", name))
	}
	fmt.Printf("Removing %s from inventory...\n", name)
	delete(inv, name)
	return nil
}

type InventoryFullError struct {
	inventorySize int
}

func (e InventoryFullError) Error() string {
	return fmt.Sprintf("Your inventory is full, Max size: %d, currrent size: %d", MaxInventorySize, e.inventorySize)
}

func updateItemQuantity(inv Inventory, name string, quantity int) {
	item, exists := inv[name]
	if !exists {
		panic(fmt.Sprintf("Attempted to update a non-existent item: %s", name))
	}
	item.quantity = quantity
	inv[name] = item
}

func main() {
	inventory := make(Inventory)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from a panic: %v\n", r)
		}
	}()

	if err := addItem(inventory, "apple", 5); err != nil {
		fmt.Printf("Error adding apple: %v\n", err)
	}

	if err := addItem(inventory, "banana", 7); err != nil {
		fmt.Printf("Error adding banana: %v\n", err)
	}

	fmt.Println("\nRemoving non-existent item from inventory: orange")
	if err := removeItem(inventory, "orange"); err != nil {
		printErrorChain(err)
	}

	updateItemQuantity(inventory, "apple", 10)

  for i := 0; i <= 9; i++ {
		itemName := fmt.Sprintf("item%d", i)
	  if err := addItem(inventory, itemName, 1); err != nil {
		  fmt.Printf("Error adding new item: %v\n", err)
	  }
  }

	fmt.Println("\nAttempting to update non-existent item: orange")
	updateItemQuantity(inventory, "orange", 5)
}