tipz.controller('connection', ['$scope', '$http', function ($scope, $http) {    
    $scope.title = "Connexion";
    
    
    $scope.connectionUserID;
    $scope.connectionPassword;
    
    $scope.connectionClick = function() {
    
    }

    
    $scope.redirection2Inscription = function() {
    	$scope.page = "inscription.html";
    }
}]);
