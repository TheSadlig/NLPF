var tipz = angular.module('Tipz', []);


tipz.controller('tipzHandler', ['$scope', '$http', function ($scope, $http) {    
    $scope.title = "Accueil";
    
    //Get the table with all project and display it
    var projectList =[];
    var projectInfo = [];
    
   $http.post("wrongfilename.htm")
    .then(data) {
    	console.log(data);
    	data.data = "toto";
        $scope.content = response.data;
    }

}]);
