var tipz = angular.module('Tipz', []);


tipz.controller('tipzHandler', ['$scope', '$http', function ($scope, $http) {    
    $scope.title = "Accueil";

    //Get the table with all project and display it
    var projectList = {content:null};
    var receivedData = {content:null};

   $http.get("test.json")
    .success(function(data, status, headers, config) {
    	console.log(data);
    	receivedData.content = data;
    })
    .error(function(data, status, headers, config) {
        console.log("Error When Charging the JSON of all project from the Server");
    });
    projectList.content = receivedData.data;
}]);
