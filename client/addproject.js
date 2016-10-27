var tipz = angular.module('Tipz', []);


tipz.controller('tipzHandler', ['$scope', '$http', function ($scope, $http) {    
    $scope.title = "Mon Project";
    
    $scope.rewards = [];
    newProject = [];
    
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
    
    $scope.deleteRewardClick = function(p) {
    	var a = $scope.rewards.indexOf("p.name");
    	var b = $scope.rewards.indexOf("p.value");
    	var c = $scope.rewards.indexOf("p.desc");
    	if ((a == b) && (a == c)) {
    		$scope.rewards.splice(a,1);
    		console.log("The reward has been removed from the list of rewards");
    	}
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
				newProject.author = "";
				newProject.mail = "";

    
			/*    $http.post("test.json")
    			.success(function(data, status, headers, config) {
    				console.log(data);
    				receivedData.content = data;
    			})
    			.error(function(data, status, headers, config) {
        			console.log("Error When Charging the JSON of all project from the Server");
    			}); */
    			console.log("Project added");
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