package iis

var getWebSiteList = ` try{
	$web=Get-Website| Select Name, ID, State, PhysicalPath, ApplicationPool
    if (!($web -is [array])){
        $web = @($web)
    }
	@{Ok=$True; Message= $web}|ConvertTo-Json
	}
catch {
	@{Ok=$False; Error= $_.Exception.Message}|ConvertTo-Json
} 
`

var getWebSite = `try{
	$web=Get-Website -Name %s| Select Name, ID, State, PhysicalPath, ApplicationPool
	@{Ok=$True; Message= $web}|ConvertTo-Json -Depth 2
	}
	catch {
		@{Ok=$False; Error= $_.Exception.Message}|ConvertTo-Json
		}`

var stopWebSite = `try{
	$web=Get-Website -Name %s| Stop-Website
	@{Ok=$True; Message= "stopped"}|ConvertTo-Json -Depth 2
	}
	catch {
		@{Ok=$False; Error= $_.Exception.Message}|ConvertTo-Json
	}`

var startWebSite = `try{
	$web=Get-Website -Name %s| Start-Website
	@{Ok=$True; Message= "started"}|ConvertTo-Json -Depth 2
	}
	catch {
		@{Ok=$False; Error= $_.Exception.Message}|ConvertTo-Json
	}`

var getAppPoolList = ` try {
    Import-Module WebAdministration
    $AppPools = Get-ChildItem IIS:\\AppPools | Select Name, State, ManagedRunTimeVersion
    @{Ok=$true; Message= $AppPools}|ConvertTo-Json -Depth 2
} catch {
    @{Ok=$false; Error= $_.Exception.Message}|ConvertTo-Json -Depth 1
}`

var startWebAppPool = ` try {
    Import-Module WebAdministration
    Start-WebAppPool -Name %s 
    @{Ok=$true; Message= "started"}|ConvertTo-Json -Depth 2
} catch {
    @{Ok=$false; Error= $_.Exception.Message}|ConvertTo-Json -Depth 1
}`

var stopWebAppPool = ` try {
    Import-Module WebAdministration
    Stop-WebAppPool -Name %s 
    @{Ok=$true; Message= "started"}|ConvertTo-Json -Depth 2
} catch {
    @{Ok=$false; Error= $_.Exception.Message}|ConvertTo-Json -Depth 1
}`
