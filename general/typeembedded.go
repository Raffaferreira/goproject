package general

import "fmt"

func InterfaceTypeEmbeddingMethodOverriding() {
	//Create an Instance of Mobile Type and Call Methods
	mobilePrd := Mobile{
		Product{
			"Apple iPhone 6s (Rose Gold, 32 GB)  (2 GB RAM)",
			"Handset, Apple EarPods with Remote and Mic, Lightning to USB Cable",
			40999, 143, 10,
		},
		[]string{"2 GB RAM", "4.7 inch Retina HD Display", "12MP Primary Camera"},
		[]string{"32 GB ROM", "4.7 inch Retina HD Display", "5MP Front"},
	}

	//Create an Instance of Shirts Type and Call Methods
	shirtPrd := Shirts{
		Product{
			"Reebok Striped Men's Round Neck Grey T-Shirt",
			"Machine Wash as per Tag, Do not Dry Clean, Do not Bleach",
			1124, 400, 25,
		},
		"XL", "Striped", "Half Sleeve", "Cotton",
	}

	//Create an Instance of Catalog Type
	category := Catalog{
		"Amazon",
		[]Information{mobilePrd, shirtPrd},
	}
	category.CatalogDetails()
}

type Information interface {
	General()
	Attributes()
	Inventory()
}

type Product struct {
	Name, Description string
	Weight, Price     float64
	Stock             int
}

func (prd Product) General() {
	fmt.Printf("\n%s", prd.Name)
	fmt.Printf("\n%s\n", prd.Description)
	fmt.Println(prd.Price)
}

func (prd Product) Attributes() {
	fmt.Println(prd.Weight)
}

type Mobile struct {
	Product
	DisplayFeatures   []string
	ProcessorFeatures []string
}

func (mob Mobile) Attributes() {
	mob.Product.Attributes()
	fmt.Println("\nDisplay Features:")
	for _, key := range mob.DisplayFeatures {
		fmt.Println(key)
	}
	fmt.Println("\nProcessor Features:")
	for _, key := range mob.ProcessorFeatures {
		fmt.Println(key)
	}
}

type Shirts struct {
	Product
	Size, Pattern, Sleeve, Fabric string
}

func (shrt Shirts) Attributes() {
	fmt.Println("\nSpecification:")
	fmt.Println(shrt.Size)
	fmt.Println(shrt.Pattern)
	fmt.Println(shrt.Sleeve)
	fmt.Println(shrt.Fabric)
}

func (prd Product) Inventory() {
	fmt.Println(prd.Stock)
}

// A struct named Catalog is declared to represent catalog of various products.
// The type of Details field uses a slice of the Information interface
type Catalog struct {
	Name     string
	Deatails []Information
}

func (c Catalog) CatalogDetails() {
	fmt.Println("**************************************************\n")
	fmt.Printf("\nStore : %s \n", c.Name)
	for _, key := range c.Deatails {
		fmt.Println("===================Products==================")
		key.General()
		key.Attributes()
		key.Inventory()
		fmt.Println("=============================================")
	}
}
