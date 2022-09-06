package main

func main() {
	a := App{}
	a.Initialize("vineethravi", "postgres")
	a.Run(":8010")
}
