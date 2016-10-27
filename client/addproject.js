tipz.controller('addProject', ['$scope', '$http', function ($scope, $http) {    
    $scope.title = "Mon Project";
    console.log("plop")
    $scope.rewards = [];
    newProject = [];
    $scope.newRewardValue = 0;
    
    $scope.newRewardClick = function() {
    	if (($scope.newRewardName != "") && ($scope.newRewardValue != "") && ($scope.newRewardDescription != "") && ($scope.newRewardName != null) && ($scope.newRewardValue != null) && ($scope.newRewardDescription != null) && ($scope.newRewardValue > 0)) {

    		$scope.rewards.push({name:$scope.newRewardName, value:$scope.newRewardValue, desc:$scope.newRewardDescription});
    		$scope.newRewardName = "";
    		$scope.newRewardValue = "";
    		$scope.newRewardDescription = "";
    		console.log("Reward added");

    	}
    	else {
    		console.log("One or several label are empty or incorrect");
    	}
    }
    
    $scope.loadFeed = function(e, p) {
        var a = 0;
        for(var i = 0; i < $scope.rewards.length; i++) {
   			if($scope.rewards[i].name === p) {
     			a = i;
   			}
		}
    	$scope.rewards.splice(a, 1);
    	console.log("The reward has been removed from the list of rewards");
    }
    
    $scope.newProjectClick = function() {
    
    	if (($scope.newProjectName != null) && ($scope.newProjectDescription != null) && ($scope.newProjectName != "") && ($scope.newProjectDescription != "")) {
    
        	if (($scope.rewards.length > 0) || ($scope.rewards == "undifined")) {

    			newProject.name = $scope.newProjectName;
    			newProject.desc = $scope.newProjectDescription;
    			newProject.rewards = $scope.rewards;
    
    			var today = new Date();
				var dd = today.getDate();
				var mm = today.getMonth()+1;
				var yyyy = today.getFullYear();
				if(dd<10) {
    				dd='0'+dd
				} 
				if(mm<10) {
    				mm='0'+mm
				} 
				today = dd+'/'+mm+'/'+yyyy;
				
				newProject.date = today;
				newProject.userID = "";
/*{ 'data' : {
        						'name' : newProject.name,
        						'desc' :  newProject.desc,
        						'rewards' : newProject.rewards,
        						'date' : newProject.date,
        						'userID' : newProject.userID
        					}}*/
        					
				$http({
        			url: 'http://localhost:9090/api/createProject',
        			method: "GET",
        			data: {test : "test"}
    			})
    			.then(function(response) {
        			console.log(response);
        			console.log("Project successfully added to database");
    			}, 
    				function(response) {
    				console.log(response);
        			console.log("Error When posting the JSON of the project to the Server");
    			});

    			console.log("Project added");
    			
    			newProject.name = "";
    			newProject.desc = "";
    			newProject.rewards = "";
    			newProject.userID = "";
			}
			else {
        		console.log("There is no Rewards");
    		}
    	}
    	else {
        	console.log("One or several label are empty");
    	}
	}
}]);