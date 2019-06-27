package MidCloudMAPEK

func Plan(chanPlanner chan CloudService, chanExecutor chan CloudService) {
	for {
		chanExecutor <- <-chanPlanner
	}
}
