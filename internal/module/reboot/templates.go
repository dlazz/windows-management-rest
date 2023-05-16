package reboot

const (
	rebootComputer = `try {
		Restart-Computer -Force;
		@{Ok= $True; Message="Restarting" }|ConvertTo-Json -Depth 2
		}
	catch {
		@{Ok= $False; Error=$_.Exception.Message }|ConvertTo-Json -Depth 2
	} `
)
