package modes

// Switch a given mode and execute a corresponding function
func SwitchMode(mode, file string) {
	switch mode {
	case "c":
		{
			CreateDocumentation(file)
		}
	case "d":
		{
			DeleteDocumentation(file)
		}
	case "e":
		{
			EditDocumentation(file)
		}
	}
}
