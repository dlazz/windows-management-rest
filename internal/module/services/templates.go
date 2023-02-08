package services

var getServiceList = `try {
    $s = Get-WmiObject win32_service| Select Name, DisplayName, Description, ProcessId, Status, StartMode, State
    @{Ok= $True; Message=$s }|ConvertTo-Json -Depth 2
    }
catch {
	@{Ok= $False; Error=$_.Exception.Message }|ConvertTo-Json -Depth 2
} `

var getService = `try {
    $s = Get-WmiObject win32_service |?{$_.Name -eq "%s"}| Select Name, DisplayName, Description, ProcessId, Status, StartMode, State
    @{Ok= $True; Message=$s }|ConvertTo-Json -Depth 2
    }
catch {
	@{Ok= $False; Error=$_.Exception.Message }|ConvertTo-Json -Depth 2
} `

var stopService = `try{
	Get-Service -Name %s|Stop-Service -Force
	@{Ok= $True; Message= "Stopped"}|ConvertTo-Json -Depth 2
}
catch {
	@{Ok= $False; Error= $_.Exception.Message }
}`

var startService = `try{
	Get-Service -Name %s|Start-Service
	@{Ok= $True; Message= "Started"}|ConvertTo-Json -Depth 2
}
catch {
	@{Ok= $False; Error= $_.Exception.Message }
}`

var restartService = `try{
	Get-Service -Name %s|Restart-Service -Force
	@{Ok= $True; Message= "Restarted"}|ConvertTo-Json -Depth 2
}
catch {
	@{Ok= $False; Error= $_.Exception.Message }
}`
