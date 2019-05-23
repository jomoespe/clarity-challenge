package main

var ProcessLog = processLog

func ExampleProcessLog() {
	config := &config{
		"../../test/input-file-10000.txt",
		int64(1565647204351),
		int64(1565687511867),
		"Aadvik",
		false,
	}

	processLog(config)
	// Output:
	//Matina
	// Zinn
	// Manit
}
