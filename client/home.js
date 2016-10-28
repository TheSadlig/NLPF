tipz.controller('home', ['$scope', '$http', function ($scope, $http) {    
    $scope.title = "Accueil";

    //Get the table with all project and display it
    $scope.projects = [];

   	$http.get('http://localhost:9090/api/getProjects')
    .success(function(data, status, headers, config) {
    	//console.log(data.data);
    	$scope.projects = data.data;
    })
    .error(function(data, status, headers, config) {
        console.log("Error When Charging the JSON of all project from the Server");
    });

	$scope.loadFeed = function(e, p) {
		console.log($scope.$parent);
		$scope.projectID = p.name;
		//console.log($scope.project);
		$scope.$parent.$parent.page = "projectdescription.html";

    }
    
}]);