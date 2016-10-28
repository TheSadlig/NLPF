tipz.controller('connection', ['$scope', '$http', function ($scope, $parent, $http) {    
    $scope.title = "Connexion";
    
    
    $scope.connectionUserID;
    $scope.connectionPassword;

        $scope.redirection2Inscription = function() {
        //console.log($scope.$parent);
    	$scope.$parent.$parent.page = "inscription.html";
    }

    $scope.connectionClick = function() {
    
    }

}]);
